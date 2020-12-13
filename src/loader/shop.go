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

const shopLoaderKey = "shopLoader"

type shopIDKey struct {
	id int64
}

func (key shopIDKey) String() string {
	return fmt.Sprintf("%s/%v", reflect.TypeOf(key).Name(), key.id) // should be global unique
}

func (key shopIDKey) Raw() interface{} {
	return key.id
}

func LoadShop(ctx context.Context, id int64) (*models.Shop, error) {
	ldr, err := getLoader(ctx, shopLoaderKey)
	if err != nil {
		return nil, err
	}

	data, err := ldr.Load(ctx, shopIDKey{id: id})()
	if err != nil {
		return nil, err
	}

	return data.(*models.Shop), nil
}

func LoadShops(ctx context.Context, ids []int64) ([]*models.Shop, error) {
	ldr, err := getLoader(ctx, shopLoaderKey)
	if err != nil {
		return nil, err
	}

	shopIDs := make(dataloader.Keys, len(ids))
	for i, id := range ids {
		shopIDs[i] = shopIDKey{id: id}
	}

	datas, errs := ldr.LoadMany(ctx, shopIDs)()
	if len(errs) != 0 {
		return nil, errs[0]
	}

	shops := make([]*models.Shop, len(datas))
	for i, data := range datas {
		shops[i] = data.(*models.Shop)
	}

	return shops, nil
}

func newShopLoader(repo *repository.Repository) dataloader.BatchFunc {
	return func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		results := make([]*dataloader.Result, len(keys))
		shopIDs := make([]int64, len(keys))

		for i, key := range keys {
			shopIDs[i] = key.(shopIDKey).id
		}

		shops, _ := repo.ShopsByIDs(ctx, shopIDs)

		for i, key := range keys {
			results[i] = &dataloader.Result{Data: nil, Error: nil}

			for _, shop := range shops {
				if key.(shopIDKey).id == shop.ID {
					results[i].Data = shop
					continue
				}
			}

			if results[i].Data == nil {
				results[i].Error = errors.New("Shop not found")
			}
		}

		return results
	}
}
