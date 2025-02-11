// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/dnsrecord"
	"github.com/gen0cide/laforge/ent/environment"
	"github.com/google/uuid"
)

// DNSRecord is the model entity for the DNSRecord schema.
type DNSRecord struct {
	config ` json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// HclID holds the value of the "hcl_id" field.
	HclID string `json:"hcl_id,omitempty" hcl:"id,label"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty" hcl:"name,attr"`
	// Values holds the value of the "values" field.
	Values []string `json:"values,omitempty" hcl:"values,optional"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty" hcl:"type,attr"`
	// Zone holds the value of the "zone" field.
	Zone string `json:"zone,omitempty" hcl:"zone,attr" `
	// Vars holds the value of the "vars" field.
	Vars map[string]string `json:"vars,omitempty" hcl:"vars,optional"`
	// Disabled holds the value of the "disabled" field.
	Disabled bool `json:"disabled,omitempty" hcl:"disabled,optional"`
	// Tags holds the value of the "tags" field.
	Tags map[string]string `json:"tags,omitempty" hcl:"tags,optional"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DNSRecordQuery when eager-loading is set.
	Edges DNSRecordEdges `json:"edges"`

	// Edges put into the main struct to be loaded via hcl
	// DNSRecordToEnvironment holds the value of the DNSRecordToEnvironment edge.
	HCLDNSRecordToEnvironment *Environment `json:"DNSRecordToEnvironment,omitempty"`
	//
	environment_environment_to_dns_record *uuid.UUID
}

// DNSRecordEdges holds the relations/edges for other nodes in the graph.
type DNSRecordEdges struct {
	// DNSRecordToEnvironment holds the value of the DNSRecordToEnvironment edge.
	DNSRecordToEnvironment *Environment `json:"DNSRecordToEnvironment,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DNSRecordToEnvironmentOrErr returns the DNSRecordToEnvironment value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DNSRecordEdges) DNSRecordToEnvironmentOrErr() (*Environment, error) {
	if e.loadedTypes[0] {
		if e.DNSRecordToEnvironment == nil {
			// The edge DNSRecordToEnvironment was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: environment.Label}
		}
		return e.DNSRecordToEnvironment, nil
	}
	return nil, &NotLoadedError{edge: "DNSRecordToEnvironment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DNSRecord) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case dnsrecord.FieldValues, dnsrecord.FieldVars, dnsrecord.FieldTags:
			values[i] = new([]byte)
		case dnsrecord.FieldDisabled:
			values[i] = new(sql.NullBool)
		case dnsrecord.FieldHclID, dnsrecord.FieldName, dnsrecord.FieldType, dnsrecord.FieldZone:
			values[i] = new(sql.NullString)
		case dnsrecord.FieldID:
			values[i] = new(uuid.UUID)
		case dnsrecord.ForeignKeys[0]: // environment_environment_to_dns_record
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type DNSRecord", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DNSRecord fields.
func (dr *DNSRecord) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case dnsrecord.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				dr.ID = *value
			}
		case dnsrecord.FieldHclID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hcl_id", values[i])
			} else if value.Valid {
				dr.HclID = value.String
			}
		case dnsrecord.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				dr.Name = value.String
			}
		case dnsrecord.FieldValues:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field values", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dr.Values); err != nil {
					return fmt.Errorf("unmarshal field values: %w", err)
				}
			}
		case dnsrecord.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				dr.Type = value.String
			}
		case dnsrecord.FieldZone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field zone", values[i])
			} else if value.Valid {
				dr.Zone = value.String
			}
		case dnsrecord.FieldVars:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field vars", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dr.Vars); err != nil {
					return fmt.Errorf("unmarshal field vars: %w", err)
				}
			}
		case dnsrecord.FieldDisabled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field disabled", values[i])
			} else if value.Valid {
				dr.Disabled = value.Bool
			}
		case dnsrecord.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &dr.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case dnsrecord.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field environment_environment_to_dns_record", values[i])
			} else if value.Valid {
				dr.environment_environment_to_dns_record = new(uuid.UUID)
				*dr.environment_environment_to_dns_record = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryDNSRecordToEnvironment queries the "DNSRecordToEnvironment" edge of the DNSRecord entity.
func (dr *DNSRecord) QueryDNSRecordToEnvironment() *EnvironmentQuery {
	return (&DNSRecordClient{config: dr.config}).QueryDNSRecordToEnvironment(dr)
}

// Update returns a builder for updating this DNSRecord.
// Note that you need to call DNSRecord.Unwrap() before calling this method if this DNSRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (dr *DNSRecord) Update() *DNSRecordUpdateOne {
	return (&DNSRecordClient{config: dr.config}).UpdateOne(dr)
}

// Unwrap unwraps the DNSRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dr *DNSRecord) Unwrap() *DNSRecord {
	tx, ok := dr.config.driver.(*txDriver)
	if !ok {
		panic("ent: DNSRecord is not a transactional entity")
	}
	dr.config.driver = tx.drv
	return dr
}

// String implements the fmt.Stringer.
func (dr *DNSRecord) String() string {
	var builder strings.Builder
	builder.WriteString("DNSRecord(")
	builder.WriteString(fmt.Sprintf("id=%v", dr.ID))
	builder.WriteString(", hcl_id=")
	builder.WriteString(dr.HclID)
	builder.WriteString(", name=")
	builder.WriteString(dr.Name)
	builder.WriteString(", values=")
	builder.WriteString(fmt.Sprintf("%v", dr.Values))
	builder.WriteString(", type=")
	builder.WriteString(dr.Type)
	builder.WriteString(", zone=")
	builder.WriteString(dr.Zone)
	builder.WriteString(", vars=")
	builder.WriteString(fmt.Sprintf("%v", dr.Vars))
	builder.WriteString(", disabled=")
	builder.WriteString(fmt.Sprintf("%v", dr.Disabled))
	builder.WriteString(", tags=")
	builder.WriteString(fmt.Sprintf("%v", dr.Tags))
	builder.WriteByte(')')
	return builder.String()
}

// DNSRecords is a parsable slice of DNSRecord.
type DNSRecords []*DNSRecord

func (dr DNSRecords) config(cfg config) {
	for _i := range dr {
		dr[_i].config = cfg
	}
}
