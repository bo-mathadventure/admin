package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Announcement holds the schema definition for the Announcement entity.
type Announcement struct {
	ent.Schema
}

// Fields of the Announcement.
func (Announcement) Fields() []ent.Field {
	return []ent.Field{
		field.String("type").Default("ban"),
		field.Text("message"),
		field.Time("createdAt").Default(time.Now),
		field.Time("validUntil").Optional(),
	}
}

// Edges of the Announcement.
func (Announcement) Edges() []ent.Edge {
	return nil
}
