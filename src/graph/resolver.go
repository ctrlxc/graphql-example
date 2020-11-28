package graph

import (
	"context"
	"app/loader"
	"app/repository"

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

func (r *Resolver) shop(ctx context.Context, id int64) (*model.Shop, error) {
	record, err := loader.LoadShop(ctx, id)

	if err != nil {
		return nil, err
	}

	return &model.Shop{
		ID:       record.ID,
		ShopName: &record.ShopName.String,
	}, nil
}

func (r *Resolver) shops(ctx context.Context, ids []int64) ([]*model.Shop, error) {
	records, err := loader.LoadShops(ctx, ids)

	if err != nil {
		return nil, err
	}

	resp := make([]*model.Shop, 0, len(records))
	for _, record := range records {
		resp = append(resp, &model.Shop{
			ID:       record.ID,
			ShopName: &record.ShopName.String,
		})
	}

	return resp, nil
}

func (r *Resolver) book(ctx context.Context, id int64) (*model.Book, error) {
	record, err := loader.LoadBook(ctx, id)

	if err != nil {
		return nil, err
	}

	return &model.Book{
		ID:        record.ID,
		BookTitle: &record.BookTitle.String,
	}, nil
}

func (r *Resolver) books(ctx context.Context, ids []int64) ([]*model.Book, error) {
	records, err := loader.LoadBooks(ctx, ids)

	if err != nil {
		return nil, err
	}

	resp := make([]*model.Book, 0, len(records))
	for _, record := range records {
		resp = append(resp, &model.Book{
			ID:        record.ID,
			BookTitle: &record.BookTitle.String,
		})
	}

	return resp, nil
}

func (r *Resolver) booksByShopID(ctx context.Context, shopID int64) ([]*model.Book, error) {
	records, err := loader.LoadBooksByShopID(ctx, shopID)

	if err != nil {
		return nil, err
	}

	resp := make([]*model.Book, 0, len(records))
	for _, record := range records {
		resp = append(resp, &model.Book{
			ID:        record.ID,
			BookTitle: &record.BookTitle.String,
		})
	}

	return resp, nil
}
