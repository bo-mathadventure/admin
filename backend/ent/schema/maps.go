package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Maps holds the schema definition for the Maps entity.
type Maps struct {
	ent.Schema
}

// Fields of the Maps.
func (Maps) Fields() []ent.Field {
	return []ent.Field{
		field.String("roomName"),
		field.String("mapUrl"),
		field.Int("policyNumber"),
		field.String("contactPage"),
		field.JSON("tags", []string{}).Default([]string{}),
		field.Bool("enableChat"),
		field.Bool("enableChatUpload"),
		field.Bool("enableChatOnlineList"),
		field.Bool("enableChatDisconnectedList"),
		field.Bool("canReport"),
		field.Time("expireOn"),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the Maps.
func (Maps) Edges() []ent.Edge {
	return nil
}
