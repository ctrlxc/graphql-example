package loader

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"app/models"
	"app/repository"

	"github.com/graph-gophers/dataloader"
)

const bookLoaderKey = "bookLoader"

type bookIDKey struct {
	id int64
}

func (key bookIDKey) String() string {
	return fmt.Sprintf("%s/%v", reflect.TypeOf(key).Name(), key.id) // should be global unique
}

func (key bookIDKey) Raw() interface{} {
	return key.id
}

func LoadBook(ctx context.Context, id int64) (*models.Book, error) {
	ldr, err := getLoader(ctx, bookLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, bookIDKey{id: id})()
	if err != nil {
		return nil, err
	}

	return data.(*models.Book), nil
}

func LoadBooks(ctx context.Context, ids []int64) ([]*models.Book, error) {
	ldr, err := getLoader(ctx, bookLoaderKey)
	if err != nil {
		return nil, err
	}

	bookIDs := make(dataloader.Keys, len(ids))
	for i, id := range ids {
		bookIDs[i] = bookIDKey{id: id}
	}

	datas, errs := ldr.LoadMany(ctx, bookIDs)()
	if len(errs) != 0 {
		return nil, errs[0]
	}

	books := make([]*models.Book, len(datas))
	for i, data := range datas {
		books[i] = data.(*models.Book)
	}

	return books, nil
}

func LoadBooksByShopID(ctx context.Context, id int64) ([]*models.Book, error) {
	ldr, err := getLoader(ctx, bookLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, shopIDKey{id: id})()
	if err != nil {
		return nil, err
	}

	return data.([]*models.Book), nil
}

func newBookLoader(repo *repository.Repository) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		bookIDs := make([]int64, 0, len(keys))
		shopIDs := make([]int64, 0, len(keys))

		for _, key := range keys {
			switch key := key.(type) {
			case bookIDKey:
				bookIDs = append(bookIDs, key.id)
			case shopIDKey:
				shopIDs = append(shopIDs, key.id)
			}
		}

		books, _ := repo.BooksByIDs(ctx, bookIDs)
		booksShops, _ := repo.BooksByShopIDs(ctx, shopIDs)

		for i, key := range keys {
			results[i] = &dataloader.Result{Data: nil, Error: nil}

			switch key := key.(type) {
			case bookIDKey:
				for _, book := range books {
					if key.id == book.ID {
						results[i].Data = book
						continue
					}
				}

				if results[i].Data == nil {
					results[i].Error = errors.New("Book not found")
				}
			case shopIDKey:
				results[i].Data = booksShops[key.id]
			}
		}

		return results
	}
}
