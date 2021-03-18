// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/filedownload"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/gen0cide/laforge/ent/tag"
)

// FileDownloadUpdate is the builder for updating FileDownload entities.
type FileDownloadUpdate struct {
	config
	hooks    []Hook
	mutation *FileDownloadMutation
}

// Where adds a new predicate for the FileDownloadUpdate builder.
func (fdu *FileDownloadUpdate) Where(ps ...predicate.FileDownload) *FileDownloadUpdate {
	fdu.mutation.predicates = append(fdu.mutation.predicates, ps...)
	return fdu
}

// SetHclID sets the "hcl_id" field.
func (fdu *FileDownloadUpdate) SetHclID(s string) *FileDownloadUpdate {
	fdu.mutation.SetHclID(s)
	return fdu
}

// SetSourceType sets the "source_type" field.
func (fdu *FileDownloadUpdate) SetSourceType(s string) *FileDownloadUpdate {
	fdu.mutation.SetSourceType(s)
	return fdu
}

// SetSource sets the "source" field.
func (fdu *FileDownloadUpdate) SetSource(s string) *FileDownloadUpdate {
	fdu.mutation.SetSource(s)
	return fdu
}

// SetDestination sets the "destination" field.
func (fdu *FileDownloadUpdate) SetDestination(s string) *FileDownloadUpdate {
	fdu.mutation.SetDestination(s)
	return fdu
}

// SetTemplate sets the "template" field.
func (fdu *FileDownloadUpdate) SetTemplate(b bool) *FileDownloadUpdate {
	fdu.mutation.SetTemplate(b)
	return fdu
}

// SetPerms sets the "perms" field.
func (fdu *FileDownloadUpdate) SetPerms(s string) *FileDownloadUpdate {
	fdu.mutation.SetPerms(s)
	return fdu
}

// SetDisabled sets the "disabled" field.
func (fdu *FileDownloadUpdate) SetDisabled(b bool) *FileDownloadUpdate {
	fdu.mutation.SetDisabled(b)
	return fdu
}

// SetMd5 sets the "md5" field.
func (fdu *FileDownloadUpdate) SetMd5(s string) *FileDownloadUpdate {
	fdu.mutation.SetMd5(s)
	return fdu
}

// SetAbsPath sets the "abs_path" field.
func (fdu *FileDownloadUpdate) SetAbsPath(s string) *FileDownloadUpdate {
	fdu.mutation.SetAbsPath(s)
	return fdu
}

// SetTags sets the "tags" field.
func (fdu *FileDownloadUpdate) SetTags(m map[string]string) *FileDownloadUpdate {
	fdu.mutation.SetTags(m)
	return fdu
}

// AddFileDownloadToTagIDs adds the "FileDownloadToTag" edge to the Tag entity by IDs.
func (fdu *FileDownloadUpdate) AddFileDownloadToTagIDs(ids ...int) *FileDownloadUpdate {
	fdu.mutation.AddFileDownloadToTagIDs(ids...)
	return fdu
}

// AddFileDownloadToTag adds the "FileDownloadToTag" edges to the Tag entity.
func (fdu *FileDownloadUpdate) AddFileDownloadToTag(t ...*Tag) *FileDownloadUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fdu.AddFileDownloadToTagIDs(ids...)
}

// AddFileDownloadToEnvironmentIDs adds the "FileDownloadToEnvironment" edge to the Environment entity by IDs.
func (fdu *FileDownloadUpdate) AddFileDownloadToEnvironmentIDs(ids ...int) *FileDownloadUpdate {
	fdu.mutation.AddFileDownloadToEnvironmentIDs(ids...)
	return fdu
}

// AddFileDownloadToEnvironment adds the "FileDownloadToEnvironment" edges to the Environment entity.
func (fdu *FileDownloadUpdate) AddFileDownloadToEnvironment(e ...*Environment) *FileDownloadUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fdu.AddFileDownloadToEnvironmentIDs(ids...)
}

// Mutation returns the FileDownloadMutation object of the builder.
func (fdu *FileDownloadUpdate) Mutation() *FileDownloadMutation {
	return fdu.mutation
}

// ClearFileDownloadToTag clears all "FileDownloadToTag" edges to the Tag entity.
func (fdu *FileDownloadUpdate) ClearFileDownloadToTag() *FileDownloadUpdate {
	fdu.mutation.ClearFileDownloadToTag()
	return fdu
}

