// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gen0cide/laforge/ent/agenttask"
	"github.com/gen0cide/laforge/ent/ansible"
	"github.com/gen0cide/laforge/ent/command"
	"github.com/gen0cide/laforge/ent/dnsrecord"
	"github.com/gen0cide/laforge/ent/filedelete"
	"github.com/gen0cide/laforge/ent/filedownload"
	"github.com/gen0cide/laforge/ent/fileextract"
	"github.com/gen0cide/laforge/ent/ginfilemiddleware"
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisioningstep"
	"github.com/gen0cide/laforge/ent/script"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/google/uuid"
)

// ProvisioningStepCreate is the builder for creating a ProvisioningStep entity.
type ProvisioningStepCreate struct {
	config
	mutation *ProvisioningStepMutation
	hooks    []Hook
}

// SetType sets the "type" field.
func (psc *ProvisioningStepCreate) SetType(pr provisioningstep.Type) *ProvisioningStepCreate {
	psc.mutation.SetType(pr)
	return psc
}

// SetStepNumber sets the "step_number" field.
func (psc *ProvisioningStepCreate) SetStepNumber(i int) *ProvisioningStepCreate {
	psc.mutation.SetStepNumber(i)
	return psc
}

// SetID sets the "id" field.
func (psc *ProvisioningStepCreate) SetID(u uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetID(u)
	return psc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableID(u *uuid.UUID) *ProvisioningStepCreate {
	if u != nil {
		psc.SetID(*u)
	}
	return psc
}

// SetProvisioningStepToStatusID sets the "ProvisioningStepToStatus" edge to the Status entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToStatusID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToStatusID(id)
	return psc
}

// SetNillableProvisioningStepToStatusID sets the "ProvisioningStepToStatus" edge to the Status entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToStatusID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToStatusID(*id)
	}
	return psc
}

// SetProvisioningStepToStatus sets the "ProvisioningStepToStatus" edge to the Status entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToStatus(s *Status) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToStatusID(s.ID)
}

// SetProvisioningStepToProvisionedHostID sets the "ProvisioningStepToProvisionedHost" edge to the ProvisionedHost entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToProvisionedHostID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToProvisionedHostID(id)
	return psc
}

// SetNillableProvisioningStepToProvisionedHostID sets the "ProvisioningStepToProvisionedHost" edge to the ProvisionedHost entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToProvisionedHostID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToProvisionedHostID(*id)
	}
	return psc
}

// SetProvisioningStepToProvisionedHost sets the "ProvisioningStepToProvisionedHost" edge to the ProvisionedHost entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToProvisionedHost(p *ProvisionedHost) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToProvisionedHostID(p.ID)
}

// SetProvisioningStepToScriptID sets the "ProvisioningStepToScript" edge to the Script entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToScriptID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToScriptID(id)
	return psc
}

// SetNillableProvisioningStepToScriptID sets the "ProvisioningStepToScript" edge to the Script entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToScriptID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToScriptID(*id)
	}
	return psc
}

// SetProvisioningStepToScript sets the "ProvisioningStepToScript" edge to the Script entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToScript(s *Script) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToScriptID(s.ID)
}

// SetProvisioningStepToCommandID sets the "ProvisioningStepToCommand" edge to the Command entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToCommandID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToCommandID(id)
	return psc
}

// SetNillableProvisioningStepToCommandID sets the "ProvisioningStepToCommand" edge to the Command entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToCommandID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToCommandID(*id)
	}
	return psc
}

// SetProvisioningStepToCommand sets the "ProvisioningStepToCommand" edge to the Command entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToCommand(c *Command) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToCommandID(c.ID)
}

// SetProvisioningStepToDNSRecordID sets the "ProvisioningStepToDNSRecord" edge to the DNSRecord entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToDNSRecordID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToDNSRecordID(id)
	return psc
}

