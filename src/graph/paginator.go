package paginator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jinzhu/gorm"
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
		p.keys = append(p.orders, Order{"id", asc})
	}

	result := p.appendPagingQuery(stmt, out).Find(out)

	// out must be a pointer or gorm will panic above
	elems := reflect.ValueOf(out).Elem()
	if elems.Kind() == reflect.Slice && elems.Len() > 0 {
		p.postProcess(out)
	}
	return result
}

func (p *Paginator) appendPagingQuery(model interface{}) *gorm.DB {
	decoder, _ := NewCursorDecoder(out, p.keys...)

	var fields []interface{}
	if p.hasAfterCursor() {
		fields = decoder.Decode(*p.cursor.After)
	} else if p.hasBeforeCursor() {
		fields = decoder.Decode(*p.cursor.Before)
	}
	if len(fields) > 0 {
		stmt = stmt.Where(
			p.getCursorQuery(),
			p.getCursorQueryArgs(fields)...,
		)
	}
	stmt = stmt.Limit(p.limit + 1)
	stmt = stmt.Order(p.getOrder())
	return stmt
}

func (p *Paginator) hasAfterCursor() bool {
	return p.cursor.After != nil
}

func (p *Paginator) hasBeforeCursor() bool {
	return !p.hasAfterCursor() && p.cursor.Before != nil
}

func (p *Paginator) getCursorQuery() string {
	qs := make([]string, len(p.tableKeys))
	op := p.getOperator()
	composite := ""
	for i, sqlKey := range p.tableKeys {
		qs[i] = fmt.Sprintf("%s%s %s ?", composite, sqlKey, op)
		composite = fmt.Sprintf("%s%s = ? AND ", composite, sqlKey)
	}
	return strings.Join(qs, " OR ")
}

func (p *Paginator) getCursorQueryArgs(fields []interface{}) (args []interface{}) {
	for i := 1; i <= len(fields); i++ {
		args = append(args, fields[:i]...)
	}
	return
}

func (p *Paginator) getOperator() string {
	if (p.hasAfterCursor() && p.order == ASC) ||
		(p.hasBeforeCursor() && p.order == DESC) {
		return ">"
	}
	return "<"
}

func (p *Paginator) getOrder() string {
	order := p.order
	if p.hasBeforeCursor() {
		order = flip(p.order)
	}
	orders := make([]string, len(p.tableKeys))
	for index, sqlKey := range p.tableKeys {
		orders[index] = fmt.Sprintf("%s %s", sqlKey, order)
	}
	return strings.Join(orders, ", ")
}

func (p *Paginator) postProcess(out interface{}) {
	elems := reflect.ValueOf(out).Elem()
	hasMore := elems.Len() > p.limit
	if hasMore {
		elems.Set(elems.Slice(0, elems.Len()-1))
	}
	if p.hasBeforeCursor() {
		elems.Set(reverse(elems))
	}
	encoder := NewCursorEncoder(p.keys...)
	if p.hasBeforeCursor() || hasMore {
		cursor := encoder.Encode(elems.Index(elems.Len() - 1))
		p.next.After = &cursor
	}
	if p.hasAfterCursor() || (hasMore && p.hasBeforeCursor()) {
		cursor := encoder.Encode(elems.Index(0))
		p.next.Before = &cursor
	}
	return
}

func reverse(v reflect.Value) reflect.Value {
	result := reflect.MakeSlice(v.Type(), 0, v.Cap())
	for i := v.Len() - 1; i >= 0; i-- {
		result = reflect.Append(result, v.Index(i))
	}
	return result
}

func flipOrder(order Order) Order {
	if order == asc {
		return desc
	}

	return asc
}
