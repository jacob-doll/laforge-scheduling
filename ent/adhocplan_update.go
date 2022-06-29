// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/adhocplan"
	"github.com/gen0cide/laforge/ent/agenttask"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/google/uuid"
)

// AdhocPlanUpdate is the builder for updating AdhocPlan entities.
type AdhocPlanUpdate struct {
	config
	hooks    []Hook
	mutation *AdhocPlanMutation
}

// Where appends a list predicates to the AdhocPlanUpdate builder.
func (apu *AdhocPlanUpdate) Where(ps ...predicate.AdhocPlan) *AdhocPlanUpdate {
	apu.mutation.Where(ps...)
	return apu
}

// AddPrevAdhocPlanIDs adds the "PrevAdhocPlan" edge to the AdhocPlan entity by IDs.
func (apu *AdhocPlanUpdate) AddPrevAdhocPlanIDs(ids ...uuid.UUID) *AdhocPlanUpdate {
	apu.mutation.AddPrevAdhocPlanIDs(ids...)
	return apu
}

// AddPrevAdhocPlan adds the "PrevAdhocPlan" edges to the AdhocPlan entity.
func (apu *AdhocPlanUpdate) AddPrevAdhocPlan(a ...*AdhocPlan) *AdhocPlanUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apu.AddPrevAdhocPlanIDs(ids...)
}

// AddNextAdhocPlanIDs adds the "NextAdhocPlan" edge to the AdhocPlan entity by IDs.
func (apu *AdhocPlanUpdate) AddNextAdhocPlanIDs(ids ...uuid.UUID) *AdhocPlanUpdate {
	apu.mutation.AddNextAdhocPlanIDs(ids...)
	return apu
}

// AddNextAdhocPlan adds the "NextAdhocPlan" edges to the AdhocPlan entity.
func (apu *AdhocPlanUpdate) AddNextAdhocPlan(a ...*AdhocPlan) *AdhocPlanUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apu.AddNextAdhocPlanIDs(ids...)
}

// SetAdhocPlanToBuildID sets the "AdhocPlanToBuild" edge to the Build entity by ID.
func (apu *AdhocPlanUpdate) SetAdhocPlanToBuildID(id uuid.UUID) *AdhocPlanUpdate {
	apu.mutation.SetAdhocPlanToBuildID(id)
	return apu
}

// SetAdhocPlanToBuild sets the "AdhocPlanToBuild" edge to the Build entity.
func (apu *AdhocPlanUpdate) SetAdhocPlanToBuild(b *Build) *AdhocPlanUpdate {
	return apu.SetAdhocPlanToBuildID(b.ID)
}

// SetAdhocPlanToStatusID sets the "AdhocPlanToStatus" edge to the Status entity by ID.
func (apu *AdhocPlanUpdate) SetAdhocPlanToStatusID(id uuid.UUID) *AdhocPlanUpdate {
	apu.mutation.SetAdhocPlanToStatusID(id)
	return apu
}

// SetAdhocPlanToStatus sets the "AdhocPlanToStatus" edge to the Status entity.
func (apu *AdhocPlanUpdate) SetAdhocPlanToStatus(s *Status) *AdhocPlanUpdate {
	return apu.SetAdhocPlanToStatusID(s.ID)
}

// SetAdhocPlanToAgentTaskID sets the "AdhocPlanToAgentTask" edge to the AgentTask entity by ID.
func (apu *AdhocPlanUpdate) SetAdhocPlanToAgentTaskID(id uuid.UUID) *AdhocPlanUpdate {
	apu.mutation.SetAdhocPlanToAgentTaskID(id)
	return apu
}

// SetAdhocPlanToAgentTask sets the "AdhocPlanToAgentTask" edge to the AgentTask entity.
func (apu *AdhocPlanUpdate) SetAdhocPlanToAgentTask(a *AgentTask) *AdhocPlanUpdate {
	return apu.SetAdhocPlanToAgentTaskID(a.ID)
}

// Mutation returns the AdhocPlanMutation object of the builder.
func (apu *AdhocPlanUpdate) Mutation() *AdhocPlanMutation {
	return apu.mutation
}

