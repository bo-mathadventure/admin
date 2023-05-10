package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique().DefaultFunc(uuid.NewString),
		field.String("email").Unique(),
		field.String("username"),
		field.String("password"),
		field.String("ssoIdentifier").Optional(),
		field.Bool("emailConfirmed").Default(true),
		field.JSON("permissions", []string{}).Default([]string{}),
		field.JSON("tags", []string{}).Default([]string{}),
		field.Time("lastLogin").Optional(),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reported", Report.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.SetNull,
		}),
		edge.To("reporter", Report.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.SetNull,
		}),
		edge.From("groups", Group.Type).Ref("users"),
		edge.To("tokens", Token.Type).Annotations(entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}
