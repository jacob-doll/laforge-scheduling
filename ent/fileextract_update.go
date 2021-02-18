// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/fileextract"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
)

// FileExtractUpdate is the builder for updating FileExtract entities.
type FileExtractUpdate struct {
	config
	hooks    []Hook
	mutation *FileExtractMutation
}

// Where adds a new predicate for the FileExtractUpdate builder.
func (feu *FileExtractUpdate) Where(ps ...predicate.FileExtract) *FileExtractUpdate {
	feu.mutation.predicates = append(feu.mutation.predicates, ps...)
	return feu
}

// SetSource sets the "source" field.
func (feu *FileExtractUpdate) SetSource(s string) *FileExtractUpdate {
	feu.mutation.SetSource(s)
	return feu
}

// SetDestination sets the "destination" field.
func (feu *FileExtractUpdate) SetDestination(s string) *FileExtractUpdate {
	feu.mutation.SetDestination(s)
	return feu
}

// SetType sets the "type" field.
func (feu *FileExtractUpdate) SetType(s string) *FileExtractUpdate {
	feu.mutation.SetType(s)
	return feu
}

// SetTags sets the "tags" field.
func (feu *FileExtractUpdate) SetTags(m map[string]string) *FileExtractUpdate {
	feu.mutation.SetTags(m)
	return feu
}

// AddFileExtractToTagIDs adds the "FileExtractToTag" edge to the Tag entity by IDs.
func (feu *FileExtractUpdate) AddFileExtractToTagIDs(ids ...int) *FileExtractUpdate {
	feu.mutation.AddFileExtractToTagIDs(ids...)
	return feu
}

// AddFileExtractToTag adds the "FileExtractToTag" edges to the Tag entity.
func (feu *FileExtractUpdate) AddFileExtractToTag(t ...*Tag) *FileExtractUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return feu.AddFileExtractToTagIDs(ids...)
}

// Mutation returns the FileExtractMutation object of the builder.
func (feu *FileExtractUpdate) Mutation() *FileExtractMutation {
	return feu.mutation
}

// ClearFileExtractToTag clears all "FileExtractToTag" edges to the Tag entity.
func (feu *FileExtractUpdate) ClearFileExtractToTag() *FileExtractUpdate {
	feu.mutation.ClearFileExtractToTag()
	return feu
}

// RemoveFileExtractToTagIDs removes the "FileExtractToTag" edge to Tag entities by IDs.
func (feu *FileExtractUpdate) RemoveFileExtractToTagIDs(ids ...int) *FileExtractUpdate {
	feu.mutation.RemoveFileExtractToTagIDs(ids...)
	return feu
}

