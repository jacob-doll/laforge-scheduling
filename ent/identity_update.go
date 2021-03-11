// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/identity"
	"github.com/gen0cide/laforge/ent/predicate"
)

// IdentityUpdate is the builder for updating Identity entities.
type IdentityUpdate struct {
	config
	hooks    []Hook
	mutation *IdentityMutation
}

// Where adds a new predicate for the IdentityUpdate builder.
func (iu *IdentityUpdate) Where(ps ...predicate.Identity) *IdentityUpdate {
	iu.mutation.predicates = append(iu.mutation.predicates, ps...)
	return iu
}

// SetHclID sets the "hcl_id" field.
func (iu *IdentityUpdate) SetHclID(s string) *IdentityUpdate {
	iu.mutation.SetHclID(s)
	return iu
}

// SetFirstName sets the "first_name" field.
func (iu *IdentityUpdate) SetFirstName(s string) *IdentityUpdate {
	iu.mutation.SetFirstName(s)
	return iu
}

// SetLastName sets the "last_name" field.
func (iu *IdentityUpdate) SetLastName(s string) *IdentityUpdate {
	iu.mutation.SetLastName(s)
	return iu
}

// SetEmail sets the "email" field.
func (iu *IdentityUpdate) SetEmail(s string) *IdentityUpdate {
	iu.mutation.SetEmail(s)
	return iu
}

// SetPassword sets the "password" field.
func (iu *IdentityUpdate) SetPassword(s string) *IdentityUpdate {
	iu.mutation.SetPassword(s)
	return iu
}

// SetDescription sets the "description" field.
func (iu *IdentityUpdate) SetDescription(s string) *IdentityUpdate {
	iu.mutation.SetDescription(s)
	return iu
}

// SetAvatarFile sets the "avatar_file" field.
func (iu *IdentityUpdate) SetAvatarFile(s string) *IdentityUpdate {
	iu.mutation.SetAvatarFile(s)
	return iu
}

// SetVars sets the "vars" field.
func (iu *IdentityUpdate) SetVars(m map[string]string) *IdentityUpdate {
	iu.mutation.SetVars(m)
	return iu
}

// SetTags sets the "tags" field.
func (iu *IdentityUpdate) SetTags(m map[string]string) *IdentityUpdate {
	iu.mutation.SetTags(m)
	return iu
}

// AddIdentityToEnvironmentIDs adds the "IdentityToEnvironment" edge to the Environment entity by IDs.
func (iu *IdentityUpdate) AddIdentityToEnvironmentIDs(ids ...int) *IdentityUpdate {
	iu.mutation.AddIdentityToEnvironmentIDs(ids...)
	return iu
}

// AddIdentityToEnvironment adds the "IdentityToEnvironment" edges to the Environment entity.
func (iu *IdentityUpdate) AddIdentityToEnvironment(e ...*Environment) *IdentityUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iu.AddIdentityToEnvironmentIDs(ids...)
}

// Mutation returns the IdentityMutation object of the builder.
func (iu *IdentityUpdate) Mutation() *IdentityMutation {
	return iu.mutation
}

// ClearIdentityToEnvironment clears all "IdentityToEnvironment" edges to the Environment entity.
func (iu *IdentityUpdate) ClearIdentityToEnvironment() *IdentityUpdate {
	iu.mutation.ClearIdentityToEnvironment()
	return iu
}

// RemoveIdentityToEnvironmentIDs removes the "IdentityToEnvironment" edge to Environment entities by IDs.
func (iu *IdentityUpdate) RemoveIdentityToEnvironmentIDs(ids ...int) *IdentityUpdate {
	iu.mutation.RemoveIdentityToEnvironmentIDs(ids...)
	return iu
}

