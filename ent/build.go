// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/build"
	"github.com/gen0cide/laforge/ent/buildcommit"
	"github.com/gen0cide/laforge/ent/competition"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/gen0cide/laforge/ent/repocommit"
	"github.com/gen0cide/laforge/ent/status"
	"github.com/google/uuid"
)

// Build is the model entity for the Build schema.
type Build struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Revision holds the value of the "revision" field.
	Revision int `json:"revision,omitempty"`
	// EnvironmentRevision holds the value of the "environment_revision" field.
	EnvironmentRevision int `json:"environment_revision,omitempty"`
	// Vars holds the value of the "vars" field.
	Vars map[string]string `json:"vars,omitempty"`
	// CompletedPlan holds the value of the "completed_plan" field.
	CompletedPlan bool `json:"completed_plan,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BuildQuery when eager-loading is set.
	Edges BuildEdges `json:"edges"`

	// Edges put into the main struct to be loaded via hcl
	// BuildToStatus holds the value of the BuildToStatus edge.
	HCLBuildToStatus *Status `json:"BuildToStatus,omitempty"`
	// BuildToEnvironment holds the value of the BuildToEnvironment edge.
	HCLBuildToEnvironment *Environment `json:"BuildToEnvironment,omitempty"`
	// BuildToCompetition holds the value of the BuildToCompetition edge.
	HCLBuildToCompetition *Competition `json:"BuildToCompetition,omitempty"`
	// BuildToLatestBuildCommit holds the value of the BuildToLatestBuildCommit edge.
	HCLBuildToLatestBuildCommit *BuildCommit `json:"BuildToLatestBuildCommit,omitempty"`
	// BuildToRepoCommit holds the value of the BuildToRepoCommit edge.
	HCLBuildToRepoCommit *RepoCommit `json:"BuildToRepoCommit,omitempty"`
	// BuildToProvisionedNetwork holds the value of the BuildToProvisionedNetwork edge.
	HCLBuildToProvisionedNetwork []*ProvisionedNetwork `json:"BuildToProvisionedNetwork,omitempty"`
	// BuildToTeam holds the value of the BuildToTeam edge.
	HCLBuildToTeam []*Team `json:"BuildToTeam,omitempty"`
	// BuildToPlan holds the value of the BuildToPlan edge.
	HCLBuildToPlan []*Plan `json:"BuildToPlan,omitempty"`
	// BuildToBuildCommits holds the value of the BuildToBuildCommits edge.
	HCLBuildToBuildCommits []*BuildCommit `json:"BuildToBuildCommits,omitempty"`
	// BuildToAdhocPlans holds the value of the BuildToAdhocPlans edge.
	HCLBuildToAdhocPlans []*AdhocPlan `json:"BuildToAdhocPlans,omitempty"`
	// BuildToAgentStatuses holds the value of the BuildToAgentStatuses edge.
	HCLBuildToAgentStatuses []*AgentStatus `json:"BuildToAgentStatuses,omitempty"`
	// BuildToServerTasks holds the value of the BuildToServerTasks edge.
	HCLBuildToServerTasks []*ServerTask `json:"BuildToServerTasks,omitempty"`
	//
	build_build_to_environment         *uuid.UUID
	build_build_to_competition         *uuid.UUID
	build_build_to_latest_build_commit *uuid.UUID
	build_build_to_repo_commit         *uuid.UUID
}

// BuildEdges holds the relations/edges for other nodes in the graph.
type BuildEdges struct {
	// BuildToStatus holds the value of the BuildToStatus edge.
	BuildToStatus *Status `json:"BuildToStatus,omitempty"`
	// BuildToEnvironment holds the value of the BuildToEnvironment edge.
	BuildToEnvironment *Environment `json:"BuildToEnvironment,omitempty"`
	// BuildToCompetition holds the value of the BuildToCompetition edge.
	BuildToCompetition *Competition `json:"BuildToCompetition,omitempty"`
	// BuildToLatestBuildCommit holds the value of the BuildToLatestBuildCommit edge.
	BuildToLatestBuildCommit *BuildCommit `json:"BuildToLatestBuildCommit,omitempty"`
	// BuildToRepoCommit holds the value of the BuildToRepoCommit edge.
	BuildToRepoCommit *RepoCommit `json:"BuildToRepoCommit,omitempty"`
	// BuildToProvisionedNetwork holds the value of the BuildToProvisionedNetwork edge.
	BuildToProvisionedNetwork []*ProvisionedNetwork `json:"BuildToProvisionedNetwork,omitempty"`
	// BuildToTeam holds the value of the BuildToTeam edge.
	BuildToTeam []*Team `json:"BuildToTeam,omitempty"`
	// BuildToPlan holds the value of the BuildToPlan edge.
	BuildToPlan []*Plan `json:"BuildToPlan,omitempty"`
	// BuildToBuildCommits holds the value of the BuildToBuildCommits edge.
	BuildToBuildCommits []*BuildCommit `json:"BuildToBuildCommits,omitempty"`
	// BuildToAdhocPlans holds the value of the BuildToAdhocPlans edge.
	BuildToAdhocPlans []*AdhocPlan `json:"BuildToAdhocPlans,omitempty"`
	// BuildToAgentStatuses holds the value of the BuildToAgentStatuses edge.
	BuildToAgentStatuses []*AgentStatus `json:"BuildToAgentStatuses,omitempty"`
	// BuildToServerTasks holds the value of the BuildToServerTasks edge.
	BuildToServerTasks []*ServerTask `json:"BuildToServerTasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [12]bool
}

