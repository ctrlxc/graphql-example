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

// Shop is an object representing the database table.
type Shop struct {
	ID       int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	ShopName null.String `boil:"shop_name" json:"shop_name,omitempty" toml:"shop_name" yaml:"shop_name,omitempty"`

	R *shopR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L shopL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ShopColumns = struct {
	ID       string
	ShopName string
}{
	ID:       "id",
	ShopName: "shop_name",
}

// Generated where

var ShopWhere = struct {
	ID       whereHelperint64
	ShopName whereHelpernull_String
}{
	ID:       whereHelperint64{field: "\"graphql\".\"shop\".\"id\""},
	ShopName: whereHelpernull_String{field: "\"graphql\".\"shop\".\"shop_name\""},
}

// ShopRels is where relationship names are stored.
var ShopRels = struct {
	Stocks string
}{
	Stocks: "Stocks",
}

// shopR is where relationships are stored.
type shopR struct {
	Stocks StockSlice `boil:"Stocks" json:"Stocks" toml:"Stocks" yaml:"Stocks"`
}

// NewStruct creates a new relationship struct
func (*shopR) NewStruct() *shopR {
	return &shopR{}
}

// shopL is where Load methods for each relationship are stored.
type shopL struct{}

var (
	shopAllColumns            = []string{"id", "shop_name"}
	shopColumnsWithoutDefault = []string{"shop_name"}
	shopColumnsWithDefault    = []string{"id"}
	shopPrimaryKeyColumns     = []string{"id"}
)

type (
	// ShopSlice is an alias for a slice of pointers to Shop.
	// This should generally be used opposed to []Shop.
	ShopSlice []*Shop
	// ShopHook is the signature for custom Shop hook methods
	ShopHook func(context.Context, boil.ContextExecutor, *Shop) error

	shopQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	shopType                 = reflect.TypeOf(&Shop{})
	shopMapping              = queries.MakeStructMapping(shopType)
	shopPrimaryKeyMapping, _ = queries.BindMapping(shopType, shopMapping, shopPrimaryKeyColumns)
	shopInsertCacheMut       sync.RWMutex
	shopInsertCache          = make(map[string]insertCache)
	shopUpdateCacheMut       sync.RWMutex
	shopUpdateCache          = make(map[string]updateCache)
	shopUpsertCacheMut       sync.RWMutex
	shopUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var shopBeforeInsertHooks []ShopHook
var shopBeforeUpdateHooks []ShopHook
var shopBeforeDeleteHooks []ShopHook
var shopBeforeUpsertHooks []ShopHook

var shopAfterInsertHooks []ShopHook
var shopAfterSelectHooks []ShopHook
var shopAfterUpdateHooks []ShopHook
var shopAfterDeleteHooks []ShopHook
var shopAfterUpsertHooks []ShopHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Shop) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Shop) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Shop) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Shop) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Shop) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Shop) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Shop) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Shop) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Shop) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range shopAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddShopHook registers your hook function for all future operations.
func AddShopHook(hookPoint boil.HookPoint, shopHook ShopHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		shopBeforeInsertHooks = append(shopBeforeInsertHooks, shopHook)
	case boil.BeforeUpdateHook:
		shopBeforeUpdateHooks = append(shopBeforeUpdateHooks, shopHook)
	case boil.BeforeDeleteHook:
		shopBeforeDeleteHooks = append(shopBeforeDeleteHooks, shopHook)
	case boil.BeforeUpsertHook:
		shopBeforeUpsertHooks = append(shopBeforeUpsertHooks, shopHook)
	case boil.AfterInsertHook:
		shopAfterInsertHooks = append(shopAfterInsertHooks, shopHook)
	case boil.AfterSelectHook:
		shopAfterSelectHooks = append(shopAfterSelectHooks, shopHook)
	case boil.AfterUpdateHook:
		shopAfterUpdateHooks = append(shopAfterUpdateHooks, shopHook)
	case boil.AfterDeleteHook:
		shopAfterDeleteHooks = append(shopAfterDeleteHooks, shopHook)
	case boil.AfterUpsertHook:
		shopAfterUpsertHooks = append(shopAfterUpsertHooks, shopHook)
	}
}

// One returns a single shop record from the query.
func (q shopQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Shop, error) {
	o := &Shop{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for shop")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Shop records from the query.
func (q shopQuery) All(ctx context.Context, exec boil.ContextExecutor) (ShopSlice, error) {
	var o []*Shop

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Shop slice")
	}

	if len(shopAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Shop records in the query.
func (q shopQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count shop rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q shopQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if shop exists")
	}

	return count > 0, nil
}

// Stocks retrieves all the stock's Stocks with an executor.
func (o *Shop) Stocks(mods ...qm.QueryMod) stockQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"graphql\".\"stock\".\"shop_id\"=?", o.ID),
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
func (shopL) LoadStocks(ctx context.Context, e boil.ContextExecutor, singular bool, maybeShop interface{}, mods queries.Applicator) error {
	var slice []*Shop
	var object *Shop

	if singular {
		object = maybeShop.(*Shop)
	} else {
		slice = *maybeShop.(*[]*Shop)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &shopR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &shopR{}
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
		qm.WhereIn(`graphql.stock.shop_id in ?`, args...),
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
			foreign.R.Shop = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.ShopID) {
				local.R.Stocks = append(local.R.Stocks, foreign)
				if foreign.R == nil {
					foreign.R = &stockR{}
				}
				foreign.R.Shop = local
				break
			}
		}
	}

	return nil
}

