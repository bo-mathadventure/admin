package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").DefaultFunc(uuid.NewString),
		field.String("action"),
		field.Time("validUntil").Default(func() time.Time {
			return time.Now().Add(time.Duration(24) * time.Hour)
		}),
		field.Bool("send").Default(false),
		field.Time("createdAt").Default(time.Now),
		field.JSON("data", map[string]interface{}{}).Default(func() map[string]interface{} {
			return map[string]interface{}{}
		}),
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("tokens").Unique(),
	}
}
