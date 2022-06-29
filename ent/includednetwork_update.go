// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/gen0cide/laforge/ent/includednetwork"
	"github.com/gen0cide/laforge/ent/network"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/google/uuid"
)

// IncludedNetworkUpdate is the builder for updating IncludedNetwork entities.
type IncludedNetworkUpdate struct {
	config
	hooks    []Hook
	mutation *IncludedNetworkMutation
}

// Where appends a list predicates to the IncludedNetworkUpdate builder.
func (inu *IncludedNetworkUpdate) Where(ps ...predicate.IncludedNetwork) *IncludedNetworkUpdate {
	inu.mutation.Where(ps...)
	return inu
}

// SetName sets the "name" field.
func (inu *IncludedNetworkUpdate) SetName(s string) *IncludedNetworkUpdate {
	inu.mutation.SetName(s)
	return inu
}

// SetHosts sets the "hosts" field.
func (inu *IncludedNetworkUpdate) SetHosts(s []string) *IncludedNetworkUpdate {
	inu.mutation.SetHosts(s)
	return inu
}

// AddIncludedNetworkToTagIDs adds the "IncludedNetworkToTag" edge to the Tag entity by IDs.
func (inu *IncludedNetworkUpdate) AddIncludedNetworkToTagIDs(ids ...uuid.UUID) *IncludedNetworkUpdate {
	inu.mutation.AddIncludedNetworkToTagIDs(ids...)
	return inu
}

// AddIncludedNetworkToTag adds the "IncludedNetworkToTag" edges to the Tag entity.
func (inu *IncludedNetworkUpdate) AddIncludedNetworkToTag(t ...*Tag) *IncludedNetworkUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return inu.AddIncludedNetworkToTagIDs(ids...)
}

// AddIncludedNetworkToHostIDs adds the "IncludedNetworkToHost" edge to the Host entity by IDs.
func (inu *IncludedNetworkUpdate) AddIncludedNetworkToHostIDs(ids ...uuid.UUID) *IncludedNetworkUpdate {
	inu.mutation.AddIncludedNetworkToHostIDs(ids...)
	return inu
}

// AddIncludedNetworkToHost adds the "IncludedNetworkToHost" edges to the Host entity.
func (inu *IncludedNetworkUpdate) AddIncludedNetworkToHost(h ...*Host) *IncludedNetworkUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return inu.AddIncludedNetworkToHostIDs(ids...)
}

// SetIncludedNetworkToNetworkID sets the "IncludedNetworkToNetwork" edge to the Network entity by ID.
func (inu *IncludedNetworkUpdate) SetIncludedNetworkToNetworkID(id uuid.UUID) *IncludedNetworkUpdate {
	inu.mutation.SetIncludedNetworkToNetworkID(id)
	return inu
}

// SetNillableIncludedNetworkToNetworkID sets the "IncludedNetworkToNetwork" edge to the Network entity by ID if the given value is not nil.
func (inu *IncludedNetworkUpdate) SetNillableIncludedNetworkToNetworkID(id *uuid.UUID) *IncludedNetworkUpdate {
	if id != nil {
		inu = inu.SetIncludedNetworkToNetworkID(*id)
	}
	return inu
}

// SetIncludedNetworkToNetwork sets the "IncludedNetworkToNetwork" edge to the Network entity.
func (inu *IncludedNetworkUpdate) SetIncludedNetworkToNetwork(n *Network) *IncludedNetworkUpdate {
	return inu.SetIncludedNetworkToNetworkID(n.ID)
}

// AddIncludedNetworkToEnvironmentIDs adds the "IncludedNetworkToEnvironment" edge to the Environment entity by IDs.
func (inu *IncludedNetworkUpdate) AddIncludedNetworkToEnvironmentIDs(ids ...uuid.UUID) *IncludedNetworkUpdate {
	inu.mutation.AddIncludedNetworkToEnvironmentIDs(ids...)
	return inu
}

// AddIncludedNetworkToEnvironment adds the "IncludedNetworkToEnvironment" edges to the Environment entity.
func (inu *IncludedNetworkUpdate) AddIncludedNetworkToEnvironment(e ...*Environment) *IncludedNetworkUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return inu.AddIncludedNetworkToEnvironmentIDs(ids...)
}

