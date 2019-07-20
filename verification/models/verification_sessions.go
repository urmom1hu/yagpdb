// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// VerificationSession is an object representing the database table.
type VerificationSession struct {
	Token     string    `boil:"token" json:"token" toml:"token" yaml:"token"`
	UserID    int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GuildID   int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	SolvedAt  null.Time `boil:"solved_at" json:"solved_at,omitempty" toml:"solved_at" yaml:"solved_at,omitempty"`
	ExpiredAt null.Time `boil:"expired_at" json:"expired_at,omitempty" toml:"expired_at" yaml:"expired_at,omitempty"`

	R *verificationSessionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L verificationSessionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var VerificationSessionColumns = struct {
	Token     string
	UserID    string
	GuildID   string
	CreatedAt string
	SolvedAt  string
	ExpiredAt string
}{
	Token:     "token",
	UserID:    "user_id",
	GuildID:   "guild_id",
	CreatedAt: "created_at",
	SolvedAt:  "solved_at",
	ExpiredAt: "expired_at",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var VerificationSessionWhere = struct {
	Token     whereHelperstring
	UserID    whereHelperint64
	GuildID   whereHelperint64
	CreatedAt whereHelpertime_Time
	SolvedAt  whereHelpernull_Time
	ExpiredAt whereHelpernull_Time
}{
	Token:     whereHelperstring{field: "\"verification_sessions\".\"token\""},
	UserID:    whereHelperint64{field: "\"verification_sessions\".\"user_id\""},
	GuildID:   whereHelperint64{field: "\"verification_sessions\".\"guild_id\""},
	CreatedAt: whereHelpertime_Time{field: "\"verification_sessions\".\"created_at\""},
	SolvedAt:  whereHelpernull_Time{field: "\"verification_sessions\".\"solved_at\""},
	ExpiredAt: whereHelpernull_Time{field: "\"verification_sessions\".\"expired_at\""},
}

// VerificationSessionRels is where relationship names are stored.
var VerificationSessionRels = struct {
}{}

// verificationSessionR is where relationships are stored.
type verificationSessionR struct {
}

// NewStruct creates a new relationship struct
func (*verificationSessionR) NewStruct() *verificationSessionR {
	return &verificationSessionR{}
}

// verificationSessionL is where Load methods for each relationship are stored.
type verificationSessionL struct{}

var (
	verificationSessionAllColumns            = []string{"token", "user_id", "guild_id", "created_at", "solved_at", "expired_at"}
	verificationSessionColumnsWithoutDefault = []string{"token", "user_id", "guild_id", "created_at", "solved_at", "expired_at"}
	verificationSessionColumnsWithDefault    = []string{}
	verificationSessionPrimaryKeyColumns     = []string{"token"}
)

type (
	// VerificationSessionSlice is an alias for a slice of pointers to VerificationSession.
	// This should generally be used opposed to []VerificationSession.
	VerificationSessionSlice []*VerificationSession

	verificationSessionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	verificationSessionType                 = reflect.TypeOf(&VerificationSession{})
	verificationSessionMapping              = queries.MakeStructMapping(verificationSessionType)
	verificationSessionPrimaryKeyMapping, _ = queries.BindMapping(verificationSessionType, verificationSessionMapping, verificationSessionPrimaryKeyColumns)
	verificationSessionInsertCacheMut       sync.RWMutex
	verificationSessionInsertCache          = make(map[string]insertCache)
	verificationSessionUpdateCacheMut       sync.RWMutex
	verificationSessionUpdateCache          = make(map[string]updateCache)
	verificationSessionUpsertCacheMut       sync.RWMutex
	verificationSessionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single verificationSession record from the query using the global executor.
func (q verificationSessionQuery) OneG(ctx context.Context) (*VerificationSession, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single verificationSession record from the query.
func (q verificationSessionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*VerificationSession, error) {
	o := &VerificationSession{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for verification_sessions")
	}

	return o, nil
}

// AllG returns all VerificationSession records from the query using the global executor.
func (q verificationSessionQuery) AllG(ctx context.Context) (VerificationSessionSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all VerificationSession records from the query.
func (q verificationSessionQuery) All(ctx context.Context, exec boil.ContextExecutor) (VerificationSessionSlice, error) {
	var o []*VerificationSession

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to VerificationSession slice")
	}

	return o, nil
}

// CountG returns the count of all VerificationSession records in the query, and panics on error.
func (q verificationSessionQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all VerificationSession records in the query.
func (q verificationSessionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count verification_sessions rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q verificationSessionQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q verificationSessionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if verification_sessions exists")
	}

	return count > 0, nil
}

// VerificationSessions retrieves all the records using an executor.
func VerificationSessions(mods ...qm.QueryMod) verificationSessionQuery {
	mods = append(mods, qm.From("\"verification_sessions\""))
	return verificationSessionQuery{NewQuery(mods...)}
}

// FindVerificationSessionG retrieves a single record by ID.
func FindVerificationSessionG(ctx context.Context, token string, selectCols ...string) (*VerificationSession, error) {
	return FindVerificationSession(ctx, boil.GetContextDB(), token, selectCols...)
}

// FindVerificationSession retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindVerificationSession(ctx context.Context, exec boil.ContextExecutor, token string, selectCols ...string) (*VerificationSession, error) {
	verificationSessionObj := &VerificationSession{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"verification_sessions\" where \"token\"=$1", sel,
	)

	q := queries.Raw(query, token)

	err := q.Bind(ctx, exec, verificationSessionObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from verification_sessions")
	}

	return verificationSessionObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *VerificationSession) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *VerificationSession) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no verification_sessions provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(verificationSessionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	verificationSessionInsertCacheMut.RLock()
	cache, cached := verificationSessionInsertCache[key]
	verificationSessionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			verificationSessionAllColumns,
			verificationSessionColumnsWithDefault,
			verificationSessionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(verificationSessionType, verificationSessionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(verificationSessionType, verificationSessionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"verification_sessions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"verification_sessions\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into verification_sessions")
	}

	if !cached {
		verificationSessionInsertCacheMut.Lock()
		verificationSessionInsertCache[key] = cache
		verificationSessionInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single VerificationSession record using the global executor.
// See Update for more documentation.
func (o *VerificationSession) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the VerificationSession.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *VerificationSession) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	verificationSessionUpdateCacheMut.RLock()
	cache, cached := verificationSessionUpdateCache[key]
	verificationSessionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			verificationSessionAllColumns,
			verificationSessionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update verification_sessions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"verification_sessions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, verificationSessionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(verificationSessionType, verificationSessionMapping, append(wl, verificationSessionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update verification_sessions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for verification_sessions")
	}

	if !cached {
		verificationSessionUpdateCacheMut.Lock()
		verificationSessionUpdateCache[key] = cache
		verificationSessionUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q verificationSessionQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q verificationSessionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for verification_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for verification_sessions")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o VerificationSessionSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o VerificationSessionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), verificationSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"verification_sessions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, verificationSessionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in verificationSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all verificationSession")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *VerificationSession) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *VerificationSession) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no verification_sessions provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(verificationSessionColumnsWithDefault, o)

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

	verificationSessionUpsertCacheMut.RLock()
	cache, cached := verificationSessionUpsertCache[key]
	verificationSessionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			verificationSessionAllColumns,
			verificationSessionColumnsWithDefault,
			verificationSessionColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			verificationSessionAllColumns,
			verificationSessionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert verification_sessions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(verificationSessionPrimaryKeyColumns))
			copy(conflict, verificationSessionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"verification_sessions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(verificationSessionType, verificationSessionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(verificationSessionType, verificationSessionMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
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
		return errors.Wrap(err, "models: unable to upsert verification_sessions")
	}

	if !cached {
		verificationSessionUpsertCacheMut.Lock()
		verificationSessionUpsertCache[key] = cache
		verificationSessionUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single VerificationSession record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *VerificationSession) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single VerificationSession record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *VerificationSession) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no VerificationSession provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), verificationSessionPrimaryKeyMapping)
	sql := "DELETE FROM \"verification_sessions\" WHERE \"token\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from verification_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for verification_sessions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q verificationSessionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no verificationSessionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from verification_sessions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for verification_sessions")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o VerificationSessionSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o VerificationSessionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), verificationSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"verification_sessions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, verificationSessionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from verificationSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for verification_sessions")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *VerificationSession) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no VerificationSession provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *VerificationSession) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindVerificationSession(ctx, exec, o.Token)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VerificationSessionSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty VerificationSessionSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *VerificationSessionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := VerificationSessionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), verificationSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"verification_sessions\".* FROM \"verification_sessions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, verificationSessionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in VerificationSessionSlice")
	}

	*o = slice

	return nil
}

// VerificationSessionExistsG checks if the VerificationSession row exists.
func VerificationSessionExistsG(ctx context.Context, token string) (bool, error) {
	return VerificationSessionExists(ctx, boil.GetContextDB(), token)
}

// VerificationSessionExists checks if the VerificationSession row exists.
func VerificationSessionExists(ctx context.Context, exec boil.ContextExecutor, token string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"verification_sessions\" where \"token\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, token)
	}

	row := exec.QueryRowContext(ctx, sql, token)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if verification_sessions exists")
	}

	return exists, nil
}