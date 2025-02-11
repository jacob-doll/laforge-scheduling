// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/gen0cide/laforge/ent/hostdependency"
	"github.com/gen0cide/laforge/ent/network"
	"github.com/google/uuid"
)

// HostDependencyCreate is the builder for creating a HostDependency entity.
type HostDependencyCreate struct {
	config
	mutation *HostDependencyMutation
	hooks    []Hook
}

// SetHostID sets the "host_id" field.
func (hdc *HostDependencyCreate) SetHostID(s string) *HostDependencyCreate {
	hdc.mutation.SetHostID(s)
	return hdc
}

// SetNetworkID sets the "network_id" field.
func (hdc *HostDependencyCreate) SetNetworkID(s string) *HostDependencyCreate {
	hdc.mutation.SetNetworkID(s)
	return hdc
}

// SetID sets the "id" field.
func (hdc *HostDependencyCreate) SetID(u uuid.UUID) *HostDependencyCreate {
	hdc.mutation.SetID(u)
	return hdc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (hdc *HostDependencyCreate) SetNillableID(u *uuid.UUID) *HostDependencyCreate {
	if u != nil {
		hdc.SetID(*u)
	}
	return hdc
}

// SetHostDependencyToDependOnHostID sets the "HostDependencyToDependOnHost" edge to the Host entity by ID.
func (hdc *HostDependencyCreate) SetHostDependencyToDependOnHostID(id uuid.UUID) *HostDependencyCreate {
	hdc.mutation.SetHostDependencyToDependOnHostID(id)
	return hdc
}

// SetNillableHostDependencyToDependOnHostID sets the "HostDependencyToDependOnHost" edge to the Host entity by ID if the given value is not nil.
func (hdc *HostDependencyCreate) SetNillableHostDependencyToDependOnHostID(id *uuid.UUID) *HostDependencyCreate {
	if id != nil {
		hdc = hdc.SetHostDependencyToDependOnHostID(*id)
	}
	return hdc
}

// SetHostDependencyToDependOnHost sets the "HostDependencyToDependOnHost" edge to the Host entity.
func (hdc *HostDependencyCreate) SetHostDependencyToDependOnHost(h *Host) *HostDependencyCreate {
	return hdc.SetHostDependencyToDependOnHostID(h.ID)
}

// SetHostDependencyToDependByHostID sets the "HostDependencyToDependByHost" edge to the Host entity by ID.
func (hdc *HostDependencyCreate) SetHostDependencyToDependByHostID(id uuid.UUID) *HostDependencyCreate {
	hdc.mutation.SetHostDependencyToDependByHostID(id)
	return hdc
}

// SetNillableHostDependencyToDependByHostID sets the "HostDependencyToDependByHost" edge to the Host entity by ID if the given value is not nil.
func (hdc *HostDependencyCreate) SetNillableHostDependencyToDependByHostID(id *uuid.UUID) *HostDependencyCreate {
	if id != nil {
		hdc = hdc.SetHostDependencyToDependByHostID(*id)
	}
	return hdc
}

// SetHostDependencyToDependByHost sets the "HostDependencyToDependByHost" edge to the Host entity.
func (hdc *HostDependencyCreate) SetHostDependencyToDependByHost(h *Host) *HostDependencyCreate {
	return hdc.SetHostDependencyToDependByHostID(h.ID)
}

// SetHostDependencyToNetworkID sets the "HostDependencyToNetwork" edge to the Network entity by ID.
func (hdc *HostDependencyCreate) SetHostDependencyToNetworkID(id uuid.UUID) *HostDependencyCreate {
	hdc.mutation.SetHostDependencyToNetworkID(id)
	return hdc
}

// SetNillableHostDependencyToNetworkID sets the "HostDependencyToNetwork" edge to the Network entity by ID if the given value is not nil.
func (hdc *HostDependencyCreate) SetNillableHostDependencyToNetworkID(id *uuid.UUID) *HostDependencyCreate {
	if id != nil {
		hdc = hdc.SetHostDependencyToNetworkID(*id)
	}
	return hdc
}

// SetHostDependencyToNetwork sets the "HostDependencyToNetwork" edge to the Network entity.
func (hdc *HostDependencyCreate) SetHostDependencyToNetwork(n *Network) *HostDependencyCreate {
	return hdc.SetHostDependencyToNetworkID(n.ID)
}

// SetHostDependencyToEnvironmentID sets the "HostDependencyToEnvironment" edge to the Environment entity by ID.
func (hdc *HostDependencyCreate) SetHostDependencyToEnvironmentID(id uuid.UUID) *HostDependencyCreate {
	hdc.mutation.SetHostDependencyToEnvironmentID(id)
	return hdc
}

// SetNillableHostDependencyToEnvironmentID sets the "HostDependencyToEnvironment" edge to the Environment entity by ID if the given value is not nil.
func (hdc *HostDependencyCreate) SetNillableHostDependencyToEnvironmentID(id *uuid.UUID) *HostDependencyCreate {
	if id != nil {
		hdc = hdc.SetHostDependencyToEnvironmentID(*id)
	}
	return hdc
}

// SetHostDependencyToEnvironment sets the "HostDependencyToEnvironment" edge to the Environment entity.
func (hdc *HostDependencyCreate) SetHostDependencyToEnvironment(e *Environment) *HostDependencyCreate {
	return hdc.SetHostDependencyToEnvironmentID(e.ID)
}

// Mutation returns the HostDependencyMutation object of the builder.
func (hdc *HostDependencyCreate) Mutation() *HostDependencyMutation {
	return hdc.mutation
}

// Save creates the HostDependency in the database.
func (hdc *HostDependencyCreate) Save(ctx context.Context) (*HostDependency, error) {
	var (
		err  error
		node *HostDependency
	)
	hdc.defaults()
	if len(hdc.hooks) == 0 {
		if err = hdc.check(); err != nil {
			return nil, err
		}
		node, err = hdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HostDependencyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hdc.check(); err != nil {
				return nil, err
			}
			hdc.mutation = mutation
			if node, err = hdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(hdc.hooks) - 1; i >= 0; i-- {
			if hdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hdc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, hdc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*HostDependency)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from HostDependencyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hdc *HostDependencyCreate) SaveX(ctx context.Context) *HostDependency {
	v, err := hdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hdc *HostDependencyCreate) Exec(ctx context.Context) error {
	_, err := hdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdc *HostDependencyCreate) ExecX(ctx context.Context) {
	if err := hdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hdc *HostDependencyCreate) defaults() {
	if _, ok := hdc.mutation.ID(); !ok {
		v := hostdependency.DefaultID()
		hdc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hdc *HostDependencyCreate) check() error {
	if _, ok := hdc.mutation.HostID(); !ok {
		return &ValidationError{Name: "host_id", err: errors.New(`ent: missing required field "HostDependency.host_id"`)}
	}
	if _, ok := hdc.mutation.NetworkID(); !ok {
		return &ValidationError{Name: "network_id", err: errors.New(`ent: missing required field "HostDependency.network_id"`)}
	}
	return nil
}

func (hdc *HostDependencyCreate) sqlSave(ctx context.Context) (*HostDependency, error) {
	_node, _spec := hdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hdc.driver, _spec); err != nil {
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

func (hdc *HostDependencyCreate) createSpec() (*HostDependency, *sqlgraph.CreateSpec) {
	var (
		_node = &HostDependency{config: hdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: hostdependency.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hostdependency.FieldID,
			},
		}
	)
	if id, ok := hdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := hdc.mutation.HostID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hostdependency.FieldHostID,
		})
		_node.HostID = value
	}
	if value, ok := hdc.mutation.NetworkID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: hostdependency.FieldNetworkID,
		})
		_node.NetworkID = value
	}
	if nodes := hdc.mutation.HostDependencyToDependOnHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostdependency.HostDependencyToDependOnHostTable,
			Columns: []string{hostdependency.HostDependencyToDependOnHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.host_dependency_host_dependency_to_depend_on_host = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hdc.mutation.HostDependencyToDependByHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostdependency.HostDependencyToDependByHostTable,
			Columns: []string{hostdependency.HostDependencyToDependByHostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: host.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.host_dependency_host_dependency_to_depend_by_host = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hdc.mutation.HostDependencyToNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostdependency.HostDependencyToNetworkTable,
			Columns: []string{hostdependency.HostDependencyToNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: network.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.host_dependency_host_dependency_to_network = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hdc.mutation.HostDependencyToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hostdependency.HostDependencyToEnvironmentTable,
			Columns: []string{hostdependency.HostDependencyToEnvironmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.environment_environment_to_host_dependency = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HostDependencyCreateBulk is the builder for creating many HostDependency entities in bulk.
type HostDependencyCreateBulk struct {
	config
	builders []*HostDependencyCreate
}

// Save creates the HostDependency entities in the database.
func (hdcb *HostDependencyCreateBulk) Save(ctx context.Context) ([]*HostDependency, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hdcb.builders))
	nodes := make([]*HostDependency, len(hdcb.builders))
	mutators := make([]Mutator, len(hdcb.builders))
	for i := range hdcb.builders {
		func(i int, root context.Context) {
			builder := hdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HostDependencyMutation)
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
					_, err = mutators[i+1].Mutate(root, hdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hdcb *HostDependencyCreateBulk) SaveX(ctx context.Context) []*HostDependency {
	v, err := hdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hdcb *HostDependencyCreateBulk) Exec(ctx context.Context) error {
	_, err := hdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hdcb *HostDependencyCreateBulk) ExecX(ctx context.Context) {
	if err := hdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
