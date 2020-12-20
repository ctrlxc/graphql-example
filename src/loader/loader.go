package loader

import (
	"app/repository"
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

type Loaders struct {
	batchFuncs map[string]dataloader.BatchFunc
}

func NewLoaders(repo *repository.Repository) *Loaders {
	return &Loaders{
		batchFuncs: map[string]dataloader.BatchFunc{
			shopLoaderKey: newShopLoader(repo),
			bookLoaderKey: newBookLoader(repo),
		},
	}
}

func (c *Loaders) Attach(ctx context.Context) context.Context {
	for key, batchFn := range c.batchFuncs {
		ctx = context.WithValue(ctx, key, dataloader.NewBatchedLoader(batchFn))
	}

	return ctx
}

func getLoader(ctx context.Context, key string) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(key).(*dataloader.Loader)

	if !ok {
		return nil, fmt.Errorf("no loader: %s", key)
	}

	return ldr, nil
}
