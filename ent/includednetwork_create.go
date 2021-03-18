// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/includednetwork"
	"github.com/gen0cide/laforge/ent/tag"
)

// IncludedNetworkCreate is the builder for creating a IncludedNetwork entity.
type IncludedNetworkCreate struct {
	config
	mutation *IncludedNetworkMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (inc *IncludedNetworkCreate) SetName(s string) *IncludedNetworkCreate {
	inc.mutation.SetName(s)
	return inc
}

// SetHosts sets the "hosts" field.
func (inc *IncludedNetworkCreate) SetHosts(s []string) *IncludedNetworkCreate {
	inc.mutation.SetHosts(s)
	return inc
}

// AddIncludedNetworkToTagIDs adds the "IncludedNetworkToTag" edge to the Tag entity by IDs.
func (inc *IncludedNetworkCreate) AddIncludedNetworkToTagIDs(ids ...int) *IncludedNetworkCreate {
	inc.mutation.AddIncludedNetworkToTagIDs(ids...)
	return inc
}

// AddIncludedNetworkToTag adds the "IncludedNetworkToTag" edges to the Tag entity.
func (inc *IncludedNetworkCreate) AddIncludedNetworkToTag(t ...*Tag) *IncludedNetworkCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return inc.AddIncludedNetworkToTagIDs(ids...)
}

// AddIncludedNetworkToEnvironmentIDs adds the "IncludedNetworkToEnvironment" edge to the Environment entity by IDs.
func (inc *IncludedNetworkCreate) AddIncludedNetworkToEnvironmentIDs(ids ...int) *IncludedNetworkCreate {
	inc.mutation.AddIncludedNetworkToEnvironmentIDs(ids...)
	return inc
}

// AddIncludedNetworkToEnvironment adds the "IncludedNetworkToEnvironment" edges to the Environment entity.
func (inc *IncludedNetworkCreate) AddIncludedNetworkToEnvironment(e ...*Environment) *IncludedNetworkCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return inc.AddIncludedNetworkToEnvironmentIDs(ids...)
}

// Mutation returns the IncludedNetworkMutation object of the builder.
func (inc *IncludedNetworkCreate) Mutation() *IncludedNetworkMutation {
	return inc.mutation
}

// Save creates the IncludedNetwork in the database.
func (inc *IncludedNetworkCreate) Save(ctx context.Context) (*IncludedNetwork, error) {
	var (
		err  error
		node *IncludedNetwork
	)
	if len(inc.hooks) == 0 {
		if err = inc.check(); err != nil {
			return nil, err
		}
		node, err = inc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IncludedNetworkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = inc.check(); err != nil {
				return nil, err
			}
			inc.mutation = mutation
			node, err = inc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(inc.hooks) - 1; i >= 0; i-- {
			mut = inc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, inc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (inc *IncludedNetworkCreate) SaveX(ctx context.Context) *IncludedNetwork {
	v, err := inc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (inc *IncludedNetworkCreate) check() error {
	if _, ok := inc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := inc.mutation.Hosts(); !ok {
		return &ValidationError{Name: "hosts", err: errors.New("ent: missing required field \"hosts\"")}
	}
	return nil
}

func (inc *IncludedNetworkCreate) sqlSave(ctx context.Context) (*IncludedNetwork, error) {
	_node, _spec := inc.createSpec()
	if err := sqlgraph.CreateNode(ctx, inc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (inc *IncludedNetworkCreate) createSpec() (*IncludedNetwork, *sqlgraph.CreateSpec) {
	var (
		_node = &IncludedNetwork{config: inc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: includednetwork.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: includednetwork.FieldID,
			},
		}
	)
	if value, ok := inc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: includednetwork.FieldName,
		})
		_node.Name = value
	}
	if value, ok := inc.mutation.Hosts(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: includednetwork.FieldHosts,
		})
		_node.Hosts = value
	}
	if nodes := inc.mutation.IncludedNetworkToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   includednetwork.IncludedNetworkToTagTable,
			Columns: []string{includednetwork.IncludedNetworkToTagColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := inc.mutation.IncludedNetworkToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   includednetwork.IncludedNetworkToEnvironmentTable,
			Columns: includednetwork.IncludedNetworkToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IncludedNetworkCreateBulk is the builder for creating many IncludedNetwork entities in bulk.
type IncludedNetworkCreateBulk struct {
	config
	builders []*IncludedNetworkCreate
}

// Save creates the IncludedNetwork entities in the database.
func (incb *IncludedNetworkCreateBulk) Save(ctx context.Context) ([]*IncludedNetwork, error) {
	specs := make([]*sqlgraph.CreateSpec, len(incb.builders))
	nodes := make([]*IncludedNetwork, len(incb.builders))
	mutators := make([]Mutator, len(incb.builders))
	for i := range incb.builders {
		func(i int, root context.Context) {
			builder := incb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IncludedNetworkMutation)
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
					_, err = mutators[i+1].Mutate(root, incb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, incb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, incb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (incb *IncludedNetworkCreateBulk) SaveX(ctx context.Context) []*IncludedNetwork {
	v, err := incb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