// BuildToStatusOrErr returns the BuildToStatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BuildEdges) BuildToStatusOrErr() (*Status, error) {
	if e.loadedTypes[0] {
		if e.BuildToStatus == nil {
			// The edge BuildToStatus was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: status.Label}
		}
		return e.BuildToStatus, nil
	}
	return nil, &NotLoadedError{edge: "BuildToStatus"}
}

// BuildToEnvironmentOrErr returns the BuildToEnvironment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BuildEdges) BuildToEnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[1] {
		if e.BuildToEnvironment == nil {
			// The edge BuildToEnvironment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.BuildToEnvironment, nil
	}
	return nil, &NotLoadedError{edge: "BuildToEnvironment"}
}

// BuildToCompetitionOrErr returns the BuildToCompetition value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BuildEdges) BuildToCompetitionOrErr() (*Competition, error) {
	if e.loadedTypes[2] {
		if e.BuildToCompetition == nil {
			// The edge BuildToCompetition was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: competition.Label}
		}
		return e.BuildToCompetition, nil
	}
	return nil, &NotLoadedError{edge: "BuildToCompetition"}
}

// BuildToLatestBuildCommitOrErr returns the BuildToLatestBuildCommit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BuildEdges) BuildToLatestBuildCommitOrErr() (*BuildCommit, error) {
	if e.loadedTypes[3] {
		if e.BuildToLatestBuildCommit == nil {
			// The edge BuildToLatestBuildCommit was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: buildcommit.Label}
		}
		return e.BuildToLatestBuildCommit, nil
	}
	return nil, &NotLoadedError{edge: "BuildToLatestBuildCommit"}
}

// BuildToRepoCommitOrErr returns the BuildToRepoCommit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BuildEdges) BuildToRepoCommitOrErr() (*RepoCommit, error) {
	if e.loadedTypes[4] {
		if e.BuildToRepoCommit == nil {
			// The edge BuildToRepoCommit was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: repocommit.Label}
		}
		return e.BuildToRepoCommit, nil
	}
	return nil, &NotLoadedError{edge: "BuildToRepoCommit"}
}

// BuildToProvisionedNetworkOrErr returns the BuildToProvisionedNetwork value or an error if the edge
// was not loaded in eager-loading.
func (e BuildEdges) BuildToProvisionedNetworkOrErr() ([]*ProvisionedNetwork, error) {
	if e.loadedTypes[5] {
		return e.BuildToProvisionedNetwork, nil
	}
	return nil, &NotLoadedError{edge: "BuildToProvisionedNetwork"}
}

// BuildToTeamOrErr returns the BuildToTeam value or an error if the edge
// was not loaded in eager-loading.
func (e BuildEdges) BuildToTeamOrErr() ([]*Team, error) {
	if e.loadedTypes[6] {
		return e.BuildToTeam, nil
	}
	return nil, &NotLoadedError{edge: "BuildToTeam"}
}

