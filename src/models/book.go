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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Book is an object representing the database table.
type Book struct {
	ID        int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	BookTitle null.String `boil:"book_title" json:"book_title,omitempty" toml:"book_title" yaml:"book_title,omitempty"`

	R *bookR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bookL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BookColumns = struct {
	ID        string
	BookTitle string
}{
	ID:        "id",
	BookTitle: "book_title",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var BookWhere = struct {
	ID        whereHelperint64
	BookTitle whereHelpernull_String
}{
	ID:        whereHelperint64{field: "\"graphql\".\"book\".\"id\""},
	BookTitle: whereHelpernull_String{field: "\"graphql\".\"book\".\"book_title\""},
}

// BookRels is where relationship names are stored.
var BookRels = struct {
	Stocks string
}{
	Stocks: "Stocks",
}

// bookR is where relationships are stored.
type bookR struct {
	Stocks StockSlice `boil:"Stocks" json:"Stocks" toml:"Stocks" yaml:"Stocks"`
}

// NewStruct creates a new relationship struct
func (*bookR) NewStruct() *bookR {
	return &bookR{}
}

// bookL is where Load methods for each relationship are stored.
type bookL struct{}

var (
	bookAllColumns            = []string{"id", "book_title"}
	bookColumnsWithoutDefault = []string{"book_title"}
	bookColumnsWithDefault    = []string{"id"}
	bookPrimaryKeyColumns     = []string{"id"}
)

type (
	// BookSlice is an alias for a slice of pointers to Book.
	// This should generally be used opposed to []Book.
	BookSlice []*Book
	// BookHook is the signature for custom Book hook methods
	BookHook func(context.Context, boil.ContextExecutor, *Book) error

	bookQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bookType                 = reflect.TypeOf(&Book{})
	bookMapping              = queries.MakeStructMapping(bookType)
	bookPrimaryKeyMapping, _ = queries.BindMapping(bookType, bookMapping, bookPrimaryKeyColumns)
	bookInsertCacheMut       sync.RWMutex
	bookInsertCache          = make(map[string]insertCache)
	bookUpdateCacheMut       sync.RWMutex
	bookUpdateCache          = make(map[string]updateCache)
	bookUpsertCacheMut       sync.RWMutex
	bookUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bookBeforeInsertHooks []BookHook
var bookBeforeUpdateHooks []BookHook
var bookBeforeDeleteHooks []BookHook
var bookBeforeUpsertHooks []BookHook

var bookAfterInsertHooks []BookHook
var bookAfterSelectHooks []BookHook
var bookAfterUpdateHooks []BookHook
var bookAfterDeleteHooks []BookHook
var bookAfterUpsertHooks []BookHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Book) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Book) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Book) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Book) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Book) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Book) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Book) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Book) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Book) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range bookAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBookHook registers your hook function for all future operations.
func AddBookHook(hookPoint boil.HookPoint, bookHook BookHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		bookBeforeInsertHooks = append(bookBeforeInsertHooks, bookHook)
	case boil.BeforeUpdateHook:
		bookBeforeUpdateHooks = append(bookBeforeUpdateHooks, bookHook)
	case boil.BeforeDeleteHook:
		bookBeforeDeleteHooks = append(bookBeforeDeleteHooks, bookHook)
	case boil.BeforeUpsertHook:
		bookBeforeUpsertHooks = append(bookBeforeUpsertHooks, bookHook)
	case boil.AfterInsertHook:
		bookAfterInsertHooks = append(bookAfterInsertHooks, bookHook)
	case boil.AfterSelectHook:
		bookAfterSelectHooks = append(bookAfterSelectHooks, bookHook)
	case boil.AfterUpdateHook:
		bookAfterUpdateHooks = append(bookAfterUpdateHooks, bookHook)
	case boil.AfterDeleteHook:
		bookAfterDeleteHooks = append(bookAfterDeleteHooks, bookHook)
	case boil.AfterUpsertHook:
		bookAfterUpsertHooks = append(bookAfterUpsertHooks, bookHook)
	}
}

// One returns a single book record from the query.
func (q bookQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Book, error) {
	o := &Book{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for book")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Book records from the query.
func (q bookQuery) All(ctx context.Context, exec boil.ContextExecutor) (BookSlice, error) {
	var o []*Book

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Book slice")
	}

	if len(bookAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Book records in the query.
func (q bookQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count book rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bookQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if book exists")
	}

	return count > 0, nil
}

// Stocks retrieves all the stock's Stocks with an executor.
func (o *Book) Stocks(mods ...qm.QueryMod) stockQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"graphql\".\"stock\".\"book_id\"=?", o.ID),
	)

	query := Stocks(queryMods...)
	queries.SetFrom(query.Query, "\"graphql\".\"stock\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"graphql\".\"stock\".*"})
	}

	return query
}