// SetNillableProvisioningStepToDNSRecordID sets the "ProvisioningStepToDNSRecord" edge to the DNSRecord entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToDNSRecordID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToDNSRecordID(*id)
	}
	return psc
}

// SetProvisioningStepToDNSRecord sets the "ProvisioningStepToDNSRecord" edge to the DNSRecord entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToDNSRecord(d *DNSRecord) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToDNSRecordID(d.ID)
}

// SetProvisioningStepToFileDeleteID sets the "ProvisioningStepToFileDelete" edge to the FileDelete entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToFileDeleteID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToFileDeleteID(id)
	return psc
}

// SetNillableProvisioningStepToFileDeleteID sets the "ProvisioningStepToFileDelete" edge to the FileDelete entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToFileDeleteID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToFileDeleteID(*id)
	}
	return psc
}

// SetProvisioningStepToFileDelete sets the "ProvisioningStepToFileDelete" edge to the FileDelete entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToFileDelete(f *FileDelete) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToFileDeleteID(f.ID)
}

// SetProvisioningStepToFileDownloadID sets the "ProvisioningStepToFileDownload" edge to the FileDownload entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToFileDownloadID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToFileDownloadID(id)
	return psc
}

// SetNillableProvisioningStepToFileDownloadID sets the "ProvisioningStepToFileDownload" edge to the FileDownload entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToFileDownloadID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToFileDownloadID(*id)
	}
	return psc
}

// SetProvisioningStepToFileDownload sets the "ProvisioningStepToFileDownload" edge to the FileDownload entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToFileDownload(f *FileDownload) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToFileDownloadID(f.ID)
}

// SetProvisioningStepToFileExtractID sets the "ProvisioningStepToFileExtract" edge to the FileExtract entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToFileExtractID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToFileExtractID(id)
	return psc
}

// SetNillableProvisioningStepToFileExtractID sets the "ProvisioningStepToFileExtract" edge to the FileExtract entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToFileExtractID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToFileExtractID(*id)
	}
	return psc
}

// SetProvisioningStepToFileExtract sets the "ProvisioningStepToFileExtract" edge to the FileExtract entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToFileExtract(f *FileExtract) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToFileExtractID(f.ID)
}

// SetProvisioningStepToAnsibleID sets the "ProvisioningStepToAnsible" edge to the Ansible entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToAnsibleID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToAnsibleID(id)
	return psc
}

// SetNillableProvisioningStepToAnsibleID sets the "ProvisioningStepToAnsible" edge to the Ansible entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToAnsibleID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToAnsibleID(*id)
	}
	return psc
}

// SetProvisioningStepToAnsible sets the "ProvisioningStepToAnsible" edge to the Ansible entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToAnsible(a *Ansible) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToAnsibleID(a.ID)
}

// SetProvisioningStepToPlanID sets the "ProvisioningStepToPlan" edge to the Plan entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToPlanID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToPlanID(id)
	return psc
}

// SetNillableProvisioningStepToPlanID sets the "ProvisioningStepToPlan" edge to the Plan entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToPlanID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToPlanID(*id)
	}
	return psc
}

// SetProvisioningStepToPlan sets the "ProvisioningStepToPlan" edge to the Plan entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToPlan(p *Plan) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToPlanID(p.ID)
}

// AddProvisioningStepToAgentTaskIDs adds the "ProvisioningStepToAgentTask" edge to the AgentTask entity by IDs.
func (psc *ProvisioningStepCreate) AddProvisioningStepToAgentTaskIDs(ids ...uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.AddProvisioningStepToAgentTaskIDs(ids...)
	return psc
}

// AddProvisioningStepToAgentTask adds the "ProvisioningStepToAgentTask" edges to the AgentTask entity.
func (psc *ProvisioningStepCreate) AddProvisioningStepToAgentTask(a ...*AgentTask) *ProvisioningStepCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return psc.AddProvisioningStepToAgentTaskIDs(ids...)
}

