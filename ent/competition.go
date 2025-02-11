// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/competition"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/google/uuid"
)

// Competition is the model entity for the Competition schema.
type Competition struct {
	config ` json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// HclID holds the value of the "hcl_id" field.
	HclID string `json:"hcl_id,omitempty" hcl:"id,label"`
	// RootPassword holds the value of the "root_password" field.
	RootPassword string `json:"root_password,omitempty" hcl:"root_password,attr"`
	// Config holds the value of the "config" field.
	Config map[string]string `json:"config,omitempty" hcl:"config,optional"`
	// Tags holds the value of the "tags" field.
	Tags map[string]string `json:"tags,omitempty" hcl:"tags,optional"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompetitionQuery when eager-loading is set.
	Edges CompetitionEdges `json:"edges"`

	// Edges put into the main struct to be loaded via hcl
	// CompetitionToDNS holds the value of the CompetitionToDNS edge.
	HCLCompetitionToDNS []*DNS `json:"CompetitionToDNS,omitempty" hcl:"dns,block"`
	// CompetitionToEnvironment holds the value of the CompetitionToEnvironment edge.
	HCLCompetitionToEnvironment *Environment `json:"CompetitionToEnvironment,omitempty"`
	// CompetitionToBuild holds the value of the CompetitionToBuild edge.
	HCLCompetitionToBuild []*Build `json:"CompetitionToBuild,omitempty"`
	//
	environment_environment_to_competition *uuid.UUID
}

// CompetitionEdges holds the relations/edges for other nodes in the graph.
type CompetitionEdges struct {
	// CompetitionToDNS holds the value of the CompetitionToDNS edge.
	CompetitionToDNS []*DNS `json:"CompetitionToDNS,omitempty" hcl:"dns,block"`
	// CompetitionToEnvironment holds the value of the CompetitionToEnvironment edge.
	CompetitionToEnvironment *Environment `json:"CompetitionToEnvironment,omitempty"`
	// CompetitionToBuild holds the value of the CompetitionToBuild edge.
	CompetitionToBuild []*Build `json:"CompetitionToBuild,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CompetitionToDNSOrErr returns the CompetitionToDNS value or an error if the edge
// was not loaded in eager-loading.
func (e CompetitionEdges) CompetitionToDNSOrErr() ([]*DNS, error) {
	if e.loadedTypes[0] {
		return e.CompetitionToDNS, nil
	}
	return nil, &NotLoadedError{edge: "CompetitionToDNS"}
}

// CompetitionToEnvironmentOrErr returns the CompetitionToEnvironment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CompetitionEdges) CompetitionToEnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[1] {
		if e.CompetitionToEnvironment == nil {
			// The edge CompetitionToEnvironment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.CompetitionToEnvironment, nil
	}
	return nil, &NotLoadedError{edge: "CompetitionToEnvironment"}
}

// CompetitionToBuildOrErr returns the CompetitionToBuild value or an error if the edge
// was not loaded in eager-loading.
func (e CompetitionEdges) CompetitionToBuildOrErr() ([]*Build, error) {
	if e.loadedTypes[2] {
		return e.CompetitionToBuild, nil
	}
	return nil, &NotLoadedError{edge: "CompetitionToBuild"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Competition) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case competition.FieldConfig, competition.FieldTags:
			values[i] = new([]byte)
		case competition.FieldHclID, competition.FieldRootPassword:
			values[i] = new(sql.NullString)
		case competition.FieldID:
			values[i] = new(uuid.UUID)
		case competition.ForeignKeys[0]: // environment_environment_to_competition
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Competition", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Competition fields.
func (c *Competition) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case competition.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case competition.FieldHclID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hcl_id", values[i])
			} else if value.Valid {
				c.HclID = value.String
			}
		case competition.FieldRootPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field root_password", values[i])
			} else if value.Valid {
				c.RootPassword = value.String
			}
		case competition.FieldConfig:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field config", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Config); err != nil {
					return fmt.Errorf("unmarshal field config: %w", err)
				}
			}
		case competition.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &c.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case competition.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field environment_environment_to_competition", values[i])
			} else if value.Valid {
				c.environment_environment_to_competition = new(uuid.UUID)
				*c.environment_environment_to_competition = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryCompetitionToDNS queries the "CompetitionToDNS" edge of the Competition entity.
func (c *Competition) QueryCompetitionToDNS() *DNSQuery {
	return (&CompetitionClient{config: c.config}).QueryCompetitionToDNS(c)
}

// QueryCompetitionToEnvironment queries the "CompetitionToEnvironment" edge of the Competition entity.
func (c *Competition) QueryCompetitionToEnvironment() *EnvironmentQuery {
	return (&CompetitionClient{config: c.config}).QueryCompetitionToEnvironment(c)
}

// QueryCompetitionToBuild queries the "CompetitionToBuild" edge of the Competition entity.
func (c *Competition) QueryCompetitionToBuild() *BuildQuery {
	return (&CompetitionClient{config: c.config}).QueryCompetitionToBuild(c)
}

// Update returns a builder for updating this Competition.
// Note that you need to call Competition.Unwrap() before calling this method if this Competition
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Competition) Update() *CompetitionUpdateOne {
	return (&CompetitionClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Competition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Competition) Unwrap() *Competition {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Competition is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Competition) String() string {
	var builder strings.Builder
	builder.WriteString("Competition(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", hcl_id=")
	builder.WriteString(c.HclID)
	builder.WriteString(", root_password=")
	builder.WriteString(c.RootPassword)
	builder.WriteString(", config=")
	builder.WriteString(fmt.Sprintf("%v", c.Config))
	builder.WriteString(", tags=")
	builder.WriteString(fmt.Sprintf("%v", c.Tags))
	builder.WriteByte(')')
	return builder.String()
}

// Competitions is a parsable slice of Competition.
type Competitions []*Competition

func (c Competitions) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
