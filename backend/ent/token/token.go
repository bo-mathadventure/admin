// Code generated by ent, DO NOT EDIT.

package token

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the token type in the database.
	Label = "token"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// FieldValidUntil holds the string denoting the validuntil field in the database.
	FieldValidUntil = "valid_until"
	// FieldSend holds the string denoting the send field in the database.
	FieldSend = "send"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the token in the database.
	Table = "tokens"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "tokens"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_tokens"
)

// Columns holds all SQL columns for token fields.
var Columns = []string{
	FieldID,
	FieldToken,
	FieldAction,
	FieldValidUntil,
	FieldSend,
	FieldCreatedAt,
	FieldData,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_tokens",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultToken holds the default value on creation for the "token" field.
	DefaultToken func() string
	// DefaultValidUntil holds the default value on creation for the "validUntil" field.
	DefaultValidUntil func() time.Time
	// DefaultSend holds the default value on creation for the "send" field.
	DefaultSend bool
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultData holds the default value on creation for the "data" field.
	DefaultData func() map[string]interface{}
)

// OrderOption defines the ordering options for the Token queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByAction orders the results by the action field.
func ByAction(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAction, opts...).ToFunc()
}

// ByValidUntil orders the results by the validUntil field.
func ByValidUntil(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidUntil, opts...).ToFunc()
}

// BySend orders the results by the send field.
func BySend(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSend, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
