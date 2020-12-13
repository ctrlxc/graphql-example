// Code generated by SQLBoiler 4.3.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"app/util"
	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Stock is an object representing the database table.
type Stock struct {
	ID        int64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ShopID    null.Int64 `boil:"shop_id" json:"shop_id,omitempty" toml:"shop_id" yaml:"shop_id,omitempty"`
	BookID    null.Int64 `boil:"book_id" json:"book_id,omitempty" toml:"book_id" yaml:"book_id,omitempty"`
	CreatedAt time.Time  `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time  `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *stockR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stockL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StockColumns = struct {
	ID        string
	ShopID    string
	BookID    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	ShopID:    "shop_id",
	BookID:    "book_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var StockWhere = struct {
	ID        whereHelperint64
	ShopID    whereHelpernull_Int64
	BookID    whereHelpernull_Int64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"graphql\".\"stock\".\"id\""},
	ShopID:    whereHelpernull_Int64{field: "\"graphql\".\"stock\".\"shop_id\""},
	BookID:    whereHelpernull_Int64{field: "\"graphql\".\"stock\".\"book_id\""},
	CreatedAt: whereHelpertime_Time{field: "\"graphql\".\"stock\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"graphql\".\"stock\".\"updated_at\""},
}

// StockRels is where relationship names are stored.
var StockRels = struct {
	Book string
	Shop string
}{
	Book: "Book",
	Shop: "Shop",
}

// stockR is where relationships are stored.
type stockR struct {
	Book *Book `boil:"Book" json:"Book" toml:"Book" yaml:"Book"`
	Shop *Shop `boil:"Shop" json:"Shop" toml:"Shop" yaml:"Shop"`
}

// NewStruct creates a new relationship struct
func (*stockR) NewStruct() *stockR {
	return &stockR{}
}

// stockL is where Load methods for each relationship are stored.
type stockL struct{}

var (
	stockAllColumns            = []string{"id", "shop_id", "book_id", "created_at", "updated_at"}
	stockColumnsWithoutDefault = []string{"shop_id", "book_id", "created_at", "updated_at"}
	stockColumnsWithDefault    = []string{"id"}
	stockPrimaryKeyColumns     = []string{"id"}
)

type (
	// StockSlice is an alias for a slice of pointers to Stock.
	// This should generally be used opposed to []Stock.
	StockSlice []*Stock
	// StockHook is the signature for custom Stock hook methods
	StockHook func(context.Context, boil.ContextExecutor, *Stock) error

	stockQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stockType                 = reflect.TypeOf(&Stock{})
	stockMapping              = queries.MakeStructMapping(stockType)
	stockPrimaryKeyMapping, _ = queries.BindMapping(stockType, stockMapping, stockPrimaryKeyColumns)
	stockInsertCacheMut       sync.RWMutex
	stockInsertCache          = make(map[string]insertCache)
	stockUpdateCacheMut       sync.RWMutex
	stockUpdateCache          = make(map[string]updateCache)
	stockUpsertCacheMut       sync.RWMutex
	stockUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var stockBeforeInsertHooks []StockHook
var stockBeforeUpdateHooks []StockHook
var stockBeforeDeleteHooks []StockHook
var stockBeforeUpsertHooks []StockHook

var stockAfterInsertHooks []StockHook
var stockAfterSelectHooks []StockHook
var stockAfterUpdateHooks []StockHook
var stockAfterDeleteHooks []StockHook
var stockAfterUpsertHooks []StockHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Stock) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Stock) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Stock) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Stock) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Stock) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Stock) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Stock) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Stock) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Stock) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range stockAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddStockHook registers your hook function for all future operations.
func AddStockHook(hookPoint boil.HookPoint, stockHook StockHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		stockBeforeInsertHooks = append(stockBeforeInsertHooks, stockHook)
	case boil.BeforeUpdateHook:
		stockBeforeUpdateHooks = append(stockBeforeUpdateHooks, stockHook)
	case boil.BeforeDeleteHook:
		stockBeforeDeleteHooks = append(stockBeforeDeleteHooks, stockHook)
	case boil.BeforeUpsertHook:
		stockBeforeUpsertHooks = append(stockBeforeUpsertHooks, stockHook)
	case boil.AfterInsertHook:
		stockAfterInsertHooks = append(stockAfterInsertHooks, stockHook)
	case boil.AfterSelectHook:
		stockAfterSelectHooks = append(stockAfterSelectHooks, stockHook)
	case boil.AfterUpdateHook:
		stockAfterUpdateHooks = append(stockAfterUpdateHooks, stockHook)
	case boil.AfterDeleteHook:
		stockAfterDeleteHooks = append(stockAfterDeleteHooks, stockHook)
	case boil.AfterUpsertHook:
		stockAfterUpsertHooks = append(stockAfterUpsertHooks, stockHook)
	}
}

