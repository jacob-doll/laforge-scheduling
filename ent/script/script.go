// Code generated by entc, DO NOT EDIT.

package script

const (
	// Label holds the string label denoting the script type in the database.
	Label = "script"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHclID holds the string denoting the hcl_id field in the database.
	FieldHclID = "hcl_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldSourceType holds the string denoting the source_type field in the database.
	FieldSourceType = "source_type"
	// FieldCooldown holds the string denoting the cooldown field in the database.
	FieldCooldown = "cooldown"
	// FieldTimeout holds the string denoting the timeout field in the database.
	FieldTimeout = "timeout"
	// FieldIgnoreErrors holds the string denoting the ignore_errors field in the database.
	FieldIgnoreErrors = "ignore_errors"
	// FieldArgs holds the string denoting the args field in the database.
	FieldArgs = "args"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// FieldVars holds the string denoting the vars field in the database.
	FieldVars = "vars"
	// FieldAbsPath holds the string denoting the abs_path field in the database.
	FieldAbsPath = "abs_path"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"

	// EdgeScriptToTag holds the string denoting the scripttotag edge name in mutations.
	EdgeScriptToTag = "ScriptToTag"
	// EdgeScriptToUser holds the string denoting the scripttouser edge name in mutations.
	EdgeScriptToUser = "ScriptToUser"
	// EdgeScriptToFinding holds the string denoting the scripttofinding edge name in mutations.
	EdgeScriptToFinding = "ScriptToFinding"
	// EdgeScriptToEnvironment holds the string denoting the scripttoenvironment edge name in mutations.
	EdgeScriptToEnvironment = "ScriptToEnvironment"

	// Table holds the table name of the script in the database.
	Table = "scripts"
	// ScriptToTagTable is the table the holds the ScriptToTag relation/edge.
	ScriptToTagTable = "tags"
	// ScriptToTagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	ScriptToTagInverseTable = "tags"
	// ScriptToTagColumn is the table column denoting the ScriptToTag relation/edge.
	ScriptToTagColumn = "script_script_to_tag"
	// ScriptToUserTable is the table the holds the ScriptToUser relation/edge.
	ScriptToUserTable = "users"
	// ScriptToUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ScriptToUserInverseTable = "users"
	// ScriptToUserColumn is the table column denoting the ScriptToUser relation/edge.
	ScriptToUserColumn = "script_script_to_user"
	// ScriptToFindingTable is the table the holds the ScriptToFinding relation/edge. The primary key declared below.
	ScriptToFindingTable = "script_ScriptToFinding"
	// ScriptToFindingInverseTable is the table name for the Finding entity.
	// It exists in this package in order to avoid circular dependency with the "finding" package.
	ScriptToFindingInverseTable = "findings"
	// ScriptToEnvironmentTable is the table the holds the ScriptToEnvironment relation/edge. The primary key declared below.
	ScriptToEnvironmentTable = "environment_EnvironmentToScript"
	// ScriptToEnvironmentInverseTable is the table name for the Environment entity.
	// It exists in this package in order to avoid circular dependency with the "environment" package.
	ScriptToEnvironmentInverseTable = "environments"
)

// Columns holds all SQL columns for script fields.
var Columns = []string{
	FieldID,
	FieldHclID,
	FieldName,
	FieldLanguage,
	FieldDescription,
	FieldSource,
	FieldSourceType,
	FieldCooldown,
	FieldTimeout,
	FieldIgnoreErrors,
	FieldArgs,
	FieldDisabled,
	FieldVars,
	FieldAbsPath,
	FieldTags,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Script type.
var ForeignKeys = []string{
	"provisioning_step_provisioning_step_to_script",
}

var (
	// ScriptToFindingPrimaryKey and ScriptToFindingColumn2 are the table columns denoting the
	// primary key for the ScriptToFinding relation (M2M).
	ScriptToFindingPrimaryKey = []string{"script_id", "finding_id"}
	// ScriptToEnvironmentPrimaryKey and ScriptToEnvironmentColumn2 are the table columns denoting the
	// primary key for the ScriptToEnvironment relation (M2M).
	ScriptToEnvironmentPrimaryKey = []string{"environment_id", "script_id"}
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
