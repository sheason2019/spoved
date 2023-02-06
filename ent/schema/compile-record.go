package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type CompileRecord struct {
	ent.Schema
}

// Fields of the User.
func (CompileRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("version"),
		field.String("image"),
		field.Time("created_at"),
		// 0表示执行中 1表示成功 -1表示失败
		field.Int("status_code"),
		field.String("output"),
		field.String("branch"),
	}
}

// Edges of the User.
func (CompileRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("operator", User.Type).Ref("compile_records"),
		edge.From("project", Project.Type).Ref("compile_records"),
	}
}

func (CompileRecord) Indexes() []ent.Index {
	return []ent.Index{}
}
