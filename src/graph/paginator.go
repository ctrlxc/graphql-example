package graph

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	defaultLimit = 10
	defaultOrder = asc
)

type Paginator struct {
	cursor Cursor
	orders []Order
	limit  int
}

type Cursor struct {
	After  *string
	Before *string
}

func (c *Cursor) Cursor() *string {
	if c.After != nil {
		return c.After
	}

	return c.Before
}

type CursorItem struct {
	Field     string    `json:"field"`
	Direction Direction `json:"direction"`
	Value     string    `json:"value"`
	Type      reflect.Type
}

type Direction string

const (
	asc  Direction = "ASC"
	desc Direction = "DESC"
)

type Order struct {
	Field     string
	Direction Direction
}

func NewPaginator() *Paginator {
	return &Paginator{}
}

func (p *Paginator) Paginate(model interface{}, after *string, before *string, limit int, orders ...Order) {
	p.cursor.After = after
	p.cursor.Before = before
	p.limit = limit
	p.orders = orders

	if p.limit == 0 {
		p.limit = defaultLimit
	}

	if len(p.orders) == 0 || p.orders[len(p.orders)-1].Field != "id" {
		p.orders = append(p.orders, Order{"id", asc})
	}

	where := p.where()
	orders := p.orderBy()
	limit := p.limit

}

func (p *Paginator) where() string {
	// cursorItems = decoder.Decode(*p.cursor.Cursor()) @@@
	cursorItems := make([]CursorItem, 0)

	qs := make([]string, len(cursorItems))
	composite := ""

	for i, c := range cursorItems {
		op := p.operator(c.Direction)

		qs[i] = fmt.Sprintf("%s%s %s ?", composite, c.Field, op)
		composite = fmt.Sprintf("%s%s = ? AND ", composite, c.Field)
	}

	return strings.Join(qs, " OR ")
}

func (p *Paginator) operator(d Direction) string {
	if (p.hasAfterCursor() && d == asc) ||
		(p.hasBeforeCursor() && d == desc) {
		return ">"
	}

	return "<"
}

func (p *Paginator) orderBy() string {
	orders := make([]string, len(p.orders))

	for i, o := range p.orders {
		dir := p.direction(o.Direction)
		orders[i] = fmt.Sprintf("%s %s", c.Field, dir)
	}

	return strings.Join(orders, ", ")
}

func (p *Paginator) direction(d Direction) Direction {
	if p.hasAfterCursor() {
		return d
	}

	if d == asc {
		return desc
	}

	return asc
}

// func (p *Paginator) appendPagingQuery(model interface{}) *gorm.DB {
// 	decoder, _ := NewCursorDecoder(out, p.keys...)

// 	var fields []interface{}
// 	if p.hasAfterCursor() {
// 		fields = decoder.Decode(*p.cursor.After)
// 	} else if p.hasBeforeCursor() {
// 		fields = decoder.Decode(*p.cursor.Before)
// 	}
// 	if len(fields) > 0 {
// 		stmt = stmt.Where(
// 			p.getCursorQuery(),
// 			p.getCursorQueryArgs(fields)...,
// 		)
// 	}
// 	stmt = stmt.Limit(p.limit + 1)
// 	stmt = stmt.Order(p.getOrder())
// 	return stmt
// }

func (p *Paginator) hasAfterCursor() bool {
	return p.cursor.After != nil
}

func (p *Paginator) hasBeforeCursor() bool {
	return !p.hasAfterCursor() && p.cursor.Before != nil
}

// func (p *Paginator) getCursorQuery() string {
// 	qs := make([]string, len(p.tableKeys))
// 	op := p.getOperator()
// 	composite := ""
// 	for i, sqlKey := range p.tableKeys {
// 		qs[i] = fmt.Sprintf("%s%s %s ?", composite, sqlKey, op)
// 		composite = fmt.Sprintf("%s%s = ? AND ", composite, sqlKey)
// 	}
// 	return strings.Join(qs, " OR ")
// }

// func (p *Paginator) getCursorQueryArgs(fields []interface{}) (args []interface{}) {
// 	for i := 1; i <= len(fields); i++ {
// 		args = append(args, fields[:i]...)
// 	}
// 	return
// }

// func (p *Paginator) getOperator() string {
// 	if (p.hasAfterCursor() && p.order == ASC) ||
// 		(p.hasBeforeCursor() && p.order == DESC) {
// 		return ">"
// 	}
// 	return "<"
// }

// func (p *Paginator) postProcess(out interface{}) {
// 	elems := reflect.ValueOf(out).Elem()
// 	hasMore := elems.Len() > p.limit
// 	if hasMore {
// 		elems.Set(elems.Slice(0, elems.Len()-1))
// 	}
// 	if p.hasBeforeCursor() {
// 		elems.Set(reverse(elems))
// 	}
// 	encoder := NewCursorEncoder(p.keys...)
// 	if p.hasBeforeCursor() || hasMore {
// 		cursor := encoder.Encode(elems.Index(elems.Len() - 1))
// 		p.next.After = &cursor
// 	}
// 	if p.hasAfterCursor() || (hasMore && p.hasBeforeCursor()) {
// 		cursor := encoder.Encode(elems.Index(0))
// 		p.next.Before = &cursor
// 	}
// 	return
// }

// func reverse(v reflect.Value) reflect.Value {
// 	result := reflect.MakeSlice(v.Type(), 0, v.Cap())
// 	for i := v.Len() - 1; i >= 0; i-- {
// 		result = reflect.Append(result, v.Index(i))
// 	}
// 	return result
// }
