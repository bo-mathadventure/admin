// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/bo-mathadventure/admin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUUID, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// SsoIdentifier applies equality check predicate on the "ssoIdentifier" field. It's identical to SsoIdentifierEQ.
func SsoIdentifier(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSsoIdentifier, v))
}

// EmailConfirmed applies equality check predicate on the "emailConfirmed" field. It's identical to EmailConfirmedEQ.
func EmailConfirmed(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailConfirmed, v))
}

// LastLogin applies equality check predicate on the "lastLogin" field. It's identical to LastLoginEQ.
func LastLogin(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLogin, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUUID, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// SsoIdentifierEQ applies the EQ predicate on the "ssoIdentifier" field.
func SsoIdentifierEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSsoIdentifier, v))
}

// SsoIdentifierNEQ applies the NEQ predicate on the "ssoIdentifier" field.
func SsoIdentifierNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSsoIdentifier, v))
}

// SsoIdentifierIn applies the In predicate on the "ssoIdentifier" field.
func SsoIdentifierIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldSsoIdentifier, vs...))
}

// SsoIdentifierNotIn applies the NotIn predicate on the "ssoIdentifier" field.
func SsoIdentifierNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSsoIdentifier, vs...))
}

// SsoIdentifierGT applies the GT predicate on the "ssoIdentifier" field.
func SsoIdentifierGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldSsoIdentifier, v))
}

// SsoIdentifierGTE applies the GTE predicate on the "ssoIdentifier" field.
func SsoIdentifierGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldSsoIdentifier, v))
}

// SsoIdentifierLT applies the LT predicate on the "ssoIdentifier" field.
func SsoIdentifierLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldSsoIdentifier, v))
}

// SsoIdentifierLTE applies the LTE predicate on the "ssoIdentifier" field.
func SsoIdentifierLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldSsoIdentifier, v))
}

// SsoIdentifierContains applies the Contains predicate on the "ssoIdentifier" field.
func SsoIdentifierContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldSsoIdentifier, v))
}

// SsoIdentifierHasPrefix applies the HasPrefix predicate on the "ssoIdentifier" field.
func SsoIdentifierHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldSsoIdentifier, v))
}

// SsoIdentifierHasSuffix applies the HasSuffix predicate on the "ssoIdentifier" field.
func SsoIdentifierHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldSsoIdentifier, v))
}

// SsoIdentifierIsNil applies the IsNil predicate on the "ssoIdentifier" field.
func SsoIdentifierIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldSsoIdentifier))
}

// SsoIdentifierNotNil applies the NotNil predicate on the "ssoIdentifier" field.
func SsoIdentifierNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldSsoIdentifier))
}

// SsoIdentifierEqualFold applies the EqualFold predicate on the "ssoIdentifier" field.
func SsoIdentifierEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldSsoIdentifier, v))
}

// SsoIdentifierContainsFold applies the ContainsFold predicate on the "ssoIdentifier" field.
func SsoIdentifierContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldSsoIdentifier, v))
}

// EmailConfirmedEQ applies the EQ predicate on the "emailConfirmed" field.
func EmailConfirmedEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmailConfirmed, v))
}

// EmailConfirmedNEQ applies the NEQ predicate on the "emailConfirmed" field.
func EmailConfirmedNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmailConfirmed, v))
}

// LastLoginEQ applies the EQ predicate on the "lastLogin" field.
func LastLoginEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastLogin, v))
}

// LastLoginNEQ applies the NEQ predicate on the "lastLogin" field.
func LastLoginNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastLogin, v))
}

// LastLoginIn applies the In predicate on the "lastLogin" field.
func LastLoginIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastLogin, vs...))
}

// LastLoginNotIn applies the NotIn predicate on the "lastLogin" field.
func LastLoginNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastLogin, vs...))
}

// LastLoginGT applies the GT predicate on the "lastLogin" field.
func LastLoginGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastLogin, v))
}

// LastLoginGTE applies the GTE predicate on the "lastLogin" field.
func LastLoginGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastLogin, v))
}

// LastLoginLT applies the LT predicate on the "lastLogin" field.
func LastLoginLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastLogin, v))
}

// LastLoginLTE applies the LTE predicate on the "lastLogin" field.
func LastLoginLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastLogin, v))
}

// LastLoginIsNil applies the IsNil predicate on the "lastLogin" field.
func LastLoginIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldLastLogin))
}

// LastLoginNotNil applies the NotNil predicate on the "lastLogin" field.
func LastLoginNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldLastLogin))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldCreatedAt, v))
}

// HasReported applies the HasEdge predicate on the "reported" edge.
func HasReported() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReportedTable, ReportedColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReportedWith applies the HasEdge predicate on the "reported" edge with a given conditions (other predicates).
func HasReportedWith(preds ...predicate.Report) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newReportedStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReporter applies the HasEdge predicate on the "reporter" edge.
func HasReporter() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ReporterTable, ReporterColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReporterWith applies the HasEdge predicate on the "reporter" edge with a given conditions (other predicates).
func HasReporterWith(preds ...predicate.Report) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newReporterStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newGroupsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTokens applies the HasEdge predicate on the "tokens" edge.
func HasTokens() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TokensTable, TokensColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTokensWith applies the HasEdge predicate on the "tokens" edge with a given conditions (other predicates).
func HasTokensWith(preds ...predicate.Token) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newTokensStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
