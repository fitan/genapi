package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Server holds the schema definition for the Server entity.
type Server struct {
	ent.Schema
}

//func (Server) Mixin() []ent.Mixin {
//	return []ent.Mixin{
//		mixin.Time{},
//	}
//}

// Fields of the Server.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		field.String("ip").Unique(),
		field.Enum("machine_type").Values("physical", "virtual"),
		field.Enum("platform_type").Values("zstack", "k8s", "openstack"),
		field.Enum("system_type").Values("linux", "windows"),
	}
}

// Edges of the Server.
func (Server) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", ServiceTree.Type).Ref("servers").Unique(),
	}
}
