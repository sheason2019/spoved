// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompileRecordsColumns holds the columns for the "compile_records" table.
	CompileRecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "version", Type: field.TypeString},
		{Name: "image", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "status_code", Type: field.TypeInt},
		{Name: "output", Type: field.TypeString},
		{Name: "branch", Type: field.TypeString},
	}
	// CompileRecordsTable holds the schema information for the "compile_records" table.
	CompileRecordsTable = &schema.Table{
		Name:       "compile_records",
		Columns:    CompileRecordsColumns,
		PrimaryKey: []*schema.Column{CompileRecordsColumns[0]},
	}
	// ProjectsColumns holds the columns for the "projects" table.
	ProjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "project_name", Type: field.TypeString},
		{Name: "describe", Type: field.TypeString},
		{Name: "git_url", Type: field.TypeString},
		{Name: "dir_path", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// ProjectsTable holds the schema information for the "projects" table.
	ProjectsTable = &schema.Table{
		Name:       "projects",
		Columns:    ProjectsColumns,
		PrimaryKey: []*schema.Column{ProjectsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "project_project_name",
				Unique:  false,
				Columns: []*schema.Column{ProjectsColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "password_salt", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_username",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1]},
			},
		},
	}
	// ProjectCompileRecordsColumns holds the columns for the "project_compile_records" table.
	ProjectCompileRecordsColumns = []*schema.Column{
		{Name: "project_id", Type: field.TypeInt},
		{Name: "compile_record_id", Type: field.TypeInt},
	}
	// ProjectCompileRecordsTable holds the schema information for the "project_compile_records" table.
	ProjectCompileRecordsTable = &schema.Table{
		Name:       "project_compile_records",
		Columns:    ProjectCompileRecordsColumns,
		PrimaryKey: []*schema.Column{ProjectCompileRecordsColumns[0], ProjectCompileRecordsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "project_compile_records_project_id",
				Columns:    []*schema.Column{ProjectCompileRecordsColumns[0]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "project_compile_records_compile_record_id",
				Columns:    []*schema.Column{ProjectCompileRecordsColumns[1]},
				RefColumns: []*schema.Column{CompileRecordsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserProjectsColumns holds the columns for the "user_projects" table.
	UserProjectsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "project_id", Type: field.TypeInt},
	}
	// UserProjectsTable holds the schema information for the "user_projects" table.
	UserProjectsTable = &schema.Table{
		Name:       "user_projects",
		Columns:    UserProjectsColumns,
		PrimaryKey: []*schema.Column{UserProjectsColumns[0], UserProjectsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_projects_user_id",
				Columns:    []*schema.Column{UserProjectsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_projects_project_id",
				Columns:    []*schema.Column{UserProjectsColumns[1]},
				RefColumns: []*schema.Column{ProjectsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserCompileRecordsColumns holds the columns for the "user_compile_records" table.
	UserCompileRecordsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "compile_record_id", Type: field.TypeInt},
	}
	// UserCompileRecordsTable holds the schema information for the "user_compile_records" table.
	UserCompileRecordsTable = &schema.Table{
		Name:       "user_compile_records",
		Columns:    UserCompileRecordsColumns,
		PrimaryKey: []*schema.Column{UserCompileRecordsColumns[0], UserCompileRecordsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_compile_records_user_id",
				Columns:    []*schema.Column{UserCompileRecordsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_compile_records_compile_record_id",
				Columns:    []*schema.Column{UserCompileRecordsColumns[1]},
				RefColumns: []*schema.Column{CompileRecordsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompileRecordsTable,
		ProjectsTable,
		UsersTable,
		ProjectCompileRecordsTable,
		UserProjectsTable,
		UserCompileRecordsTable,
	}
)

func init() {
	ProjectCompileRecordsTable.ForeignKeys[0].RefTable = ProjectsTable
	ProjectCompileRecordsTable.ForeignKeys[1].RefTable = CompileRecordsTable
	UserProjectsTable.ForeignKeys[0].RefTable = UsersTable
	UserProjectsTable.ForeignKeys[1].RefTable = ProjectsTable
	UserCompileRecordsTable.ForeignKeys[0].RefTable = UsersTable
	UserCompileRecordsTable.ForeignKeys[1].RefTable = CompileRecordsTable
}
