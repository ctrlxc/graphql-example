package pagination

import (
	"fmt"
	"strings"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const (
	defaultLimit = 10 // FIXME: from settings file
)

type Pagination struct {
	After  *string
	Before *string
	First  *int
	Last   *int
	Orders []*Order
}

type Direction string

const (
	Asc  Direction = "ASC"
	Desc Direction = "DESC"
)

type Order struct {
	Field     string
	Direction Direction
}

func (p *Pagination) Queries(mods ...qm.QueryMod) []qm.QueryMod {
	return append(mods, []qm.QueryMod{
		p.QueryWhere(),
		p.QueryOrderBy(),
		p.QueryLimit(),
	}...)
}

// @@@ not support nullable column
func (p *Pagination) QueryWhere() qm.QueryMod {
	cursorstr := p.validCursor()

	if cursorstr == nil {
		return qm.And("1 = 1")
	}

	cursor, err := p.CursorDecode(*cursorstr)

	if err != nil {
		panic(fmt.Errorf("unknown cursor: %s", *cursorstr))
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

func (p *Pagination) QueryOrderBy() qm.QueryMod {
	if len(p.Orders) == 0 || p.Orders[len(p.Orders)-1].Field != "id" {
		p.Orders = append(p.Orders, &Order{"id", Asc})
	}

	orders := make([]string, len(p.Orders))

	for i, o := range p.Orders {
		dir := p.direction(o.Direction)
		orders[i] = fmt.Sprintf("%s %s", o.Field, dir)
	}

	orderBy := strings.Join(orders, ", ")

	return qm.OrderBy(orderBy)
}

func (p *Pagination) QueryLimit() qm.QueryMod {
	return qm.Limit(p.Limit() + 1)
}

func (p *Pagination) Limit() int {
	limit := defaultLimit

	if p.IsAfter() && p.First != nil && *p.First > 0 {
		limit = *p.First
	} else if !p.IsAfter() && p.Last != nil && *p.Last > 0 {
		limit = *p.Last
	}

	return limit
}

func (p *Pagination) operator(d Direction) string {
	if (p.IsAfter() && d == Asc) ||
		(!p.IsAfter() && d == Desc) {
		return ">"
	}

	return "<"
}

func (p *Pagination) direction(d Direction) Direction {
	if p.IsAfter() {
		return d
	}

	if d == Asc {
		return Desc
	}

	return Asc
}

func (p *Pagination) IsAfter() bool {
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

func (p *Pagination) hasAfterCursor() bool {
	return p.After != nil
}

func (p *Pagination) hasBeforeCursor() bool {
	return !p.hasAfterCursor() && p.Before != nil
}

func (p *Pagination) validCursor() *string {
	if p.hasAfterCursor() {
		return p.After
	}

	if p.hasBeforeCursor() {
		return p.Before
	}

	return nil
}