// One returns a single stock record from the query.
func (q stockQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Stock, error) {
	o := &Stock{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stock")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Stock records from the query.
func (q stockQuery) All(ctx context.Context, exec boil.ContextExecutor) (StockSlice, error) {
	var o []*Stock

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Stock slice")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Stock records in the query.
func (q stockQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stock rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q stockQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stock exists")
	}

	return count > 0, nil
}

// Book pointed to by the foreign key.
func (o *Stock) Book(mods ...qm.QueryMod) bookQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.BookID),
	}

	queryMods = append(queryMods, mods...)

	query := Books(queryMods...)
	queries.SetFrom(query.Query, "\"graphql\".\"book\"")

	return query
}

// Shop pointed to by the foreign key.
func (o *Stock) Shop(mods ...qm.QueryMod) shopQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.ShopID),
	}

	queryMods = append(queryMods, mods...)

	query := Shops(queryMods...)
	queries.SetFrom(query.Query, "\"graphql\".\"shop\"")

	return query
}

// LoadBook allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stockL) LoadBook(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStock interface{}, mods queries.Applicator) error {
	var slice []*Stock
	var object *Stock

	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*[]*Stock)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stockR{}
		}
		if !queries.IsNil(object.BookID) {
			args = append(args, object.BookID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stockR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.BookID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.BookID) {
				args = append(args, obj.BookID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`graphql.book`),
		qm.WhereIn(`graphql.book.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Book")
	}

	var resultSlice []*Book
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Book")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for book")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for book")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Book = foreign
		if foreign.R == nil {
			foreign.R = &bookR{}
		}
		foreign.R.Stocks = append(foreign.R.Stocks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.BookID, foreign.ID) {
				local.R.Book = foreign
				if foreign.R == nil {
					foreign.R = &bookR{}
				}
				foreign.R.Stocks = append(foreign.R.Stocks, local)
				break
			}
		}
	}

	return nil
}

// LoadShop allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (stockL) LoadShop(ctx context.Context, e boil.ContextExecutor, singular bool, maybeStock interface{}, mods queries.Applicator) error {
	var slice []*Stock
	var object *Stock

	if singular {
		object = maybeStock.(*Stock)
	} else {
		slice = *maybeStock.(*[]*Stock)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &stockR{}
		}
		if !queries.IsNil(object.ShopID) {
			args = append(args, object.ShopID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &stockR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ShopID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ShopID) {
				args = append(args, obj.ShopID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`graphql.shop`),
		qm.WhereIn(`graphql.shop.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Shop")
	}

	var resultSlice []*Shop
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Shop")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for shop")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for shop")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Shop = foreign
		if foreign.R == nil {
			foreign.R = &shopR{}
		}
		foreign.R.Stocks = append(foreign.R.Stocks, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ShopID, foreign.ID) {
				local.R.Shop = foreign
				if foreign.R == nil {
					foreign.R = &shopR{}
				}
				foreign.R.Stocks = append(foreign.R.Stocks, local)
				break
			}
		}
	}

	return nil
}

// SetBook of the stock to the related item.
// Sets o.R.Book to related.
// Adds o to related.R.Stocks.
func (o *Stock) SetBook(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Book) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"graphql\".\"stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"book_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.BookID, related.ID)
	if o.R == nil {
		o.R = &stockR{
			Book: related,
		}
	} else {
		o.R.Book = related
	}

	if related.R == nil {
		related.R = &bookR{
			Stocks: StockSlice{o},
		}
	} else {
		related.R.Stocks = append(related.R.Stocks, o)
	}

	return nil
}

