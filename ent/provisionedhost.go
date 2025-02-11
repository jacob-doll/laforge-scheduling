// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/ginfilemiddleware"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/gen0cide/laforge/ent/plan"
	"github.com/gen0cide/laforge/ent/provisionedhost"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/google/uuid"
)

// ProvisionedHost is the model entity for the ProvisionedHost schema.
type ProvisionedHost struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SubnetIP holds the value of the "subnet_ip" field.
	SubnetIP string `json:"subnet_ip,omitempty"`
	// AddonType holds the value of the "addon_type" field.
	AddonType *provisionedhost.AddonType `json:"addon_type,omitempty"`
	// Vars holds the value of the "vars" field.
	Vars map[string]string `json:"vars,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProvisionedHostQuery when eager-loading is set.
	Edges ProvisionedHostEdges `json:"edges"`

	// Edges put into the main struct to be loaded via hcl
	// ProvisionedHostToStatus holds the value of the ProvisionedHostToStatus edge.
	HCLProvisionedHostToStatus *Status `json:"ProvisionedHostToStatus,omitempty"`
	// ProvisionedHostToProvisionedNetwork holds the value of the ProvisionedHostToProvisionedNetwork edge.
	HCLProvisionedHostToProvisionedNetwork *ProvisionedNetwork `json:"ProvisionedHostToProvisionedNetwork,omitempty"`
	// ProvisionedHostToHost holds the value of the ProvisionedHostToHost edge.
	HCLProvisionedHostToHost *Host `json:"ProvisionedHostToHost,omitempty"`
	// ProvisionedHostToEndStepPlan holds the value of the ProvisionedHostToEndStepPlan edge.
	HCLProvisionedHostToEndStepPlan *Plan `json:"ProvisionedHostToEndStepPlan,omitempty"`
	// ProvisionedHostToBuild holds the value of the ProvisionedHostToBuild edge.
	HCLProvisionedHostToBuild *Build `json:"ProvisionedHostToBuild,omitempty"`
	// ProvisionedHostToProvisionedScheduleStep holds the value of the ProvisionedHostToProvisionedScheduleStep edge.
	HCLProvisionedHostToProvisionedScheduleStep []*ProvisionedScheduleStep `json:"ProvisionedHostToProvisionedScheduleStep,omitempty"`
	// ProvisionedHostToProvisioningStep holds the value of the ProvisionedHostToProvisioningStep edge.
	HCLProvisionedHostToProvisioningStep []*ProvisioningStep `json:"ProvisionedHostToProvisioningStep,omitempty"`
	// ProvisionedHostToAgentStatus holds the value of the ProvisionedHostToAgentStatus edge.
	HCLProvisionedHostToAgentStatus []*AgentStatus `json:"ProvisionedHostToAgentStatus,omitempty"`
	// ProvisionedHostToAgentTask holds the value of the ProvisionedHostToAgentTask edge.
	HCLProvisionedHostToAgentTask []*AgentTask `json:"ProvisionedHostToAgentTask,omitempty"`
	// ProvisionedHostToPlan holds the value of the ProvisionedHostToPlan edge.
	HCLProvisionedHostToPlan *Plan `json:"ProvisionedHostToPlan,omitempty"`
	// ProvisionedHostToGinFileMiddleware holds the value of the ProvisionedHostToGinFileMiddleware edge.
	HCLProvisionedHostToGinFileMiddleware *GinFileMiddleware `json:"ProvisionedHostToGinFileMiddleware,omitempty"`
	//
	gin_file_middleware_gin_file_middleware_to_provisioned_host *uuid.UUID
	plan_plan_to_provisioned_host                               *uuid.UUID
	provisioned_host_provisioned_host_to_provisioned_network    *uuid.UUID
	provisioned_host_provisioned_host_to_host                   *uuid.UUID
	provisioned_host_provisioned_host_to_end_step_plan          *uuid.UUID
	provisioned_host_provisioned_host_to_build                  *uuid.UUID
}

// ProvisionedHostEdges holds the relations/edges for other nodes in the graph.
type ProvisionedHostEdges struct {
	// ProvisionedHostToStatus holds the value of the ProvisionedHostToStatus edge.
	ProvisionedHostToStatus *Status `json:"ProvisionedHostToStatus,omitempty"`
	// ProvisionedHostToProvisionedNetwork holds the value of the ProvisionedHostToProvisionedNetwork edge.
	ProvisionedHostToProvisionedNetwork *ProvisionedNetwork `json:"ProvisionedHostToProvisionedNetwork,omitempty"`
	// ProvisionedHostToHost holds the value of the ProvisionedHostToHost edge.
	ProvisionedHostToHost *Host `json:"ProvisionedHostToHost,omitempty"`
	// ProvisionedHostToEndStepPlan holds the value of the ProvisionedHostToEndStepPlan edge.
	ProvisionedHostToEndStepPlan *Plan `json:"ProvisionedHostToEndStepPlan,omitempty"`
	// ProvisionedHostToBuild holds the value of the ProvisionedHostToBuild edge.
	ProvisionedHostToBuild *Build `json:"ProvisionedHostToBuild,omitempty"`
	// ProvisionedHostToProvisionedScheduleStep holds the value of the ProvisionedHostToProvisionedScheduleStep edge.
	ProvisionedHostToProvisionedScheduleStep []*ProvisionedScheduleStep `json:"ProvisionedHostToProvisionedScheduleStep,omitempty"`
	// ProvisionedHostToProvisioningStep holds the value of the ProvisionedHostToProvisioningStep edge.
	ProvisionedHostToProvisioningStep []*ProvisioningStep `json:"ProvisionedHostToProvisioningStep,omitempty"`
	// ProvisionedHostToAgentStatus holds the value of the ProvisionedHostToAgentStatus edge.
	ProvisionedHostToAgentStatus []*AgentStatus `json:"ProvisionedHostToAgentStatus,omitempty"`
	// ProvisionedHostToAgentTask holds the value of the ProvisionedHostToAgentTask edge.
	ProvisionedHostToAgentTask []*AgentTask `json:"ProvisionedHostToAgentTask,omitempty"`
	// ProvisionedHostToPlan holds the value of the ProvisionedHostToPlan edge.
	ProvisionedHostToPlan *Plan `json:"ProvisionedHostToPlan,omitempty"`
	// ProvisionedHostToGinFileMiddleware holds the value of the ProvisionedHostToGinFileMiddleware edge.
	ProvisionedHostToGinFileMiddleware *GinFileMiddleware `json:"ProvisionedHostToGinFileMiddleware,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [11]bool
}

