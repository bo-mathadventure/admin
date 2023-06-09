// Code generated by ent, DO NOT EDIT.

package textures

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the textures type in the database.
	Label = "textures"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTexture holds the string denoting the texture field in the database.
	FieldTexture = "texture"
	// FieldLayer holds the string denoting the layer field in the database.
	FieldLayer = "layer"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the textures in the database.
	Table = "textures"
)

// Columns holds all SQL columns for textures fields.
var Columns = []string{
	FieldID,
	FieldTexture,
	FieldLayer,
	FieldURL,
	FieldTags,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the Textures queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTexture orders the results by the texture field.
func ByTexture(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTexture, opts...).ToFunc()
}

// ByLayer orders the results by the layer field.
func ByLayer(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLayer, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}