// BuildToPlanOrErr returns the BuildToPlan value or an error if the edge
// was not loaded in eager-loading.
func (e BuildEdges) BuildToPlanOrErr() ([]*Plan, error) {
	if e.loadedTypes[7] {
		return e.BuildToPlan, nil
	}
	return nil, &NotLoadedError{edge: "BuildToPlan"}
}

// BuildToBuildCommitsOrErr returns the BuildToBuildCommits value or an error if the edge
// was not loaded in eager-loading.
func (e BuildEdges) BuildToBuildCommitsOrErr() ([]*BuildCommit, error) {
	if e.loadedTypes[8] {
		return e.BuildToBuildCommits, nil
	}
	return nil, &NotLoadedError{edge: "BuildToBuildCommits"}
}

// BuildToAdhocPlansOrErr returns the BuildToAdhocPlans value or an error if the edge
// was not loaded in eager-loading.
func (e BuildEdges) BuildToAdhocPlansOrErr() ([]*AdhocPlan, error) {
	if e.loadedTypes[9] {
		return e.BuildToAdhocPlans, nil
	}
	return nil, &NotLoadedError{edge: "BuildToAdhocPlans"}
}

// BuildToAgentStatusesOrErr returns the BuildToAgentStatuses value or an error if the edge
// was not loaded in eager-loading.
func (e BuildEdges) BuildToAgentStatusesOrErr() ([]*AgentStatus, error) {
	if e.loadedTypes[10] {
		return e.BuildToAgentStatuses, nil
	}
	return nil, &NotLoadedError{edge: "BuildToAgentStatuses"}
}