// ProvisionedHostToStatusOrErr returns the ProvisionedHostToStatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProvisionedHostEdges) ProvisionedHostToStatusOrErr() (*Status, error) {
	if e.loadedTypes[0] {
		if e.ProvisionedHostToStatus == nil {
			// The edge ProvisionedHostToStatus was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: status.Label}
		}
		return e.ProvisionedHostToStatus, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToStatus"}
}

// ProvisionedHostToProvisionedNetworkOrErr returns the ProvisionedHostToProvisionedNetwork value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProvisionedHostEdges) ProvisionedHostToProvisionedNetworkOrErr() (*ProvisionedNetwork, error) {
	if e.loadedTypes[1] {
		if e.ProvisionedHostToProvisionedNetwork == nil {
			// The edge ProvisionedHostToProvisionedNetwork was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: provisionednetwork.Label}
		}
		return e.ProvisionedHostToProvisionedNetwork, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToProvisionedNetwork"}
}

// ProvisionedHostToHostOrErr returns the ProvisionedHostToHost value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProvisionedHostEdges) ProvisionedHostToHostOrErr() (*Host, error) {
	if e.loadedTypes[2] {
		if e.ProvisionedHostToHost == nil {
			// The edge ProvisionedHostToHost was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: host.Label}
		}
		return e.ProvisionedHostToHost, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToHost"}
}

// ProvisionedHostToEndStepPlanOrErr returns the ProvisionedHostToEndStepPlan value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProvisionedHostEdges) ProvisionedHostToEndStepPlanOrErr() (*Plan, error) {
	if e.loadedTypes[3] {
		if e.ProvisionedHostToEndStepPlan == nil {
			// The edge ProvisionedHostToEndStepPlan was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: plan.Label}
		}
		return e.ProvisionedHostToEndStepPlan, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToEndStepPlan"}
}

// ProvisionedHostToBuildOrErr returns the ProvisionedHostToBuild value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProvisionedHostEdges) ProvisionedHostToBuildOrErr() (*Build, error) {
	if e.loadedTypes[4] {
		if e.ProvisionedHostToBuild == nil {
			// The edge ProvisionedHostToBuild was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: build.Label}
		}
		return e.ProvisionedHostToBuild, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToBuild"}
}

// ProvisionedHostToProvisionedScheduleStepOrErr returns the ProvisionedHostToProvisionedScheduleStep value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedHostEdges) ProvisionedHostToProvisionedScheduleStepOrErr() ([]*ProvisionedScheduleStep, error) {
	if e.loadedTypes[5] {
		return e.ProvisionedHostToProvisionedScheduleStep, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToProvisionedScheduleStep"}
}