// SetProvisioningStepToGinFileMiddlewareID sets the "ProvisioningStepToGinFileMiddleware" edge to the GinFileMiddleware entity by ID.
func (psc *ProvisioningStepCreate) SetProvisioningStepToGinFileMiddlewareID(id uuid.UUID) *ProvisioningStepCreate {
	psc.mutation.SetProvisioningStepToGinFileMiddlewareID(id)
	return psc
}

// SetNillableProvisioningStepToGinFileMiddlewareID sets the "ProvisioningStepToGinFileMiddleware" edge to the GinFileMiddleware entity by ID if the given value is not nil.
func (psc *ProvisioningStepCreate) SetNillableProvisioningStepToGinFileMiddlewareID(id *uuid.UUID) *ProvisioningStepCreate {
	if id != nil {
		psc = psc.SetProvisioningStepToGinFileMiddlewareID(*id)
	}
	return psc
}

// SetProvisioningStepToGinFileMiddleware sets the "ProvisioningStepToGinFileMiddleware" edge to the GinFileMiddleware entity.
func (psc *ProvisioningStepCreate) SetProvisioningStepToGinFileMiddleware(g *GinFileMiddleware) *ProvisioningStepCreate {
	return psc.SetProvisioningStepToGinFileMiddlewareID(g.ID)
}

// Mutation returns the ProvisioningStepMutation object of the builder.
func (psc *ProvisioningStepCreate) Mutation() *ProvisioningStepMutation {
	return psc.mutation
}

// Save creates the ProvisioningStep in the database.
func (psc *ProvisioningStepCreate) Save(ctx context.Context) (*ProvisioningStep, error) {
	var (
		err  error
		node *ProvisioningStep
	)
	psc.defaults()
	if len(psc.hooks) == 0 {
		if err = psc.check(); err != nil {
			return nil, err
		}
		node, err = psc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProvisioningStepMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = psc.check(); err != nil {
				return nil, err
			}
			psc.mutation = mutation
			if node, err = psc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(psc.hooks) - 1; i >= 0; i-- {
			if psc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = psc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, psc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (psc *ProvisioningStepCreate) SaveX(ctx context.Context) *ProvisioningStep {
	v, err := psc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psc *ProvisioningStepCreate) Exec(ctx context.Context) error {
	_, err := psc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psc *ProvisioningStepCreate) ExecX(ctx context.Context) {
	if err := psc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psc *ProvisioningStepCreate) defaults() {
	if _, ok := psc.mutation.ID(); !ok {
		v := provisioningstep.DefaultID()
		psc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psc *ProvisioningStepCreate) check() error {
	if _, ok := psc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "ProvisioningStep.type"`)}
	}
	if v, ok := psc.mutation.GetType(); ok {
		if err := provisioningstep.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "ProvisioningStep.type": %w`, err)}
		}
	}
	if _, ok := psc.mutation.StepNumber(); !ok {
		return &ValidationError{Name: "step_number", err: errors.New(`ent: missing required field "ProvisioningStep.step_number"`)}
	}
	return nil
}

