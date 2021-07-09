package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fitan/genapi/pkg/gen_mgr"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User)Annotations() []schema.Annotation {
	return []schema.Annotation{
		gen_mgr.RestNodeOp{
			Paging: gen_mgr.Paging{
				Open:     true,
				Must:     false,
				MaxLimit: 100,
			},
			Order:  gen_mgr.Order{
				DefaultAcsOrder:   nil,
				DefaultDescOrder:  nil,
				OpenOptionalOrder: false,
				OptionalOrder:     nil,
			},
			Method: gen_mgr.NodeMethod{},
		},
	}

}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("age1").Optional().Annotations(gen_mgr.RestFieldOp{
			FieldQueryable:   gen_mgr.FieldQueryable{
				EQ: gen_mgr.GenRestTrue,
			},
		}),
		field.Enum("en").Values("1", "2", "3").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),
	}
}
