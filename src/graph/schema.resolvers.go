package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	generated1 "app/graph/generated"
	model1 "app/graph/model"
	"context"
)

func (r *queryResolver) Shop(ctx context.Context, id int64) (*model1.Shop, error) {
	return r.shop(ctx, id)
}

func (r *queryResolver) Shops(ctx context.Context, ids []int64) ([]*model1.Shop, error) {
	return r.shops(ctx, ids)
}

func (r *queryResolver) Book(ctx context.Context, id int64) (*model1.Book, error) {
	return r.book(ctx, id)
}

func (r *queryResolver) Books(ctx context.Context, ids []int64) ([]*model1.Book, error) {
	return r.books(ctx, ids)
}

func (r *queryResolver) BooksByShopID(ctx context.Context, shopID int64) ([]*model1.Book, error) {
	return r.booksByShopID(ctx, shopID)
}

func (r *shopResolver) Books(ctx context.Context, obj *model1.Shop) ([]*model1.Book, error) {
	return r.booksByShopID(ctx, obj.ID)
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

// Shop returns generated1.ShopResolver implementation.
func (r *Resolver) Shop() generated1.ShopResolver { return &shopResolver{r} }

type queryResolver struct{ *Resolver }
type shopResolver struct{ *Resolver }
