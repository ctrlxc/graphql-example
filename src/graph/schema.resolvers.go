package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/graph/model"
	"context"
)

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	return r.node(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]model.Node, error) {
	return r.nodes(ctx, ids)
}

func (r *queryResolver) Shops(ctx context.Context, after *string, before *string, first *int, last *int, query string, orderBy []*model.ShopOrder) (*model.ShopConnection, error) {
	return r.shops(ctx, after, before, first, last, query, orderBy)
}

func (r *queryResolver) Books(ctx context.Context, after *string, before *string, first *int, last *int, query string, orderBy []*model.BookOrder) (*model.BookConnection, error) {
	return r.books(ctx, after, before, first, last, query, orderBy)
}

func (r *shopResolver) Books(ctx context.Context, obj *model.Shop) ([]*model.Book, error) {
	return r.booksByShopID(ctx, obj.ID)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Shop returns generated.ShopResolver implementation.
func (r *Resolver) Shop() generated.ShopResolver { return &shopResolver{r} }

type queryResolver struct{ *Resolver }
type shopResolver struct{ *Resolver }
