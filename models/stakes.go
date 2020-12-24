// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Stake is an object representing the database table.
type Stake struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Height    int64     `boil:"height" json:"height" toml:"height" yaml:"height"`
	VendorID  int       `boil:"vendor_id" json:"vendorID" toml:"vendorID" yaml:"vendorID"`
	PostID    string    `boil:"post_id" json:"postID" toml:"postID" yaml:"postID"`
	Delegator string    `boil:"delegator" json:"delegator" toml:"delegator" yaml:"delegator"`
	Validator string    `boil:"validator" json:"validator" toml:"validator" yaml:"validator"`
	Amount    int64     `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	CreatedAt time.Time `boil:"created_at" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time `boil:"updated_at" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`
	DeletedAt null.Time `boil:"deleted_at" json:"deletedAt,omitempty" toml:"deletedAt" yaml:"deletedAt,omitempty"`

	R *stakeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L stakeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var StakeColumns = struct {
	ID        string
	Height    string
	VendorID  string
	PostID    string
	Delegator string
	Validator string
	Amount    string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}{
	ID:        "id",
	Height:    "height",
	VendorID:  "vendor_id",
	PostID:    "post_id",
	Delegator: "delegator",
	Validator: "validator",
	Amount:    "amount",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// Generated where

var StakeWhere = struct {
	ID        whereHelperint
	Height    whereHelperint64
	VendorID  whereHelperint
	PostID    whereHelperstring
	Delegator whereHelperstring
	Validator whereHelperstring
	Amount    whereHelperint64
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	DeletedAt whereHelpernull_Time
}{
	ID:        whereHelperint{field: "\"stakes\".\"id\""},
	Height:    whereHelperint64{field: "\"stakes\".\"height\""},
	VendorID:  whereHelperint{field: "\"stakes\".\"vendor_id\""},
	PostID:    whereHelperstring{field: "\"stakes\".\"post_id\""},
	Delegator: whereHelperstring{field: "\"stakes\".\"delegator\""},
	Validator: whereHelperstring{field: "\"stakes\".\"validator\""},
	Amount:    whereHelperint64{field: "\"stakes\".\"amount\""},
	CreatedAt: whereHelpertime_Time{field: "\"stakes\".\"created_at\""},
	UpdatedAt: whereHelpertime_Time{field: "\"stakes\".\"updated_at\""},
	DeletedAt: whereHelpernull_Time{field: "\"stakes\".\"deleted_at\""},
}

// StakeRels is where relationship names are stored.
var StakeRels = struct {
}{}

// stakeR is where relationships are stored.
type stakeR struct {
}

// NewStruct creates a new relationship struct
func (*stakeR) NewStruct() *stakeR {
	return &stakeR{}
}

// stakeL is where Load methods for each relationship are stored.
type stakeL struct{}

var (
	stakeAllColumns            = []string{"id", "height", "vendor_id", "post_id", "delegator", "validator", "amount", "created_at", "updated_at", "deleted_at"}
	stakeColumnsWithoutDefault = []string{"height", "vendor_id", "post_id", "delegator", "validator", "amount", "deleted_at"}
	stakeColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	stakePrimaryKeyColumns     = []string{"id"}
)

type (
	// StakeSlice is an alias for a slice of pointers to Stake.
	// This should generally be used opposed to []Stake.
	StakeSlice []*Stake

	stakeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	stakeType                 = reflect.TypeOf(&Stake{})
	stakeMapping              = queries.MakeStructMapping(stakeType)
	stakePrimaryKeyMapping, _ = queries.BindMapping(stakeType, stakeMapping, stakePrimaryKeyColumns)
	stakeInsertCacheMut       sync.RWMutex
	stakeInsertCache          = make(map[string]insertCache)
	stakeUpdateCacheMut       sync.RWMutex
	stakeUpdateCache          = make(map[string]updateCache)
	stakeUpsertCacheMut       sync.RWMutex
	stakeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single stake record from the query.
func (q stakeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Stake, error) {
	o := &Stake{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for stakes")
	}

	return o, nil
}

// All returns all Stake records from the query.
func (q stakeQuery) All(ctx context.Context, exec boil.ContextExecutor) (StakeSlice, error) {
	var o []*Stake

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Stake slice")
	}

	return o, nil
}

// Count returns the count of all Stake records in the query.
func (q stakeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count stakes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q stakeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if stakes exists")
	}

	return count > 0, nil
}

// Stakes retrieves all the records using an executor.
func Stakes(mods ...qm.QueryMod) stakeQuery {
	mods = append(mods, qm.From("\"stakes\""))
	return stakeQuery{NewQuery(mods...)}
}

// FindStake retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindStake(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Stake, error) {
	stakeObj := &Stake{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"stakes\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, stakeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from stakes")
	}

	return stakeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Stake) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stakes provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(stakeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	stakeInsertCacheMut.RLock()
	cache, cached := stakeInsertCache[key]
	stakeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			stakeAllColumns,
			stakeColumnsWithDefault,
			stakeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(stakeType, stakeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(stakeType, stakeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"stakes\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"stakes\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into stakes")
	}

	if !cached {
		stakeInsertCacheMut.Lock()
		stakeInsertCache[key] = cache
		stakeInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Stake.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Stake) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	stakeUpdateCacheMut.RLock()
	cache, cached := stakeUpdateCache[key]
	stakeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			stakeAllColumns,
			stakePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update stakes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"stakes\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, stakePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(stakeType, stakeMapping, append(wl, stakePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update stakes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for stakes")
	}

	if !cached {
		stakeUpdateCacheMut.Lock()
		stakeUpdateCache[key] = cache
		stakeUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q stakeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for stakes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for stakes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o StakeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stakePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"stakes\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, stakePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in stake slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all stake")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Stake) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no stakes provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(stakeColumnsWithDefault, o)

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

	stakeUpsertCacheMut.RLock()
	cache, cached := stakeUpsertCache[key]
	stakeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			stakeAllColumns,
			stakeColumnsWithDefault,
			stakeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			stakeAllColumns,
			stakePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert stakes, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(stakePrimaryKeyColumns))
			copy(conflict, stakePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"stakes\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(stakeType, stakeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(stakeType, stakeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert stakes")
	}

	if !cached {
		stakeUpsertCacheMut.Lock()
		stakeUpsertCache[key] = cache
		stakeUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Stake record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Stake) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Stake provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), stakePrimaryKeyMapping)
	sql := "DELETE FROM \"stakes\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from stakes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for stakes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q stakeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no stakeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stakes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stakes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o StakeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stakePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"stakes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stakePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from stake slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for stakes")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Stake) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindStake(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *StakeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := StakeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), stakePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"stakes\".* FROM \"stakes\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, stakePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in StakeSlice")
	}

	*o = slice

	return nil
}

// StakeExists checks if the Stake row exists.
func StakeExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"stakes\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if stakes exists")
	}

	return exists, nil
}
