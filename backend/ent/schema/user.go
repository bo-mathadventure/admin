package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique(),
		field.String("email").Unique(),
		field.String("username"),
		field.String("password"),
		field.String("token"),
		field.String("vCardURL").Optional(),
		field.JSON("permissions", []string{}).Default([]string{}),
		field.JSON("tags", []string{}).Default([]string{}),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reported", Report.Type),
		edge.To("reporter", Report.Type),
	}
}
