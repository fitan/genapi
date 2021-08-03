package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ServiceTree struct {
	ent.Schema
}

func (ServiceTree) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.String("note"),
		field.Enum("type").Values("project", "service"),
	}
}

func (ServiceTree) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("service", ServiceTree.Type).From("project").Unique(),
	}
}