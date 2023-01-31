// Code generated by ent, DO NOT EDIT.

package compilerecord

const (
	// Label holds the string label denoting the compilerecord type in the database.
	Label = "compile_record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldImage holds the string denoting the image field in the database.
	FieldImage = "image"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldStatusCode holds the string denoting the status_code field in the database.
	FieldStatusCode = "status_code"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldBranch holds the string denoting the branch field in the database.
	FieldBranch = "branch"
	// EdgeOperator holds the string denoting the operator edge name in mutations.
	EdgeOperator = "operator"
	// EdgeProject holds the string denoting the project edge name in mutations.
	EdgeProject = "project"
	// Table holds the table name of the compilerecord in the database.
	Table = "compile_records"
	// OperatorTable is the table that holds the operator relation/edge. The primary key declared below.
	OperatorTable = "user_compile_records"
	// OperatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OperatorInverseTable = "users"
	// ProjectTable is the table that holds the project relation/edge. The primary key declared below.
	ProjectTable = "project_compile_records"
	// ProjectInverseTable is the table name for the Project entity.
	// It exists in this package in order to avoid circular dependency with the "project" package.
	ProjectInverseTable = "projects"
)

// Columns holds all SQL columns for compilerecord fields.
var Columns = []string{
	FieldID,
	FieldVersion,
	FieldImage,
	FieldCreatedAt,
	FieldStatusCode,
	FieldOutput,
	FieldBranch,
}

var (
	// OperatorPrimaryKey and OperatorColumn2 are the table columns denoting the
	// primary key for the operator relation (M2M).
	OperatorPrimaryKey = []string{"user_id", "compile_record_id"}
	// ProjectPrimaryKey and ProjectColumn2 are the table columns denoting the
	// primary key for the project relation (M2M).
	ProjectPrimaryKey = []string{"project_id", "compile_record_id"}
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