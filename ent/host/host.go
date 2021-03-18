// Code generated by entc, DO NOT EDIT.

package host

const (
	// Label holds the string label denoting the host type in the database.
	Label = "host"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHclID holds the string denoting the hcl_id field in the database.
	FieldHclID = "hcl_id"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldOS holds the string denoting the os field in the database.
	FieldOS = "os"
	// FieldLastOctet holds the string denoting the last_octet field in the database.
	FieldLastOctet = "last_octet"
	// FieldInstanceSize holds the string denoting the instance_size field in the database.
	FieldInstanceSize = "instance_size"
	// FieldAllowMACChanges holds the string denoting the allow_mac_changes field in the database.
	FieldAllowMACChanges = "allow_mac_changes"
	// FieldExposedTCPPorts holds the string denoting the exposed_tcp_ports field in the database.
	FieldExposedTCPPorts = "exposed_tcp_ports"
	// FieldExposedUDPPorts holds the string denoting the exposed_udp_ports field in the database.
	FieldExposedUDPPorts = "exposed_udp_ports"
	// FieldOverridePassword holds the string denoting the override_password field in the database.
	FieldOverridePassword = "override_password"
	// FieldVars holds the string denoting the vars field in the database.
	FieldVars = "vars"
	// FieldUserGroups holds the string denoting the user_groups field in the database.
	FieldUserGroups = "user_groups"
	// FieldProvisionSteps holds the string denoting the provision_steps field in the database.
	FieldProvisionSteps = "provision_steps"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"

	// EdgeHostToDisk holds the string denoting the hosttodisk edge name in mutations.
	EdgeHostToDisk = "HostToDisk"
	// EdgeHostToUser holds the string denoting the hosttouser edge name in mutations.
	EdgeHostToUser = "HostToUser"
	// EdgeHostToTag holds the string denoting the hosttotag edge name in mutations.
	EdgeHostToTag = "HostToTag"
	// EdgeHostToEnvironment holds the string denoting the hosttoenvironment edge name in mutations.
	EdgeHostToEnvironment = "HostToEnvironment"
	// EdgeHostToIncludedNetwork holds the string denoting the hosttoincludednetwork edge name in mutations.
	EdgeHostToIncludedNetwork = "HostToIncludedNetwork"
	// EdgeDependOnHostToHostDependency holds the string denoting the dependonhosttohostdependency edge name in mutations.
	EdgeDependOnHostToHostDependency = "DependOnHostToHostDependency"
	// EdgeDependByHostToHostDependency holds the string denoting the dependbyhosttohostdependency edge name in mutations.
	EdgeDependByHostToHostDependency = "DependByHostToHostDependency"

	// Table holds the table name of the host in the database.
	Table = "hosts"
	// HostToDiskTable is the table the holds the HostToDisk relation/edge. The primary key declared below.
	HostToDiskTable = "host_HostToDisk"
	// HostToDiskInverseTable is the table name for the Disk entity.
	// It exists in this package in order to avoid circular dependency with the "disk" package.
	HostToDiskInverseTable = "disks"
	// HostToUserTable is the table the holds the HostToUser relation/edge.
	HostToUserTable = "users"
	// HostToUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	HostToUserInverseTable = "users"
	// HostToUserColumn is the table column denoting the HostToUser relation/edge.
	HostToUserColumn = "host_host_to_user"
	// HostToTagTable is the table the holds the HostToTag relation/edge.
	HostToTagTable = "tags"
	// HostToTagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	HostToTagInverseTable = "tags"
	// HostToTagColumn is the table column denoting the HostToTag relation/edge.
	HostToTagColumn = "host_host_to_tag"
	// HostToEnvironmentTable is the table the holds the HostToEnvironment relation/edge. The primary key declared below.
	HostToEnvironmentTable = "environment_EnvironmentToHost"
	// HostToEnvironmentInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	HostToEnvironmentInverseTable = "environments"
	// HostToIncludedNetworkTable is the table the holds the HostToIncludedNetwork relation/edge. The primary key declared below.
	HostToIncludedNetworkTable = "included_network_IncludedNetworkToHost"
	// HostToIncludedNetworkInverseTable is the table name for the IncludedNetwork entity.
	// It exists in this package in order to avoid circular dependency with the "includednetwork" package.
	HostToIncludedNetworkInverseTable = "included_networks"
	// DependOnHostToHostDependencyTable is the table the holds the DependOnHostToHostDependency relation/edge. The primary key declared below.
	DependOnHostToHostDependencyTable = "host_dependency_HostDependencyToDependOnHost"
	// DependOnHostToHostDependencyInverseTable is the table name for the HostDependency entity.
	// It exists in this package in order to avoid circular dependency with the "hostdependency" package.
	DependOnHostToHostDependencyInverseTable = "host_dependencies"
	// DependByHostToHostDependencyTable is the table the holds the DependByHostToHostDependency relation/edge. The primary key declared below.
	DependByHostToHostDependencyTable = "host_dependency_HostDependencyToDependByHost"
	// DependByHostToHostDependencyInverseTable is the table name for the HostDependency entity.
	// It exists in this package in order to avoid circular dependency with the "hostdependency" package.
	DependByHostToHostDependencyInverseTable = "host_dependencies"
)

// Columns holds all SQL columns for host fields.
var Columns = []string{
	FieldID,
	FieldHclID,
	FieldHostname,
	FieldDescription,
	FieldOS,
	FieldLastOctet,
	FieldInstanceSize,
	FieldAllowMACChanges,
	FieldExposedTCPPorts,
	FieldExposedUDPPorts,
	FieldOverridePassword,
	FieldVars,
	FieldUserGroups,
	FieldProvisionSteps,
	FieldTags,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Host type.
var ForeignKeys = []string{
	"finding_finding_to_host",
	"provisioned_host_provisioned_host_to_host",
}

var (
	// HostToDiskPrimaryKey and HostToDiskColumn2 are the table columns denoting the
	// primary key for the HostToDisk relation (M2M).
	HostToDiskPrimaryKey = []string{"host_id", "disk_id"}
	// HostToEnvironmentPrimaryKey and HostToEnvironmentColumn2 are the table columns denoting the
	// primary key for the HostToEnvironment relation (M2M).
	HostToEnvironmentPrimaryKey = []string{"environment_id", "host_id"}
	// HostToIncludedNetworkPrimaryKey and HostToIncludedNetworkColumn2 are the table columns denoting the
	// primary key for the HostToIncludedNetwork relation (M2M).
	HostToIncludedNetworkPrimaryKey = []string{"included_network_id", "host_id"}
	// DependOnHostToHostDependencyPrimaryKey and DependOnHostToHostDependencyColumn2 are the table columns denoting the
	// primary key for the DependOnHostToHostDependency relation (M2M).
	DependOnHostToHostDependencyPrimaryKey = []string{"host_dependency_id", "host_id"}
	// DependByHostToHostDependencyPrimaryKey and DependByHostToHostDependencyColumn2 are the table columns denoting the
	// primary key for the DependByHostToHostDependency relation (M2M).
	DependByHostToHostDependencyPrimaryKey = []string{"host_dependency_id", "host_id"}
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
