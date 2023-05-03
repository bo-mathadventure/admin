package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	// reportedUserUuid reportedUserComment reporterUserUuid roomUrl
	return []ent.Field{
		field.Text("reportedUserComment"),
		field.String("roomUrl"),
		field.Bool("hide").Default(false),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("reportedUser", User.Type).Ref("reported").Unique(),
		edge.From("reporterUser", User.Type).Ref("reporter").Unique(),
	}
}
