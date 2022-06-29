// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/fileextract"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/google/uuid"
)

// FileExtractQuery is the builder for querying FileExtract entities.
type FileExtractQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.FileExtract
	// eager-loading edges.
	withFileExtractToEnvironment *EnvironmentQuery
	withFKs                      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FileExtractQuery builder.
func (feq *FileExtractQuery) Where(ps ...predicate.FileExtract) *FileExtractQuery {
	feq.predicates = append(feq.predicates, ps...)
	return feq
}

// Limit adds a limit step to the query.
func (feq *FileExtractQuery) Limit(limit int) *FileExtractQuery {
	feq.limit = &limit
	return feq
}

// Offset adds an offset step to the query.
func (feq *FileExtractQuery) Offset(offset int) *FileExtractQuery {
	feq.offset = &offset
	return feq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (feq *FileExtractQuery) Unique(unique bool) *FileExtractQuery {
	feq.unique = &unique
	return feq
}

// Order adds an order step to the query.
func (feq *FileExtractQuery) Order(o ...OrderFunc) *FileExtractQuery {
	feq.order = append(feq.order, o...)
	return feq
}

// QueryFileExtractToEnvironment chains the current query on the "FileExtractToEnvironment" edge.
func (feq *FileExtractQuery) QueryFileExtractToEnvironment() *EnvironmentQuery {
	query := &EnvironmentQuery{config: feq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := feq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := feq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(fileextract.Table, fileextract.FieldID, selector),
			sqlgraph.To(environment.Table, environment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, fileextract.FileExtractToEnvironmentTable, fileextract.FileExtractToEnvironmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(feq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FileExtract entity from the query.
// Returns a *NotFoundError when no FileExtract was found.
func (feq *FileExtractQuery) First(ctx context.Context) (*FileExtract, error) {
	nodes, err := feq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fileextract.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (feq *FileExtractQuery) FirstX(ctx context.Context) *FileExtract {
	node, err := feq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FileExtract ID from the query.
// Returns a *NotFoundError when no FileExtract ID was found.
func (feq *FileExtractQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = feq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fileextract.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (feq *FileExtractQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := feq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FileExtract entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FileExtract entity is found.
// Returns a *NotFoundError when no FileExtract entities are found.
func (feq *FileExtractQuery) Only(ctx context.Context) (*FileExtract, error) {
	nodes, err := feq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fileextract.Label}
	default:
		return nil, &NotSingularError{fileextract.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (feq *FileExtractQuery) OnlyX(ctx context.Context) *FileExtract {
	node, err := feq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FileExtract ID in the query.
// Returns a *NotSingularError when more than one FileExtract ID is found.
// Returns a *NotFoundError when no entities are found.
func (feq *FileExtractQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = feq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fileextract.Label}
	default:
		err = &NotSingularError{fileextract.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (feq *FileExtractQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := feq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FileExtracts.
func (feq *FileExtractQuery) All(ctx context.Context) ([]*FileExtract, error) {
	if err := feq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return feq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (feq *FileExtractQuery) AllX(ctx context.Context) []*FileExtract {
	nodes, err := feq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FileExtract IDs.
func (feq *FileExtractQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := feq.Select(fileextract.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (feq *FileExtractQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := feq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (feq *FileExtractQuery) Count(ctx context.Context) (int, error) {
	if err := feq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return feq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (feq *FileExtractQuery) CountX(ctx context.Context) int {
	count, err := feq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (feq *FileExtractQuery) Exist(ctx context.Context) (bool, error) {
	if err := feq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return feq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (feq *FileExtractQuery) ExistX(ctx context.Context) bool {
	exist, err := feq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FileExtractQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (feq *FileExtractQuery) Clone() *FileExtractQuery {
	if feq == nil {
		return nil
	}
	return &FileExtractQuery{
		config:                       feq.config,
		limit:                        feq.limit,
		offset:                       feq.offset,
		order:                        append([]OrderFunc{}, feq.order...),
		predicates:                   append([]predicate.FileExtract{}, feq.predicates...),
		withFileExtractToEnvironment: feq.withFileExtractToEnvironment.Clone(),
		// clone intermediate query.
		sql:    feq.sql.Clone(),
		path:   feq.path,
		unique: feq.unique,
	}
}

// WithFileExtractToEnvironment tells the query-builder to eager-load the nodes that are connected to
// the "FileExtractToEnvironment" edge. The optional arguments are used to configure the query builder of the edge.
func (feq *FileExtractQuery) WithFileExtractToEnvironment(opts ...func(*EnvironmentQuery)) *FileExtractQuery {
	query := &EnvironmentQuery{config: feq.config}
	for _, opt := range opts {
		opt(query)
	}
	feq.withFileExtractToEnvironment = query
	return feq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HclID string `json:"hcl_id,omitempty" hcl:"id,label"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FileExtract.Query().
//		GroupBy(fileextract.FieldHclID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (feq *FileExtractQuery) GroupBy(field string, fields ...string) *FileExtractGroupBy {
	group := &FileExtractGroupBy{config: feq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := feq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return feq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HclID string `json:"hcl_id,omitempty" hcl:"id,label"`
//	}
//
//	client.FileExtract.Query().
//		Select(fileextract.FieldHclID).
//		Scan(ctx, &v)
//
func (feq *FileExtractQuery) Select(fields ...string) *FileExtractSelect {
	feq.fields = append(feq.fields, fields...)
	return &FileExtractSelect{FileExtractQuery: feq}
}

func (feq *FileExtractQuery) prepareQuery(ctx context.Context) error {
	for _, f := range feq.fields {
		if !fileextract.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if feq.path != nil {
		prev, err := feq.path(ctx)
		if err != nil {
			return err
		}
		feq.sql = prev
	}
	return nil
}

func (feq *FileExtractQuery) sqlAll(ctx context.Context) ([]*FileExtract, error) {
	var (
		nodes       = []*FileExtract{}
		withFKs     = feq.withFKs
		_spec       = feq.querySpec()
		loadedTypes = [1]bool{
			feq.withFileExtractToEnvironment != nil,
		}
	)
	if feq.withFileExtractToEnvironment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, fileextract.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &FileExtract{config: feq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, feq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := feq.withFileExtractToEnvironment; query != nil {
		ids := make([]uuid.UUID, 0, len(nodes))
		nodeids := make(map[uuid.UUID][]*FileExtract)
		for i := range nodes {
			if nodes[i].environment_environment_to_file_extract == nil {
				continue
			}
			fk := *nodes[i].environment_environment_to_file_extract
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(environment.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "environment_environment_to_file_extract" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.FileExtractToEnvironment = n
			}
		}
	}

	return nodes, nil
}

func (feq *FileExtractQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := feq.querySpec()
	_spec.Node.Columns = feq.fields
	if len(feq.fields) > 0 {
		_spec.Unique = feq.unique != nil && *feq.unique
	}
	return sqlgraph.CountNodes(ctx, feq.driver, _spec)
}

func (feq *FileExtractQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := feq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (feq *FileExtractQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fileextract.Table,
			Columns: fileextract.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fileextract.FieldID,
			},
		},
		From:   feq.sql,
		Unique: true,
	}
	if unique := feq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := feq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fileextract.FieldID)
		for i := range fields {
			if fields[i] != fileextract.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := feq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := feq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := feq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := feq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (feq *FileExtractQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(feq.driver.Dialect())
	t1 := builder.Table(fileextract.Table)
	columns := feq.fields
	if len(columns) == 0 {
		columns = fileextract.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if feq.sql != nil {
		selector = feq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if feq.unique != nil && *feq.unique {
		selector.Distinct()
	}
	for _, p := range feq.predicates {
		p(selector)
	}
	for _, p := range feq.order {
		p(selector)
	}
	if offset := feq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := feq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FileExtractGroupBy is the group-by builder for FileExtract entities.
type FileExtractGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fegb *FileExtractGroupBy) Aggregate(fns ...AggregateFunc) *FileExtractGroupBy {
	fegb.fns = append(fegb.fns, fns...)
	return fegb
}

// Scan applies the group-by query and scans the result into the given value.
func (fegb *FileExtractGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := fegb.path(ctx)
	if err != nil {
		return err
	}
	fegb.sql = query
	return fegb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fegb *FileExtractGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := fegb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (fegb *FileExtractGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(fegb.fields) > 1 {
		return nil, errors.New("ent: FileExtractGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := fegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fegb *FileExtractGroupBy) StringsX(ctx context.Context) []string {
	v, err := fegb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fegb *FileExtractGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fegb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fileextract.Label}
	default:
		err = fmt.Errorf("ent: FileExtractGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fegb *FileExtractGroupBy) StringX(ctx context.Context) string {
	v, err := fegb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (fegb *FileExtractGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(fegb.fields) > 1 {
		return nil, errors.New("ent: FileExtractGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := fegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fegb *FileExtractGroupBy) IntsX(ctx context.Context) []int {
	v, err := fegb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fegb *FileExtractGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fegb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fileextract.Label}
	default:
		err = fmt.Errorf("ent: FileExtractGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fegb *FileExtractGroupBy) IntX(ctx context.Context) int {
	v, err := fegb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (fegb *FileExtractGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(fegb.fields) > 1 {
		return nil, errors.New("ent: FileExtractGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := fegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fegb *FileExtractGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := fegb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fegb *FileExtractGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fegb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fileextract.Label}
	default:
		err = fmt.Errorf("ent: FileExtractGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fegb *FileExtractGroupBy) Float64X(ctx context.Context) float64 {
	v, err := fegb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (fegb *FileExtractGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(fegb.fields) > 1 {
		return nil, errors.New("ent: FileExtractGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := fegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fegb *FileExtractGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := fegb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (fegb *FileExtractGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fegb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fileextract.Label}
	default:
		err = fmt.Errorf("ent: FileExtractGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fegb *FileExtractGroupBy) BoolX(ctx context.Context) bool {
	v, err := fegb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fegb *FileExtractGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range fegb.fields {
		if !fileextract.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fegb *FileExtractGroupBy) sqlQuery() *sql.Selector {
	selector := fegb.sql.Select()
	aggregation := make([]string, 0, len(fegb.fns))
	for _, fn := range fegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fegb.fields)+len(fegb.fns))
		for _, f := range fegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fegb.fields...)...)
}

// FileExtractSelect is the builder for selecting fields of FileExtract entities.
type FileExtractSelect struct {
	*FileExtractQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fes *FileExtractSelect) Scan(ctx context.Context, v interface{}) error {
	if err := fes.prepareQuery(ctx); err != nil {
		return err
	}
	fes.sql = fes.FileExtractQuery.sqlQuery(ctx)
	return fes.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (fes *FileExtractSelect) ScanX(ctx context.Context, v interface{}) {
	if err := fes.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (fes *FileExtractSelect) Strings(ctx context.Context) ([]string, error) {
	if len(fes.fields) > 1 {
		return nil, errors.New("ent: FileExtractSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := fes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (fes *FileExtractSelect) StringsX(ctx context.Context) []string {
	v, err := fes.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (fes *FileExtractSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = fes.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fileextract.Label}
	default:
		err = fmt.Errorf("ent: FileExtractSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (fes *FileExtractSelect) StringX(ctx context.Context) string {
	v, err := fes.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (fes *FileExtractSelect) Ints(ctx context.Context) ([]int, error) {
	if len(fes.fields) > 1 {
		return nil, errors.New("ent: FileExtractSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := fes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (fes *FileExtractSelect) IntsX(ctx context.Context) []int {
	v, err := fes.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (fes *FileExtractSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = fes.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fileextract.Label}
	default:
		err = fmt.Errorf("ent: FileExtractSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (fes *FileExtractSelect) IntX(ctx context.Context) int {
	v, err := fes.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (fes *FileExtractSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(fes.fields) > 1 {
		return nil, errors.New("ent: FileExtractSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := fes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (fes *FileExtractSelect) Float64sX(ctx context.Context) []float64 {
	v, err := fes.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (fes *FileExtractSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = fes.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fileextract.Label}
	default:
		err = fmt.Errorf("ent: FileExtractSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (fes *FileExtractSelect) Float64X(ctx context.Context) float64 {
	v, err := fes.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (fes *FileExtractSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(fes.fields) > 1 {
		return nil, errors.New("ent: FileExtractSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := fes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (fes *FileExtractSelect) BoolsX(ctx context.Context) []bool {
	v, err := fes.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (fes *FileExtractSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = fes.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{fileextract.Label}
	default:
		err = fmt.Errorf("ent: FileExtractSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (fes *FileExtractSelect) BoolX(ctx context.Context) bool {
	v, err := fes.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (fes *FileExtractSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := fes.sql.Query()
	if err := fes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