// RemoveIdentityToEnvironment removes "IdentityToEnvironment" edges to Environment entities.
func (iu *IdentityUpdate) RemoveIdentityToEnvironment(e ...*Environment) *IdentityUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iu.RemoveIdentityToEnvironmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *IdentityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IdentityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *IdentityUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *IdentityUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *IdentityUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *IdentityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   identity.Table,
			Columns: identity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: identity.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldHclID,
		})
	}
	if value, ok := iu.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldFirstName,
		})
	}
	if value, ok := iu.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldLastName,
		})
	}
	if value, ok := iu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldEmail,
		})
	}
	if value, ok := iu.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldPassword,
		})
	}
	if value, ok := iu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldDescription,
		})
	}
	if value, ok := iu.mutation.AvatarFile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldAvatarFile,
		})
	}
	if value, ok := iu.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: identity.FieldVars,
		})
	}
	if value, ok := iu.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: identity.FieldTags,
		})
	}
	if iu.mutation.IdentityToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identity.IdentityToEnvironmentTable,
			Columns: identity.IdentityToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedIdentityToEnvironmentIDs(); len(nodes) > 0 && !iu.mutation.IdentityToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identity.IdentityToEnvironmentTable,
			Columns: identity.IdentityToEnvironmentPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.IdentityToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identity.IdentityToEnvironmentTable,
			Columns: identity.IdentityToEnvironmentPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{identity.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// IdentityUpdateOne is the builder for updating a single Identity entity.
type IdentityUpdateOne struct {
	config
	hooks    []Hook
	mutation *IdentityMutation
}

// SetHclID sets the "hcl_id" field.
func (iuo *IdentityUpdateOne) SetHclID(s string) *IdentityUpdateOne {
	iuo.mutation.SetHclID(s)
	return iuo
}

// SetFirstName sets the "first_name" field.
func (iuo *IdentityUpdateOne) SetFirstName(s string) *IdentityUpdateOne {
	iuo.mutation.SetFirstName(s)
	return iuo
}

// SetLastName sets the "last_name" field.
func (iuo *IdentityUpdateOne) SetLastName(s string) *IdentityUpdateOne {
	iuo.mutation.SetLastName(s)
	return iuo
}

// SetEmail sets the "email" field.
func (iuo *IdentityUpdateOne) SetEmail(s string) *IdentityUpdateOne {
	iuo.mutation.SetEmail(s)
	return iuo
}

// SetPassword sets the "password" field.
func (iuo *IdentityUpdateOne) SetPassword(s string) *IdentityUpdateOne {
	iuo.mutation.SetPassword(s)
	return iuo
}

// SetDescription sets the "description" field.
func (iuo *IdentityUpdateOne) SetDescription(s string) *IdentityUpdateOne {
	iuo.mutation.SetDescription(s)
	return iuo
}

// SetAvatarFile sets the "avatar_file" field.
func (iuo *IdentityUpdateOne) SetAvatarFile(s string) *IdentityUpdateOne {
	iuo.mutation.SetAvatarFile(s)
	return iuo
}

// SetVars sets the "vars" field.
func (iuo *IdentityUpdateOne) SetVars(m map[string]string) *IdentityUpdateOne {
	iuo.mutation.SetVars(m)
	return iuo
}

// SetTags sets the "tags" field.
func (iuo *IdentityUpdateOne) SetTags(m map[string]string) *IdentityUpdateOne {
	iuo.mutation.SetTags(m)
	return iuo
}

// AddIdentityToEnvironmentIDs adds the "IdentityToEnvironment" edge to the Environment entity by IDs.
func (iuo *IdentityUpdateOne) AddIdentityToEnvironmentIDs(ids ...int) *IdentityUpdateOne {
	iuo.mutation.AddIdentityToEnvironmentIDs(ids...)
	return iuo
}

// AddIdentityToEnvironment adds the "IdentityToEnvironment" edges to the Environment entity.
func (iuo *IdentityUpdateOne) AddIdentityToEnvironment(e ...*Environment) *IdentityUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iuo.AddIdentityToEnvironmentIDs(ids...)
}

// Mutation returns the IdentityMutation object of the builder.
func (iuo *IdentityUpdateOne) Mutation() *IdentityMutation {
	return iuo.mutation
}

// ClearIdentityToEnvironment clears all "IdentityToEnvironment" edges to the Environment entity.
func (iuo *IdentityUpdateOne) ClearIdentityToEnvironment() *IdentityUpdateOne {
	iuo.mutation.ClearIdentityToEnvironment()
	return iuo
}

// RemoveIdentityToEnvironmentIDs removes the "IdentityToEnvironment" edge to Environment entities by IDs.
func (iuo *IdentityUpdateOne) RemoveIdentityToEnvironmentIDs(ids ...int) *IdentityUpdateOne {
	iuo.mutation.RemoveIdentityToEnvironmentIDs(ids...)
	return iuo
}

// RemoveIdentityToEnvironment removes "IdentityToEnvironment" edges to Environment entities.
func (iuo *IdentityUpdateOne) RemoveIdentityToEnvironment(e ...*Environment) *IdentityUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return iuo.RemoveIdentityToEnvironmentIDs(ids...)
}

// Save executes the query and returns the updated Identity entity.
func (iuo *IdentityUpdateOne) Save(ctx context.Context) (*Identity, error) {
	var (
		err  error
		node *Identity
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IdentityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *IdentityUpdateOne) SaveX(ctx context.Context) *Identity {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *IdentityUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *IdentityUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *IdentityUpdateOne) sqlSave(ctx context.Context) (_node *Identity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   identity.Table,
			Columns: identity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: identity.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Identity.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldHclID,
		})
	}
	if value, ok := iuo.mutation.FirstName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldFirstName,
		})
	}
	if value, ok := iuo.mutation.LastName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldLastName,
		})
	}
	if value, ok := iuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldEmail,
		})
	}
	if value, ok := iuo.mutation.Password(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldPassword,
		})
	}
	if value, ok := iuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldDescription,
		})
	}
	if value, ok := iuo.mutation.AvatarFile(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: identity.FieldAvatarFile,
		})
	}
	if value, ok := iuo.mutation.Vars(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: identity.FieldVars,
		})
	}
	if value, ok := iuo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: identity.FieldTags,
		})
	}
	if iuo.mutation.IdentityToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identity.IdentityToEnvironmentTable,
			Columns: identity.IdentityToEnvironmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: environment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedIdentityToEnvironmentIDs(); len(nodes) > 0 && !iuo.mutation.IdentityToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identity.IdentityToEnvironmentTable,
			Columns: identity.IdentityToEnvironmentPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.IdentityToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   identity.IdentityToEnvironmentTable,
			Columns: identity.IdentityToEnvironmentPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Identity{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{identity.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
