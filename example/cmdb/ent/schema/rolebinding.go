package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// RoleBinding holds the schema definition for the RoleBinding entity.
type RoleBinding struct {
	ent.Schema
}

func (RoleBinding) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the RoleBinding.
func (RoleBinding) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").Values("admin", "user"),
	}
}

// Edges of the RoleBinding.
func (RoleBinding) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("project", Project.Type).Ref("role_bindings").Required().Unique(),
		//edge.From("service", Service.Type).Ref("role_bindings").Unique(),
		//edge.From("user", User.Type).Ref("role_bindings").Required().Unique(),
	}
}
