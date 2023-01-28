// Code generated by ent, DO NOT EDIT.

package project

const (
	// Label holds the string label denoting the project type in the database.
	Label = "project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProjectName holds the string denoting the project_name field in the database.
	FieldProjectName = "project_name"
	// FieldDescribe holds the string denoting the describe field in the database.
	FieldDescribe = "describe"
	// FieldGitURL holds the string denoting the git_url field in the database.
	FieldGitURL = "git_url"
	// FieldDirPath holds the string denoting the dir_path field in the database.
	FieldDirPath = "dir_path"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// Table holds the table name of the project in the database.
	Table = "projects"
	// CreatorTable is the table that holds the creator relation/edge. The primary key declared below.
	CreatorTable = "user_projects"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
)

// Columns holds all SQL columns for project fields.
var Columns = []string{
	FieldID,
	FieldProjectName,
	FieldDescribe,
	FieldGitURL,
	FieldDirPath,
	FieldCreatedAt,
}

var (
	// CreatorPrimaryKey and CreatorColumn2 are the table columns denoting the
	// primary key for the creator relation (M2M).
	CreatorPrimaryKey = []string{"user_id", "project_id"}
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