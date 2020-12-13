package paginator

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	defaultLimit = 10
)

type Paginator struct {
	After  *string
	Before *string
	First  *int
	Last   *int
	Orders []*PaginatorOrder
}

type PaginatorDirection string

const (
	Asc  PaginatorDirection = "ASC"
	Desc PaginatorDirection = "DESC"
)

type PaginatorOrder struct {
	Field     string
	Direction PaginatorDirection
}

func (p *Paginator) Queries() []qm.QueryMod {
	return []qm.QueryMod{
		p.QueryWhere(),
		p.QueryOrderBy(),
		p.QueryLimit(),
	}
}

// @@@ not support nullable column
func (p *Paginator) QueryWhere() qm.QueryMod {
	cursorstr := p.validCursor()

	if cursorstr == nil {
		return qm.And("1 = 1")
	}

	cursor, err := p.CursorDecode(*cursorstr)

	if err != nil {
		return nil // @@@ err?
	}

	queries := make([]string, len(cursor.Items))
	binds := make([]interface{}, 0)

	composite := ""

	for i, c := range cursor.Items {
		op := p.operator(c.Direction)

		queries[i] = fmt.Sprintf("%s%s %s ?", composite, c.Field, op)
		for _, v := range cursor.Items[:i+1] {
			binds = append(binds, v.Value)
		}

		composite = fmt.Sprintf("%s%s = ? AND ", composite, c.Field)
	}

	query := strings.Join(queries, " OR ")

	return qm.Where(query, binds...)
}

func (p *Paginator) QueryOrderBy() qm.QueryMod {
	if len(p.Orders) == 0 || p.Orders[len(p.Orders)-1].Field != "id" {
		p.Orders = append(p.Orders, &PaginatorOrder{"id", Asc})
	}

	orders := make([]string, len(p.Orders))

	for i, o := range p.Orders {
		dir := p.direction(o.Direction)
		orders[i] = fmt.Sprintf("%s %s", o.Field, dir)
	}

	orderBy := strings.Join(orders, ", ")

	return qm.OrderBy(orderBy)
}

func (p *Paginator) QueryLimit() qm.QueryMod {
	return qm.Limit(p.Limit() + 1)
}

func (p *Paginator) Limit() int {
	limit := defaultLimit

	if p.IsAfter() && p.First != nil && *p.First > 0 {
		limit = *p.First
	} else if !p.IsAfter() && p.Last != nil && *p.Last > 0 {
		limit = *p.Last
	}

	return limit
}

func (p *Paginator) operator(d PaginatorDirection) string {
	if (p.IsAfter() && d == Asc) ||
		(!p.IsAfter() && d == Desc) {
		return ">"
	}

	return "<"
}

func (p *Paginator) direction(d PaginatorDirection) PaginatorDirection {
	if p.IsAfter() {
		return d
	}

	if d == Asc {
		return Desc
	}

	return Asc
}

func (p *Paginator) IsAfter() bool {
	if p.hasAfterCursor() {
		return true
	}

	if !p.hasBeforeCursor() {
		if p.First != nil && *p.First > 0 {
			return true
		}

		if p.Last == nil || *p.Last <= 0 {
			return true
		}
	}

	return false
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
	return p.After != nil
}

func (p *Paginator) hasBeforeCursor() bool {
	return !p.hasAfterCursor() && p.Before != nil
}

func (p *Paginator) validCursor() *string {
	if p.hasAfterCursor() {
		return p.After
	}

	if p.hasBeforeCursor() {
		return p.Before
	}

	return nil
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

type Cursor struct {
	Items []*CursorItem `json:"items"`
}

type CursorItem struct {
	Field     string             `json:"field"`
	Direction PaginatorDirection `json:"direction"`
	Value     interface{}        `json:"value"`
	ValueType string             `json:"value_type"`
}

func (p *Paginator) CursorEncode(cursor *Cursor) (string, error) {
	b, err := json.Marshal(cursor)

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func (p *Paginator) CursorDecode(cursorstr string) (*Cursor, error) {
	b, err := base64.StdEncoding.DecodeString(cursorstr)

	if err != nil {
		return nil, err
	}

	cursor := Cursor{}

	err = json.Unmarshal(b, &cursor)

	if err != nil {
		return nil, err
	}

	// for _, item := range cursor.Items {
	// 	if item.ValueType == "Time" {
	// 		if v, ok := (item.Value.(string)); ok {
	// 			item.Value, _ = time.Parse("", v)
	// 		}
	// 	}
	// }

	return &cursor, nil
}

func (p *Paginator) CreateCursor(v interface{}) *Cursor {
	rv := reflectValue(v)

	for rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}

	items := make([]*CursorItem, len(p.Orders))

	for i, o := range p.Orders {
		camelField := strmangle.TitleCase(o.Field)

		f := rv.FieldByName(camelField)

		items[i] = &CursorItem{
			Field:     o.Field,
			Direction: o.Direction,
			Value:     f.Interface(),
			ValueType: f.Type().Name(),
		}
	}

	return &Cursor{Items: items}
}

func (p *Paginator) CreateEncodedCursor(v interface{}) (string, error) {
	c := p.CreateCursor(v)
	return p.CursorEncode(c)
}

func reflectValue(v interface{}) reflect.Value {
	rv, ok := v.(reflect.Value)

	if !ok {
		return reflect.ValueOf(v)
	}

	return rv
}