// AddStocks adds the given related objects to the existing relationships
// of the shop, optionally inserting them as new records.
// Appends related to o.R.Stocks.
// Sets related.R.Shop appropriately.
func (o *Shop) AddStocks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Stock) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.ShopID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"graphql\".\"stock\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"shop_id"}),
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

			queries.Assign(&rel.ShopID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &shopR{
			Stocks: related,
		}
	} else {
		o.R.Stocks = append(o.R.Stocks, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &stockR{
				Shop: o,
			}
		} else {
			rel.R.Shop = o
		}
	}
	return nil
}

// SetStocks removes all previously related items of the
// shop replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Shop's Stocks accordingly.
// Replaces o.R.Stocks with related.
// Sets related.R.Shop's Stocks accordingly.
func (o *Shop) SetStocks(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Stock) error {
	query := "update \"graphql\".\"stock\" set \"shop_id\" = null where \"shop_id\" = $1"
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
			queries.SetScanner(&rel.ShopID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Shop = nil
		}

		o.R.Stocks = nil
	}
	return o.AddStocks(ctx, exec, insert, related...)
}

// RemoveStocks relationships from objects passed in.
// Removes related items from R.Stocks (uses pointer comparison, removal does not keep order)
// Sets related.R.Shop.
func (o *Shop) RemoveStocks(ctx context.Context, exec boil.ContextExecutor, related ...*Stock) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.ShopID, nil)
		if rel.R != nil {
			rel.R.Shop = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("shop_id")); err != nil {
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

// Shops retrieves all the records using an executor.
func Shops(mods ...qm.QueryMod) shopQuery {
	mods = append(mods, qm.From("\"graphql\".\"shop\""))
	return shopQuery{NewQuery(mods...)}
}

// FindShop retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindShop(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Shop, error) {
	shopObj := &Shop{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"graphql\".\"shop\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, shopObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from shop")
	}

	return shopObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Shop) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no shop provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shopColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	shopInsertCacheMut.RLock()
	cache, cached := shopInsertCache[key]
	shopInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			shopAllColumns,
			shopColumnsWithDefault,
			shopColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(shopType, shopMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(shopType, shopMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"graphql\".\"shop\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"graphql\".\"shop\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into shop")
	}

	if !cached {
		shopInsertCacheMut.Lock()
		shopInsertCache[key] = cache
		shopInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Shop.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Shop) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	shopUpdateCacheMut.RLock()
	cache, cached := shopUpdateCache[key]
	shopUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			shopAllColumns,
			shopPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update shop, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"graphql\".\"shop\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, shopPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(shopType, shopMapping, append(wl, shopPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update shop row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for shop")
	}

	if !cached {
		shopUpdateCacheMut.Lock()
		shopUpdateCache[key] = cache
		shopUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q shopQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for shop")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for shop")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ShopSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shopPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"graphql\".\"shop\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, shopPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in shop slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all shop")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Shop) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no shop provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(shopColumnsWithDefault, o)

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

	shopUpsertCacheMut.RLock()
	cache, cached := shopUpsertCache[key]
	shopUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			shopAllColumns,
			shopColumnsWithDefault,
			shopColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			shopAllColumns,
			shopPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert shop, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(shopPrimaryKeyColumns))
			copy(conflict, shopPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"graphql\".\"shop\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(shopType, shopMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(shopType, shopMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert shop")
	}

	if !cached {
		shopUpsertCacheMut.Lock()
		shopUpsertCache[key] = cache
		shopUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Shop record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Shop) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Shop provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), shopPrimaryKeyMapping)
	sql := "DELETE FROM \"graphql\".\"shop\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from shop")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for shop")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q shopQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no shopQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shop")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for shop")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ShopSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(shopBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shopPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"graphql\".\"shop\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, shopPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from shop slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for shop")
	}

	if len(shopAfterDeleteHooks) != 0 {
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
func (o *Shop) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindShop(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ShopSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ShopSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), shopPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"graphql\".\"shop\".* FROM \"graphql\".\"shop\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, shopPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ShopSlice")
	}

	*o = slice

	return nil
}

// ShopExists checks if the Shop row exists.
func ShopExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"graphql\".\"shop\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if shop exists")
	}

	return exists, nil
}
