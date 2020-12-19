package graph

import (
	"app/globalid"
	"app/graph/model"
	"app/loader"
	"app/pagination"
	"app/repository"
	"context"
	"fmt"
	"reflect"
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
	gid, err := globalid.FromGlobalID(id)

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
	shopId, err := globalid.FromGlobalIDInt64(id, "Shop")

	if err != nil {
		return nil, err
	}

	record, err := loader.LoadShop(ctx, shopId)

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
	paginator := pagination.NewPaginator(
		after,
		before,
		first,
		last,
		model.ShopOrderToPaginationOrders(orderBy),
	)

	shops, err := r.repo.ShopsByName(ctx, query, paginator)

	if err != nil {
		return nil, err
	}

	conn := &model.ShopConnection{}

	if len(shops) == 0 {
		return conn, nil
	}

	limit := len(shops)
	if limit > paginator.Limit() {
		limit = paginator.Limit()
	}

	conn.Edges = make([]*model.ShopEdge, limit)
	conn.Nodes = make([]*model.Shop, limit)

	for i, s := range shops[:limit] {
		cursor, _ := paginator.CreateEncodedCursor(s)

		node := &model.Shop{
			ID:        s.GlobalID(),
			ShopName:  s.ShopName.Ptr(),
			CreatedAt: &s.CreatedAt,
			UpdatedAt: &s.UpdatedAt,
		}

		pos := i
		if !paginator.IsAfter() {
			pos = len(conn.Edges) - i - 1
		}

		conn.Edges[pos] = &model.ShopEdge{Cursor: cursor, Node: node}
		conn.Nodes[pos] = node
	}

	conn.PageInfo = &model.PageInfo{
		StartCursor: &conn.Edges[0].Cursor,
		EndCursor:   &conn.Edges[len(conn.Edges)-1].Cursor,
	}

	if len(shops) > limit {
		if paginator.IsAfter() {
			conn.PageInfo.HasNextPage = true
		} else {
			conn.PageInfo.HasPreviousPage = true
		}
	}

	totalCount, err := r.repo.ShopsCountByName(ctx, query)

	if err != nil {
		return conn, err
	}

	conn.TotalCount = int(totalCount)

	return conn, nil
}

func (r *Resolver) BookByID(ctx context.Context, id string) (*model.Book, error) {
	bookId, err := globalid.FromGlobalIDInt64(id, "Book")

	if err != nil {
		return nil, err
	}

	record, err := loader.LoadBook(ctx, bookId)

	if err != nil {
		return nil, err
	}

	return &model.Book{
		ID:        record.GlobalID(),
		BookTitle: record.BookTitle.Ptr(),
		CreatedAt: &record.CreatedAt,
		UpdatedAt: &record.UpdatedAt,
	}, nil
}

func (r *Resolver) books(ctx context.Context, after *string, before *string, first *int, last *int, query string, orderBy []*model.BookOrder) (*model.BookConnection, error) {
	paginator := pagination.NewPaginator(
		after,
		before,
		first,
		last,
		model.BookOrderToPaginationOrders(orderBy),
	)

	books, err := r.repo.BoolsByTitle(ctx, query, paginator)

	if err != nil {
		return nil, err
	}

	conn := &model.BookConnection{}

	if len(books) == 0 {
		return conn, nil
	}

	limit := len(books)
	if limit > paginator.Limit() {
		limit = paginator.Limit()
	}

	conn.Edges = make([]*model.BookEdge, limit)
	conn.Nodes = make([]*model.Book, limit)

	for i, s := range books[:limit] {
		cursor, _ := paginator.CreateEncodedCursor(s)

		node := &model.Book{
			ID:        s.GlobalID(),
			BookTitle: s.BookTitle.Ptr(),
			CreatedAt: &s.CreatedAt,
			UpdatedAt: &s.UpdatedAt,
		}

		pos := i
		if !paginator.IsAfter() {
			pos = len(conn.Edges) - i - 1
		}

		conn.Edges[pos] = &model.BookEdge{Cursor: cursor, Node: node}
		conn.Nodes[pos] = node
	}

	conn.PageInfo = &model.PageInfo{
		StartCursor: &conn.Edges[0].Cursor,
		EndCursor:   &conn.Edges[len(conn.Edges)-1].Cursor,
	}

	if len(books) > limit {
		if paginator.IsAfter() {
			conn.PageInfo.HasNextPage = true
		} else {
			conn.PageInfo.HasPreviousPage = true
		}
	}

	totalCount, err := r.repo.BooksCountByTitle(ctx, query)

	if err != nil {
		return conn, err
	}

	conn.TotalCount = int(totalCount)

	return conn, nil
}

func (r *Resolver) booksByShopID(ctx context.Context, id string) ([]*model.Book, error) {
	shopId, err := globalid.FromGlobalIDInt64(id, "Shop")

	if err != nil {
		return nil, err
	}

	records, err := loader.LoadBooksByShopID(ctx, shopId)

	if err != nil {
		return nil, err
	}

	resp := make([]*model.Book, len(records))

	for i, record := range records {
		resp[i] = &model.Book{
			ID:        record.GlobalID(),
			BookTitle: &record.BookTitle.String,
		}
	}

	return resp, nil
}