// RemoveBook relationship.
// Sets o.R.Book to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Stock) RemoveBook(ctx context.Context, exec boil.ContextExecutor, related *Book) error {
	var err error

	queries.SetScanner(&o.BookID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("book_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Book = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Stocks {
		if queries.Equal(o.BookID, ri.BookID) {
			continue
		}

		ln := len(related.R.Stocks)
		if ln > 1 && i < ln-1 {
			related.R.Stocks[i] = related.R.Stocks[ln-1]
		}
		related.R.Stocks = related.R.Stocks[:ln-1]
		break
	}
	return nil
}

// SetShop of the stock to the related item.
// Sets o.R.Shop to related.
// Adds o to related.R.Stocks.
func (o *Stock) SetShop(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Shop) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"graphql\".\"stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"shop_id"}),
		strmangle.WhereClause("\"", "\"", 2, stockPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ShopID, related.ID)
	if o.R == nil {
		o.R = &stockR{
			Shop: related,
		}
	} else {
		o.R.Shop = related
	}

	if related.R == nil {
		related.R = &shopR{
			Stocks: StockSlice{o},
		}
	} else {
		related.R.Stocks = append(related.R.Stocks, o)
	}

	return nil
}

// RemoveShop relationship.
// Sets o.R.Shop to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Stock) RemoveShop(ctx context.Context, exec boil.ContextExecutor, related *Shop) error {
	var err error

	queries.SetScanner(&o.ShopID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("shop_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Shop = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Stocks {
		if queries.Equal(o.ShopID, ri.ShopID) {
			continue
		}

		ln := len(related.R.Stocks)
		if ln > 1 && i < ln-1 {
			related.R.Stocks[i] = related.R.Stocks[ln-1]
		}
		related.R.Stocks = related.R.Stocks[:ln-1]
		break
	}
	return nil
}

// Stocks retrieves all the records using an executor.
func Stocks(mods ...qm.QueryMod) stockQuery {
	mods = append(mods, qm.From("\"graphql\".\"stock\""))
	return stockQuery{NewQuery(mods...)}
}

// FindStock retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStock(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Stock, error) {
	stockObj := &Stock{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"graphql\".\"stock\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, stockObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stock")
	}

	return stockObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Stock) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stock provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stockInsertCacheMut.RLock()
	cache, cached := stockInsertCache[key]
	stockInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stockAllColumns,
			stockColumnsWithDefault,
			stockColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stockType, stockMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"graphql\".\"stock\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"graphql\".\"stock\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into stock")
	}

	if !cached {
		stockInsertCacheMut.Lock()
		stockInsertCache[key] = cache
		stockInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Stock.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Stock) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	stockUpdateCacheMut.RLock()
	cache, cached := stockUpdateCache[key]
	stockUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stockAllColumns,
			stockPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update stock, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"graphql\".\"stock\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stockPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, append(wl, stockPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update stock row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for stock")
	}

	if !cached {
		stockUpdateCacheMut.Lock()
		stockUpdateCache[key] = cache
		stockUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q stockQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for stock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for stock")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StockSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"graphql\".\"stock\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, stockPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in stock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all stock")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Stock) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stock provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(stockColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	stockUpsertCacheMut.RLock()
	cache, cached := stockUpsertCache[key]
	stockUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stockAllColumns,
			stockColumnsWithDefault,
			stockColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			stockAllColumns,
			stockPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert stock, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stockPrimaryKeyColumns))
			copy(conflict, stockPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"graphql\".\"stock\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(stockType, stockMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stockType, stockMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert stock")
	}

	if !cached {
		stockUpsertCacheMut.Lock()
		stockUpsertCache[key] = cache
		stockUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Stock record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Stock) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Stock provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stockPrimaryKeyMapping)
	sql := "DELETE FROM \"graphql\".\"stock\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from stock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for stock")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q stockQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stockQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stock")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stock")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StockSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(stockBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"graphql\".\"stock\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stockPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stock slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stock")
	}

	if len(stockAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Stock) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStock(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StockSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StockSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stockPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"graphql\".\"stock\".* FROM \"graphql\".\"stock\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stockPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StockSlice")
	}

	*o = slice

	return nil
}

// StockExists checks if the Stock row exists.
func StockExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"graphql\".\"stock\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stock exists")
	}

	return exists, nil
}

// WARNING: required ID column
func (o *Stock) GlobalID() string {
	return util.ToGlobalID("Stock", o.ID)
}
