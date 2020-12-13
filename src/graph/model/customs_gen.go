// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"app/paginator"
	"strings"
)

func BookOrderToInterface(orders []*BookOrder) []interface{} {
	ifs := make([]interface{}, len(orders))
	for i, o := range orders {
		ifs[i] = o
	}

	return ifs
}

func BookOrderToPaginatorOrder(orders []*BookOrder) []*paginator.PaginatorOrder {
	pagenatorOrders := make([]*paginator.PaginatorOrder, len(orders))

	for i, o := range orders {
		pagenatorOrders[i] = &paginator.PaginatorOrder{Field: strings.ToLower(string(*o.Field)), Direction: paginator.PaginatorDirection(*o.Direction)}
	}

	return pagenatorOrders
}
func ShopOrderToInterface(orders []*ShopOrder) []interface{} {
	ifs := make([]interface{}, len(orders))
	for i, o := range orders {
		ifs[i] = o
	}

	return ifs
}

func ShopOrderToPaginatorOrder(orders []*ShopOrder) []*paginator.PaginatorOrder {
	pagenatorOrders := make([]*paginator.PaginatorOrder, len(orders))

	for i, o := range orders {
		pagenatorOrders[i] = &paginator.PaginatorOrder{Field: strings.ToLower(string(*o.Field)), Direction: paginator.PaginatorDirection(*o.Direction)}
	}

	return pagenatorOrders
}
