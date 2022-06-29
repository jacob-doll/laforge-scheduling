// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/ginfilemiddleware"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisioningstep"
	"github.com/google/uuid"
)

// GinFileMiddlewareQuery is the builder for querying GinFileMiddleware entities.
type GinFileMiddlewareQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GinFileMiddleware
	// eager-loading edges.
	withGinFileMiddlewareToProvisionedHost  *ProvisionedHostQuery
	withGinFileMiddlewareToProvisioningStep *ProvisioningStepQuery
	withFKs                                 bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GinFileMiddlewareQuery builder.
func (gfmq *GinFileMiddlewareQuery) Where(ps ...predicate.GinFileMiddleware) *GinFileMiddlewareQuery {
	gfmq.predicates = append(gfmq.predicates, ps...)
	return gfmq
}

// Limit adds a limit step to the query.
func (gfmq *GinFileMiddlewareQuery) Limit(limit int) *GinFileMiddlewareQuery {
	gfmq.limit = &limit
	return gfmq
}

// Offset adds an offset step to the query.
func (gfmq *GinFileMiddlewareQuery) Offset(offset int) *GinFileMiddlewareQuery {
	gfmq.offset = &offset
	return gfmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gfmq *GinFileMiddlewareQuery) Unique(unique bool) *GinFileMiddlewareQuery {
	gfmq.unique = &unique
	return gfmq
}

// Order adds an order step to the query.
func (gfmq *GinFileMiddlewareQuery) Order(o ...OrderFunc) *GinFileMiddlewareQuery {
	gfmq.order = append(gfmq.order, o...)
	return gfmq
}

