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
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/repocommit"
	"github.com/gen0cide/laforge/ent/repository"
	"github.com/google/uuid"
)

// RepositoryQuery is the builder for querying Repository entities.
type RepositoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Repository
	// eager-loading edges.
	withRepositoryToEnvironment *EnvironmentQuery
	withRepositoryToRepoCommit  *RepoCommitQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RepositoryQuery builder.
func (rq *RepositoryQuery) Where(ps ...predicate.Repository) *RepositoryQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *RepositoryQuery) Limit(limit int) *RepositoryQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *RepositoryQuery) Offset(offset int) *RepositoryQuery {
	rq.offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RepositoryQuery) Unique(unique bool) *RepositoryQuery {
	rq.unique = &unique
	return rq
}

// Order adds an order step to the query.
func (rq *RepositoryQuery) Order(o ...OrderFunc) *RepositoryQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryRepositoryToEnvironment chains the current query on the "RepositoryToEnvironment" edge.
func (rq *RepositoryQuery) QueryRepositoryToEnvironment() *EnvironmentQuery {
	query := &EnvironmentQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(repository.Table, repository.FieldID, selector),
			sqlgraph.To(environment.Table, environment.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, repository.RepositoryToEnvironmentTable, repository.RepositoryToEnvironmentPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRepositoryToRepoCommit chains the current query on the "RepositoryToRepoCommit" edge.
func (rq *RepositoryQuery) QueryRepositoryToRepoCommit() *RepoCommitQuery {
	query := &RepoCommitQuery{config: rq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(repository.Table, repository.FieldID, selector),
			sqlgraph.To(repocommit.Table, repocommit.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, repository.RepositoryToRepoCommitTable, repository.RepositoryToRepoCommitColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Repository entity from the query.
// Returns a *NotFoundError when no Repository was found.
func (rq *RepositoryQuery) First(ctx context.Context) (*Repository, error) {
	nodes, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{repository.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RepositoryQuery) FirstX(ctx context.Context) *Repository {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Repository ID from the query.
// Returns a *NotFoundError when no Repository ID was found.
func (rq *RepositoryQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{repository.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RepositoryQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Repository entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Repository entity is found.
// Returns a *NotFoundError when no Repository entities are found.
func (rq *RepositoryQuery) Only(ctx context.Context) (*Repository, error) {
	nodes, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{repository.Label}
	default:
		return nil, &NotSingularError{repository.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RepositoryQuery) OnlyX(ctx context.Context) *Repository {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Repository ID in the query.
// Returns a *NotSingularError when more than one Repository ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RepositoryQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{repository.Label}
	default:
		err = &NotSingularError{repository.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RepositoryQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Repositories.
func (rq *RepositoryQuery) All(ctx context.Context) ([]*Repository, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *RepositoryQuery) AllX(ctx context.Context) []*Repository {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Repository IDs.
func (rq *RepositoryQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := rq.Select(repository.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RepositoryQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RepositoryQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RepositoryQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RepositoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RepositoryQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RepositoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RepositoryQuery) Clone() *RepositoryQuery {
	if rq == nil {
		return nil
	}
	return &RepositoryQuery{
		config:                      rq.config,
		limit:                       rq.limit,
		offset:                      rq.offset,
		order:                       append([]OrderFunc{}, rq.order...),
		predicates:                  append([]predicate.Repository{}, rq.predicates...),
		withRepositoryToEnvironment: rq.withRepositoryToEnvironment.Clone(),
		withRepositoryToRepoCommit:  rq.withRepositoryToRepoCommit.Clone(),
		// clone intermediate query.
		sql:    rq.sql.Clone(),
		path:   rq.path,
		unique: rq.unique,
	}
}

// WithRepositoryToEnvironment tells the query-builder to eager-load the nodes that are connected to
// the "RepositoryToEnvironment" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RepositoryQuery) WithRepositoryToEnvironment(opts ...func(*EnvironmentQuery)) *RepositoryQuery {
	query := &EnvironmentQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withRepositoryToEnvironment = query
	return rq
}

// WithRepositoryToRepoCommit tells the query-builder to eager-load the nodes that are connected to
// the "RepositoryToRepoCommit" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RepositoryQuery) WithRepositoryToRepoCommit(opts ...func(*RepoCommitQuery)) *RepositoryQuery {
	query := &RepoCommitQuery{config: rq.config}
	for _, opt := range opts {
		opt(query)
	}
	rq.withRepositoryToRepoCommit = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RepoURL string `json:"repo_url,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Repository.Query().
//		GroupBy(repository.FieldRepoURL).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *RepositoryQuery) GroupBy(field string, fields ...string) *RepositoryGroupBy {
	group := &RepositoryGroupBy{config: rq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RepoURL string `json:"repo_url,omitempty"`
//	}
//
//	client.Repository.Query().
//		Select(repository.FieldRepoURL).
//		Scan(ctx, &v)
//
func (rq *RepositoryQuery) Select(fields ...string) *RepositorySelect {
	rq.fields = append(rq.fields, fields...)
	return &RepositorySelect{RepositoryQuery: rq}
}

func (rq *RepositoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rq.fields {
		if !repository.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RepositoryQuery) sqlAll(ctx context.Context) ([]*Repository, error) {
	var (
		nodes       = []*Repository{}
		_spec       = rq.querySpec()
		loadedTypes = [2]bool{
			rq.withRepositoryToEnvironment != nil,
			rq.withRepositoryToRepoCommit != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Repository{config: rq.config}
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
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := rq.withRepositoryToEnvironment; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[uuid.UUID]*Repository, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.RepositoryToEnvironment = []*Environment{}
		}
		var (
			edgeids []uuid.UUID
			edges   = make(map[uuid.UUID][]*Repository)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   repository.RepositoryToEnvironmentTable,
				Columns: repository.RepositoryToEnvironmentPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(repository.RepositoryToEnvironmentPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(uuid.UUID), new(uuid.UUID)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*uuid.UUID)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*uuid.UUID)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := *eout
				inValue := *ein
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, rq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "RepositoryToEnvironment": %w`, err)
		}
		query.Where(environment.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "RepositoryToEnvironment" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.RepositoryToEnvironment = append(nodes[i].Edges.RepositoryToEnvironment, n)
			}
		}
	}

	if query := rq.withRepositoryToRepoCommit; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*Repository)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.RepositoryToRepoCommit = []*RepoCommit{}
		}
		query.withFKs = true
		query.Where(predicate.RepoCommit(func(s *sql.Selector) {
			s.Where(sql.InValues(repository.RepositoryToRepoCommitColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.repository_repository_to_repo_commit
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "repository_repository_to_repo_commit" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "repository_repository_to_repo_commit" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.RepositoryToRepoCommit = append(node.Edges.RepositoryToRepoCommit, n)
		}
	}

	return nodes, nil
}

func (rq *RepositoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.fields
	if len(rq.fields) > 0 {
		_spec.Unique = rq.unique != nil && *rq.unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RepositoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (rq *RepositoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   repository.Table,
			Columns: repository.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: repository.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if unique := rq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, repository.FieldID)
		for i := range fields {
			if fields[i] != repository.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RepositoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(repository.Table)
	columns := rq.fields
	if len(columns) == 0 {
		columns = repository.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.unique != nil && *rq.unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RepositoryGroupBy is the group-by builder for Repository entities.
type RepositoryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RepositoryGroupBy) Aggregate(fns ...AggregateFunc) *RepositoryGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rgb *RepositoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rgb *RepositoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *RepositoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RepositoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rgb *RepositoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := rgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *RepositoryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repository.Label}
	default:
		err = fmt.Errorf("ent: RepositoryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rgb *RepositoryGroupBy) StringX(ctx context.Context) string {
	v, err := rgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *RepositoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RepositoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rgb *RepositoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := rgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *RepositoryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repository.Label}
	default:
		err = fmt.Errorf("ent: RepositoryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rgb *RepositoryGroupBy) IntX(ctx context.Context) int {
	v, err := rgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *RepositoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RepositoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rgb *RepositoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *RepositoryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repository.Label}
	default:
		err = fmt.Errorf("ent: RepositoryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rgb *RepositoryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *RepositoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: RepositoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rgb *RepositoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *RepositoryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repository.Label}
	default:
		err = fmt.Errorf("ent: RepositoryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rgb *RepositoryGroupBy) BoolX(ctx context.Context) bool {
	v, err := rgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rgb *RepositoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rgb.fields {
		if !repository.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *RepositoryGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql.Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
		for _, f := range rgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rgb.fields...)...)
}

// RepositorySelect is the builder for selecting fields of Repository entities.
type RepositorySelect struct {
	*RepositoryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RepositorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	rs.sql = rs.RepositoryQuery.sqlQuery(ctx)
	return rs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rs *RepositorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := rs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rs *RepositorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RepositorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rs *RepositorySelect) StringsX(ctx context.Context) []string {
	v, err := rs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rs *RepositorySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repository.Label}
	default:
		err = fmt.Errorf("ent: RepositorySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rs *RepositorySelect) StringX(ctx context.Context) string {
	v, err := rs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rs *RepositorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RepositorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rs *RepositorySelect) IntsX(ctx context.Context) []int {
	v, err := rs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rs *RepositorySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repository.Label}
	default:
		err = fmt.Errorf("ent: RepositorySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rs *RepositorySelect) IntX(ctx context.Context) int {
	v, err := rs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rs *RepositorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RepositorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rs *RepositorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := rs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rs *RepositorySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repository.Label}
	default:
		err = fmt.Errorf("ent: RepositorySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rs *RepositorySelect) Float64X(ctx context.Context) float64 {
	v, err := rs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rs *RepositorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: RepositorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rs *RepositorySelect) BoolsX(ctx context.Context) []bool {
	v, err := rs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rs *RepositorySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{repository.Label}
	default:
		err = fmt.Errorf("ent: RepositorySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rs *RepositorySelect) BoolX(ctx context.Context) bool {
	v, err := rs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rs *RepositorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rs.sql.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