func (psc *ProvisioningStepCreate) sqlSave(ctx context.Context) (*ProvisioningStep, error) {
	_node, _spec := psc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
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

func (psc *ProvisioningStepCreate) createSpec() (*ProvisioningStep, *sqlgraph.CreateSpec) {
	var (
		_node = &ProvisioningStep{config: psc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: provisioningstep.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: provisioningstep.FieldID,
			},
		}
	)
	if id, ok := psc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := psc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: provisioningstep.FieldType,
		})
		_node.Type = value
	}
	if value, ok := psc.mutation.StepNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: provisioningstep.FieldStepNumber,
		})
		_node.StepNumber = value
	}
	if nodes := psc.mutation.ProvisioningStepToStatusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   provisioningstep.ProvisioningStepToStatusTable,
			Columns: []string{provisioningstep.ProvisioningStepToStatusColumn},
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
	if nodes := psc.mutation.ProvisioningStepToProvisionedHostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisioningstep.ProvisioningStepToProvisionedHostTable,
			Columns: []string{provisioningstep.ProvisioningStepToProvisionedHostColumn},
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
		_node.provisioning_step_provisioning_step_to_provisioned_host = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ProvisioningStepToScriptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisioningstep.ProvisioningStepToScriptTable,
			Columns: []string{provisioningstep.ProvisioningStepToScriptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: script.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioning_step_provisioning_step_to_script = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ProvisioningStepToCommandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisioningstep.ProvisioningStepToCommandTable,
			Columns: []string{provisioningstep.ProvisioningStepToCommandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: command.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioning_step_provisioning_step_to_command = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ProvisioningStepToDNSRecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisioningstep.ProvisioningStepToDNSRecordTable,
			Columns: []string{provisioningstep.ProvisioningStepToDNSRecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: dnsrecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioning_step_provisioning_step_to_dns_record = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ProvisioningStepToFileDeleteIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisioningstep.ProvisioningStepToFileDeleteTable,
			Columns: []string{provisioningstep.ProvisioningStepToFileDeleteColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: filedelete.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioning_step_provisioning_step_to_file_delete = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ProvisioningStepToFileDownloadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisioningstep.ProvisioningStepToFileDownloadTable,
			Columns: []string{provisioningstep.ProvisioningStepToFileDownloadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: filedownload.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioning_step_provisioning_step_to_file_download = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ProvisioningStepToFileExtractIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisioningstep.ProvisioningStepToFileExtractTable,
			Columns: []string{provisioningstep.ProvisioningStepToFileExtractColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: fileextract.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioning_step_provisioning_step_to_file_extract = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ProvisioningStepToAnsibleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   provisioningstep.ProvisioningStepToAnsibleTable,
			Columns: []string{provisioningstep.ProvisioningStepToAnsibleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ansible.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.provisioning_step_provisioning_step_to_ansible = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ProvisioningStepToPlanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   provisioningstep.ProvisioningStepToPlanTable,
			Columns: []string{provisioningstep.ProvisioningStepToPlanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: plan.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.plan_plan_to_provisioning_step = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ProvisioningStepToAgentTaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   provisioningstep.ProvisioningStepToAgentTaskTable,
			Columns: []string{provisioningstep.ProvisioningStepToAgentTaskColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psc.mutation.ProvisioningStepToGinFileMiddlewareIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   provisioningstep.ProvisioningStepToGinFileMiddlewareTable,
			Columns: []string{provisioningstep.ProvisioningStepToGinFileMiddlewareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ginfilemiddleware.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.gin_file_middleware_gin_file_middleware_to_provisioning_step = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProvisioningStepCreateBulk is the builder for creating many ProvisioningStep entities in bulk.
type ProvisioningStepCreateBulk struct {
	config
	builders []*ProvisioningStepCreate
}

// Save creates the ProvisioningStep entities in the database.
func (pscb *ProvisioningStepCreateBulk) Save(ctx context.Context) ([]*ProvisioningStep, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pscb.builders))
	nodes := make([]*ProvisioningStep, len(pscb.builders))
	mutators := make([]Mutator, len(pscb.builders))
	for i := range pscb.builders {
		func(i int, root context.Context) {
			builder := pscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProvisioningStepMutation)
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
					_, err = mutators[i+1].Mutate(root, pscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
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
		if _, err := mutators[0].Mutate(ctx, pscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pscb *ProvisioningStepCreateBulk) SaveX(ctx context.Context) []*ProvisioningStep {
	v, err := pscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pscb *ProvisioningStepCreateBulk) Exec(ctx context.Context) error {
	_, err := pscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pscb *ProvisioningStepCreateBulk) ExecX(ctx context.Context) {
	if err := pscb.Exec(ctx); err != nil {
		panic(err)
	}
}
