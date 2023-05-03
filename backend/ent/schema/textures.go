package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Textures holds the schema definition for the Textures entity.
type Textures struct {
	ent.Schema
}

// Fields of the Textures.
func (Textures) Fields() []ent.Field {
	return []ent.Field{
		field.String("texture").Unique(),
		field.String("layer"),
		field.String("url"),
		field.JSON("tags", []string{}).Default([]string{}),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the Textures.
func (Textures) Edges() []ent.Edge {
	return nil
}
