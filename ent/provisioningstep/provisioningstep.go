// Code generated by entc, DO NOT EDIT.

package provisioningstep

const (
	// Label holds the string label denoting the provisioningstep type in the database.
	Label = "provisioning_step"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProvisionerType holds the string denoting the provisioner_type field in the database.
	FieldProvisionerType = "provisioner_type"
	// FieldStepNumber holds the string denoting the step_number field in the database.
	FieldStepNumber = "step_number"

	// EdgeProvisioningStepToTag holds the string denoting the provisioningsteptotag edge name in mutations.
	EdgeProvisioningStepToTag = "ProvisioningStepToTag"
	// EdgeProvisioningStepToStatus holds the string denoting the provisioningsteptostatus edge name in mutations.
	EdgeProvisioningStepToStatus = "ProvisioningStepToStatus"
	// EdgeProvisioningStepToProvisionedHost holds the string denoting the provisioningsteptoprovisionedhost edge name in mutations.
	EdgeProvisioningStepToProvisionedHost = "ProvisioningStepToProvisionedHost"
	// EdgeProvisioningStepToScript holds the string denoting the provisioningsteptoscript edge name in mutations.
	EdgeProvisioningStepToScript = "ProvisioningStepToScript"
	// EdgeProvisioningStepToCommand holds the string denoting the provisioningsteptocommand edge name in mutations.
	EdgeProvisioningStepToCommand = "ProvisioningStepToCommand"
	// EdgeProvisioningStepToDNSRecord holds the string denoting the provisioningsteptodnsrecord edge name in mutations.
	EdgeProvisioningStepToDNSRecord = "ProvisioningStepToDNSRecord"
	// EdgeProvisioningStepToFileDelete holds the string denoting the provisioningsteptofiledelete edge name in mutations.
	EdgeProvisioningStepToFileDelete = "ProvisioningStepToFileDelete"
	// EdgeProvisioningStepToFileDownload holds the string denoting the provisioningsteptofiledownload edge name in mutations.
	EdgeProvisioningStepToFileDownload = "ProvisioningStepToFileDownload"
	// EdgeProvisioningStepToFileExtract holds the string denoting the provisioningsteptofileextract edge name in mutations.
	EdgeProvisioningStepToFileExtract = "ProvisioningStepToFileExtract"
	// EdgeProvisioningStepToPlan holds the string denoting the provisioningsteptoplan edge name in mutations.
	EdgeProvisioningStepToPlan = "ProvisioningStepToPlan"
	// EdgeProvisioningStepToGinFileMiddleware holds the string denoting the provisioningsteptoginfilemiddleware edge name in mutations.
	EdgeProvisioningStepToGinFileMiddleware = "ProvisioningStepToGinFileMiddleware"

	// Table holds the table name of the provisioningstep in the database.
	Table = "provisioning_steps"
	// ProvisioningStepToTagTable is the table the holds the ProvisioningStepToTag relation/edge.
	ProvisioningStepToTagTable = "tags"
	// ProvisioningStepToTagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	ProvisioningStepToTagInverseTable = "tags"
	// ProvisioningStepToTagColumn is the table column denoting the ProvisioningStepToTag relation/edge.
	ProvisioningStepToTagColumn = "provisioning_step_provisioning_step_to_tag"
	// ProvisioningStepToStatusTable is the table the holds the ProvisioningStepToStatus relation/edge.
	ProvisioningStepToStatusTable = "status"
	// ProvisioningStepToStatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	ProvisioningStepToStatusInverseTable = "status"
	// ProvisioningStepToStatusColumn is the table column denoting the ProvisioningStepToStatus relation/edge.
	ProvisioningStepToStatusColumn = "provisioning_step_provisioning_step_to_status"
	// ProvisioningStepToProvisionedHostTable is the table the holds the ProvisioningStepToProvisionedHost relation/edge. The primary key declared below.
	ProvisioningStepToProvisionedHostTable = "provisioning_step_ProvisioningStepToProvisionedHost"
	// ProvisioningStepToProvisionedHostInverseTable is the table name for the ProvisionedHost entity.
	// It exists in this package in order to avoid circular dependency with the "provisionedhost" package.
	ProvisioningStepToProvisionedHostInverseTable = "provisioned_hosts"
	// ProvisioningStepToScriptTable is the table the holds the ProvisioningStepToScript relation/edge.
	ProvisioningStepToScriptTable = "scripts"
	// ProvisioningStepToScriptInverseTable is the table name for the Script entity.
	// It exists in this package in order to avoid circular dependency with the "script" package.
	ProvisioningStepToScriptInverseTable = "scripts"
	// ProvisioningStepToScriptColumn is the table column denoting the ProvisioningStepToScript relation/edge.
	ProvisioningStepToScriptColumn = "provisioning_step_provisioning_step_to_script"
	// ProvisioningStepToCommandTable is the table the holds the ProvisioningStepToCommand relation/edge.
	ProvisioningStepToCommandTable = "commands"
	// ProvisioningStepToCommandInverseTable is the table name for the Command entity.
	// It exists in this package in order to avoid circular dependency with the "command" package.
	ProvisioningStepToCommandInverseTable = "commands"
	// ProvisioningStepToCommandColumn is the table column denoting the ProvisioningStepToCommand relation/edge.
	ProvisioningStepToCommandColumn = "provisioning_step_provisioning_step_to_command"
	// ProvisioningStepToDNSRecordTable is the table the holds the ProvisioningStepToDNSRecord relation/edge.
	ProvisioningStepToDNSRecordTable = "dns_records"
	// ProvisioningStepToDNSRecordInverseTable is the table name for the DNSRecord entity.
	// It exists in this package in order to avoid circular dependency with the "dnsrecord" package.
	ProvisioningStepToDNSRecordInverseTable = "dns_records"
	// ProvisioningStepToDNSRecordColumn is the table column denoting the ProvisioningStepToDNSRecord relation/edge.
	ProvisioningStepToDNSRecordColumn = "provisioning_step_provisioning_step_to_dns_record"
	// ProvisioningStepToFileDeleteTable is the table the holds the ProvisioningStepToFileDelete relation/edge.
	ProvisioningStepToFileDeleteTable = "file_deletes"
	// ProvisioningStepToFileDeleteInverseTable is the table name for the FileDelete entity.
	// It exists in this package in order to avoid circular dependency with the "filedelete" package.
	ProvisioningStepToFileDeleteInverseTable = "file_deletes"
	// ProvisioningStepToFileDeleteColumn is the table column denoting the ProvisioningStepToFileDelete relation/edge.
	ProvisioningStepToFileDeleteColumn = "provisioning_step_provisioning_step_to_file_delete"
	// ProvisioningStepToFileDownloadTable is the table the holds the ProvisioningStepToFileDownload relation/edge.
	ProvisioningStepToFileDownloadTable = "file_downloads"
	// ProvisioningStepToFileDownloadInverseTable is the table name for the FileDownload entity.
	// It exists in this package in order to avoid circular dependency with the "filedownload" package.
	ProvisioningStepToFileDownloadInverseTable = "file_downloads"
	// ProvisioningStepToFileDownloadColumn is the table column denoting the ProvisioningStepToFileDownload relation/edge.
	ProvisioningStepToFileDownloadColumn = "provisioning_step_provisioning_step_to_file_download"
	// ProvisioningStepToFileExtractTable is the table the holds the ProvisioningStepToFileExtract relation/edge.
	ProvisioningStepToFileExtractTable = "file_extracts"
	// ProvisioningStepToFileExtractInverseTable is the table name for the FileExtract entity.
	// It exists in this package in order to avoid circular dependency with the "fileextract" package.
	ProvisioningStepToFileExtractInverseTable = "file_extracts"
	// ProvisioningStepToFileExtractColumn is the table column denoting the ProvisioningStepToFileExtract relation/edge.
	ProvisioningStepToFileExtractColumn = "provisioning_step_provisioning_step_to_file_extract"
	// ProvisioningStepToPlanTable is the table the holds the ProvisioningStepToPlan relation/edge. The primary key declared below.
	ProvisioningStepToPlanTable = "plan_PlanToProvisioningStep"
	// ProvisioningStepToPlanInverseTable is the table name for the Plan entity.
	// It exists in this package in order to avoid circular dependency with the "plan" package.
	ProvisioningStepToPlanInverseTable = "plans"
	// ProvisioningStepToGinFileMiddlewareTable is the table the holds the ProvisioningStepToGinFileMiddleware relation/edge.
	ProvisioningStepToGinFileMiddlewareTable = "provisioning_steps"
	// ProvisioningStepToGinFileMiddlewareInverseTable is the table name for the GinFileMiddleware entity.
	// It exists in this package in order to avoid circular dependency with the "ginfilemiddleware" package.
	ProvisioningStepToGinFileMiddlewareInverseTable = "gin_file_middlewares"
	// ProvisioningStepToGinFileMiddlewareColumn is the table column denoting the ProvisioningStepToGinFileMiddleware relation/edge.
	ProvisioningStepToGinFileMiddlewareColumn = "gin_file_middleware_gin_file_middleware_to_provisioning_step"
)

// Columns holds all SQL columns for provisioningstep fields.
var Columns = []string{
	FieldID,
	FieldProvisionerType,
	FieldStepNumber,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the ProvisioningStep type.
var ForeignKeys = []string{
	"gin_file_middleware_gin_file_middleware_to_provisioning_step",
}

var (
	// ProvisioningStepToProvisionedHostPrimaryKey and ProvisioningStepToProvisionedHostColumn2 are the table columns denoting the
	// primary key for the ProvisioningStepToProvisionedHost relation (M2M).
	ProvisioningStepToProvisionedHostPrimaryKey = []string{"provisioning_step_id", "provisioned_host_id"}
	// ProvisioningStepToPlanPrimaryKey and ProvisioningStepToPlanColumn2 are the table columns denoting the
	// primary key for the ProvisioningStepToPlan relation (M2M).
	ProvisioningStepToPlanPrimaryKey = []string{"plan_id", "provisioning_step_id"}
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
