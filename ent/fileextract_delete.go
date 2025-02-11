// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/fileextract"
	"github.com/gen0cide/laforge/ent/predicate"
)

// FileExtractDelete is the builder for deleting a FileExtract entity.
type FileExtractDelete struct {
	config
	hooks    []Hook
	mutation *FileExtractMutation
}

// Where appends a list predicates to the FileExtractDelete builder.
func (fed *FileExtractDelete) Where(ps ...predicate.FileExtract) *FileExtractDelete {
	fed.mutation.Where(ps...)
	return fed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fed *FileExtractDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fed.hooks) == 0 {
		affected, err = fed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileExtractMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fed.mutation = mutation
			affected, err = fed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fed.hooks) - 1; i >= 0; i-- {
			if fed.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fed *FileExtractDelete) ExecX(ctx context.Context) int {
	n, err := fed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fed *FileExtractDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: fileextract.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: fileextract.FieldID,
			},
		},
	}
	if ps := fed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// FileExtractDeleteOne is the builder for deleting a single FileExtract entity.
type FileExtractDeleteOne struct {
	fed *FileExtractDelete
}

// Exec executes the deletion query.
func (fedo *FileExtractDeleteOne) Exec(ctx context.Context) error {
	n, err := fedo.fed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fileextract.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fedo *FileExtractDeleteOne) ExecX(ctx context.Context) {
	fedo.fed.ExecX(ctx)
}