// ClearPrevAdhocPlan clears all "PrevAdhocPlan" edges to the AdhocPlan entity.
func (apu *AdhocPlanUpdate) ClearPrevAdhocPlan() *AdhocPlanUpdate {
	apu.mutation.ClearPrevAdhocPlan()
	return apu
}

// RemovePrevAdhocPlanIDs removes the "PrevAdhocPlan" edge to AdhocPlan entities by IDs.
func (apu *AdhocPlanUpdate) RemovePrevAdhocPlanIDs(ids ...uuid.UUID) *AdhocPlanUpdate {
	apu.mutation.RemovePrevAdhocPlanIDs(ids...)
	return apu
}

// RemovePrevAdhocPlan removes "PrevAdhocPlan" edges to AdhocPlan entities.
func (apu *AdhocPlanUpdate) RemovePrevAdhocPlan(a ...*AdhocPlan) *AdhocPlanUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apu.RemovePrevAdhocPlanIDs(ids...)
}

// ClearNextAdhocPlan clears all "NextAdhocPlan" edges to the AdhocPlan entity.
func (apu *AdhocPlanUpdate) ClearNextAdhocPlan() *AdhocPlanUpdate {
	apu.mutation.ClearNextAdhocPlan()
	return apu
}

// RemoveNextAdhocPlanIDs removes the "NextAdhocPlan" edge to AdhocPlan entities by IDs.
func (apu *AdhocPlanUpdate) RemoveNextAdhocPlanIDs(ids ...uuid.UUID) *AdhocPlanUpdate {
	apu.mutation.RemoveNextAdhocPlanIDs(ids...)
	return apu
}

// RemoveNextAdhocPlan removes "NextAdhocPlan" edges to AdhocPlan entities.
func (apu *AdhocPlanUpdate) RemoveNextAdhocPlan(a ...*AdhocPlan) *AdhocPlanUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apu.RemoveNextAdhocPlanIDs(ids...)
}

// ClearAdhocPlanToBuild clears the "AdhocPlanToBuild" edge to the Build entity.
func (apu *AdhocPlanUpdate) ClearAdhocPlanToBuild() *AdhocPlanUpdate {
	apu.mutation.ClearAdhocPlanToBuild()
	return apu
}

// ClearAdhocPlanToStatus clears the "AdhocPlanToStatus" edge to the Status entity.
func (apu *AdhocPlanUpdate) ClearAdhocPlanToStatus() *AdhocPlanUpdate {
	apu.mutation.ClearAdhocPlanToStatus()
	return apu
}

