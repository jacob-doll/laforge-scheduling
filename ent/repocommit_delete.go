// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/repocommit"
)

// RepoCommitDelete is the builder for deleting a RepoCommit entity.
type RepoCommitDelete struct {
	config
	hooks    []Hook
	mutation *RepoCommitMutation
}

// Where appends a list predicates to the RepoCommitDelete builder.
func (rcd *RepoCommitDelete) Where(ps ...predicate.RepoCommit) *RepoCommitDelete {
	rcd.mutation.Where(ps...)
	return rcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rcd *RepoCommitDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rcd.hooks) == 0 {
		affected, err = rcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RepoCommitMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rcd.mutation = mutation
			affected, err = rcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rcd.hooks) - 1; i >= 0; i-- {
			if rcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcd *RepoCommitDelete) ExecX(ctx context.Context) int {
	n, err := rcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rcd *RepoCommitDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: repocommit.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: repocommit.FieldID,
			},
		},
	}
	if ps := rcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// RepoCommitDeleteOne is the builder for deleting a single RepoCommit entity.
type RepoCommitDeleteOne struct {
	rcd *RepoCommitDelete
}

// Exec executes the deletion query.
func (rcdo *RepoCommitDeleteOne) Exec(ctx context.Context) error {
	n, err := rcdo.rcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{repocommit.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rcdo *RepoCommitDeleteOne) ExecX(ctx context.Context) {
	rcdo.rcd.ExecX(ctx)
}
