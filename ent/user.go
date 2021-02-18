// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config ` json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty" hcl:"name,attr"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty" hcl:"uuid,optional"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty" hcl:"email,attr"`
	// HclID holds the value of the "hcl_id" field.
	HclID string `json:"hcl_id,omitempty" hcl:"id,label"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges                   UserEdges `json:"edges"`
	build_build_to_user     *int
	command_command_to_user *int
	finding_finding_to_user *int
	host_host_to_user       *int
	script_script_to_user   *int
	team_team_to_user       *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserToTag holds the value of the UserToTag edge.
	UserToTag []*Tag `json:"UserToTag,omitempty"`
	// UserToEnvironment holds the value of the UserToEnvironment edge.
	UserToEnvironment []*Environment `json:"UserToEnvironment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserToTagOrErr returns the UserToTag value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserToTagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.UserToTag, nil
	}
	return nil, &NotLoadedError{edge: "UserToTag"}
}

// UserToEnvironmentOrErr returns the UserToEnvironment value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserToEnvironmentOrErr() ([]*Environment, error) {
	if e.loadedTypes[1] {
		return e.UserToEnvironment, nil
	}
	return nil, &NotLoadedError{edge: "UserToEnvironment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = &sql.NullInt64{}
		case user.FieldName, user.FieldUUID, user.FieldEmail, user.FieldHclID:
			values[i] = &sql.NullString{}
		case user.ForeignKeys[0]: // build_build_to_user
			values[i] = &sql.NullInt64{}
		case user.ForeignKeys[1]: // command_command_to_user
			values[i] = &sql.NullInt64{}
		case user.ForeignKeys[2]: // finding_finding_to_user
			values[i] = &sql.NullInt64{}
		case user.ForeignKeys[3]: // host_host_to_user
			values[i] = &sql.NullInt64{}
		case user.ForeignKeys[4]: // script_script_to_user
			values[i] = &sql.NullInt64{}
		case user.ForeignKeys[5]: // team_team_to_user
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldUUID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uuid", values[i])
			} else if value.Valid {
				u.UUID = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldHclID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hcl_id", values[i])
			} else if value.Valid {
				u.HclID = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field build_build_to_user", value)
			} else if value.Valid {
				u.build_build_to_user = new(int)
				*u.build_build_to_user = int(value.Int64)
			}
		case user.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field command_command_to_user", value)
			} else if value.Valid {
				u.command_command_to_user = new(int)
				*u.command_command_to_user = int(value.Int64)
			}
		case user.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field finding_finding_to_user", value)
			} else if value.Valid {
				u.finding_finding_to_user = new(int)
				*u.finding_finding_to_user = int(value.Int64)
			}
		case user.ForeignKeys[3]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field host_host_to_user", value)
			} else if value.Valid {
				u.host_host_to_user = new(int)
				*u.host_host_to_user = int(value.Int64)
			}
		case user.ForeignKeys[4]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field script_script_to_user", value)
			} else if value.Valid {
				u.script_script_to_user = new(int)
				*u.script_script_to_user = int(value.Int64)
			}
		case user.ForeignKeys[5]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field team_team_to_user", value)
			} else if value.Valid {
				u.team_team_to_user = new(int)
				*u.team_team_to_user = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUserToTag queries the "UserToTag" edge of the User entity.
func (u *User) QueryUserToTag() *TagQuery {
	return (&UserClient{config: u.config}).QueryUserToTag(u)
}

// QueryUserToEnvironment queries the "UserToEnvironment" edge of the User entity.
func (u *User) QueryUserToEnvironment() *EnvironmentQuery {
	return (&UserClient{config: u.config}).QueryUserToEnvironment(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", uuid=")
	builder.WriteString(u.UUID)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteString(", hcl_id=")
	builder.WriteString(u.HclID)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
