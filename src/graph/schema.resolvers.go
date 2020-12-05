package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	generated1 "app/graph/generated"
	model1 "app/graph/model"
	"context"
	"fmt"
)

func (r *queryResolver) Node(ctx context.Context, id string) (model1.Node, error) {
	return r.node(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]model1.Node, error) {
	return r.nodes(ctx, ids)
}

func (r *shopResolver) Books(ctx context.Context, obj *model1.Shop) ([]*model1.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

// Shop returns generated1.ShopResolver implementation.
func (r *Resolver) Shop() generated1.ShopResolver { return &shopResolver{r} }

type queryResolver struct{ *Resolver }
type shopResolver struct{ *Resolver }
