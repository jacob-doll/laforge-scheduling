// Code generated by ent, DO NOT EDIT.

package adhocplan

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the adhocplan type in the database.
	Label = "adhoc_plan"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgePrevAdhocPlan holds the string denoting the prevadhocplan edge name in mutations.
	EdgePrevAdhocPlan = "PrevAdhocPlan"
	// EdgeNextAdhocPlan holds the string denoting the nextadhocplan edge name in mutations.
	EdgeNextAdhocPlan = "NextAdhocPlan"
	// EdgeAdhocPlanToBuild holds the string denoting the adhocplantobuild edge name in mutations.
	EdgeAdhocPlanToBuild = "AdhocPlanToBuild"
	// EdgeAdhocPlanToStatus holds the string denoting the adhocplantostatus edge name in mutations.
	EdgeAdhocPlanToStatus = "AdhocPlanToStatus"
	// EdgeAdhocPlanToAgentTask holds the string denoting the adhocplantoagenttask edge name in mutations.
	EdgeAdhocPlanToAgentTask = "AdhocPlanToAgentTask"
	// Table holds the table name of the adhocplan in the database.
	Table = "adhoc_plans"
	// PrevAdhocPlanTable is the table that holds the PrevAdhocPlan relation/edge. The primary key declared below.
	PrevAdhocPlanTable = "adhoc_plan_NextAdhocPlan"
	// NextAdhocPlanTable is the table that holds the NextAdhocPlan relation/edge. The primary key declared below.
	NextAdhocPlanTable = "adhoc_plan_NextAdhocPlan"
	// AdhocPlanToBuildTable is the table that holds the AdhocPlanToBuild relation/edge.
	AdhocPlanToBuildTable = "adhoc_plans"
	// AdhocPlanToBuildInverseTable is the table name for the Build entity.
	// It exists in this package in order to avoid circular dependency with the "build" package.
	AdhocPlanToBuildInverseTable = "builds"
	// AdhocPlanToBuildColumn is the table column denoting the AdhocPlanToBuild relation/edge.
	AdhocPlanToBuildColumn = "adhoc_plan_adhoc_plan_to_build"
	// AdhocPlanToStatusTable is the table that holds the AdhocPlanToStatus relation/edge.
	AdhocPlanToStatusTable = "status"
	// AdhocPlanToStatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	AdhocPlanToStatusInverseTable = "status"
	// AdhocPlanToStatusColumn is the table column denoting the AdhocPlanToStatus relation/edge.
	AdhocPlanToStatusColumn = "adhoc_plan_adhoc_plan_to_status"
	// AdhocPlanToAgentTaskTable is the table that holds the AdhocPlanToAgentTask relation/edge.
	AdhocPlanToAgentTaskTable = "adhoc_plans"
	// AdhocPlanToAgentTaskInverseTable is the table name for the AgentTask entity.
	// It exists in this package in order to avoid circular dependency with the "agenttask" package.
	AdhocPlanToAgentTaskInverseTable = "agent_tasks"
	// AdhocPlanToAgentTaskColumn is the table column denoting the AdhocPlanToAgentTask relation/edge.
	AdhocPlanToAgentTaskColumn = "adhoc_plan_adhoc_plan_to_agent_task"
)

// Columns holds all SQL columns for adhocplan fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "adhoc_plans"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"adhoc_plan_adhoc_plan_to_build",
	"adhoc_plan_adhoc_plan_to_agent_task",
}

var (
	// PrevAdhocPlanPrimaryKey and PrevAdhocPlanColumn2 are the table columns denoting the
	// primary key for the PrevAdhocPlan relation (M2M).
	PrevAdhocPlanPrimaryKey = []string{"adhoc_plan_id", "PrevAdhocPlan_id"}
	// NextAdhocPlanPrimaryKey and NextAdhocPlanColumn2 are the table columns denoting the
	// primary key for the NextAdhocPlan relation (M2M).
	NextAdhocPlanPrimaryKey = []string{"adhoc_plan_id", "PrevAdhocPlan_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