// QueryGinFileMiddlewareToProvisionedHost chains the current query on the "GinFileMiddlewareToProvisionedHost" edge.
func (gfmq *GinFileMiddlewareQuery) QueryGinFileMiddlewareToProvisionedHost() *ProvisionedHostQuery {
	query := &ProvisionedHostQuery{config: gfmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gfmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gfmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ginfilemiddleware.Table, ginfilemiddleware.FieldID, selector),
			sqlgraph.To(provisionedhost.Table, provisionedhost.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ginfilemiddleware.GinFileMiddlewareToProvisionedHostTable, ginfilemiddleware.GinFileMiddlewareToProvisionedHostColumn),
		)
		fromU = sqlgraph.SetNeighbors(gfmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGinFileMiddlewareToProvisioningStep chains the current query on the "GinFileMiddlewareToProvisioningStep" edge.
func (gfmq *GinFileMiddlewareQuery) QueryGinFileMiddlewareToProvisioningStep() *ProvisioningStepQuery {
	query := &ProvisioningStepQuery{config: gfmq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := gfmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := gfmq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(ginfilemiddleware.Table, ginfilemiddleware.FieldID, selector),
			sqlgraph.To(provisioningstep.Table, provisioningstep.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, ginfilemiddleware.GinFileMiddlewareToProvisioningStepTable, ginfilemiddleware.GinFileMiddlewareToProvisioningStepColumn),
		)
		fromU = sqlgraph.SetNeighbors(gfmq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GinFileMiddleware entity from the query.
// Returns a *NotFoundError when no GinFileMiddleware was found.
func (gfmq *GinFileMiddlewareQuery) First(ctx context.Context) (*GinFileMiddleware, error) {
	nodes, err := gfmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{ginfilemiddleware.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gfmq *GinFileMiddlewareQuery) FirstX(ctx context.Context) *GinFileMiddleware {
	node, err := gfmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GinFileMiddleware ID from the query.
// Returns a *NotFoundError when no GinFileMiddleware ID was found.
func (gfmq *GinFileMiddlewareQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gfmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{ginfilemiddleware.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gfmq *GinFileMiddlewareQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := gfmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GinFileMiddleware entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GinFileMiddleware entity is found.
// Returns a *NotFoundError when no GinFileMiddleware entities are found.
func (gfmq *GinFileMiddlewareQuery) Only(ctx context.Context) (*GinFileMiddleware, error) {
	nodes, err := gfmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{ginfilemiddleware.Label}
	default:
		return nil, &NotSingularError{ginfilemiddleware.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gfmq *GinFileMiddlewareQuery) OnlyX(ctx context.Context) *GinFileMiddleware {
	node, err := gfmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GinFileMiddleware ID in the query.
// Returns a *NotSingularError when more than one GinFileMiddleware ID is found.
// Returns a *NotFoundError when no entities are found.
func (gfmq *GinFileMiddlewareQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = gfmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{ginfilemiddleware.Label}
	default:
		err = &NotSingularError{ginfilemiddleware.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gfmq *GinFileMiddlewareQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := gfmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GinFileMiddlewares.
func (gfmq *GinFileMiddlewareQuery) All(ctx context.Context) ([]*GinFileMiddleware, error) {
	if err := gfmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gfmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gfmq *GinFileMiddlewareQuery) AllX(ctx context.Context) []*GinFileMiddleware {
	nodes, err := gfmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GinFileMiddleware IDs.
func (gfmq *GinFileMiddlewareQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := gfmq.Select(ginfilemiddleware.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gfmq *GinFileMiddlewareQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := gfmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gfmq *GinFileMiddlewareQuery) Count(ctx context.Context) (int, error) {
	if err := gfmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gfmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gfmq *GinFileMiddlewareQuery) CountX(ctx context.Context) int {
	count, err := gfmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gfmq *GinFileMiddlewareQuery) Exist(ctx context.Context) (bool, error) {
	if err := gfmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gfmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gfmq *GinFileMiddlewareQuery) ExistX(ctx context.Context) bool {
	exist, err := gfmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GinFileMiddlewareQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gfmq *GinFileMiddlewareQuery) Clone() *GinFileMiddlewareQuery {
	if gfmq == nil {
		return nil
	}
	return &GinFileMiddlewareQuery{
		config:                                  gfmq.config,
		limit:                                   gfmq.limit,
		offset:                                  gfmq.offset,
		order:                                   append([]OrderFunc{}, gfmq.order...),
		predicates:                              append([]predicate.GinFileMiddleware{}, gfmq.predicates...),
		withGinFileMiddlewareToProvisionedHost:  gfmq.withGinFileMiddlewareToProvisionedHost.Clone(),
		withGinFileMiddlewareToProvisioningStep: gfmq.withGinFileMiddlewareToProvisioningStep.Clone(),
		// clone intermediate query.
		sql:    gfmq.sql.Clone(),
		path:   gfmq.path,
		unique: gfmq.unique,
	}
}

// WithGinFileMiddlewareToProvisionedHost tells the query-builder to eager-load the nodes that are connected to
// the "GinFileMiddlewareToProvisionedHost" edge. The optional arguments are used to configure the query builder of the edge.
func (gfmq *GinFileMiddlewareQuery) WithGinFileMiddlewareToProvisionedHost(opts ...func(*ProvisionedHostQuery)) *GinFileMiddlewareQuery {
	query := &ProvisionedHostQuery{config: gfmq.config}
	for _, opt := range opts {
		opt(query)
	}
	gfmq.withGinFileMiddlewareToProvisionedHost = query
	return gfmq
}

// WithGinFileMiddlewareToProvisioningStep tells the query-builder to eager-load the nodes that are connected to
// the "GinFileMiddlewareToProvisioningStep" edge. The optional arguments are used to configure the query builder of the edge.
func (gfmq *GinFileMiddlewareQuery) WithGinFileMiddlewareToProvisioningStep(opts ...func(*ProvisioningStepQuery)) *GinFileMiddlewareQuery {
	query := &ProvisioningStepQuery{config: gfmq.config}
	for _, opt := range opts {
		opt(query)
	}
	gfmq.withGinFileMiddlewareToProvisioningStep = query
	return gfmq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		URLID string `json:"url_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GinFileMiddleware.Query().
//		GroupBy(ginfilemiddleware.FieldURLID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gfmq *GinFileMiddlewareQuery) GroupBy(field string, fields ...string) *GinFileMiddlewareGroupBy {
	group := &GinFileMiddlewareGroupBy{config: gfmq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gfmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gfmq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		URLID string `json:"url_id,omitempty"`
//	}
//
//	client.GinFileMiddleware.Query().
//		Select(ginfilemiddleware.FieldURLID).
//		Scan(ctx, &v)
//
func (gfmq *GinFileMiddlewareQuery) Select(fields ...string) *GinFileMiddlewareSelect {
	gfmq.fields = append(gfmq.fields, fields...)
	return &GinFileMiddlewareSelect{GinFileMiddlewareQuery: gfmq}
}

func (gfmq *GinFileMiddlewareQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gfmq.fields {
		if !ginfilemiddleware.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gfmq.path != nil {
		prev, err := gfmq.path(ctx)
		if err != nil {
			return err
		}
		gfmq.sql = prev
	}
	return nil
}

func (gfmq *GinFileMiddlewareQuery) sqlAll(ctx context.Context) ([]*GinFileMiddleware, error) {
	var (
		nodes       = []*GinFileMiddleware{}
		withFKs     = gfmq.withFKs
		_spec       = gfmq.querySpec()
		loadedTypes = [2]bool{
			gfmq.withGinFileMiddlewareToProvisionedHost != nil,
			gfmq.withGinFileMiddlewareToProvisioningStep != nil,
		}
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, ginfilemiddleware.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GinFileMiddleware{config: gfmq.config}
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
	if err := sqlgraph.QueryNodes(ctx, gfmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := gfmq.withGinFileMiddlewareToProvisionedHost; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*GinFileMiddleware)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.ProvisionedHost(func(s *sql.Selector) {
			s.Where(sql.InValues(ginfilemiddleware.GinFileMiddlewareToProvisionedHostColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.gin_file_middleware_gin_file_middleware_to_provisioned_host
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "gin_file_middleware_gin_file_middleware_to_provisioned_host" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "gin_file_middleware_gin_file_middleware_to_provisioned_host" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.GinFileMiddlewareToProvisionedHost = n
		}
	}

	if query := gfmq.withGinFileMiddlewareToProvisioningStep; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*GinFileMiddleware)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.ProvisioningStep(func(s *sql.Selector) {
			s.Where(sql.InValues(ginfilemiddleware.GinFileMiddlewareToProvisioningStepColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.gin_file_middleware_gin_file_middleware_to_provisioning_step
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "gin_file_middleware_gin_file_middleware_to_provisioning_step" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "gin_file_middleware_gin_file_middleware_to_provisioning_step" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.GinFileMiddlewareToProvisioningStep = n
		}
	}

	return nodes, nil
}

func (gfmq *GinFileMiddlewareQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gfmq.querySpec()
	_spec.Node.Columns = gfmq.fields
	if len(gfmq.fields) > 0 {
		_spec.Unique = gfmq.unique != nil && *gfmq.unique
	}
	return sqlgraph.CountNodes(ctx, gfmq.driver, _spec)
}

func (gfmq *GinFileMiddlewareQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gfmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gfmq *GinFileMiddlewareQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ginfilemiddleware.Table,
			Columns: ginfilemiddleware.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: ginfilemiddleware.FieldID,
			},
		},
		From:   gfmq.sql,
		Unique: true,
	}
	if unique := gfmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gfmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ginfilemiddleware.FieldID)
		for i := range fields {
			if fields[i] != ginfilemiddleware.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gfmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gfmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gfmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gfmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gfmq *GinFileMiddlewareQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gfmq.driver.Dialect())
	t1 := builder.Table(ginfilemiddleware.Table)
	columns := gfmq.fields
	if len(columns) == 0 {
		columns = ginfilemiddleware.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gfmq.sql != nil {
		selector = gfmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gfmq.unique != nil && *gfmq.unique {
		selector.Distinct()
	}
	for _, p := range gfmq.predicates {
		p(selector)
	}
	for _, p := range gfmq.order {
		p(selector)
	}
	if offset := gfmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gfmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GinFileMiddlewareGroupBy is the group-by builder for GinFileMiddleware entities.
type GinFileMiddlewareGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gfmgb *GinFileMiddlewareGroupBy) Aggregate(fns ...AggregateFunc) *GinFileMiddlewareGroupBy {
	gfmgb.fns = append(gfmgb.fns, fns...)
	return gfmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gfmgb *GinFileMiddlewareGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gfmgb.path(ctx)
	if err != nil {
		return err
	}
	gfmgb.sql = query
	return gfmgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gfmgb *GinFileMiddlewareGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gfmgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gfmgb *GinFileMiddlewareGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gfmgb.fields) > 1 {
		return nil, errors.New("ent: GinFileMiddlewareGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gfmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gfmgb *GinFileMiddlewareGroupBy) StringsX(ctx context.Context) []string {
	v, err := gfmgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gfmgb *GinFileMiddlewareGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gfmgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ginfilemiddleware.Label}
	default:
		err = fmt.Errorf("ent: GinFileMiddlewareGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gfmgb *GinFileMiddlewareGroupBy) StringX(ctx context.Context) string {
	v, err := gfmgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gfmgb *GinFileMiddlewareGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gfmgb.fields) > 1 {
		return nil, errors.New("ent: GinFileMiddlewareGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gfmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gfmgb *GinFileMiddlewareGroupBy) IntsX(ctx context.Context) []int {
	v, err := gfmgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gfmgb *GinFileMiddlewareGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gfmgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ginfilemiddleware.Label}
	default:
		err = fmt.Errorf("ent: GinFileMiddlewareGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gfmgb *GinFileMiddlewareGroupBy) IntX(ctx context.Context) int {
	v, err := gfmgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gfmgb *GinFileMiddlewareGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gfmgb.fields) > 1 {
		return nil, errors.New("ent: GinFileMiddlewareGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gfmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gfmgb *GinFileMiddlewareGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gfmgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gfmgb *GinFileMiddlewareGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gfmgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ginfilemiddleware.Label}
	default:
		err = fmt.Errorf("ent: GinFileMiddlewareGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gfmgb *GinFileMiddlewareGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gfmgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gfmgb *GinFileMiddlewareGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gfmgb.fields) > 1 {
		return nil, errors.New("ent: GinFileMiddlewareGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gfmgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gfmgb *GinFileMiddlewareGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gfmgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gfmgb *GinFileMiddlewareGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gfmgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ginfilemiddleware.Label}
	default:
		err = fmt.Errorf("ent: GinFileMiddlewareGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gfmgb *GinFileMiddlewareGroupBy) BoolX(ctx context.Context) bool {
	v, err := gfmgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gfmgb *GinFileMiddlewareGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gfmgb.fields {
		if !ginfilemiddleware.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gfmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gfmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gfmgb *GinFileMiddlewareGroupBy) sqlQuery() *sql.Selector {
	selector := gfmgb.sql.Select()
	aggregation := make([]string, 0, len(gfmgb.fns))
	for _, fn := range gfmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gfmgb.fields)+len(gfmgb.fns))
		for _, f := range gfmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gfmgb.fields...)...)
}

// GinFileMiddlewareSelect is the builder for selecting fields of GinFileMiddleware entities.
type GinFileMiddlewareSelect struct {
	*GinFileMiddlewareQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gfms *GinFileMiddlewareSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gfms.prepareQuery(ctx); err != nil {
		return err
	}
	gfms.sql = gfms.GinFileMiddlewareQuery.sqlQuery(ctx)
	return gfms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gfms *GinFileMiddlewareSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gfms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gfms *GinFileMiddlewareSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gfms.fields) > 1 {
		return nil, errors.New("ent: GinFileMiddlewareSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gfms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gfms *GinFileMiddlewareSelect) StringsX(ctx context.Context) []string {
	v, err := gfms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gfms *GinFileMiddlewareSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gfms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ginfilemiddleware.Label}
	default:
		err = fmt.Errorf("ent: GinFileMiddlewareSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gfms *GinFileMiddlewareSelect) StringX(ctx context.Context) string {
	v, err := gfms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gfms *GinFileMiddlewareSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gfms.fields) > 1 {
		return nil, errors.New("ent: GinFileMiddlewareSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gfms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gfms *GinFileMiddlewareSelect) IntsX(ctx context.Context) []int {
	v, err := gfms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gfms *GinFileMiddlewareSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gfms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ginfilemiddleware.Label}
	default:
		err = fmt.Errorf("ent: GinFileMiddlewareSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gfms *GinFileMiddlewareSelect) IntX(ctx context.Context) int {
	v, err := gfms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gfms *GinFileMiddlewareSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gfms.fields) > 1 {
		return nil, errors.New("ent: GinFileMiddlewareSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gfms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gfms *GinFileMiddlewareSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gfms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gfms *GinFileMiddlewareSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gfms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ginfilemiddleware.Label}
	default:
		err = fmt.Errorf("ent: GinFileMiddlewareSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gfms *GinFileMiddlewareSelect) Float64X(ctx context.Context) float64 {
	v, err := gfms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gfms *GinFileMiddlewareSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gfms.fields) > 1 {
		return nil, errors.New("ent: GinFileMiddlewareSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gfms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gfms *GinFileMiddlewareSelect) BoolsX(ctx context.Context) []bool {
	v, err := gfms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gfms *GinFileMiddlewareSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gfms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{ginfilemiddleware.Label}
	default:
		err = fmt.Errorf("ent: GinFileMiddlewareSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gfms *GinFileMiddlewareSelect) BoolX(ctx context.Context) bool {
	v, err := gfms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gfms *GinFileMiddlewareSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gfms.sql.Query()
	if err := gfms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