// RemoveFileDownloadToTagIDs removes the "FileDownloadToTag" edge to Tag entities by IDs.
func (fdu *FileDownloadUpdate) RemoveFileDownloadToTagIDs(ids ...int) *FileDownloadUpdate {
	fdu.mutation.RemoveFileDownloadToTagIDs(ids...)
	return fdu
}

// RemoveFileDownloadToTag removes "FileDownloadToTag" edges to Tag entities.
func (fdu *FileDownloadUpdate) RemoveFileDownloadToTag(t ...*Tag) *FileDownloadUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fdu.RemoveFileDownloadToTagIDs(ids...)
}

// ClearFileDownloadToEnvironment clears all "FileDownloadToEnvironment" edges to the Environment entity.
func (fdu *FileDownloadUpdate) ClearFileDownloadToEnvironment() *FileDownloadUpdate {
	fdu.mutation.ClearFileDownloadToEnvironment()
	return fdu
}

// RemoveFileDownloadToEnvironmentIDs removes the "FileDownloadToEnvironment" edge to Environment entities by IDs.
func (fdu *FileDownloadUpdate) RemoveFileDownloadToEnvironmentIDs(ids ...int) *FileDownloadUpdate {
	fdu.mutation.RemoveFileDownloadToEnvironmentIDs(ids...)
	return fdu
}

// RemoveFileDownloadToEnvironment removes "FileDownloadToEnvironment" edges to Environment entities.
func (fdu *FileDownloadUpdate) RemoveFileDownloadToEnvironment(e ...*Environment) *FileDownloadUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fdu.RemoveFileDownloadToEnvironmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fdu *FileDownloadUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fdu.hooks) == 0 {
		affected, err = fdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileDownloadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fdu.mutation = mutation
			affected, err = fdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fdu.hooks) - 1; i >= 0; i-- {
			mut = fdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fdu *FileDownloadUpdate) SaveX(ctx context.Context) int {
	affected, err := fdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fdu *FileDownloadUpdate) Exec(ctx context.Context) error {
	_, err := fdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fdu *FileDownloadUpdate) ExecX(ctx context.Context) {
	if err := fdu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fdu *FileDownloadUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filedownload.Table,
			Columns: filedownload.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filedownload.FieldID,
			},
		},
	}
	if ps := fdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fdu.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldHclID,
		})
	}
	if value, ok := fdu.mutation.SourceType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldSourceType,
		})
	}
	if value, ok := fdu.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldSource,
		})
	}
	if value, ok := fdu.mutation.Destination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldDestination,
		})
	}
	if value, ok := fdu.mutation.Template(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: filedownload.FieldTemplate,
		})
	}
	if value, ok := fdu.mutation.Perms(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldPerms,
		})
	}
	if value, ok := fdu.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: filedownload.FieldDisabled,
		})
	}
	if value, ok := fdu.mutation.Md5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldMd5,
		})
	}
	if value, ok := fdu.mutation.AbsPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldAbsPath,
		})
	}
	if value, ok := fdu.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: filedownload.FieldTags,
		})
	}
	if fdu.mutation.FileDownloadToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedownload.FileDownloadToTagTable,
			Columns: []string{filedownload.FileDownloadToTagColumn},
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
	if nodes := fdu.mutation.RemovedFileDownloadToTagIDs(); len(nodes) > 0 && !fdu.mutation.FileDownloadToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedownload.FileDownloadToTagTable,
			Columns: []string{filedownload.FileDownloadToTagColumn},
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
	if nodes := fdu.mutation.FileDownloadToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedownload.FileDownloadToTagTable,
			Columns: []string{filedownload.FileDownloadToTagColumn},
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
	if fdu.mutation.FileDownloadToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   filedownload.FileDownloadToEnvironmentTable,
			Columns: filedownload.FileDownloadToEnvironmentPrimaryKey,
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
	if nodes := fdu.mutation.RemovedFileDownloadToEnvironmentIDs(); len(nodes) > 0 && !fdu.mutation.FileDownloadToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   filedownload.FileDownloadToEnvironmentTable,
			Columns: filedownload.FileDownloadToEnvironmentPrimaryKey,
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
	if nodes := fdu.mutation.FileDownloadToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   filedownload.FileDownloadToEnvironmentTable,
			Columns: filedownload.FileDownloadToEnvironmentPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, fdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filedownload.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// FileDownloadUpdateOne is the builder for updating a single FileDownload entity.
type FileDownloadUpdateOne struct {
	config
	hooks    []Hook
	mutation *FileDownloadMutation
}

