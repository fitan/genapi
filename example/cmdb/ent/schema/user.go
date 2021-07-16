package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/fitan/genapi/pkg/gen_mgr"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		gen_mgr.RestNodeOp{
			Paging: gen_mgr.Paging{
				Open:     true,
				Must:     false,
				MaxLimit: 10,
			},
			Order: gen_mgr.Order{
				DefaultAcsOrder:   nil,
				DefaultDescOrder:  nil,
				OpenOptionalOrder: false,
				OptionalOrder:     nil,
			},
			Method: gen_mgr.NodeMethod{
				GetOne:     gen_mgr.NodeMethodOp{
					Has:       gen_mgr.GenRestTrue,
				},
				GetList:    gen_mgr.NodeMethodOp{
					Has:       gen_mgr.GenRestTrue,
				},
				CreateOne:  gen_mgr.NodeMethodOp{},
				CreateList: gen_mgr.NodeMethodOp{},
				UpdateOne:  gen_mgr.NodeMethodOp{},
				UpdateList: gen_mgr.NodeMethodOp{},
				DeleteOne:  gen_mgr.NodeMethodOp{},
				DeleteList: gen_mgr.NodeMethodOp{},
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
		field.String("name").Annotations(gen_mgr.RestFieldOp{
			FieldQueryable: gen_mgr.FieldQueryable{
				EQ:           gen_mgr.GenRestTrue,
				NEQ:          gen_mgr.GenRestTrue,
				GT:           gen_mgr.GenRestTrue,
				GTE:          gen_mgr.GenRestTrue,
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
				Order:        gen_mgr.GenRestTrue,
			},
			FieldOperability: gen_mgr.FieldOperability{
				Selete: 0,
				Create: 0,
				Update: 0,
			},
		}),
		field.String("password").Optional().Sensitive().Annotations(gen_mgr.RestFieldOp{
			FieldQueryable: gen_mgr.FieldQueryable{},
			FieldOperability: gen_mgr.FieldOperability{
				Selete: gen_mgr.GenRestFalse,
			},
		}),
		field.String("email").Annotations(gen_mgr.RestFieldOp{
			FieldQueryable: gen_mgr.FieldQueryable{
				EQ:           0,
				NEQ:          0,
				GT:           0,
				GTE:          0,
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
				Order:        gen_mgr.GenRestTrue,
			},
			FieldOperability: gen_mgr.FieldOperability{
				Selete: gen_mgr.GenRestFalse,
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