// Mutation returns the IncludedNetworkMutation object of the builder.
func (inu *IncludedNetworkUpdate) Mutation() *IncludedNetworkMutation {
	return inu.mutation
}

// ClearIncludedNetworkToTag clears all "IncludedNetworkToTag" edges to the Tag entity.
func (inu *IncludedNetworkUpdate) ClearIncludedNetworkToTag() *IncludedNetworkUpdate {
	inu.mutation.ClearIncludedNetworkToTag()
	return inu
}

// RemoveIncludedNetworkToTagIDs removes the "IncludedNetworkToTag" edge to Tag entities by IDs.
func (inu *IncludedNetworkUpdate) RemoveIncludedNetworkToTagIDs(ids ...uuid.UUID) *IncludedNetworkUpdate {
	inu.mutation.RemoveIncludedNetworkToTagIDs(ids...)
	return inu
}

// RemoveIncludedNetworkToTag removes "IncludedNetworkToTag" edges to Tag entities.
func (inu *IncludedNetworkUpdate) RemoveIncludedNetworkToTag(t ...*Tag) *IncludedNetworkUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return inu.RemoveIncludedNetworkToTagIDs(ids...)
}

// ClearIncludedNetworkToHost clears all "IncludedNetworkToHost" edges to the Host entity.
func (inu *IncludedNetworkUpdate) ClearIncludedNetworkToHost() *IncludedNetworkUpdate {
	inu.mutation.ClearIncludedNetworkToHost()
	return inu
}

// RemoveIncludedNetworkToHostIDs removes the "IncludedNetworkToHost" edge to Host entities by IDs.
func (inu *IncludedNetworkUpdate) RemoveIncludedNetworkToHostIDs(ids ...uuid.UUID) *IncludedNetworkUpdate {
	inu.mutation.RemoveIncludedNetworkToHostIDs(ids...)
	return inu
}

// RemoveIncludedNetworkToHost removes "IncludedNetworkToHost" edges to Host entities.
func (inu *IncludedNetworkUpdate) RemoveIncludedNetworkToHost(h ...*Host) *IncludedNetworkUpdate {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return inu.RemoveIncludedNetworkToHostIDs(ids...)
}

// ClearIncludedNetworkToNetwork clears the "IncludedNetworkToNetwork" edge to the Network entity.
func (inu *IncludedNetworkUpdate) ClearIncludedNetworkToNetwork() *IncludedNetworkUpdate {
	inu.mutation.ClearIncludedNetworkToNetwork()
	return inu
}

// ClearIncludedNetworkToEnvironment clears all "IncludedNetworkToEnvironment" edges to the Environment entity.
func (inu *IncludedNetworkUpdate) ClearIncludedNetworkToEnvironment() *IncludedNetworkUpdate {
	inu.mutation.ClearIncludedNetworkToEnvironment()
	return inu
}

// RemoveIncludedNetworkToEnvironmentIDs removes the "IncludedNetworkToEnvironment" edge to Environment entities by IDs.
func (inu *IncludedNetworkUpdate) RemoveIncludedNetworkToEnvironmentIDs(ids ...uuid.UUID) *IncludedNetworkUpdate {
	inu.mutation.RemoveIncludedNetworkToEnvironmentIDs(ids...)
	return inu
}