// SetHclID sets the "hcl_id" field.
func (fduo *FileDownloadUpdateOne) SetHclID(s string) *FileDownloadUpdateOne {
	fduo.mutation.SetHclID(s)
	return fduo
}

// SetSourceType sets the "source_type" field.
func (fduo *FileDownloadUpdateOne) SetSourceType(s string) *FileDownloadUpdateOne {
	fduo.mutation.SetSourceType(s)
	return fduo
}

// SetSource sets the "source" field.
func (fduo *FileDownloadUpdateOne) SetSource(s string) *FileDownloadUpdateOne {
	fduo.mutation.SetSource(s)
	return fduo
}

// SetDestination sets the "destination" field.
func (fduo *FileDownloadUpdateOne) SetDestination(s string) *FileDownloadUpdateOne {
	fduo.mutation.SetDestination(s)
	return fduo
}

// SetTemplate sets the "template" field.
func (fduo *FileDownloadUpdateOne) SetTemplate(b bool) *FileDownloadUpdateOne {
	fduo.mutation.SetTemplate(b)
	return fduo
}

// SetPerms sets the "perms" field.
func (fduo *FileDownloadUpdateOne) SetPerms(s string) *FileDownloadUpdateOne {
	fduo.mutation.SetPerms(s)
	return fduo
}

// SetDisabled sets the "disabled" field.
func (fduo *FileDownloadUpdateOne) SetDisabled(b bool) *FileDownloadUpdateOne {
	fduo.mutation.SetDisabled(b)
	return fduo
}

// SetMd5 sets the "md5" field.
func (fduo *FileDownloadUpdateOne) SetMd5(s string) *FileDownloadUpdateOne {
	fduo.mutation.SetMd5(s)
	return fduo
}

// SetAbsPath sets the "abs_path" field.
func (fduo *FileDownloadUpdateOne) SetAbsPath(s string) *FileDownloadUpdateOne {
	fduo.mutation.SetAbsPath(s)
	return fduo
}

// SetTags sets the "tags" field.
func (fduo *FileDownloadUpdateOne) SetTags(m map[string]string) *FileDownloadUpdateOne {
	fduo.mutation.SetTags(m)
	return fduo
}

// AddFileDownloadToTagIDs adds the "FileDownloadToTag" edge to the Tag entity by IDs.
func (fduo *FileDownloadUpdateOne) AddFileDownloadToTagIDs(ids ...int) *FileDownloadUpdateOne {
	fduo.mutation.AddFileDownloadToTagIDs(ids...)
	return fduo
}

// AddFileDownloadToTag adds the "FileDownloadToTag" edges to the Tag entity.
func (fduo *FileDownloadUpdateOne) AddFileDownloadToTag(t ...*Tag) *FileDownloadUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fduo.AddFileDownloadToTagIDs(ids...)
}

// AddFileDownloadToEnvironmentIDs adds the "FileDownloadToEnvironment" edge to the Environment entity by IDs.
func (fduo *FileDownloadUpdateOne) AddFileDownloadToEnvironmentIDs(ids ...int) *FileDownloadUpdateOne {
	fduo.mutation.AddFileDownloadToEnvironmentIDs(ids...)
	return fduo
}

// AddFileDownloadToEnvironment adds the "FileDownloadToEnvironment" edges to the Environment entity.
func (fduo *FileDownloadUpdateOne) AddFileDownloadToEnvironment(e ...*Environment) *FileDownloadUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fduo.AddFileDownloadToEnvironmentIDs(ids...)
}

// Mutation returns the FileDownloadMutation object of the builder.
func (fduo *FileDownloadUpdateOne) Mutation() *FileDownloadMutation {
	return fduo.mutation
}

// ClearFileDownloadToTag clears all "FileDownloadToTag" edges to the Tag entity.
func (fduo *FileDownloadUpdateOne) ClearFileDownloadToTag() *FileDownloadUpdateOne {
	fduo.mutation.ClearFileDownloadToTag()
	return fduo
}

// RemoveFileDownloadToTagIDs removes the "FileDownloadToTag" edge to Tag entities by IDs.
func (fduo *FileDownloadUpdateOne) RemoveFileDownloadToTagIDs(ids ...int) *FileDownloadUpdateOne {
	fduo.mutation.RemoveFileDownloadToTagIDs(ids...)
	return fduo
}

