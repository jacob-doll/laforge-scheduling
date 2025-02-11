// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/adhocplan"
	"github.com/gen0cide/laforge/ent/agenttask"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/google/uuid"
)

// AdhocPlan is the model entity for the AdhocPlan schema.
type AdhocPlan struct {
	config
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AdhocPlanQuery when eager-loading is set.
	Edges AdhocPlanEdges `json:"edges"`

	// Edges put into the main struct to be loaded via hcl
	// PrevAdhocPlan holds the value of the PrevAdhocPlan edge.
	HCLPrevAdhocPlan []*AdhocPlan `json:"PrevAdhocPlan,omitempty"`
	// NextAdhocPlan holds the value of the NextAdhocPlan edge.
	HCLNextAdhocPlan []*AdhocPlan `json:"NextAdhocPlan,omitempty"`
	// AdhocPlanToBuild holds the value of the AdhocPlanToBuild edge.
	HCLAdhocPlanToBuild *Build `json:"AdhocPlanToBuild,omitempty"`
	// AdhocPlanToStatus holds the value of the AdhocPlanToStatus edge.
	HCLAdhocPlanToStatus *Status `json:"AdhocPlanToStatus,omitempty"`
	// AdhocPlanToAgentTask holds the value of the AdhocPlanToAgentTask edge.
	HCLAdhocPlanToAgentTask *AgentTask `json:"AdhocPlanToAgentTask,omitempty"`
	//
	adhoc_plan_adhoc_plan_to_build      *uuid.UUID
	adhoc_plan_adhoc_plan_to_agent_task *uuid.UUID
}