// ClearAdhocPlanToAgentTask clears the "AdhocPlanToAgentTask" edge to the AgentTask entity.
func (apu *AdhocPlanUpdate) ClearAdhocPlanToAgentTask() *AdhocPlanUpdate {
	apu.mutation.ClearAdhocPlanToAgentTask()
	return apu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (apu *AdhocPlanUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(apu.hooks) == 0 {
		if err = apu.check(); err != nil {
			return 0, err
		}
		affected, err = apu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdhocPlanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = apu.check(); err != nil {
				return 0, err
			}
			apu.mutation = mutation
			affected, err = apu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(apu.hooks) - 1; i >= 0; i-- {
			if apu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = apu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, apu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (apu *AdhocPlanUpdate) SaveX(ctx context.Context) int {
	affected, err := apu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (apu *AdhocPlanUpdate) Exec(ctx context.Context) error {
	_, err := apu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apu *AdhocPlanUpdate) ExecX(ctx context.Context) {
	if err := apu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apu *AdhocPlanUpdate) check() error {
	if _, ok := apu.mutation.AdhocPlanToBuildID(); apu.mutation.AdhocPlanToBuildCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AdhocPlan.AdhocPlanToBuild"`)
	}
	if _, ok := apu.mutation.AdhocPlanToStatusID(); apu.mutation.AdhocPlanToStatusCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AdhocPlan.AdhocPlanToStatus"`)
	}
	if _, ok := apu.mutation.AdhocPlanToAgentTaskID(); apu.mutation.AdhocPlanToAgentTaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AdhocPlan.AdhocPlanToAgentTask"`)
	}
	return nil
}

func (apu *AdhocPlanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adhocplan.Table,
			Columns: adhocplan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: adhocplan.FieldID,
			},
		},
	}
	if ps := apu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if apu.mutation.PrevAdhocPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   adhocplan.PrevAdhocPlanTable,
			Columns: adhocplan.PrevAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.RemovedPrevAdhocPlanIDs(); len(nodes) > 0 && !apu.mutation.PrevAdhocPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   adhocplan.PrevAdhocPlanTable,
			Columns: adhocplan.PrevAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.PrevAdhocPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   adhocplan.PrevAdhocPlanTable,
			Columns: adhocplan.PrevAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apu.mutation.NextAdhocPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adhocplan.NextAdhocPlanTable,
			Columns: adhocplan.NextAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.RemovedNextAdhocPlanIDs(); len(nodes) > 0 && !apu.mutation.NextAdhocPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adhocplan.NextAdhocPlanTable,
			Columns: adhocplan.NextAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.NextAdhocPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adhocplan.NextAdhocPlanTable,
			Columns: adhocplan.NextAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apu.mutation.AdhocPlanToBuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToBuildTable,
			Columns: []string{adhocplan.AdhocPlanToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.AdhocPlanToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToBuildTable,
			Columns: []string{adhocplan.AdhocPlanToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apu.mutation.AdhocPlanToStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToStatusTable,
			Columns: []string{adhocplan.AdhocPlanToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.AdhocPlanToStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToStatusTable,
			Columns: []string{adhocplan.AdhocPlanToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apu.mutation.AdhocPlanToAgentTaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToAgentTaskTable,
			Columns: []string{adhocplan.AdhocPlanToAgentTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: agenttask.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apu.mutation.AdhocPlanToAgentTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToAgentTaskTable,
			Columns: []string{adhocplan.AdhocPlanToAgentTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: agenttask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, apu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adhocplan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AdhocPlanUpdateOne is the builder for updating a single AdhocPlan entity.
type AdhocPlanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdhocPlanMutation
}

// AddPrevAdhocPlanIDs adds the "PrevAdhocPlan" edge to the AdhocPlan entity by IDs.
func (apuo *AdhocPlanUpdateOne) AddPrevAdhocPlanIDs(ids ...uuid.UUID) *AdhocPlanUpdateOne {
	apuo.mutation.AddPrevAdhocPlanIDs(ids...)
	return apuo
}

// AddPrevAdhocPlan adds the "PrevAdhocPlan" edges to the AdhocPlan entity.
func (apuo *AdhocPlanUpdateOne) AddPrevAdhocPlan(a ...*AdhocPlan) *AdhocPlanUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apuo.AddPrevAdhocPlanIDs(ids...)
}

// AddNextAdhocPlanIDs adds the "NextAdhocPlan" edge to the AdhocPlan entity by IDs.
func (apuo *AdhocPlanUpdateOne) AddNextAdhocPlanIDs(ids ...uuid.UUID) *AdhocPlanUpdateOne {
	apuo.mutation.AddNextAdhocPlanIDs(ids...)
	return apuo
}

// AddNextAdhocPlan adds the "NextAdhocPlan" edges to the AdhocPlan entity.
func (apuo *AdhocPlanUpdateOne) AddNextAdhocPlan(a ...*AdhocPlan) *AdhocPlanUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apuo.AddNextAdhocPlanIDs(ids...)
}

// SetAdhocPlanToBuildID sets the "AdhocPlanToBuild" edge to the Build entity by ID.
func (apuo *AdhocPlanUpdateOne) SetAdhocPlanToBuildID(id uuid.UUID) *AdhocPlanUpdateOne {
	apuo.mutation.SetAdhocPlanToBuildID(id)
	return apuo
}

// SetAdhocPlanToBuild sets the "AdhocPlanToBuild" edge to the Build entity.
func (apuo *AdhocPlanUpdateOne) SetAdhocPlanToBuild(b *Build) *AdhocPlanUpdateOne {
	return apuo.SetAdhocPlanToBuildID(b.ID)
}

// SetAdhocPlanToStatusID sets the "AdhocPlanToStatus" edge to the Status entity by ID.
func (apuo *AdhocPlanUpdateOne) SetAdhocPlanToStatusID(id uuid.UUID) *AdhocPlanUpdateOne {
	apuo.mutation.SetAdhocPlanToStatusID(id)
	return apuo
}

// SetAdhocPlanToStatus sets the "AdhocPlanToStatus" edge to the Status entity.
func (apuo *AdhocPlanUpdateOne) SetAdhocPlanToStatus(s *Status) *AdhocPlanUpdateOne {
	return apuo.SetAdhocPlanToStatusID(s.ID)
}

// SetAdhocPlanToAgentTaskID sets the "AdhocPlanToAgentTask" edge to the AgentTask entity by ID.
func (apuo *AdhocPlanUpdateOne) SetAdhocPlanToAgentTaskID(id uuid.UUID) *AdhocPlanUpdateOne {
	apuo.mutation.SetAdhocPlanToAgentTaskID(id)
	return apuo
}

// SetAdhocPlanToAgentTask sets the "AdhocPlanToAgentTask" edge to the AgentTask entity.
func (apuo *AdhocPlanUpdateOne) SetAdhocPlanToAgentTask(a *AgentTask) *AdhocPlanUpdateOne {
	return apuo.SetAdhocPlanToAgentTaskID(a.ID)
}

// Mutation returns the AdhocPlanMutation object of the builder.
func (apuo *AdhocPlanUpdateOne) Mutation() *AdhocPlanMutation {
	return apuo.mutation
}

// ClearPrevAdhocPlan clears all "PrevAdhocPlan" edges to the AdhocPlan entity.
func (apuo *AdhocPlanUpdateOne) ClearPrevAdhocPlan() *AdhocPlanUpdateOne {
	apuo.mutation.ClearPrevAdhocPlan()
	return apuo
}

// RemovePrevAdhocPlanIDs removes the "PrevAdhocPlan" edge to AdhocPlan entities by IDs.
func (apuo *AdhocPlanUpdateOne) RemovePrevAdhocPlanIDs(ids ...uuid.UUID) *AdhocPlanUpdateOne {
	apuo.mutation.RemovePrevAdhocPlanIDs(ids...)
	return apuo
}

// RemovePrevAdhocPlan removes "PrevAdhocPlan" edges to AdhocPlan entities.
func (apuo *AdhocPlanUpdateOne) RemovePrevAdhocPlan(a ...*AdhocPlan) *AdhocPlanUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apuo.RemovePrevAdhocPlanIDs(ids...)
}

// ClearNextAdhocPlan clears all "NextAdhocPlan" edges to the AdhocPlan entity.
func (apuo *AdhocPlanUpdateOne) ClearNextAdhocPlan() *AdhocPlanUpdateOne {
	apuo.mutation.ClearNextAdhocPlan()
	return apuo
}

// RemoveNextAdhocPlanIDs removes the "NextAdhocPlan" edge to AdhocPlan entities by IDs.
func (apuo *AdhocPlanUpdateOne) RemoveNextAdhocPlanIDs(ids ...uuid.UUID) *AdhocPlanUpdateOne {
	apuo.mutation.RemoveNextAdhocPlanIDs(ids...)
	return apuo
}

// RemoveNextAdhocPlan removes "NextAdhocPlan" edges to AdhocPlan entities.
func (apuo *AdhocPlanUpdateOne) RemoveNextAdhocPlan(a ...*AdhocPlan) *AdhocPlanUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return apuo.RemoveNextAdhocPlanIDs(ids...)
}

// ClearAdhocPlanToBuild clears the "AdhocPlanToBuild" edge to the Build entity.
func (apuo *AdhocPlanUpdateOne) ClearAdhocPlanToBuild() *AdhocPlanUpdateOne {
	apuo.mutation.ClearAdhocPlanToBuild()
	return apuo
}

// ClearAdhocPlanToStatus clears the "AdhocPlanToStatus" edge to the Status entity.
func (apuo *AdhocPlanUpdateOne) ClearAdhocPlanToStatus() *AdhocPlanUpdateOne {
	apuo.mutation.ClearAdhocPlanToStatus()
	return apuo
}

// ClearAdhocPlanToAgentTask clears the "AdhocPlanToAgentTask" edge to the AgentTask entity.
func (apuo *AdhocPlanUpdateOne) ClearAdhocPlanToAgentTask() *AdhocPlanUpdateOne {
	apuo.mutation.ClearAdhocPlanToAgentTask()
	return apuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (apuo *AdhocPlanUpdateOne) Select(field string, fields ...string) *AdhocPlanUpdateOne {
	apuo.fields = append([]string{field}, fields...)
	return apuo
}

// Save executes the query and returns the updated AdhocPlan entity.
func (apuo *AdhocPlanUpdateOne) Save(ctx context.Context) (*AdhocPlan, error) {
	var (
		err  error
		node *AdhocPlan
	)
	if len(apuo.hooks) == 0 {
		if err = apuo.check(); err != nil {
			return nil, err
		}
		node, err = apuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdhocPlanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = apuo.check(); err != nil {
				return nil, err
			}
			apuo.mutation = mutation
			node, err = apuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(apuo.hooks) - 1; i >= 0; i-- {
			if apuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = apuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, apuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (apuo *AdhocPlanUpdateOne) SaveX(ctx context.Context) *AdhocPlan {
	node, err := apuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (apuo *AdhocPlanUpdateOne) Exec(ctx context.Context) error {
	_, err := apuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (apuo *AdhocPlanUpdateOne) ExecX(ctx context.Context) {
	if err := apuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (apuo *AdhocPlanUpdateOne) check() error {
	if _, ok := apuo.mutation.AdhocPlanToBuildID(); apuo.mutation.AdhocPlanToBuildCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AdhocPlan.AdhocPlanToBuild"`)
	}
	if _, ok := apuo.mutation.AdhocPlanToStatusID(); apuo.mutation.AdhocPlanToStatusCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AdhocPlan.AdhocPlanToStatus"`)
	}
	if _, ok := apuo.mutation.AdhocPlanToAgentTaskID(); apuo.mutation.AdhocPlanToAgentTaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "AdhocPlan.AdhocPlanToAgentTask"`)
	}
	return nil
}

func (apuo *AdhocPlanUpdateOne) sqlSave(ctx context.Context) (_node *AdhocPlan, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   adhocplan.Table,
			Columns: adhocplan.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: adhocplan.FieldID,
			},
		},
	}
	id, ok := apuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AdhocPlan.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := apuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, adhocplan.FieldID)
		for _, f := range fields {
			if !adhocplan.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != adhocplan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := apuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if apuo.mutation.PrevAdhocPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   adhocplan.PrevAdhocPlanTable,
			Columns: adhocplan.PrevAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.RemovedPrevAdhocPlanIDs(); len(nodes) > 0 && !apuo.mutation.PrevAdhocPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   adhocplan.PrevAdhocPlanTable,
			Columns: adhocplan.PrevAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.PrevAdhocPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   adhocplan.PrevAdhocPlanTable,
			Columns: adhocplan.PrevAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apuo.mutation.NextAdhocPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adhocplan.NextAdhocPlanTable,
			Columns: adhocplan.NextAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.RemovedNextAdhocPlanIDs(); len(nodes) > 0 && !apuo.mutation.NextAdhocPlanCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adhocplan.NextAdhocPlanTable,
			Columns: adhocplan.NextAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.NextAdhocPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   adhocplan.NextAdhocPlanTable,
			Columns: adhocplan.NextAdhocPlanPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: adhocplan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apuo.mutation.AdhocPlanToBuildCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToBuildTable,
			Columns: []string{adhocplan.AdhocPlanToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.AdhocPlanToBuildIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToBuildTable,
			Columns: []string{adhocplan.AdhocPlanToBuildColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: build.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apuo.mutation.AdhocPlanToStatusCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToStatusTable,
			Columns: []string{adhocplan.AdhocPlanToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.AdhocPlanToStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToStatusTable,
			Columns: []string{adhocplan.AdhocPlanToStatusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: status.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if apuo.mutation.AdhocPlanToAgentTaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToAgentTaskTable,
			Columns: []string{adhocplan.AdhocPlanToAgentTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: agenttask.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := apuo.mutation.AdhocPlanToAgentTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   adhocplan.AdhocPlanToAgentTaskTable,
			Columns: []string{adhocplan.AdhocPlanToAgentTaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: agenttask.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &AdhocPlan{config: apuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, apuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{adhocplan.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