// RemoveFileExtractToTag removes "FileExtractToTag" edges to Tag entities.
func (feu *FileExtractUpdate) RemoveFileExtractToTag(t ...*Tag) *FileExtractUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return feu.RemoveFileExtractToTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (feu *FileExtractUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(feu.hooks) == 0 {
		affected, err = feu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileExtractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			feu.mutation = mutation
			affected, err = feu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(feu.hooks) - 1; i >= 0; i-- {
			mut = feu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, feu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (feu *FileExtractUpdate) SaveX(ctx context.Context) int {
	affected, err := feu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (feu *FileExtractUpdate) Exec(ctx context.Context) error {
	_, err := feu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (feu *FileExtractUpdate) ExecX(ctx context.Context) {
	if err := feu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (feu *FileExtractUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fileextract.Table,
			Columns: fileextract.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fileextract.FieldID,
			},
		},
	}
	if ps := feu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := feu.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileextract.FieldSource,
		})
	}
	if value, ok := feu.mutation.Destination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileextract.FieldDestination,
		})
	}
	if value, ok := feu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileextract.FieldType,
		})
	}
	if value, ok := feu.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: fileextract.FieldTags,
		})
	}
	if feu.mutation.FileExtractToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fileextract.FileExtractToTagTable,
			Columns: []string{fileextract.FileExtractToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := feu.mutation.RemovedFileExtractToTagIDs(); len(nodes) > 0 && !feu.mutation.FileExtractToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fileextract.FileExtractToTagTable,
			Columns: []string{fileextract.FileExtractToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := feu.mutation.FileExtractToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fileextract.FileExtractToTagTable,
			Columns: []string{fileextract.FileExtractToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, feu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fileextract.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FileExtractUpdateOne is the builder for updating a single FileExtract entity.
type FileExtractUpdateOne struct {
	config
	hooks    []Hook
	mutation *FileExtractMutation
}

// SetSource sets the "source" field.
func (feuo *FileExtractUpdateOne) SetSource(s string) *FileExtractUpdateOne {
	feuo.mutation.SetSource(s)
	return feuo
}

// SetDestination sets the "destination" field.
func (feuo *FileExtractUpdateOne) SetDestination(s string) *FileExtractUpdateOne {
	feuo.mutation.SetDestination(s)
	return feuo
}

// SetType sets the "type" field.
func (feuo *FileExtractUpdateOne) SetType(s string) *FileExtractUpdateOne {
	feuo.mutation.SetType(s)
	return feuo
}

// SetTags sets the "tags" field.
func (feuo *FileExtractUpdateOne) SetTags(m map[string]string) *FileExtractUpdateOne {
	feuo.mutation.SetTags(m)
	return feuo
}

// AddFileExtractToTagIDs adds the "FileExtractToTag" edge to the Tag entity by IDs.
func (feuo *FileExtractUpdateOne) AddFileExtractToTagIDs(ids ...int) *FileExtractUpdateOne {
	feuo.mutation.AddFileExtractToTagIDs(ids...)
	return feuo
}

// AddFileExtractToTag adds the "FileExtractToTag" edges to the Tag entity.
func (feuo *FileExtractUpdateOne) AddFileExtractToTag(t ...*Tag) *FileExtractUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return feuo.AddFileExtractToTagIDs(ids...)
}

// Mutation returns the FileExtractMutation object of the builder.
func (feuo *FileExtractUpdateOne) Mutation() *FileExtractMutation {
	return feuo.mutation
}

// ClearFileExtractToTag clears all "FileExtractToTag" edges to the Tag entity.
func (feuo *FileExtractUpdateOne) ClearFileExtractToTag() *FileExtractUpdateOne {
	feuo.mutation.ClearFileExtractToTag()
	return feuo
}

// RemoveFileExtractToTagIDs removes the "FileExtractToTag" edge to Tag entities by IDs.
func (feuo *FileExtractUpdateOne) RemoveFileExtractToTagIDs(ids ...int) *FileExtractUpdateOne {
	feuo.mutation.RemoveFileExtractToTagIDs(ids...)
	return feuo
}

// RemoveFileExtractToTag removes "FileExtractToTag" edges to Tag entities.
func (feuo *FileExtractUpdateOne) RemoveFileExtractToTag(t ...*Tag) *FileExtractUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return feuo.RemoveFileExtractToTagIDs(ids...)
}

// Save executes the query and returns the updated FileExtract entity.
func (feuo *FileExtractUpdateOne) Save(ctx context.Context) (*FileExtract, error) {
	var (
		err  error
		node *FileExtract
	)
	if len(feuo.hooks) == 0 {
		node, err = feuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileExtractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			feuo.mutation = mutation
			node, err = feuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(feuo.hooks) - 1; i >= 0; i-- {
			mut = feuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, feuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (feuo *FileExtractUpdateOne) SaveX(ctx context.Context) *FileExtract {
	node, err := feuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (feuo *FileExtractUpdateOne) Exec(ctx context.Context) error {
	_, err := feuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (feuo *FileExtractUpdateOne) ExecX(ctx context.Context) {
	if err := feuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (feuo *FileExtractUpdateOne) sqlSave(ctx context.Context) (_node *FileExtract, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   fileextract.Table,
			Columns: fileextract.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fileextract.FieldID,
			},
		},
	}
	id, ok := feuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing FileExtract.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := feuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := feuo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileextract.FieldSource,
		})
	}
	if value, ok := feuo.mutation.Destination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileextract.FieldDestination,
		})
	}
	if value, ok := feuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileextract.FieldType,
		})
	}
	if value, ok := feuo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: fileextract.FieldTags,
		})
	}
	if feuo.mutation.FileExtractToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fileextract.FileExtractToTagTable,
			Columns: []string{fileextract.FileExtractToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := feuo.mutation.RemovedFileExtractToTagIDs(); len(nodes) > 0 && !feuo.mutation.FileExtractToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fileextract.FileExtractToTagTable,
			Columns: []string{fileextract.FileExtractToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := feuo.mutation.FileExtractToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fileextract.FileExtractToTagTable,
			Columns: []string{fileextract.FileExtractToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FileExtract{config: feuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, feuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fileextract.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
