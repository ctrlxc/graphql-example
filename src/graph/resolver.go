package graph

import (
	"app/loader"
	"app/repository"
	"context"
	"fmt"

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

func (r *Resolver) node(ctx context.Context, id string) (model.Node, error) {
	gid, err := fromGlobalID(id)

	if err != nil {
		return nil, err
	}

	var node model.Node = nil

	switch gid.Type {
	case "Shop":
		node, err = r.shop(ctx, id)
	case "Book":
		node, err = r.book(ctx, id)
	default:
		return nil, fmt.Errorf("unknown type. %s", gid.Type)
	}

	return node, err
}

func (r *Resolver) nodes(ctx context.Context, ids []string) ([]model.Node, error) {
	nodes := make([]model.Node, len(ids))

	for i, id := range ids {
		node, err := r.node(ctx, id)

		if err != nil {
			return nil, err
		}

		nodes[i] = node
	}

	return nodes, nil
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

// func (r *Resolver) shops(ctx context.Context, after *string, before *string, first int, last *int, query string, orderBy []*model1.ShopOrder) ([]*model.Shop, error) {
// 	realids, err := fromGlobalIDInt64s(ids, "Shop")

// 	if err != nil {
// 		return nil, err
// 	}

// 	records, err := loader.LoadShops(ctx, realids)

// 	if err != nil {
// 		return nil, err
// 	}

// 	resp := make([]*model.Shop, 0, len(records))
// 	for _, record := range records {
// 		resp = append(resp, &model.Shop{
// 			ID:       toGlobalIDInt64("Shop", record.ID),
// 			ShopName: record.ShopName.Ptr(),
// 		})
// 	}

// 	return resp, nil
// }

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

// func (r *Resolver) books(ctx context.Context, after *string, before *string, first int, last *int, query string, orderBy []*model1.BookOrder) ([]*model.Book, error) {
// 	realids, err := fromGlobalIDInt64s(ids, "Book")

// 	if err != nil {
// 		return nil, err
// 	}

// 	records, err := loader.LoadBooks(ctx, realids)

// 	if err != nil {
// 		return nil, err
// 	}

// 	resp := make([]*model.Book, 0, len(records))
// 	for _, record := range records {
// 		resp = append(resp, &model.Book{
// 			ID:        toGlobalIDInt64("Book", record.ID),
// 			BookTitle: record.BookTitle.Ptr(),
// 		})
// 	}

// 	return resp, nil
// }

func (r *Resolver) booksByShopID(ctx context.Context, after *string, before *string, first int, last *int, id string, orderBy []*model.BookOrder) ([]*model.Book, error) {
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