// RemoveIncludedNetworkToEnvironment removes "IncludedNetworkToEnvironment" edges to Environment entities.
func (inu *IncludedNetworkUpdate) RemoveIncludedNetworkToEnvironment(e ...*Environment) *IncludedNetworkUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return inu.RemoveIncludedNetworkToEnvironmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (inu *IncludedNetworkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(inu.hooks) == 0 {
		affected, err = inu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IncludedNetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			inu.mutation = mutation
			affected, err = inu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(inu.hooks) - 1; i >= 0; i-- {
			if inu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = inu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, inu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (inu *IncludedNetworkUpdate) SaveX(ctx context.Context) int {
	affected, err := inu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (inu *IncludedNetworkUpdate) Exec(ctx context.Context) error {
	_, err := inu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (inu *IncludedNetworkUpdate) ExecX(ctx context.Context) {
	if err := inu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (inu *IncludedNetworkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   includednetwork.Table,
			Columns: includednetwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: includednetwork.FieldID,
			},
		},
	}
	if ps := inu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := inu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: includednetwork.FieldName,
		})
	}
	if value, ok := inu.mutation.Hosts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: includednetwork.FieldHosts,
		})
	}
	if inu.mutation.IncludedNetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.RemovedIncludedNetworkToTagIDs(); len(nodes) > 0 && !inu.mutation.IncludedNetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.IncludedNetworkToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if inu.mutation.IncludedNetworkToHostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToHostTable,
			Columns: includednetwork.IncludedNetworkToHostPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: host.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.RemovedIncludedNetworkToHostIDs(); len(nodes) > 0 && !inu.mutation.IncludedNetworkToHostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToHostTable,
			Columns: includednetwork.IncludedNetworkToHostPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.IncludedNetworkToHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToHostTable,
			Columns: includednetwork.IncludedNetworkToHostPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if inu.mutation.IncludedNetworkToNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToNetworkTable,
			Columns: []string{includednetwork.IncludedNetworkToNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: network.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.IncludedNetworkToNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToNetworkTable,
			Columns: []string{includednetwork.IncludedNetworkToNetworkColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if inu.mutation.IncludedNetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.RemovedIncludedNetworkToEnvironmentIDs(); len(nodes) > 0 && !inu.mutation.IncludedNetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inu.mutation.IncludedNetworkToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, inu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{includednetwork.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// IncludedNetworkUpdateOne is the builder for updating a single IncludedNetwork entity.
type IncludedNetworkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IncludedNetworkMutation
}

// SetName sets the "name" field.
func (inuo *IncludedNetworkUpdateOne) SetName(s string) *IncludedNetworkUpdateOne {
	inuo.mutation.SetName(s)
	return inuo
}

// SetHosts sets the "hosts" field.
func (inuo *IncludedNetworkUpdateOne) SetHosts(s []string) *IncludedNetworkUpdateOne {
	inuo.mutation.SetHosts(s)
	return inuo
}

// AddIncludedNetworkToTagIDs adds the "IncludedNetworkToTag" edge to the Tag entity by IDs.
func (inuo *IncludedNetworkUpdateOne) AddIncludedNetworkToTagIDs(ids ...uuid.UUID) *IncludedNetworkUpdateOne {
	inuo.mutation.AddIncludedNetworkToTagIDs(ids...)
	return inuo
}

// AddIncludedNetworkToTag adds the "IncludedNetworkToTag" edges to the Tag entity.
func (inuo *IncludedNetworkUpdateOne) AddIncludedNetworkToTag(t ...*Tag) *IncludedNetworkUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return inuo.AddIncludedNetworkToTagIDs(ids...)
}

// AddIncludedNetworkToHostIDs adds the "IncludedNetworkToHost" edge to the Host entity by IDs.
func (inuo *IncludedNetworkUpdateOne) AddIncludedNetworkToHostIDs(ids ...uuid.UUID) *IncludedNetworkUpdateOne {
	inuo.mutation.AddIncludedNetworkToHostIDs(ids...)
	return inuo
}

// AddIncludedNetworkToHost adds the "IncludedNetworkToHost" edges to the Host entity.
func (inuo *IncludedNetworkUpdateOne) AddIncludedNetworkToHost(h ...*Host) *IncludedNetworkUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return inuo.AddIncludedNetworkToHostIDs(ids...)
}

// SetIncludedNetworkToNetworkID sets the "IncludedNetworkToNetwork" edge to the Network entity by ID.
func (inuo *IncludedNetworkUpdateOne) SetIncludedNetworkToNetworkID(id uuid.UUID) *IncludedNetworkUpdateOne {
	inuo.mutation.SetIncludedNetworkToNetworkID(id)
	return inuo
}