// ProvisionedHostToProvisioningStepOrErr returns the ProvisionedHostToProvisioningStep value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedHostEdges) ProvisionedHostToProvisioningStepOrErr() ([]*ProvisioningStep, error) {
	if e.loadedTypes[6] {
		return e.ProvisionedHostToProvisioningStep, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToProvisioningStep"}
}

// ProvisionedHostToAgentStatusOrErr returns the ProvisionedHostToAgentStatus value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedHostEdges) ProvisionedHostToAgentStatusOrErr() ([]*AgentStatus, error) {
	if e.loadedTypes[7] {
		return e.ProvisionedHostToAgentStatus, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToAgentStatus"}
}

// ProvisionedHostToAgentTaskOrErr returns the ProvisionedHostToAgentTask value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedHostEdges) ProvisionedHostToAgentTaskOrErr() ([]*AgentTask, error) {
	if e.loadedTypes[8] {
		return e.ProvisionedHostToAgentTask, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToAgentTask"}
}

// ProvisionedHostToPlanOrErr returns the ProvisionedHostToPlan value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProvisionedHostEdges) ProvisionedHostToPlanOrErr() (*Plan, error) {
	if e.loadedTypes[9] {
		if e.ProvisionedHostToPlan == nil {
			// The edge ProvisionedHostToPlan was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: plan.Label}
		}
		return e.ProvisionedHostToPlan, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToPlan"}
}

// ProvisionedHostToGinFileMiddlewareOrErr returns the ProvisionedHostToGinFileMiddleware value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProvisionedHostEdges) ProvisionedHostToGinFileMiddlewareOrErr() (*GinFileMiddleware, error) {
	if e.loadedTypes[10] {
		if e.ProvisionedHostToGinFileMiddleware == nil {
			// The edge ProvisionedHostToGinFileMiddleware was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: ginfilemiddleware.Label}
		}
		return e.ProvisionedHostToGinFileMiddleware, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedHostToGinFileMiddleware"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProvisionedHost) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case provisionedhost.FieldVars:
			values[i] = new([]byte)
		case provisionedhost.FieldSubnetIP, provisionedhost.FieldAddonType:
			values[i] = new(sql.NullString)
		case provisionedhost.FieldID:
			values[i] = new(uuid.UUID)
		case provisionedhost.ForeignKeys[0]: // gin_file_middleware_gin_file_middleware_to_provisioned_host
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case provisionedhost.ForeignKeys[1]: // plan_plan_to_provisioned_host
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case provisionedhost.ForeignKeys[2]: // provisioned_host_provisioned_host_to_provisioned_network
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case provisionedhost.ForeignKeys[3]: // provisioned_host_provisioned_host_to_host
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case provisionedhost.ForeignKeys[4]: // provisioned_host_provisioned_host_to_end_step_plan
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case provisionedhost.ForeignKeys[5]: // provisioned_host_provisioned_host_to_build
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type ProvisionedHost", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProvisionedHost fields.
func (ph *ProvisionedHost) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case provisionedhost.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ph.ID = *value
			}
		case provisionedhost.FieldSubnetIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subnet_ip", values[i])
			} else if value.Valid {
				ph.SubnetIP = value.String
			}
		case provisionedhost.FieldAddonType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field addon_type", values[i])
			} else if value.Valid {
				ph.AddonType = new(provisionedhost.AddonType)
				*ph.AddonType = provisionedhost.AddonType(value.String)
			}
		case provisionedhost.FieldVars:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field vars", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ph.Vars); err != nil {
					return fmt.Errorf("unmarshal field vars: %w", err)
				}
			}
		case provisionedhost.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field gin_file_middleware_gin_file_middleware_to_provisioned_host", values[i])
			} else if value.Valid {
				ph.gin_file_middleware_gin_file_middleware_to_provisioned_host = new(uuid.UUID)
				*ph.gin_file_middleware_gin_file_middleware_to_provisioned_host = *value.S.(*uuid.UUID)
			}
		case provisionedhost.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field plan_plan_to_provisioned_host", values[i])
			} else if value.Valid {
				ph.plan_plan_to_provisioned_host = new(uuid.UUID)
				*ph.plan_plan_to_provisioned_host = *value.S.(*uuid.UUID)
			}
		case provisionedhost.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field provisioned_host_provisioned_host_to_provisioned_network", values[i])
			} else if value.Valid {
				ph.provisioned_host_provisioned_host_to_provisioned_network = new(uuid.UUID)
				*ph.provisioned_host_provisioned_host_to_provisioned_network = *value.S.(*uuid.UUID)
			}
		case provisionedhost.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field provisioned_host_provisioned_host_to_host", values[i])
			} else if value.Valid {
				ph.provisioned_host_provisioned_host_to_host = new(uuid.UUID)
				*ph.provisioned_host_provisioned_host_to_host = *value.S.(*uuid.UUID)
			}
		case provisionedhost.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field provisioned_host_provisioned_host_to_end_step_plan", values[i])
			} else if value.Valid {
				ph.provisioned_host_provisioned_host_to_end_step_plan = new(uuid.UUID)
				*ph.provisioned_host_provisioned_host_to_end_step_plan = *value.S.(*uuid.UUID)
			}
		case provisionedhost.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field provisioned_host_provisioned_host_to_build", values[i])
			} else if value.Valid {
				ph.provisioned_host_provisioned_host_to_build = new(uuid.UUID)
				*ph.provisioned_host_provisioned_host_to_build = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryProvisionedHostToStatus queries the "ProvisionedHostToStatus" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToStatus() *StatusQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToStatus(ph)
}

