package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type Project struct {
	ent.Schema
}

// Fields of the User.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("project_name"),
		field.String("describe"),
		field.Time("created_at"),
	}
}

// Edges of the User.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("git_repo", GitRepo.Type).Ref("projects"),
		edge.From("creator", User.Type).Ref("projects"),
	}
}

func (Project) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("project_name"),
	}
}
