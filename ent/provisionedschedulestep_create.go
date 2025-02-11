// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/agenttask"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisionedschedulestep"
	"github.com/gen0cide/laforge/ent/schedulestep"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/google/uuid"
)

// ProvisionedScheduleStepCreate is the builder for creating a ProvisionedScheduleStep entity.
type ProvisionedScheduleStepCreate struct {
	config
	mutation *ProvisionedScheduleStepMutation
	hooks    []Hook
}

// SetRunTime sets the "run_time" field.
func (pssc *ProvisionedScheduleStepCreate) SetRunTime(t time.Time) *ProvisionedScheduleStepCreate {
	pssc.mutation.SetRunTime(t)
	return pssc
}

// SetID sets the "id" field.
func (pssc *ProvisionedScheduleStepCreate) SetID(u uuid.UUID) *ProvisionedScheduleStepCreate {
	pssc.mutation.SetID(u)
	return pssc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pssc *ProvisionedScheduleStepCreate) SetNillableID(u *uuid.UUID) *ProvisionedScheduleStepCreate {
	if u != nil {
		pssc.SetID(*u)
	}
	return pssc
}

// SetProvisionedScheduleStepToStatusID sets the "ProvisionedScheduleStepToStatus" edge to the Status entity by ID.
func (pssc *ProvisionedScheduleStepCreate) SetProvisionedScheduleStepToStatusID(id uuid.UUID) *ProvisionedScheduleStepCreate {
	pssc.mutation.SetProvisionedScheduleStepToStatusID(id)
	return pssc
}

// SetNillableProvisionedScheduleStepToStatusID sets the "ProvisionedScheduleStepToStatus" edge to the Status entity by ID if the given value is not nil.
func (pssc *ProvisionedScheduleStepCreate) SetNillableProvisionedScheduleStepToStatusID(id *uuid.UUID) *ProvisionedScheduleStepCreate {
	if id != nil {
		pssc = pssc.SetProvisionedScheduleStepToStatusID(*id)
	}
	return pssc
}

// SetProvisionedScheduleStepToStatus sets the "ProvisionedScheduleStepToStatus" edge to the Status entity.
func (pssc *ProvisionedScheduleStepCreate) SetProvisionedScheduleStepToStatus(s *Status) *ProvisionedScheduleStepCreate {
	return pssc.SetProvisionedScheduleStepToStatusID(s.ID)
}

// SetProvisionedScheduleStepToScheduleStepID sets the "ProvisionedScheduleStepToScheduleStep" edge to the ScheduleStep entity by ID.
func (pssc *ProvisionedScheduleStepCreate) SetProvisionedScheduleStepToScheduleStepID(id uuid.UUID) *ProvisionedScheduleStepCreate {
	pssc.mutation.SetProvisionedScheduleStepToScheduleStepID(id)
	return pssc
}

// SetProvisionedScheduleStepToScheduleStep sets the "ProvisionedScheduleStepToScheduleStep" edge to the ScheduleStep entity.
func (pssc *ProvisionedScheduleStepCreate) SetProvisionedScheduleStepToScheduleStep(s *ScheduleStep) *ProvisionedScheduleStepCreate {
	return pssc.SetProvisionedScheduleStepToScheduleStepID(s.ID)
}

// SetProvisionedScheduleStepToProvisionedHostID sets the "ProvisionedScheduleStepToProvisionedHost" edge to the ProvisionedHost entity by ID.
func (pssc *ProvisionedScheduleStepCreate) SetProvisionedScheduleStepToProvisionedHostID(id uuid.UUID) *ProvisionedScheduleStepCreate {
	pssc.mutation.SetProvisionedScheduleStepToProvisionedHostID(id)
	return pssc
}

// SetProvisionedScheduleStepToProvisionedHost sets the "ProvisionedScheduleStepToProvisionedHost" edge to the ProvisionedHost entity.
func (pssc *ProvisionedScheduleStepCreate) SetProvisionedScheduleStepToProvisionedHost(p *ProvisionedHost) *ProvisionedScheduleStepCreate {
	return pssc.SetProvisionedScheduleStepToProvisionedHostID(p.ID)
}

// SetProvisionedScheduleStepToAgentTaskID sets the "ProvisionedScheduleStepToAgentTask" edge to the AgentTask entity by ID.
func (pssc *ProvisionedScheduleStepCreate) SetProvisionedScheduleStepToAgentTaskID(id uuid.UUID) *ProvisionedScheduleStepCreate {
	pssc.mutation.SetProvisionedScheduleStepToAgentTaskID(id)
	return pssc
}

// SetNillableProvisionedScheduleStepToAgentTaskID sets the "ProvisionedScheduleStepToAgentTask" edge to the AgentTask entity by ID if the given value is not nil.
func (pssc *ProvisionedScheduleStepCreate) SetNillableProvisionedScheduleStepToAgentTaskID(id *uuid.UUID) *ProvisionedScheduleStepCreate {
	if id != nil {
		pssc = pssc.SetProvisionedScheduleStepToAgentTaskID(*id)
	}
	return pssc
}

// SetProvisionedScheduleStepToAgentTask sets the "ProvisionedScheduleStepToAgentTask" edge to the AgentTask entity.
func (pssc *ProvisionedScheduleStepCreate) SetProvisionedScheduleStepToAgentTask(a *AgentTask) *ProvisionedScheduleStepCreate {
	return pssc.SetProvisionedScheduleStepToAgentTaskID(a.ID)
}