// RemoveFileDownloadToTag removes "FileDownloadToTag" edges to Tag entities.
func (fduo *FileDownloadUpdateOne) RemoveFileDownloadToTag(t ...*Tag) *FileDownloadUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return fduo.RemoveFileDownloadToTagIDs(ids...)
}

// ClearFileDownloadToEnvironment clears all "FileDownloadToEnvironment" edges to the Environment entity.
func (fduo *FileDownloadUpdateOne) ClearFileDownloadToEnvironment() *FileDownloadUpdateOne {
	fduo.mutation.ClearFileDownloadToEnvironment()
	return fduo
}

// RemoveFileDownloadToEnvironmentIDs removes the "FileDownloadToEnvironment" edge to Environment entities by IDs.
func (fduo *FileDownloadUpdateOne) RemoveFileDownloadToEnvironmentIDs(ids ...int) *FileDownloadUpdateOne {
	fduo.mutation.RemoveFileDownloadToEnvironmentIDs(ids...)
	return fduo
}

// RemoveFileDownloadToEnvironment removes "FileDownloadToEnvironment" edges to Environment entities.
func (fduo *FileDownloadUpdateOne) RemoveFileDownloadToEnvironment(e ...*Environment) *FileDownloadUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fduo.RemoveFileDownloadToEnvironmentIDs(ids...)
}

// Save executes the query and returns the updated FileDownload entity.
func (fduo *FileDownloadUpdateOne) Save(ctx context.Context) (*FileDownload, error) {
	var (
		err  error
		node *FileDownload
	)
	if len(fduo.hooks) == 0 {
		node, err = fduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileDownloadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fduo.mutation = mutation
			node, err = fduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fduo.hooks) - 1; i >= 0; i-- {
			mut = fduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fduo *FileDownloadUpdateOne) SaveX(ctx context.Context) *FileDownload {
	node, err := fduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fduo *FileDownloadUpdateOne) Exec(ctx context.Context) error {
	_, err := fduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fduo *FileDownloadUpdateOne) ExecX(ctx context.Context) {
	if err := fduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fduo *FileDownloadUpdateOne) sqlSave(ctx context.Context) (_node *FileDownload, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   filedownload.Table,
			Columns: filedownload.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: filedownload.FieldID,
			},
		},
	}
	id, ok := fduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing FileDownload.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := fduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fduo.mutation.HclID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldHclID,
		})
	}
	if value, ok := fduo.mutation.SourceType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldSourceType,
		})
	}
	if value, ok := fduo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldSource,
		})
	}
	if value, ok := fduo.mutation.Destination(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldDestination,
		})
	}
	if value, ok := fduo.mutation.Template(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: filedownload.FieldTemplate,
		})
	}
	if value, ok := fduo.mutation.Perms(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldPerms,
		})
	}
	if value, ok := fduo.mutation.Disabled(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: filedownload.FieldDisabled,
		})
	}
	if value, ok := fduo.mutation.Md5(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldMd5,
		})
	}
	if value, ok := fduo.mutation.AbsPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: filedownload.FieldAbsPath,
		})
	}
	if value, ok := fduo.mutation.Tags(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: filedownload.FieldTags,
		})
	}
	if fduo.mutation.FileDownloadToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedownload.FileDownloadToTagTable,
			Columns: []string{filedownload.FileDownloadToTagColumn},
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
	if nodes := fduo.mutation.RemovedFileDownloadToTagIDs(); len(nodes) > 0 && !fduo.mutation.FileDownloadToTagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedownload.FileDownloadToTagTable,
			Columns: []string{filedownload.FileDownloadToTagColumn},
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
	if nodes := fduo.mutation.FileDownloadToTagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filedownload.FileDownloadToTagTable,
			Columns: []string{filedownload.FileDownloadToTagColumn},
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
	if fduo.mutation.FileDownloadToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   filedownload.FileDownloadToEnvironmentTable,
			Columns: filedownload.FileDownloadToEnvironmentPrimaryKey,
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
	if nodes := fduo.mutation.RemovedFileDownloadToEnvironmentIDs(); len(nodes) > 0 && !fduo.mutation.FileDownloadToEnvironmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   filedownload.FileDownloadToEnvironmentTable,
			Columns: filedownload.FileDownloadToEnvironmentPrimaryKey,
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
	if nodes := fduo.mutation.FileDownloadToEnvironmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   filedownload.FileDownloadToEnvironmentTable,
			Columns: filedownload.FileDownloadToEnvironmentPrimaryKey,
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
	_node = &FileDownload{config: fduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filedownload.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
