package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/fitan/genapi/pkg/gen_mgr"
	"time"
)

// RoleBinding holds the schema definition for the RoleBinding entity.
type RoleBinding struct {
	ent.Schema
}

// Fields of the RoleBinding.
func (RoleBinding) Fields() []ent.Field {
	return []ent.Field{
		field.String("role_name").Unique().Annotations(gen_mgr.RestFieldOp{
			FieldQueryable:   gen_mgr.FieldQueryable{EQ: gen_mgr.GenRestTrue, ContainsFold: gen_mgr.GenRestTrue},
			FieldOperability: gen_mgr.FieldOperability{},
		}),
		field.String("role_id").Unique().Annotations(gen_mgr.RestFieldOp{
			FieldQueryable:   gen_mgr.FieldQueryable{EQ: gen_mgr.GenRestTrue},
			FieldOperability: gen_mgr.FieldOperability{},
		}),
		field.Bool("status").Annotations(gen_mgr.RestFieldOp{
			FieldQueryable:   gen_mgr.FieldQueryable{EQ: gen_mgr.GenRestTrue},
			FieldOperability: gen_mgr.FieldOperability{},
		}),
		field.Time("created_at").Default(time.Now).Immutable().Optional().Nillable(),
		field.Text("note"),
		field.Strings("permissions"),
	}
}

// Edges of the RoleBinding.
func (RoleBinding) Edges() []ent.Edge {
	return []ent.Edge{}
}