// QueryProvisionedHostToProvisionedNetwork queries the "ProvisionedHostToProvisionedNetwork" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToProvisionedNetwork() *ProvisionedNetworkQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToProvisionedNetwork(ph)
}

// QueryProvisionedHostToHost queries the "ProvisionedHostToHost" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToHost() *HostQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToHost(ph)
}

// QueryProvisionedHostToEndStepPlan queries the "ProvisionedHostToEndStepPlan" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToEndStepPlan() *PlanQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToEndStepPlan(ph)
}

// QueryProvisionedHostToBuild queries the "ProvisionedHostToBuild" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToBuild() *BuildQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToBuild(ph)
}

// QueryProvisionedHostToProvisionedScheduleStep queries the "ProvisionedHostToProvisionedScheduleStep" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToProvisionedScheduleStep() *ProvisionedScheduleStepQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToProvisionedScheduleStep(ph)
}

// QueryProvisionedHostToProvisioningStep queries the "ProvisionedHostToProvisioningStep" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToProvisioningStep() *ProvisioningStepQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToProvisioningStep(ph)
}

// QueryProvisionedHostToAgentStatus queries the "ProvisionedHostToAgentStatus" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToAgentStatus() *AgentStatusQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToAgentStatus(ph)
}

// QueryProvisionedHostToAgentTask queries the "ProvisionedHostToAgentTask" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToAgentTask() *AgentTaskQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToAgentTask(ph)
}

// QueryProvisionedHostToPlan queries the "ProvisionedHostToPlan" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToPlan() *PlanQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToPlan(ph)
}

// QueryProvisionedHostToGinFileMiddleware queries the "ProvisionedHostToGinFileMiddleware" edge of the ProvisionedHost entity.
func (ph *ProvisionedHost) QueryProvisionedHostToGinFileMiddleware() *GinFileMiddlewareQuery {
	return (&ProvisionedHostClient{config: ph.config}).QueryProvisionedHostToGinFileMiddleware(ph)
}

// Update returns a builder for updating this ProvisionedHost.
// Note that you need to call ProvisionedHost.Unwrap() before calling this method if this ProvisionedHost
// was returned from a transaction, and the transaction was committed or rolled back.
func (ph *ProvisionedHost) Update() *ProvisionedHostUpdateOne {
	return (&ProvisionedHostClient{config: ph.config}).UpdateOne(ph)
}

// Unwrap unwraps the ProvisionedHost entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ph *ProvisionedHost) Unwrap() *ProvisionedHost {
	tx, ok := ph.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProvisionedHost is not a transactional entity")
	}
	ph.config.driver = tx.drv
	return ph
}

// String implements the fmt.Stringer.
func (ph *ProvisionedHost) String() string {
	var builder strings.Builder
	builder.WriteString("ProvisionedHost(")
	builder.WriteString(fmt.Sprintf("id=%v", ph.ID))
	builder.WriteString(", subnet_ip=")
	builder.WriteString(ph.SubnetIP)
	if v := ph.AddonType; v != nil {
		builder.WriteString(", addon_type=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", vars=")
	builder.WriteString(fmt.Sprintf("%v", ph.Vars))
	builder.WriteByte(')')
	return builder.String()
}

// ProvisionedHosts is a parsable slice of ProvisionedHost.
type ProvisionedHosts []*ProvisionedHost

func (ph ProvisionedHosts) config(cfg config) {
	for _i := range ph {
		ph[_i].config = cfg
	}
}