// SetNillableIncludedNetworkToNetworkID sets the "IncludedNetworkToNetwork" edge to the Network entity by ID if the given value is not nil.
func (inuo *IncludedNetworkUpdateOne) SetNillableIncludedNetworkToNetworkID(id *uuid.UUID) *IncludedNetworkUpdateOne {
	if id != nil {
		inuo = inuo.SetIncludedNetworkToNetworkID(*id)
	}
	return inuo
}

// SetIncludedNetworkToNetwork sets the "IncludedNetworkToNetwork" edge to the Network entity.
func (inuo *IncludedNetworkUpdateOne) SetIncludedNetworkToNetwork(n *Network) *IncludedNetworkUpdateOne {
	return inuo.SetIncludedNetworkToNetworkID(n.ID)
}

// AddIncludedNetworkToEnvironmentIDs adds the "IncludedNetworkToEnvironment" edge to the Environment entity by IDs.
func (inuo *IncludedNetworkUpdateOne) AddIncludedNetworkToEnvironmentIDs(ids ...uuid.UUID) *IncludedNetworkUpdateOne {
	inuo.mutation.AddIncludedNetworkToEnvironmentIDs(ids...)
	return inuo
}

// AddIncludedNetworkToEnvironment adds the "IncludedNetworkToEnvironment" edges to the Environment entity.
func (inuo *IncludedNetworkUpdateOne) AddIncludedNetworkToEnvironment(e ...*Environment) *IncludedNetworkUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return inuo.AddIncludedNetworkToEnvironmentIDs(ids...)
}

// Mutation returns the IncludedNetworkMutation object of the builder.
func (inuo *IncludedNetworkUpdateOne) Mutation() *IncludedNetworkMutation {
	return inuo.mutation
}

// ClearIncludedNetworkToTag clears all "IncludedNetworkToTag" edges to the Tag entity.
func (inuo *IncludedNetworkUpdateOne) ClearIncludedNetworkToTag() *IncludedNetworkUpdateOne {
	inuo.mutation.ClearIncludedNetworkToTag()
	return inuo
}

// RemoveIncludedNetworkToTagIDs removes the "IncludedNetworkToTag" edge to Tag entities by IDs.
func (inuo *IncludedNetworkUpdateOne) RemoveIncludedNetworkToTagIDs(ids ...uuid.UUID) *IncludedNetworkUpdateOne {
	inuo.mutation.RemoveIncludedNetworkToTagIDs(ids...)
	return inuo
}

// RemoveIncludedNetworkToTag removes "IncludedNetworkToTag" edges to Tag entities.
func (inuo *IncludedNetworkUpdateOne) RemoveIncludedNetworkToTag(t ...*Tag) *IncludedNetworkUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return inuo.RemoveIncludedNetworkToTagIDs(ids...)
}

// ClearIncludedNetworkToHost clears all "IncludedNetworkToHost" edges to the Host entity.
func (inuo *IncludedNetworkUpdateOne) ClearIncludedNetworkToHost() *IncludedNetworkUpdateOne {
	inuo.mutation.ClearIncludedNetworkToHost()
	return inuo
}

// RemoveIncludedNetworkToHostIDs removes the "IncludedNetworkToHost" edge to Host entities by IDs.
func (inuo *IncludedNetworkUpdateOne) RemoveIncludedNetworkToHostIDs(ids ...uuid.UUID) *IncludedNetworkUpdateOne {
	inuo.mutation.RemoveIncludedNetworkToHostIDs(ids...)
	return inuo
}