// BuildToServerTasksOrErr returns the BuildToServerTasks value or an error if the edge
// was not loaded in eager-loading.
func (e BuildEdges) BuildToServerTasksOrErr() ([]*ServerTask, error) {
	if e.loadedTypes[11] {
		return e.BuildToServerTasks, nil
	}
	return nil, &NotLoadedError{edge: "BuildToServerTasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Build) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case build.FieldVars:
			values[i] = new([]byte)
		case build.FieldCompletedPlan:
			values[i] = new(sql.NullBool)
		case build.FieldRevision, build.FieldEnvironmentRevision:
			values[i] = new(sql.NullInt64)
		case build.FieldID:
			values[i] = new(uuid.UUID)
		case build.ForeignKeys[0]: // build_build_to_environment
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case build.ForeignKeys[1]: // build_build_to_competition
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case build.ForeignKeys[2]: // build_build_to_latest_build_commit
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case build.ForeignKeys[3]: // build_build_to_repo_commit
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Build", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Build fields.
func (b *Build) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case build.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case build.FieldRevision:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field revision", values[i])
			} else if value.Valid {
				b.Revision = int(value.Int64)
			}
		case build.FieldEnvironmentRevision:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field environment_revision", values[i])
			} else if value.Valid {
				b.EnvironmentRevision = int(value.Int64)
			}
		case build.FieldVars:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field vars", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &b.Vars); err != nil {
					return fmt.Errorf("unmarshal field vars: %w", err)
				}
			}
		case build.FieldCompletedPlan:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field completed_plan", values[i])
			} else if value.Valid {
				b.CompletedPlan = value.Bool
			}
		case build.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field build_build_to_environment", values[i])
			} else if value.Valid {
				b.build_build_to_environment = new(uuid.UUID)
				*b.build_build_to_environment = *value.S.(*uuid.UUID)
			}
		case build.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field build_build_to_competition", values[i])
			} else if value.Valid {
				b.build_build_to_competition = new(uuid.UUID)
				*b.build_build_to_competition = *value.S.(*uuid.UUID)
			}
		case build.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field build_build_to_latest_build_commit", values[i])
			} else if value.Valid {
				b.build_build_to_latest_build_commit = new(uuid.UUID)
				*b.build_build_to_latest_build_commit = *value.S.(*uuid.UUID)
			}
		case build.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field build_build_to_repo_commit", values[i])
			} else if value.Valid {
				b.build_build_to_repo_commit = new(uuid.UUID)
				*b.build_build_to_repo_commit = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryBuildToStatus queries the "BuildToStatus" edge of the Build entity.
func (b *Build) QueryBuildToStatus() *StatusQuery {
	return (&BuildClient{config: b.config}).QueryBuildToStatus(b)
}

// QueryBuildToEnvironment queries the "BuildToEnvironment" edge of the Build entity.
func (b *Build) QueryBuildToEnvironment() *EnvironmentQuery {
	return (&BuildClient{config: b.config}).QueryBuildToEnvironment(b)
}

// QueryBuildToCompetition queries the "BuildToCompetition" edge of the Build entity.
func (b *Build) QueryBuildToCompetition() *CompetitionQuery {
	return (&BuildClient{config: b.config}).QueryBuildToCompetition(b)
}

// QueryBuildToLatestBuildCommit queries the "BuildToLatestBuildCommit" edge of the Build entity.
func (b *Build) QueryBuildToLatestBuildCommit() *BuildCommitQuery {
	return (&BuildClient{config: b.config}).QueryBuildToLatestBuildCommit(b)
}

// QueryBuildToRepoCommit queries the "BuildToRepoCommit" edge of the Build entity.
func (b *Build) QueryBuildToRepoCommit() *RepoCommitQuery {
	return (&BuildClient{config: b.config}).QueryBuildToRepoCommit(b)
}

// QueryBuildToProvisionedNetwork queries the "BuildToProvisionedNetwork" edge of the Build entity.
func (b *Build) QueryBuildToProvisionedNetwork() *ProvisionedNetworkQuery {
	return (&BuildClient{config: b.config}).QueryBuildToProvisionedNetwork(b)
}

// QueryBuildToTeam queries the "BuildToTeam" edge of the Build entity.
func (b *Build) QueryBuildToTeam() *TeamQuery {
	return (&BuildClient{config: b.config}).QueryBuildToTeam(b)
}

// QueryBuildToPlan queries the "BuildToPlan" edge of the Build entity.
func (b *Build) QueryBuildToPlan() *PlanQuery {
	return (&BuildClient{config: b.config}).QueryBuildToPlan(b)
}

// QueryBuildToBuildCommits queries the "BuildToBuildCommits" edge of the Build entity.
func (b *Build) QueryBuildToBuildCommits() *BuildCommitQuery {
	return (&BuildClient{config: b.config}).QueryBuildToBuildCommits(b)
}

// QueryBuildToAdhocPlans queries the "BuildToAdhocPlans" edge of the Build entity.
func (b *Build) QueryBuildToAdhocPlans() *AdhocPlanQuery {
	return (&BuildClient{config: b.config}).QueryBuildToAdhocPlans(b)
}

// QueryBuildToAgentStatuses queries the "BuildToAgentStatuses" edge of the Build entity.
func (b *Build) QueryBuildToAgentStatuses() *AgentStatusQuery {
	return (&BuildClient{config: b.config}).QueryBuildToAgentStatuses(b)
}

// QueryBuildToServerTasks queries the "BuildToServerTasks" edge of the Build entity.
func (b *Build) QueryBuildToServerTasks() *ServerTaskQuery {
	return (&BuildClient{config: b.config}).QueryBuildToServerTasks(b)
}

// Update returns a builder for updating this Build.
// Note that you need to call Build.Unwrap() before calling this method if this Build
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Build) Update() *BuildUpdateOne {
	return (&BuildClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Build entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Build) Unwrap() *Build {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Build is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Build) String() string {
	var builder strings.Builder
	builder.WriteString("Build(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", revision=")
	builder.WriteString(fmt.Sprintf("%v", b.Revision))
	builder.WriteString(", environment_revision=")
	builder.WriteString(fmt.Sprintf("%v", b.EnvironmentRevision))
	builder.WriteString(", vars=")
	builder.WriteString(fmt.Sprintf("%v", b.Vars))
	builder.WriteString(", completed_plan=")
	builder.WriteString(fmt.Sprintf("%v", b.CompletedPlan))
	builder.WriteByte(')')
	return builder.String()
}

// Builds is a parsable slice of Build.
type Builds []*Build

func (b Builds) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
