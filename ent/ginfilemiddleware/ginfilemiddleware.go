// Code generated by entc, DO NOT EDIT.

package ginfilemiddleware

const (
	// Label holds the string label denoting the ginfilemiddleware type in the database.
	Label = "gin_file_middleware"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURLPath holds the string denoting the url_path field in the database.
	FieldURLPath = "url_path"
	// FieldFilePath holds the string denoting the file_path field in the database.
	FieldFilePath = "file_path"
	// FieldAccessed holds the string denoting the accessed field in the database.
	FieldAccessed = "accessed"

	// EdgeGinFileMiddlewareToProvisionedHost holds the string denoting the ginfilemiddlewaretoprovisionedhost edge name in mutations.
	EdgeGinFileMiddlewareToProvisionedHost = "GinFileMiddlewareToProvisionedHost"
	// EdgeGinFileMiddlewareToProvisioningStep holds the string denoting the ginfilemiddlewaretoprovisioningstep edge name in mutations.
	EdgeGinFileMiddlewareToProvisioningStep = "GinFileMiddlewareToProvisioningStep"

	// Table holds the table name of the ginfilemiddleware in the database.
	Table = "gin_file_middlewares"
	// GinFileMiddlewareToProvisionedHostTable is the table the holds the GinFileMiddlewareToProvisionedHost relation/edge.
	GinFileMiddlewareToProvisionedHostTable = "provisioned_hosts"
	// GinFileMiddlewareToProvisionedHostInverseTable is the table name for the ProvisionedHost entity.
	// It exists in this package in order to avoid circular dependency with the "provisionedhost" package.
	GinFileMiddlewareToProvisionedHostInverseTable = "provisioned_hosts"
	// GinFileMiddlewareToProvisionedHostColumn is the table column denoting the GinFileMiddlewareToProvisionedHost relation/edge.
	GinFileMiddlewareToProvisionedHostColumn = "gin_file_middleware_gin_file_middleware_to_provisioned_host"
	// GinFileMiddlewareToProvisioningStepTable is the table the holds the GinFileMiddlewareToProvisioningStep relation/edge.
	GinFileMiddlewareToProvisioningStepTable = "provisioning_steps"
	// GinFileMiddlewareToProvisioningStepInverseTable is the table name for the ProvisioningStep entity.
	// It exists in this package in order to avoid circular dependency with the "provisioningstep" package.
	GinFileMiddlewareToProvisioningStepInverseTable = "provisioning_steps"
	// GinFileMiddlewareToProvisioningStepColumn is the table column denoting the GinFileMiddlewareToProvisioningStep relation/edge.
	GinFileMiddlewareToProvisioningStepColumn = "gin_file_middleware_gin_file_middleware_to_provisioning_step"
)

// Columns holds all SQL columns for ginfilemiddleware fields.
var Columns = []string{
	FieldID,
	FieldURLPath,
	FieldFilePath,
	FieldAccessed,
}

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
	// DefaultAccessed holds the default value on creation for the "accessed" field.
	DefaultAccessed bool
)
