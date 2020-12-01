package graph

import (
	"app/loader"
	"app/repository"
	"context"

	// "errors"
	"app/graph/model"
	// "app/models"

	_ "github.com/lib/pq"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	repo *repository.Repository
}

func NewResolver(repo *repository.Repository) *Resolver {
	return &Resolver{repo}
}

func (r *Resolver) shop(ctx context.Context, id string) (*model.Shop, error) {
	realid, err := fromGlobalIDInt64(id, "Shop")

	if err != nil {
		return nil, err
	}

	record, err := loader.LoadShop(ctx, realid)

	if err != nil {
		return nil, err
	}

	return &model.Shop{
		ID:       toGlobalIDInt64("Shop", record.ID),
		ShopName: record.ShopName.Ptr(),
	}, nil
}

func (r *Resolver) shops(ctx context.Context, ids []string) ([]*model.Shop, error) {
	realids, err := fromGlobalIDInt64s(ids, "Shop")

	if err != nil {
		return nil, err
	}

	records, err := loader.LoadShops(ctx, realids)

	if err != nil {
		return nil, err
	}

	resp := make([]*model.Shop, 0, len(records))
	for _, record := range records {
		resp = append(resp, &model.Shop{
			ID:       toGlobalIDInt64("Shop", record.ID),
			ShopName: record.ShopName.Ptr(),
		})
	}

	return resp, nil
}

func (r *Resolver) book(ctx context.Context, id string) (*model.Book, error) {
	realid, err := fromGlobalIDInt64(id, "Book")

	if err != nil {
		return nil, err
	}

	record, err := loader.LoadBook(ctx, realid)

	if err != nil {
		return nil, err
	}

	return &model.Book{
		ID:        toGlobalIDInt64("Book", record.ID),
		BookTitle: record.BookTitle.Ptr(),
	}, nil
}

func (r *Resolver) books(ctx context.Context, ids []string) ([]*model.Book, error) {
	realids, err := fromGlobalIDInt64s(ids, "Book")

	if err != nil {
		return nil, err
	}

	records, err := loader.LoadBooks(ctx, realids)

	if err != nil {
		return nil, err
	}

	resp := make([]*model.Book, 0, len(records))
	for _, record := range records {
		resp = append(resp, &model.Book{
			ID:        toGlobalIDInt64("Book", record.ID),
			BookTitle: record.BookTitle.Ptr(),
		})
	}

	return resp, nil
}

func (r *Resolver) booksByShopID(ctx context.Context, id string) ([]*model.Book, error) {
	realid, err := fromGlobalIDInt64(id, "Shop")

	if err != nil {
		return nil, err
	}

	records, err := loader.LoadBooksByShopID(ctx, realid)

	if err != nil {
		return nil, err
	}

	resp := make([]*model.Book, 0, len(records))
	for _, record := range records {
		resp = append(resp, &model.Book{
			ID:        toGlobalIDInt64("Book", record.ID),
			BookTitle: &record.BookTitle.String,
		})
	}

	return resp, nil
}
