// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package datamodel

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

// InsertAll inserts all rows with the specified column values, using an executor.
func (o ReportSlice) InsertAll(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}
	var sql string
	vals := []interface{}{}
	for i, row := range o {
		if !boil.TimestampsAreSkipped(ctx) {
			currTime := time.Now().In(boil.GetLocation())

			if queries.MustTime(row.CreatedAt).IsZero() {
				queries.SetScanner(&row.CreatedAt, currTime)
			}
			if queries.MustTime(row.UpdatedAt).IsZero() {
				queries.SetScanner(&row.UpdatedAt, currTime)
			}
		}

		if err := row.doBeforeInsertHooks(ctx, exec); err != nil {
			return err
		}

		nzDefaults := queries.NonZeroDefaultSet(reportColumnsWithDefault, row)
		wl, _ := columns.InsertColumnSet(
			reportAllColumns,
			reportColumnsWithDefault,
			reportColumnsWithoutDefault,
			nzDefaults,
		)
		if i == 0 {
			sql = "INSERT INTO `reports` " + "(`" + strings.Join(wl, "`,`") + "`)" + " VALUES "
		}
		sql += strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), len(vals)+1, len(wl))
		if i != len(o)-1 {
			sql += ","
		}
		valMapping, err := queries.BindMapping(reportType, reportMapping, wl)
		if err != nil {
			return err
		}
		value := reflect.Indirect(reflect.ValueOf(row))
		vals = append(vals, queries.ValuesFromMapping(value, valMapping)...)
	}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, vals...)
	}

	_, err := exec.ExecContext(ctx, sql, vals...)
	if err != nil {
		return errors.Wrap(err, "datamodel: unable to insert into reports")
	}

	return nil
}

