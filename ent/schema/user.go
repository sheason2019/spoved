package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username"),
		field.String("password_hash"),
		field.String("password_salt"),
		field.Time("created_at"),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("projects", Project.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("username").Unique(),
	}
}
