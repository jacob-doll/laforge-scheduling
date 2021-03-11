// Code generated by entc, DO NOT EDIT.

package environment

const (
	// Label holds the string label denoting the environment type in the database.
	Label = "environment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHclID holds the string denoting the hcl_id field in the database.
	FieldHclID = "hcl_id"
	// FieldCompetitionID holds the string denoting the competition_id field in the database.
	FieldCompetitionID = "competition_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldBuilder holds the string denoting the builder field in the database.
	FieldBuilder = "builder"
	// FieldTeamCount holds the string denoting the team_count field in the database.
	FieldTeamCount = "team_count"
	// FieldRevision holds the string denoting the revision field in the database.
	FieldRevision = "revision"
	// FieldAdminCidrs holds the string denoting the admin_cidrs field in the database.
	FieldAdminCidrs = "admin_cidrs"
	// FieldExposedVdiPorts holds the string denoting the exposed_vdi_ports field in the database.
	FieldExposedVdiPorts = "exposed_vdi_ports"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"

	// EdgeEnvironmentToTag holds the string denoting the environmenttotag edge name in mutations.
	EdgeEnvironmentToTag = "EnvironmentToTag"
	// EdgeEnvironmentToUser holds the string denoting the environmenttouser edge name in mutations.
	EdgeEnvironmentToUser = "EnvironmentToUser"
	// EdgeEnvironmentToHost holds the string denoting the environmenttohost edge name in mutations.
	EdgeEnvironmentToHost = "EnvironmentToHost"
	// EdgeEnvironmentToCompetition holds the string denoting the environmenttocompetition edge name in mutations.
	EdgeEnvironmentToCompetition = "EnvironmentToCompetition"
	// EdgeEnvironmentToBuild holds the string denoting the environmenttobuild edge name in mutations.
	EdgeEnvironmentToBuild = "EnvironmentToBuild"
	// EdgeEnvironmentToIdentity holds the string denoting the environmenttoidentity edge name in mutations.
	EdgeEnvironmentToIdentity = "EnvironmentToIdentity"
	// EdgeEnvironmentToIncludedNetwork holds the string denoting the environmenttoincludednetwork edge name in mutations.
	EdgeEnvironmentToIncludedNetwork = "EnvironmentToIncludedNetwork"
	// EdgeEnvironmentToNetwork holds the string denoting the environmenttonetwork edge name in mutations.
	EdgeEnvironmentToNetwork = "EnvironmentToNetwork"
	// EdgeEnvironmentToTeam holds the string denoting the environmenttoteam edge name in mutations.
	EdgeEnvironmentToTeam = "EnvironmentToTeam"

	// Table holds the table name of the environment in the database.
	Table = "environments"
	// EnvironmentToTagTable is the table the holds the EnvironmentToTag relation/edge.
	EnvironmentToTagTable = "tags"
	// EnvironmentToTagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	EnvironmentToTagInverseTable = "tags"
	// EnvironmentToTagColumn is the table column denoting the EnvironmentToTag relation/edge.
	EnvironmentToTagColumn = "environment_environment_to_tag"
	// EnvironmentToUserTable is the table the holds the EnvironmentToUser relation/edge. The primary key declared below.
	EnvironmentToUserTable = "environment_EnvironmentToUser"
	// EnvironmentToUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	EnvironmentToUserInverseTable = "users"
	// EnvironmentToHostTable is the table the holds the EnvironmentToHost relation/edge. The primary key declared below.
	EnvironmentToHostTable = "environment_EnvironmentToHost"
	// EnvironmentToHostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	EnvironmentToHostInverseTable = "hosts"
	// EnvironmentToCompetitionTable is the table the holds the EnvironmentToCompetition relation/edge. The primary key declared below.
	EnvironmentToCompetitionTable = "environment_EnvironmentToCompetition"
	// EnvironmentToCompetitionInverseTable is the table name for the Competition entity.
	// It exists in this package in order to avoid circular dependency with the "competition" package.
	EnvironmentToCompetitionInverseTable = "competitions"
	// EnvironmentToBuildTable is the table the holds the EnvironmentToBuild relation/edge. The primary key declared below.
	EnvironmentToBuildTable = "environment_EnvironmentToBuild"
	// EnvironmentToBuildInverseTable is the table name for the Build entity.
	// It exists in this package in order to avoid circular dependency with the "build" package.
	EnvironmentToBuildInverseTable = "builds"
	// EnvironmentToIdentityTable is the table the holds the EnvironmentToIdentity relation/edge. The primary key declared below.
	EnvironmentToIdentityTable = "environment_EnvironmentToIdentity"
	// EnvironmentToIdentityInverseTable is the table name for the Identity entity.
	// It exists in this package in order to avoid circular dependency with the "identity" package.
	EnvironmentToIdentityInverseTable = "identities"
	// EnvironmentToIncludedNetworkTable is the table the holds the EnvironmentToIncludedNetwork relation/edge. The primary key declared below.
	EnvironmentToIncludedNetworkTable = "included_network_IncludedNetworkToEnvironment"
	// EnvironmentToIncludedNetworkInverseTable is the table name for the IncludedNetwork entity.
	// It exists in this package in order to avoid circular dependency with the "includednetwork" package.
	EnvironmentToIncludedNetworkInverseTable = "included_networks"
	// EnvironmentToNetworkTable is the table the holds the EnvironmentToNetwork relation/edge. The primary key declared below.
	EnvironmentToNetworkTable = "network_NetworkToEnvironment"
	// EnvironmentToNetworkInverseTable is the table name for the Network entity.
	// It exists in this package in order to avoid circular dependency with the "network" package.
	EnvironmentToNetworkInverseTable = "networks"
	// EnvironmentToTeamTable is the table the holds the EnvironmentToTeam relation/edge. The primary key declared below.
	EnvironmentToTeamTable = "team_TeamToEnvironment"
	// EnvironmentToTeamInverseTable is the table name for the Team entity.
	// It exists in this package in order to avoid circular dependency with the "team" package.
	EnvironmentToTeamInverseTable = "teams"
)

