package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type DeployRecord struct {
	ent.Schema
}

// Fields of the User.
func (DeployRecord) Fields() []ent.Field {
	return []ent.Field{
		field.String("image"),
		field.Time("created_at"),
		// 0表示启动中 1表示正在执行 -1表示停机
		field.Int("status_code"),
		field.String("container_hash"),
	}
}

// Edges of the User.
func (DeployRecord) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("operator", User.Type).Ref("deploy_records"),
		edge.From("compile_record", CompileRecord.Type).Ref("deploy_records"),
	}
}

func (DeployRecord) Indexes() []ent.Index {
	return []ent.Index{}
}