// RemoveIncludedNetworkToHost removes "IncludedNetworkToHost" edges to Host entities.
func (inuo *IncludedNetworkUpdateOne) RemoveIncludedNetworkToHost(h ...*Host) *IncludedNetworkUpdateOne {
	ids := make([]uuid.UUID, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return inuo.RemoveIncludedNetworkToHostIDs(ids...)
}

// ClearIncludedNetworkToNetwork clears the "IncludedNetworkToNetwork" edge to the Network entity.
func (inuo *IncludedNetworkUpdateOne) ClearIncludedNetworkToNetwork() *IncludedNetworkUpdateOne {
	inuo.mutation.ClearIncludedNetworkToNetwork()
	return inuo
}

// ClearIncludedNetworkToEnvironment clears all "IncludedNetworkToEnvironment" edges to the Environment entity.
func (inuo *IncludedNetworkUpdateOne) ClearIncludedNetworkToEnvironment() *IncludedNetworkUpdateOne {
	inuo.mutation.ClearIncludedNetworkToEnvironment()
	return inuo
}

// RemoveIncludedNetworkToEnvironmentIDs removes the "IncludedNetworkToEnvironment" edge to Environment entities by IDs.
func (inuo *IncludedNetworkUpdateOne) RemoveIncludedNetworkToEnvironmentIDs(ids ...uuid.UUID) *IncludedNetworkUpdateOne {
	inuo.mutation.RemoveIncludedNetworkToEnvironmentIDs(ids...)
	return inuo
}

// RemoveIncludedNetworkToEnvironment removes "IncludedNetworkToEnvironment" edges to Environment entities.
func (inuo *IncludedNetworkUpdateOne) RemoveIncludedNetworkToEnvironment(e ...*Environment) *IncludedNetworkUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return inuo.RemoveIncludedNetworkToEnvironmentIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (inuo *IncludedNetworkUpdateOne) Select(field string, fields ...string) *IncludedNetworkUpdateOne {
	inuo.fields = append([]string{field}, fields...)
	return inuo
}

// Save executes the query and returns the updated IncludedNetwork entity.
func (inuo *IncludedNetworkUpdateOne) Save(ctx context.Context) (*IncludedNetwork, error) {
	var (
		err  error
		node *IncludedNetwork
	)
	if len(inuo.hooks) == 0 {
		node, err = inuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IncludedNetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			inuo.mutation = mutation
			node, err = inuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(inuo.hooks) - 1; i >= 0; i-- {
			if inuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = inuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, inuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (inuo *IncludedNetworkUpdateOne) SaveX(ctx context.Context) *IncludedNetwork {
	node, err := inuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (inuo *IncludedNetworkUpdateOne) Exec(ctx context.Context) error {
	_, err := inuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (inuo *IncludedNetworkUpdateOne) ExecX(ctx context.Context) {
	if err := inuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (inuo *IncludedNetworkUpdateOne) sqlSave(ctx context.Context) (_node *IncludedNetwork, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   includednetwork.Table,
			Columns: includednetwork.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: includednetwork.FieldID,
			},
		},
	}
	id, ok := inuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IncludedNetwork.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := inuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, includednetwork.FieldID)
		for _, f := range fields {
			if !includednetwork.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != includednetwork.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := inuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := inuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: includednetwork.FieldName,
		})
	}
	if value, ok := inuo.mutation.Hosts(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: includednetwork.FieldHosts,
		})
	}
	if inuo.mutation.IncludedNetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.RemovedIncludedNetworkToTagIDs(); len(nodes) > 0 && !inuo.mutation.IncludedNetworkToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.IncludedNetworkToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if inuo.mutation.IncludedNetworkToHostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToHostTable,
			Columns: includednetwork.IncludedNetworkToHostPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: host.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.RemovedIncludedNetworkToHostIDs(); len(nodes) > 0 && !inuo.mutation.IncludedNetworkToHostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToHostTable,
			Columns: includednetwork.IncludedNetworkToHostPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.IncludedNetworkToHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToHostTable,
			Columns: includednetwork.IncludedNetworkToHostPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if inuo.mutation.IncludedNetworkToNetworkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToNetworkTable,
			Columns: []string{includednetwork.IncludedNetworkToNetworkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: network.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.IncludedNetworkToNetworkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToNetworkTable,
			Columns: []string{includednetwork.IncludedNetworkToNetworkColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if inuo.mutation.IncludedNetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.RemovedIncludedNetworkToEnvironmentIDs(); len(nodes) > 0 && !inuo.mutation.IncludedNetworkToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := inuo.mutation.IncludedNetworkToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IncludedNetwork{config: inuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, inuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{includednetwork.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
