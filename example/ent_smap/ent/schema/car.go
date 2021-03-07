package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("registered_at").StructTag(`time_format:"2006-01-02 15:04:05" time_utc:"1"`),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return nil
}
