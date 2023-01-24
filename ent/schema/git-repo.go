package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type GitRepo struct {
	ent.Schema
}

// Fields of the User.
func (GitRepo) Fields() []ent.Field {
	return []ent.Field{
		field.String("git_url"),
	}
}

// Edges of the User.
func (GitRepo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("projects", Project.Type),
	}
}

func (GitRepo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("git_url").Unique(),
	}
}
