package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fitan/genapi/pkg"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		pkg.RestNodeOp{
			Paging: pkg.Paging{
				Open:     true,
				Must:     false,
				MaxLimit: 10,
			},
			Order: pkg.Order{
				DefaultAcsOrder:   nil,
				DefaultDescOrder:  nil,
				OpenOptionalOrder: false,
				OptionalOrder:     nil,
			},
			Method: pkg.NodeMethod{
				Get:    pkg.GenRestTrue,
				Create: pkg.GenRestTrue,
				Update: pkg.GenRestFalse,
				Delete: pkg.GenRestFalse,
			},
		},
	}
}

//func (User) Mixin() []ent.Mixin {
//	return []ent.Mixin{
//		mixin.Time{},
//	}
//}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(pkg.RestFieldOp{
			FieldQueryable: pkg.FieldQueryable{
				EQ:           pkg.GenRestTrue,
				NEQ:          pkg.GenRestTrue,
				GT:           pkg.GenRestTrue,
				GTE:          pkg.GenRestTrue,
				LT:           0,
				LTE:          0,
				IsNil:        0,
				NotNil:       0,
				EqualFold:    0,
				Contains:     0,
				ContainsFold: 0,
				HasPrefix:    0,
				HasSuffix:    0,
				In:           0,
				NotIn:        0,
			},
			FieldOperability: pkg.FieldOperability{
				Selete: 0,
				Create: 0,
				Update: 0,
			},
		}),
		field.String("password").Optional().Sensitive().Annotations(pkg.RestFieldOp{
			FieldQueryable: pkg.FieldQueryable{},
			FieldOperability: pkg.FieldOperability{
				Selete: pkg.GenRestFalse,
			},
		}),
		field.String("email").Annotations(pkg.RestFieldOp{
			FieldQueryable: pkg.FieldQueryable{},
			FieldOperability: pkg.FieldOperability{
				Selete: pkg.GenRestFalse,
			},
		}),
		field.String("phone").Comment("这是我的电话"),
		field.Enum("role").Values("user", "admin", "tourist"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role_bindings", RoleBinding.Type).Annotations(entsql.Annotation{}),
		edge.To("alert", Alert.Type).Unique(),
	}
}