// Report is an object representing the database table.
type Report struct {
	ID           string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Comment      null.String `boil:"comment" json:"comment,omitempty" toml:"comment" yaml:"comment,omitempty"`
	ThumbnailURL null.String `boil:"thumbnail_url" json:"thumbnail_url,omitempty" toml:"thumbnail_url" yaml:"thumbnail_url,omitempty"`
	UserID       string      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	RecipeID     string      `boil:"recipe_id" json:"recipe_id" toml:"recipe_id" yaml:"recipe_id"`
	CreatedAt    null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt    null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *reportR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L reportL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ReportColumns = struct {
	ID           string
	Comment      string
	ThumbnailURL string
	UserID       string
	RecipeID     string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	Comment:      "comment",
	ThumbnailURL: "thumbnail_url",
	UserID:       "user_id",
	RecipeID:     "recipe_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

var ReportTableColumns = struct {
	ID           string
	Comment      string
	ThumbnailURL string
	UserID       string
	RecipeID     string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "reports.id",
	Comment:      "reports.comment",
	ThumbnailURL: "reports.thumbnail_url",
	UserID:       "reports.user_id",
	RecipeID:     "reports.recipe_id",
	CreatedAt:    "reports.created_at",
	UpdatedAt:    "reports.updated_at",
}

// Generated where

var ReportWhere = struct {
	ID           whereHelperstring
	Comment      whereHelpernull_String
	ThumbnailURL whereHelpernull_String
	UserID       whereHelperstring
	RecipeID     whereHelperstring
	CreatedAt    whereHelpernull_Time
	UpdatedAt    whereHelpernull_Time
}{
	ID:           whereHelperstring{field: "`reports`.`id`"},
	Comment:      whereHelpernull_String{field: "`reports`.`comment`"},
	ThumbnailURL: whereHelpernull_String{field: "`reports`.`thumbnail_url`"},
	UserID:       whereHelperstring{field: "`reports`.`user_id`"},
	RecipeID:     whereHelperstring{field: "`reports`.`recipe_id`"},
	CreatedAt:    whereHelpernull_Time{field: "`reports`.`created_at`"},
	UpdatedAt:    whereHelpernull_Time{field: "`reports`.`updated_at`"},
}

// ReportRels is where relationship names are stored.
var ReportRels = struct {
	Recipe string
	User   string
}{
	Recipe: "Recipe",
	User:   "User",
}

// reportR is where relationships are stored.
type reportR struct {
	Recipe *Recipe `boil:"Recipe" json:"Recipe" toml:"Recipe" yaml:"Recipe"`
	User   *User   `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*reportR) NewStruct() *reportR {
	return &reportR{}
}

func (r *reportR) GetRecipe() *Recipe {
	if r == nil {
		return nil
	}
	return r.Recipe
}

func (r *reportR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// reportL is where Load methods for each relationship are stored.
type reportL struct{}

var (
	reportAllColumns            = []string{"id", "comment", "thumbnail_url", "user_id", "recipe_id", "created_at", "updated_at"}
	reportColumnsWithoutDefault = []string{"id", "comment", "thumbnail_url", "user_id", "recipe_id"}
	reportColumnsWithDefault    = []string{"created_at", "updated_at"}
	reportPrimaryKeyColumns     = []string{"id"}
	reportGeneratedColumns      = []string{}
)

type (
	// ReportSlice is an alias for a slice of pointers to Report.
	// This should almost always be used instead of []Report.
	ReportSlice []*Report
	// ReportHook is the signature for custom Report hook methods
	ReportHook func(context.Context, boil.ContextExecutor, *Report) error

	reportQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	reportType                 = reflect.TypeOf(&Report{})
	reportMapping              = queries.MakeStructMapping(reportType)
	reportPrimaryKeyMapping, _ = queries.BindMapping(reportType, reportMapping, reportPrimaryKeyColumns)
	reportInsertCacheMut       sync.RWMutex
	reportInsertCache          = make(map[string]insertCache)
	reportUpdateCacheMut       sync.RWMutex
	reportUpdateCache          = make(map[string]updateCache)
	reportUpsertCacheMut       sync.RWMutex
	reportUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var reportAfterSelectHooks []ReportHook

var reportBeforeInsertHooks []ReportHook
var reportAfterInsertHooks []ReportHook

var reportBeforeUpdateHooks []ReportHook
var reportAfterUpdateHooks []ReportHook

var reportBeforeDeleteHooks []ReportHook
var reportAfterDeleteHooks []ReportHook

var reportBeforeUpsertHooks []ReportHook
var reportAfterUpsertHooks []ReportHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Report) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reportAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Report) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reportBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Report) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reportAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Report) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reportBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Report) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reportAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Report) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reportBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Report) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reportAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Report) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reportBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Report) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range reportAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddReportHook registers your hook function for all future operations.
func AddReportHook(hookPoint boil.HookPoint, reportHook ReportHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		reportAfterSelectHooks = append(reportAfterSelectHooks, reportHook)
	case boil.BeforeInsertHook:
		reportBeforeInsertHooks = append(reportBeforeInsertHooks, reportHook)
	case boil.AfterInsertHook:
		reportAfterInsertHooks = append(reportAfterInsertHooks, reportHook)
	case boil.BeforeUpdateHook:
		reportBeforeUpdateHooks = append(reportBeforeUpdateHooks, reportHook)
	case boil.AfterUpdateHook:
		reportAfterUpdateHooks = append(reportAfterUpdateHooks, reportHook)
	case boil.BeforeDeleteHook:
		reportBeforeDeleteHooks = append(reportBeforeDeleteHooks, reportHook)
	case boil.AfterDeleteHook:
		reportAfterDeleteHooks = append(reportAfterDeleteHooks, reportHook)
	case boil.BeforeUpsertHook:
		reportBeforeUpsertHooks = append(reportBeforeUpsertHooks, reportHook)
	case boil.AfterUpsertHook:
		reportAfterUpsertHooks = append(reportAfterUpsertHooks, reportHook)
	}
}

// One returns a single report record from the query.
func (q reportQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Report, error) {
	o := &Report{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datamodel: failed to execute a one query for reports")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Report records from the query.
func (q reportQuery) All(ctx context.Context, exec boil.ContextExecutor) (ReportSlice, error) {
	var o []*Report

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "datamodel: failed to assign all query results to Report slice")
	}

	if len(reportAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Report records in the query.
func (q reportQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: failed to count reports rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q reportQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "datamodel: failed to check if reports exists")
	}

	return count > 0, nil
}

// Recipe pointed to by the foreign key.
func (o *Report) Recipe(mods ...qm.QueryMod) recipeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.RecipeID),
	}

	queryMods = append(queryMods, mods...)

	return Recipes(queryMods...)
}

// User pointed to by the foreign key.
func (o *Report) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadRecipe allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (reportL) LoadRecipe(ctx context.Context, e boil.ContextExecutor, singular bool, maybeReport interface{}, mods queries.Applicator) error {
	var slice []*Report
	var object *Report

	if singular {
		var ok bool
		object, ok = maybeReport.(*Report)
		if !ok {
			object = new(Report)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeReport)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeReport))
			}
		}
	} else {
		s, ok := maybeReport.(*[]*Report)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeReport)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeReport))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &reportR{}
		}
		args = append(args, object.RecipeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &reportR{}
			}

			for _, a := range args {
				if a == obj.RecipeID {
					continue Outer
				}
			}

			args = append(args, obj.RecipeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`recipes`),
		qm.WhereIn(`recipes.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Recipe")
	}

	var resultSlice []*Recipe
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Recipe")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for recipes")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for recipes")
	}

	if len(recipeAfterSelectHooks) != 0 {
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
		object.R.Recipe = foreign
		if foreign.R == nil {
			foreign.R = &recipeR{}
		}
		foreign.R.Reports = append(foreign.R.Reports, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.RecipeID == foreign.ID {
				local.R.Recipe = foreign
				if foreign.R == nil {
					foreign.R = &recipeR{}
				}
				foreign.R.Reports = append(foreign.R.Reports, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (reportL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeReport interface{}, mods queries.Applicator) error {
	var slice []*Report
	var object *Report

	if singular {
		var ok bool
		object, ok = maybeReport.(*Report)
		if !ok {
			object = new(Report)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeReport)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeReport))
			}
		}
	} else {
		s, ok := maybeReport.(*[]*Report)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeReport)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeReport))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &reportR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &reportR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Reports = append(foreign.R.Reports, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Reports = append(foreign.R.Reports, local)
				break
			}
		}
	}

	return nil
}

