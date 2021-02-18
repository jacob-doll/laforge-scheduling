// Code generated by entc, DO NOT EDIT.

package provisionedhost

const (
	// Label holds the string label denoting the provisionedhost type in the database.
	Label = "provisioned_host"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSubnetIP holds the string denoting the subnet_ip field in the database.
	FieldSubnetIP = "subnet_ip"

	// EdgeProvisionedHostToTag holds the string denoting the provisionedhosttotag edge name in mutations.
	EdgeProvisionedHostToTag = "ProvisionedHostToTag"
	// EdgeProvisionedHostToStatus holds the string denoting the provisionedhosttostatus edge name in mutations.
	EdgeProvisionedHostToStatus = "ProvisionedHostToStatus"
	// EdgeProvisionedHostToProvisionedNetwork holds the string denoting the provisionedhosttoprovisionednetwork edge name in mutations.
	EdgeProvisionedHostToProvisionedNetwork = "ProvisionedHostToProvisionedNetwork"
	// EdgeProvisionedHostToHost holds the string denoting the provisionedhosttohost edge name in mutations.
	EdgeProvisionedHostToHost = "ProvisionedHostToHost"
	// EdgeProvisionedHostToProvisioningStep holds the string denoting the provisionedhosttoprovisioningstep edge name in mutations.
	EdgeProvisionedHostToProvisioningStep = "ProvisionedHostToProvisioningStep"
	// EdgeProvisionedHostToAgentStatus holds the string denoting the provisionedhosttoagentstatus edge name in mutations.
	EdgeProvisionedHostToAgentStatus = "ProvisionedHostToAgentStatus"

	// Table holds the table name of the provisionedhost in the database.
	Table = "provisioned_hosts"
	// ProvisionedHostToTagTable is the table the holds the ProvisionedHostToTag relation/edge.
	ProvisionedHostToTagTable = "tags"
	// ProvisionedHostToTagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	ProvisionedHostToTagInverseTable = "tags"
	// ProvisionedHostToTagColumn is the table column denoting the ProvisionedHostToTag relation/edge.
	ProvisionedHostToTagColumn = "provisioned_host_provisioned_host_to_tag"
	// ProvisionedHostToStatusTable is the table the holds the ProvisionedHostToStatus relation/edge.
	ProvisionedHostToStatusTable = "status"
	// ProvisionedHostToStatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	ProvisionedHostToStatusInverseTable = "status"
	// ProvisionedHostToStatusColumn is the table column denoting the ProvisionedHostToStatus relation/edge.
	ProvisionedHostToStatusColumn = "provisioned_host_provisioned_host_to_status"
	// ProvisionedHostToProvisionedNetworkTable is the table the holds the ProvisionedHostToProvisionedNetwork relation/edge. The primary key declared below.
	ProvisionedHostToProvisionedNetworkTable = "provisioned_host_ProvisionedHostToProvisionedNetwork"
	// ProvisionedHostToProvisionedNetworkInverseTable is the table name for the ProvisionedNetwork entity.
	// It exists in this package in order to avoid circular dependency with the "provisionednetwork" package.
	ProvisionedHostToProvisionedNetworkInverseTable = "provisioned_networks"
	// ProvisionedHostToHostTable is the table the holds the ProvisionedHostToHost relation/edge.
	ProvisionedHostToHostTable = "hosts"
	// ProvisionedHostToHostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	ProvisionedHostToHostInverseTable = "hosts"
	// ProvisionedHostToHostColumn is the table column denoting the ProvisionedHostToHost relation/edge.
	ProvisionedHostToHostColumn = "provisioned_host_provisioned_host_to_host"
	// ProvisionedHostToProvisioningStepTable is the table the holds the ProvisionedHostToProvisioningStep relation/edge. The primary key declared below.
	ProvisionedHostToProvisioningStepTable = "provisioning_step_ProvisioningStepToProvisionedHost"
	// ProvisionedHostToProvisioningStepInverseTable is the table name for the ProvisioningStep entity.
	// It exists in this package in order to avoid circular dependency with the "provisioningstep" package.
	ProvisionedHostToProvisioningStepInverseTable = "provisioning_steps"
	// ProvisionedHostToAgentStatusTable is the table the holds the ProvisionedHostToAgentStatus relation/edge. The primary key declared below.
	ProvisionedHostToAgentStatusTable = "agent_status_AgentStatusToProvisionedHost"
	// ProvisionedHostToAgentStatusInverseTable is the table name for the AgentStatus entity.
	// It exists in this package in order to avoid circular dependency with the "agentstatus" package.
	ProvisionedHostToAgentStatusInverseTable = "agent_status"
)

// Columns holds all SQL columns for provisionedhost fields.
var Columns = []string{
	FieldID,
	FieldSubnetIP,
}

var (
	// ProvisionedHostToProvisionedNetworkPrimaryKey and ProvisionedHostToProvisionedNetworkColumn2 are the table columns denoting the
	// primary key for the ProvisionedHostToProvisionedNetwork relation (M2M).
	ProvisionedHostToProvisionedNetworkPrimaryKey = []string{"provisioned_host_id", "provisioned_network_id"}
	// ProvisionedHostToProvisioningStepPrimaryKey and ProvisionedHostToProvisioningStepColumn2 are the table columns denoting the
	// primary key for the ProvisionedHostToProvisioningStep relation (M2M).
	ProvisionedHostToProvisioningStepPrimaryKey = []string{"provisioning_step_id", "provisioned_host_id"}
	// ProvisionedHostToAgentStatusPrimaryKey and ProvisionedHostToAgentStatusColumn2 are the table columns denoting the
	// primary key for the ProvisionedHostToAgentStatus relation (M2M).
	ProvisionedHostToAgentStatusPrimaryKey = []string{"agent_status_id", "provisioned_host_id"}
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
