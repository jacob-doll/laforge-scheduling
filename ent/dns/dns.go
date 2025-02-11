// Code generated by ent, DO NOT EDIT.

package dns

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the dns type in the database.
	Label = "dns"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHclID holds the string denoting the hcl_id field in the database.
	FieldHclID = "hcl_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldRootDomain holds the string denoting the root_domain field in the database.
	FieldRootDomain = "root_domain"
	// FieldDNSServers holds the string denoting the dns_servers field in the database.
	FieldDNSServers = "dns_servers"
	// FieldNtpServers holds the string denoting the ntp_servers field in the database.
	FieldNtpServers = "ntp_servers"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// EdgeDNSToEnvironment holds the string denoting the dnstoenvironment edge name in mutations.
	EdgeDNSToEnvironment = "DNSToEnvironment"
	// EdgeDNSToCompetition holds the string denoting the dnstocompetition edge name in mutations.
	EdgeDNSToCompetition = "DNSToCompetition"
	// Table holds the table name of the dns in the database.
	Table = "dn_ss"
	// DNSToEnvironmentTable is the table that holds the DNSToEnvironment relation/edge. The primary key declared below.
	DNSToEnvironmentTable = "environment_EnvironmentToDNS"
	// DNSToEnvironmentInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	DNSToEnvironmentInverseTable = "environments"
	// DNSToCompetitionTable is the table that holds the DNSToCompetition relation/edge. The primary key declared below.
	DNSToCompetitionTable = "competition_CompetitionToDNS"
	// DNSToCompetitionInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	DNSToCompetitionInverseTable = "competitions"
)

// Columns holds all SQL columns for dns fields.
var Columns = []string{
	FieldID,
	FieldHclID,
	FieldType,
	FieldRootDomain,
	FieldDNSServers,
	FieldNtpServers,
	FieldConfig,
}

var (
	// DNSToEnvironmentPrimaryKey and DNSToEnvironmentColumn2 are the table columns denoting the
	// primary key for the DNSToEnvironment relation (M2M).
	DNSToEnvironmentPrimaryKey = []string{"environment_id", "dns_id"}
	// DNSToCompetitionPrimaryKey and DNSToCompetitionColumn2 are the table columns denoting the
	// primary key for the DNSToCompetition relation (M2M).
	DNSToCompetitionPrimaryKey = []string{"competition_id", "dns_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