// SetRecipe of the report to the related item.
// Sets o.R.Recipe to related.
// Adds o to related.R.Reports.
func (o *Report) SetRecipe(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Recipe) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `reports` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"recipe_id"}),
		strmangle.WhereClause("`", "`", 0, reportPrimaryKeyColumns),
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

	o.RecipeID = related.ID
	if o.R == nil {
		o.R = &reportR{
			Recipe: related,
		}
	} else {
		o.R.Recipe = related
	}

	if related.R == nil {
		related.R = &recipeR{
			Reports: ReportSlice{o},
		}
	} else {
		related.R.Reports = append(related.R.Reports, o)
	}

	return nil
}

// SetUser of the report to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Reports.
func (o *Report) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `reports` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, reportPrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &reportR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Reports: ReportSlice{o},
		}
	} else {
		related.R.Reports = append(related.R.Reports, o)
	}

	return nil
}

// Reports retrieves all the records using an executor.
func Reports(mods ...qm.QueryMod) reportQuery {
	mods = append(mods, qm.From("`reports`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`reports`.*"})
	}

	return reportQuery{q}
}

// FindReport retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReport(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Report, error) {
	reportObj := &Report{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `reports` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, reportObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "datamodel: unable to select from reports")
	}

	if err = reportObj.doAfterSelectHooks(ctx, exec); err != nil {
		return reportObj, err
	}

	return reportObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Report) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("datamodel: no reports provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(reportColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	reportInsertCacheMut.RLock()
	cache, cached := reportInsertCache[key]
	reportInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			reportAllColumns,
			reportColumnsWithDefault,
			reportColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(reportType, reportMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(reportType, reportMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `reports` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `reports` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `reports` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, reportPrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "datamodel: unable to insert into reports")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "datamodel: unable to populate default values for reports")
	}

CacheNoHooks:
	if !cached {
		reportInsertCacheMut.Lock()
		reportInsertCache[key] = cache
		reportInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Report.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Report) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	reportUpdateCacheMut.RLock()
	cache, cached := reportUpdateCache[key]
	reportUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			reportAllColumns,
			reportPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("datamodel: unable to update reports, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `reports` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, reportPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(reportType, reportMapping, append(wl, reportPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "datamodel: unable to update reports row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: failed to get rows affected by update for reports")
	}

	if !cached {
		reportUpdateCacheMut.Lock()
		reportUpdateCache[key] = cache
		reportUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q reportQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: unable to update all for reports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: unable to retrieve rows affected for reports")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReportSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("datamodel: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `reports` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, reportPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: unable to update all in report slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: unable to retrieve rows affected all in update all report")
	}
	return rowsAff, nil
}

// Delete deletes a single Report record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Report) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("datamodel: no Report provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), reportPrimaryKeyMapping)
	sql := "DELETE FROM `reports` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: unable to delete from reports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: failed to get rows affected by delete for reports")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q reportQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("datamodel: no reportQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: unable to delete all from reports")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: failed to get rows affected by deleteall for reports")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReportSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(reportBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `reports` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, reportPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: unable to delete all from report slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "datamodel: failed to get rows affected by deleteall for reports")
	}

	if len(reportAfterDeleteHooks) != 0 {
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
func (o *Report) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindReport(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReportSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ReportSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reportPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `reports`.* FROM `reports` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, reportPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "datamodel: unable to reload all in ReportSlice")
	}

	*o = slice

	return nil
}

// ReportExists checks if the Report row exists.
func ReportExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `reports` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "datamodel: unable to check if reports exists")
	}

	return exists, nil
}

// Exists checks if the Report row exists.
func (o *Report) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return ReportExists(ctx, exec, o.ID)
}

var mySQLReportUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Report) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("datamodel: no reports provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(reportColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLReportUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	reportUpsertCacheMut.RLock()
	cache, cached := reportUpsertCache[key]
	reportUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			reportAllColumns,
			reportColumnsWithDefault,
			reportColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			reportAllColumns,
			reportPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("datamodel: unable to upsert reports, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`reports`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `reports` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(reportType, reportMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(reportType, reportMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "datamodel: unable to upsert for reports")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(reportType, reportMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "datamodel: unable to retrieve unique values for reports")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "datamodel: unable to populate default values for reports")
	}

CacheNoHooks:
	if !cached {
		reportUpsertCacheMut.Lock()
		reportUpsertCache[key] = cache
		reportUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
