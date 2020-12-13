package graph

import (
	"app/loader"
	"app/models"
	"app/paginator"
	"app/repository"
	"app/util"
	"context"
	"fmt"
	"reflect"

	// "errors"
	"app/graph/model"

	// "app/models"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	gid, err := util.FromGlobalID(id)

	if err != nil {
		return nil, err
	}

	v := reflect.ValueOf(r)
	m := v.MethodByName(gid.Type + "ByID")

	if m.Kind() != reflect.Func {
		return nil, fmt.Errorf("unknown type. %s", gid.Type)
	}

	argv := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(id)}

	result := m.Call(argv)

	var node model.Node

	for _, r := range result {
		switch rv := r.Interface().(type) {
		case model.Node:
			node = rv
		case error:
			err = rv
		}
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

func (r *Resolver) ShopByID(ctx context.Context, id string) (*model.Shop, error) {
	realid, err := util.FromGlobalIDInt64(id, "Shop")

	if err != nil {
		return nil, err
	}

	record, err := loader.LoadShop(ctx, realid)

	if err != nil {
		return nil, err
	}

	return &model.Shop{
		ID:        record.GlobalID(),
		ShopName:  record.ShopName.Ptr(),
		CreatedAt: &record.CreatedAt,
		UpdatedAt: &record.UpdatedAt,
	}, nil
}

func (r *Resolver) shops(ctx context.Context, after *string, before *string, first *int, last *int, query string, orderBy []*model.ShopOrder) (*model.ShopConnection, error) {
	pageOrders := model.ShopOrderToPaginatorOrder(orderBy)
	pagenator := paginator.Paginator{after, before, first, last, pageOrders}

	shops, err := models.Shops(
		qm.Where(fmt.Sprintf("%s like ?", models.ShopColumns.ShopName), fmt.Sprintf("%%%s%%", query)),
		pagenator.QueryWhere(),
		pagenator.QueryOrderBy(),
		pagenator.QueryLimit(),
	).All(ctx, r.repo.Db)

	if err != nil {
		return nil, err
	}

	conn := &model.ShopConnection{
		PageInfo: &model.PageInfo{},
	}

	if len(shops) == 0 {
		return conn, nil
	}

	limit := len(shops)
	if limit > pagenator.Limit() {
		limit = pagenator.Limit()
	}

	conn.Edges = make([]*model.ShopEdge, limit)
	conn.Nodes = make([]*model.Shop, limit)

	for i, s := range shops[:limit] {
		cursor, _ := pagenator.CreateEncodedCursor(s)

		node := &model.Shop{
			ID:       s.GlobalID(),
			ShopName: s.ShopName.Ptr(),
		}

		pos := i
		if !pagenator.IsAfter() {
			pos = len(conn.Edges) - i - 1
		}

		conn.Edges[pos] = &model.ShopEdge{Cursor: cursor, Node: node}
		conn.Nodes[pos] = node
	}

	conn.PageInfo.StartCursor = &conn.Edges[0].Cursor
	conn.PageInfo.EndCursor = &conn.Edges[len(conn.Edges)-1].Cursor

	if len(shops) > limit {
		if pagenator.IsAfter() {
			conn.PageInfo.HasNextPage = true
		} else {
			conn.PageInfo.HasPreviousPage = true
		}
	}

	return conn, nil
}

func (r *Resolver) BookByID(ctx context.Context, id string) (*model.Book, error) {
	realid, err := util.FromGlobalIDInt64(id, "Book")

	if err != nil {
		return nil, err
	}

	record, err := loader.LoadBook(ctx, realid)

	if err != nil {
		return nil, err
	}

	return &model.Book{
		ID:        record.GlobalID(),
		BookTitle: record.BookTitle.Ptr(),
	}, nil
}

// func (r *Resolver) books(ctx context.Context, after *string, before *string, first int, last *int, query string, orderBy []*model.BookOrder) ([]*model.Book, error) {
// 	realids, err := util.FromGlobalIDInt64s(ids, "Book")

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

func (r *Resolver) booksByShopID(ctx context.Context, id string) ([]*model.Book, error) {
	realid, err := util.FromGlobalIDInt64(id, "Shop")

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
			ID:        record.GlobalID(),
			BookTitle: &record.BookTitle.String,
		})
	}

	return resp, nil
}