// Columns holds all SQL columns for environment fields.
var Columns = []string{
	FieldID,
	FieldHclID,
	FieldCompetitionID,
	FieldName,
	FieldDescription,
	FieldBuilder,
	FieldTeamCount,
	FieldRevision,
	FieldAdminCidrs,
	FieldExposedVdiPorts,
	FieldConfig,
	FieldTags,
}

var (
	// EnvironmentToUserPrimaryKey and EnvironmentToUserColumn2 are the table columns denoting the
	// primary key for the EnvironmentToUser relation (M2M).
	EnvironmentToUserPrimaryKey = []string{"environment_id", "user_id"}
	// EnvironmentToHostPrimaryKey and EnvironmentToHostColumn2 are the table columns denoting the
	// primary key for the EnvironmentToHost relation (M2M).
	EnvironmentToHostPrimaryKey = []string{"environment_id", "host_id"}
	// EnvironmentToCompetitionPrimaryKey and EnvironmentToCompetitionColumn2 are the table columns denoting the
	// primary key for the EnvironmentToCompetition relation (M2M).
	EnvironmentToCompetitionPrimaryKey = []string{"environment_id", "competition_id"}
	// EnvironmentToBuildPrimaryKey and EnvironmentToBuildColumn2 are the table columns denoting the
	// primary key for the EnvironmentToBuild relation (M2M).
	EnvironmentToBuildPrimaryKey = []string{"environment_id", "build_id"}
	// EnvironmentToIdentityPrimaryKey and EnvironmentToIdentityColumn2 are the table columns denoting the
	// primary key for the EnvironmentToIdentity relation (M2M).
	EnvironmentToIdentityPrimaryKey = []string{"environment_id", "identity_id"}
	// EnvironmentToIncludedNetworkPrimaryKey and EnvironmentToIncludedNetworkColumn2 are the table columns denoting the
	// primary key for the EnvironmentToIncludedNetwork relation (M2M).
	EnvironmentToIncludedNetworkPrimaryKey = []string{"included_network_id", "environment_id"}
	// EnvironmentToNetworkPrimaryKey and EnvironmentToNetworkColumn2 are the table columns denoting the
	// primary key for the EnvironmentToNetwork relation (M2M).
	EnvironmentToNetworkPrimaryKey = []string{"network_id", "environment_id"}
	// EnvironmentToTeamPrimaryKey and EnvironmentToTeamColumn2 are the table columns denoting the
	// primary key for the EnvironmentToTeam relation (M2M).
	EnvironmentToTeamPrimaryKey = []string{"team_id", "environment_id"}
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