// LoadStocks allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (bookL) LoadStocks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeBook interface{}, mods queries.Applicator) error {
	var slice []*Book
	var object *Book

	if singular {
		object = maybeBook.(*Book)
	} else {
		slice = *maybeBook.(*[]*Book)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &bookR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &bookR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`graphql.stock`),
		qm.WhereIn(`graphql.stock.book_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load stock")
	}

	var resultSlice []*Stock
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice stock")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on stock")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for stock")
	}

	if len(stockAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Stocks = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &stockR{}
			}
			foreign.R.Book = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.BookID) {
				local.R.Stocks = append(local.R.Stocks, foreign)
				if foreign.R == nil {
					foreign.R = &stockR{}
				}
				foreign.R.Book = local
				break
			}
		}
	}

	return nil
}

// AddStocks adds the given related objects to the existing relationships
// of the book, optionally inserting them as new records.
// Appends related to o.R.Stocks.
// Sets related.R.Book appropriately.
func (o *Book) AddStocks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Stock) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.BookID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"graphql\".\"stock\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"book_id"}),
				strmangle.WhereClause("\"", "\"", 2, stockPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.BookID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &bookR{
			Stocks: related,
		}
	} else {
		o.R.Stocks = append(o.R.Stocks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stockR{
				Book: o,
			}
		} else {
			rel.R.Book = o
		}
	}
	return nil
}

// SetStocks removes all previously related items of the
// book replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Book's Stocks accordingly.
// Replaces o.R.Stocks with related.
// Sets related.R.Book's Stocks accordingly.
func (o *Book) SetStocks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Stock) error {
	query := "update \"graphql\".\"stock\" set \"book_id\" = null where \"book_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Stocks {
			queries.SetScanner(&rel.BookID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Book = nil
		}

		o.R.Stocks = nil
	}
	return o.AddStocks(ctx, exec, insert, related...)
}

// RemoveStocks relationships from objects passed in.
// Removes related items from R.Stocks (uses pointer comparison, removal does not keep order)
// Sets related.R.Book.
func (o *Book) RemoveStocks(ctx context.Context, exec boil.ContextExecutor, related ...*Stock) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.BookID, nil)
		if rel.R != nil {
			rel.R.Book = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("book_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Stocks {
			if rel != ri {
				continue
			}

			ln := len(o.R.Stocks)
			if ln > 1 && i < ln-1 {
				o.R.Stocks[i] = o.R.Stocks[ln-1]
			}
			o.R.Stocks = o.R.Stocks[:ln-1]
			break
		}
	}

	return nil
}

// Books retrieves all the records using an executor.
func Books(mods ...qm.QueryMod) bookQuery {
	mods = append(mods, qm.From("\"graphql\".\"book\""))
	return bookQuery{NewQuery(mods...)}
}

// FindBook retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBook(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Book, error) {
	bookObj := &Book{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"graphql\".\"book\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, bookObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from book")
	}

	return bookObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Book) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no book provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bookInsertCacheMut.RLock()
	cache, cached := bookInsertCache[key]
	bookInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bookAllColumns,
			bookColumnsWithDefault,
			bookColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bookType, bookMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bookType, bookMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"graphql\".\"book\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"graphql\".\"book\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into book")
	}

	if !cached {
		bookInsertCacheMut.Lock()
		bookInsertCache[key] = cache
		bookInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Book.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Book) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bookUpdateCacheMut.RLock()
	cache, cached := bookUpdateCache[key]
	bookUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bookAllColumns,
			bookPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update book, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"graphql\".\"book\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bookPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bookType, bookMapping, append(wl, bookPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update book row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for book")
	}

	if !cached {
		bookUpdateCacheMut.Lock()
		bookUpdateCache[key] = cache
		bookUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q bookQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for book")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for book")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BookSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"graphql\".\"book\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bookPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in book slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all book")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Book) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no book provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bookColumnsWithDefault, o)

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

	bookUpsertCacheMut.RLock()
	cache, cached := bookUpsertCache[key]
	bookUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bookAllColumns,
			bookColumnsWithDefault,
			bookColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bookAllColumns,
			bookPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert book, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bookPrimaryKeyColumns))
			copy(conflict, bookPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"graphql\".\"book\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bookType, bookMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bookType, bookMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert book")
	}

	if !cached {
		bookUpsertCacheMut.Lock()
		bookUpsertCache[key] = cache
		bookUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Book record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Book) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Book provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bookPrimaryKeyMapping)
	sql := "DELETE FROM \"graphql\".\"book\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from book")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for book")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bookQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bookQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from book")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for book")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BookSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bookBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"graphql\".\"book\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from book slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for book")
	}

	if len(bookAfterDeleteHooks) != 0 {
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
func (o *Book) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBook(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BookSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BookSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bookPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"graphql\".\"book\".* FROM \"graphql\".\"book\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bookPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BookSlice")
	}

	*o = slice

	return nil
}

// BookExists checks if the Book row exists.
func BookExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"graphql\".\"book\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if book exists")
	}

	return exists, nil
}