// AdhocPlanEdges holds the relations/edges for other nodes in the graph.
type AdhocPlanEdges struct {
	// PrevAdhocPlan holds the value of the PrevAdhocPlan edge.
	PrevAdhocPlan []*AdhocPlan `json:"PrevAdhocPlan,omitempty"`
	// NextAdhocPlan holds the value of the NextAdhocPlan edge.
	NextAdhocPlan []*AdhocPlan `json:"NextAdhocPlan,omitempty"`
	// AdhocPlanToBuild holds the value of the AdhocPlanToBuild edge.
	AdhocPlanToBuild *Build `json:"AdhocPlanToBuild,omitempty"`
	// AdhocPlanToStatus holds the value of the AdhocPlanToStatus edge.
	AdhocPlanToStatus *Status `json:"AdhocPlanToStatus,omitempty"`
	// AdhocPlanToAgentTask holds the value of the AdhocPlanToAgentTask edge.
	AdhocPlanToAgentTask *AgentTask `json:"AdhocPlanToAgentTask,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// PrevAdhocPlanOrErr returns the PrevAdhocPlan value or an error if the edge
// was not loaded in eager-loading.
func (e AdhocPlanEdges) PrevAdhocPlanOrErr() ([]*AdhocPlan, error) {
	if e.loadedTypes[0] {
		return e.PrevAdhocPlan, nil
	}
	return nil, &NotLoadedError{edge: "PrevAdhocPlan"}
}

// NextAdhocPlanOrErr returns the NextAdhocPlan value or an error if the edge
// was not loaded in eager-loading.
func (e AdhocPlanEdges) NextAdhocPlanOrErr() ([]*AdhocPlan, error) {
	if e.loadedTypes[1] {
		return e.NextAdhocPlan, nil
	}
	return nil, &NotLoadedError{edge: "NextAdhocPlan"}
}

// AdhocPlanToBuildOrErr returns the AdhocPlanToBuild value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdhocPlanEdges) AdhocPlanToBuildOrErr() (*Build, error) {
	if e.loadedTypes[2] {
		if e.AdhocPlanToBuild == nil {
			// The edge AdhocPlanToBuild was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: build.Label}
		}
		return e.AdhocPlanToBuild, nil
	}
	return nil, &NotLoadedError{edge: "AdhocPlanToBuild"}
}

// AdhocPlanToStatusOrErr returns the AdhocPlanToStatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdhocPlanEdges) AdhocPlanToStatusOrErr() (*Status, error) {
	if e.loadedTypes[3] {
		if e.AdhocPlanToStatus == nil {
			// The edge AdhocPlanToStatus was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: status.Label}
		}
		return e.AdhocPlanToStatus, nil
	}
	return nil, &NotLoadedError{edge: "AdhocPlanToStatus"}
}

// AdhocPlanToAgentTaskOrErr returns the AdhocPlanToAgentTask value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AdhocPlanEdges) AdhocPlanToAgentTaskOrErr() (*AgentTask, error) {
	if e.loadedTypes[4] {
		if e.AdhocPlanToAgentTask == nil {
			// The edge AdhocPlanToAgentTask was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: agenttask.Label}
		}
		return e.AdhocPlanToAgentTask, nil
	}
	return nil, &NotLoadedError{edge: "AdhocPlanToAgentTask"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AdhocPlan) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case adhocplan.FieldID:
			values[i] = new(uuid.UUID)
		case adhocplan.ForeignKeys[0]: // adhoc_plan_adhoc_plan_to_build
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case adhocplan.ForeignKeys[1]: // adhoc_plan_adhoc_plan_to_agent_task
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type AdhocPlan", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AdhocPlan fields.
func (ap *AdhocPlan) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case adhocplan.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ap.ID = *value
			}
		case adhocplan.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field adhoc_plan_adhoc_plan_to_build", values[i])
			} else if value.Valid {
				ap.adhoc_plan_adhoc_plan_to_build = new(uuid.UUID)
				*ap.adhoc_plan_adhoc_plan_to_build = *value.S.(*uuid.UUID)
			}
		case adhocplan.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field adhoc_plan_adhoc_plan_to_agent_task", values[i])
			} else if value.Valid {
				ap.adhoc_plan_adhoc_plan_to_agent_task = new(uuid.UUID)
				*ap.adhoc_plan_adhoc_plan_to_agent_task = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryPrevAdhocPlan queries the "PrevAdhocPlan" edge of the AdhocPlan entity.
func (ap *AdhocPlan) QueryPrevAdhocPlan() *AdhocPlanQuery {
	return (&AdhocPlanClient{config: ap.config}).QueryPrevAdhocPlan(ap)
}

// QueryNextAdhocPlan queries the "NextAdhocPlan" edge of the AdhocPlan entity.
func (ap *AdhocPlan) QueryNextAdhocPlan() *AdhocPlanQuery {
	return (&AdhocPlanClient{config: ap.config}).QueryNextAdhocPlan(ap)
}

// QueryAdhocPlanToBuild queries the "AdhocPlanToBuild" edge of the AdhocPlan entity.
func (ap *AdhocPlan) QueryAdhocPlanToBuild() *BuildQuery {
	return (&AdhocPlanClient{config: ap.config}).QueryAdhocPlanToBuild(ap)
}

// QueryAdhocPlanToStatus queries the "AdhocPlanToStatus" edge of the AdhocPlan entity.
func (ap *AdhocPlan) QueryAdhocPlanToStatus() *StatusQuery {
	return (&AdhocPlanClient{config: ap.config}).QueryAdhocPlanToStatus(ap)
}

// QueryAdhocPlanToAgentTask queries the "AdhocPlanToAgentTask" edge of the AdhocPlan entity.
func (ap *AdhocPlan) QueryAdhocPlanToAgentTask() *AgentTaskQuery {
	return (&AdhocPlanClient{config: ap.config}).QueryAdhocPlanToAgentTask(ap)
}

// Update returns a builder for updating this AdhocPlan.
// Note that you need to call AdhocPlan.Unwrap() before calling this method if this AdhocPlan
// was returned from a transaction, and the transaction was committed or rolled back.
func (ap *AdhocPlan) Update() *AdhocPlanUpdateOne {
	return (&AdhocPlanClient{config: ap.config}).UpdateOne(ap)
}

// Unwrap unwraps the AdhocPlan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ap *AdhocPlan) Unwrap() *AdhocPlan {
	tx, ok := ap.config.driver.(*txDriver)
	if !ok {
		panic("ent: AdhocPlan is not a transactional entity")
	}
	ap.config.driver = tx.drv
	return ap
}

// String implements the fmt.Stringer.
func (ap *AdhocPlan) String() string {
	var builder strings.Builder
	builder.WriteString("AdhocPlan(")
	builder.WriteString(fmt.Sprintf("id=%v", ap.ID))
	builder.WriteByte(')')
	return builder.String()
}

// AdhocPlans is a parsable slice of AdhocPlan.
type AdhocPlans []*AdhocPlan

func (ap AdhocPlans) config(cfg config) {
	for _i := range ap {
		ap[_i].config = cfg
	}
}
