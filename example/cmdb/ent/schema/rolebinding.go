package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// RoleBinding holds the schema definition for the RoleBinding entity.
type RoleBinding struct {
	ent.Schema
}

// Fields of the RoleBinding.
func (RoleBinding) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_name").Unique(),
		field.String("role_id").Unique(),
		field.Bool("status"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Text("note"),
		field.Strings("permissions"),
	}
}

// Edges of the RoleBinding.
func (RoleBinding) Edges() []ent.Edge {
	return []ent.Edge{}
}
