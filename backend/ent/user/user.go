// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldToken holds the string denoting the token field in the database.
	FieldToken = "token"
	// FieldVCardURL holds the string denoting the vcardurl field in the database.
	FieldVCardURL = "v_card_url"
	// FieldPermissions holds the string denoting the permissions field in the database.
	FieldPermissions = "permissions"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// EdgeReported holds the string denoting the reported edge name in mutations.
	EdgeReported = "reported"
	// EdgeReporter holds the string denoting the reporter edge name in mutations.
	EdgeReporter = "reporter"
	// Table holds the table name of the user in the database.
	Table = "users"
	// ReportedTable is the table that holds the reported relation/edge.
	ReportedTable = "reports"
	// ReportedInverseTable is the table name for the Report entity.
	// It exists in this package in order to avoid circular dependency with the "report" package.
	ReportedInverseTable = "reports"
	// ReportedColumn is the table column denoting the reported relation/edge.
	ReportedColumn = "user_reported"
	// ReporterTable is the table that holds the reporter relation/edge.
	ReporterTable = "reports"
	// ReporterInverseTable is the table name for the Report entity.
	// It exists in this package in order to avoid circular dependency with the "report" package.
	ReporterInverseTable = "reports"
	// ReporterColumn is the table column denoting the reporter relation/edge.
	ReporterColumn = "user_reporter"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldEmail,
	FieldUsername,
	FieldPassword,
	FieldToken,
	FieldVCardURL,
	FieldPermissions,
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
	// DefaultPermissions holds the default value on creation for the "permissions" field.
	DefaultPermissions []string
	// DefaultTags holds the default value on creation for the "tags" field.
	DefaultTags []string
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUUID orders the results by the uuid field.
func ByUUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUUID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByToken orders the results by the token field.
func ByToken(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToken, opts...).ToFunc()
}

// ByVCardURL orders the results by the vCardURL field.
func ByVCardURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVCardURL, opts...).ToFunc()
}

// ByCreatedAt orders the results by the createdAt field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByReportedCount orders the results by reported count.
func ByReportedCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReportedStep(), opts...)
	}
}

// ByReported orders the results by reported terms.
func ByReported(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReportedStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByReporterCount orders the results by reporter count.
func ByReporterCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReporterStep(), opts...)
	}
}

// ByReporter orders the results by reporter terms.
func ByReporter(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReporterStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newReportedStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReportedInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReportedTable, ReportedColumn),
	)
}
func newReporterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReporterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReporterTable, ReporterColumn),
	)
}