package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Ban holds the schema definition for the Ban entity.
type Ban struct {
	ent.Schema
}

// Fields of the Ban.
func (Ban) Fields() []ent.Field {
	return []ent.Field{
		field.String("check"),
		field.String("message"),
		field.Time("validUntil"),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the Ban.
func (Ban) Edges() []ent.Edge {
	return nil
}