// Mutation returns the ProvisionedScheduleStepMutation object of the builder.
func (pssc *ProvisionedScheduleStepCreate) Mutation() *ProvisionedScheduleStepMutation {
	return pssc.mutation
}

// Save creates the ProvisionedScheduleStep in the database.
func (pssc *ProvisionedScheduleStepCreate) Save(ctx context.Context) (*ProvisionedScheduleStep, error) {
	var (
		err  error
		node *ProvisionedScheduleStep
	)
	pssc.defaults()
	if len(pssc.hooks) == 0 {
		if err = pssc.check(); err != nil {
			return nil, err
		}
		node, err = pssc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvisionedScheduleStepMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pssc.check(); err != nil {
				return nil, err
			}
			pssc.mutation = mutation
			if node, err = pssc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pssc.hooks) - 1; i >= 0; i-- {
			if pssc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pssc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pssc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ProvisionedScheduleStep)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ProvisionedScheduleStepMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pssc *ProvisionedScheduleStepCreate) SaveX(ctx context.Context) *ProvisionedScheduleStep {
	v, err := pssc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pssc *ProvisionedScheduleStepCreate) Exec(ctx context.Context) error {
	_, err := pssc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pssc *ProvisionedScheduleStepCreate) ExecX(ctx context.Context) {
	if err := pssc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pssc *ProvisionedScheduleStepCreate) defaults() {
	if _, ok := pssc.mutation.ID(); !ok {
		v := provisionedschedulestep.DefaultID()
		pssc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pssc *ProvisionedScheduleStepCreate) check() error {
	if _, ok := pssc.mutation.RunTime(); !ok {
		return &ValidationError{Name: "run_time", err: errors.New(`ent: missing required field "ProvisionedScheduleStep.run_time"`)}
	}
	if _, ok := pssc.mutation.ProvisionedScheduleStepToScheduleStepID(); !ok {
		return &ValidationError{Name: "ProvisionedScheduleStepToScheduleStep", err: errors.New(`ent: missing required edge "ProvisionedScheduleStep.ProvisionedScheduleStepToScheduleStep"`)}
	}
	if _, ok := pssc.mutation.ProvisionedScheduleStepToProvisionedHostID(); !ok {
		return &ValidationError{Name: "ProvisionedScheduleStepToProvisionedHost", err: errors.New(`ent: missing required edge "ProvisionedScheduleStep.ProvisionedScheduleStepToProvisionedHost"`)}
	}
	return nil
}

func (pssc *ProvisionedScheduleStepCreate) sqlSave(ctx context.Context) (*ProvisionedScheduleStep, error) {
	_node, _spec := pssc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pssc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pssc *ProvisionedScheduleStepCreate) createSpec() (*ProvisionedScheduleStep, *sqlgraph.CreateSpec) {
	var (
		_node = &ProvisionedScheduleStep{config: pssc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: provisionedschedulestep.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: provisionedschedulestep.FieldID,
			},
		}
	)
	if id, ok := pssc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pssc.mutation.RunTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: provisionedschedulestep.FieldRunTime,
		})
		_node.RunTime = value
	}
	if nodes := pssc.mutation.ProvisionedScheduleStepToStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   provisionedschedulestep.ProvisionedScheduleStepToStatusTable,
			Columns: []string{provisionedschedulestep.ProvisionedScheduleStepToStatusColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pssc.mutation.ProvisionedScheduleStepToScheduleStepIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provisionedschedulestep.ProvisionedScheduleStepToScheduleStepTable,
			Columns: []string{provisionedschedulestep.ProvisionedScheduleStepToScheduleStepColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: schedulestep.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.schedule_step_schedule_step_to_provisioned_schedule_step = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pssc.mutation.ProvisionedScheduleStepToProvisionedHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   provisionedschedulestep.ProvisionedScheduleStepToProvisionedHostTable,
			Columns: []string{provisionedschedulestep.ProvisionedScheduleStepToProvisionedHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: provisionedhost.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioned_host_provisioned_host_to_provisioned_schedule_step = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pssc.mutation.ProvisionedScheduleStepToAgentTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   provisionedschedulestep.ProvisionedScheduleStepToAgentTaskTable,
			Columns: []string{provisionedschedulestep.ProvisionedScheduleStepToAgentTaskColumn},
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
		_node.agent_task_agent_task_to_provisioned_schedule_step = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProvisionedScheduleStepCreateBulk is the builder for creating many ProvisionedScheduleStep entities in bulk.
type ProvisionedScheduleStepCreateBulk struct {
	config
	builders []*ProvisionedScheduleStepCreate
}

// Save creates the ProvisionedScheduleStep entities in the database.
func (psscb *ProvisionedScheduleStepCreateBulk) Save(ctx context.Context) ([]*ProvisionedScheduleStep, error) {
	specs := make([]*sqlgraph.CreateSpec, len(psscb.builders))
	nodes := make([]*ProvisionedScheduleStep, len(psscb.builders))
	mutators := make([]Mutator, len(psscb.builders))
	for i := range psscb.builders {
		func(i int, root context.Context) {
			builder := psscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProvisionedScheduleStepMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, psscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, psscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, psscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (psscb *ProvisionedScheduleStepCreateBulk) SaveX(ctx context.Context) []*ProvisionedScheduleStep {
	v, err := psscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psscb *ProvisionedScheduleStepCreateBulk) Exec(ctx context.Context) error {
	_, err := psscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psscb *ProvisionedScheduleStepCreateBulk) ExecX(ctx context.Context) {
	if err := psscb.Exec(ctx); err != nil {
		panic(err)
	}
}
