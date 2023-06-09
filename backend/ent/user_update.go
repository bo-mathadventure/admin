// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/bo-mathadventure/admin/ent/group"
	"github.com/bo-mathadventure/admin/ent/predicate"
	"github.com/bo-mathadventure/admin/ent/report"
	"github.com/bo-mathadventure/admin/ent/token"
	"github.com/bo-mathadventure/admin/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUUID sets the "uuid" field.
func (uu *UserUpdate) SetUUID(s string) *UserUpdate {
	uu.mutation.SetUUID(s)
	return uu
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUUID(s *string) *UserUpdate {
	if s != nil {
		uu.SetUUID(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetSsoIdentifier sets the "ssoIdentifier" field.
func (uu *UserUpdate) SetSsoIdentifier(s string) *UserUpdate {
	uu.mutation.SetSsoIdentifier(s)
	return uu
}

// SetNillableSsoIdentifier sets the "ssoIdentifier" field if the given value is not nil.
func (uu *UserUpdate) SetNillableSsoIdentifier(s *string) *UserUpdate {
	if s != nil {
		uu.SetSsoIdentifier(*s)
	}
	return uu
}

// ClearSsoIdentifier clears the value of the "ssoIdentifier" field.
func (uu *UserUpdate) ClearSsoIdentifier() *UserUpdate {
	uu.mutation.ClearSsoIdentifier()
	return uu
}

// SetEmailConfirmed sets the "emailConfirmed" field.
func (uu *UserUpdate) SetEmailConfirmed(b bool) *UserUpdate {
	uu.mutation.SetEmailConfirmed(b)
	return uu
}

// SetNillableEmailConfirmed sets the "emailConfirmed" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmailConfirmed(b *bool) *UserUpdate {
	if b != nil {
		uu.SetEmailConfirmed(*b)
	}
	return uu
}

// SetPermissions sets the "permissions" field.
func (uu *UserUpdate) SetPermissions(s []string) *UserUpdate {
	uu.mutation.SetPermissions(s)
	return uu
}

// AppendPermissions appends s to the "permissions" field.
func (uu *UserUpdate) AppendPermissions(s []string) *UserUpdate {
	uu.mutation.AppendPermissions(s)
	return uu
}

// SetTags sets the "tags" field.
func (uu *UserUpdate) SetTags(s []string) *UserUpdate {
	uu.mutation.SetTags(s)
	return uu
}

// AppendTags appends s to the "tags" field.
func (uu *UserUpdate) AppendTags(s []string) *UserUpdate {
	uu.mutation.AppendTags(s)
	return uu
}

// SetLastLogin sets the "lastLogin" field.
func (uu *UserUpdate) SetLastLogin(t time.Time) *UserUpdate {
	uu.mutation.SetLastLogin(t)
	return uu
}

// SetNillableLastLogin sets the "lastLogin" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastLogin(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetLastLogin(*t)
	}
	return uu
}

// ClearLastLogin clears the value of the "lastLogin" field.
func (uu *UserUpdate) ClearLastLogin() *UserUpdate {
	uu.mutation.ClearLastLogin()
	return uu
}

// SetCreatedAt sets the "createdAt" field.
func (uu *UserUpdate) SetCreatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetCreatedAt(t)
	return uu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreatedAt(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetCreatedAt(*t)
	}
	return uu
}

// AddReportedIDs adds the "reported" edge to the Report entity by IDs.
func (uu *UserUpdate) AddReportedIDs(ids ...int) *UserUpdate {
	uu.mutation.AddReportedIDs(ids...)
	return uu
}

// AddReported adds the "reported" edges to the Report entity.
func (uu *UserUpdate) AddReported(r ...*Report) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddReportedIDs(ids...)
}

// AddReporterIDs adds the "reporter" edge to the Report entity by IDs.
func (uu *UserUpdate) AddReporterIDs(ids ...int) *UserUpdate {
	uu.mutation.AddReporterIDs(ids...)
	return uu
}

// AddReporter adds the "reporter" edges to the Report entity.
func (uu *UserUpdate) AddReporter(r ...*Report) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddReporterIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uu *UserUpdate) AddGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.AddGroupIDs(ids...)
	return uu
}

// AddGroups adds the "groups" edges to the Group entity.
func (uu *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupIDs(ids...)
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (uu *UserUpdate) AddTokenIDs(ids ...int) *UserUpdate {
	uu.mutation.AddTokenIDs(ids...)
	return uu
}

// AddTokens adds the "tokens" edges to the Token entity.
func (uu *UserUpdate) AddTokens(t ...*Token) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTokenIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearReported clears all "reported" edges to the Report entity.
func (uu *UserUpdate) ClearReported() *UserUpdate {
	uu.mutation.ClearReported()
	return uu
}

// RemoveReportedIDs removes the "reported" edge to Report entities by IDs.
func (uu *UserUpdate) RemoveReportedIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveReportedIDs(ids...)
	return uu
}

// RemoveReported removes "reported" edges to Report entities.
func (uu *UserUpdate) RemoveReported(r ...*Report) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveReportedIDs(ids...)
}

// ClearReporter clears all "reporter" edges to the Report entity.
func (uu *UserUpdate) ClearReporter() *UserUpdate {
	uu.mutation.ClearReporter()
	return uu
}

// RemoveReporterIDs removes the "reporter" edge to Report entities by IDs.
func (uu *UserUpdate) RemoveReporterIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveReporterIDs(ids...)
	return uu
}

// RemoveReporter removes "reporter" edges to Report entities.
func (uu *UserUpdate) RemoveReporter(r ...*Report) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveReporterIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uu *UserUpdate) ClearGroups() *UserUpdate {
	uu.mutation.ClearGroups()
	return uu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uu *UserUpdate) RemoveGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveGroupIDs(ids...)
	return uu
}

// RemoveGroups removes "groups" edges to Group entities.
func (uu *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
}

// ClearTokens clears all "tokens" edges to the Token entity.
func (uu *UserUpdate) ClearTokens() *UserUpdate {
	uu.mutation.ClearTokens()
	return uu
}

// RemoveTokenIDs removes the "tokens" edge to Token entities by IDs.
func (uu *UserUpdate) RemoveTokenIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveTokenIDs(ids...)
	return uu
}

// RemoveTokens removes "tokens" edges to Token entities.
func (uu *UserUpdate) RemoveTokens(t ...*Token) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTokenIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UUID(); ok {
		_spec.SetField(user.FieldUUID, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.SsoIdentifier(); ok {
		_spec.SetField(user.FieldSsoIdentifier, field.TypeString, value)
	}
	if uu.mutation.SsoIdentifierCleared() {
		_spec.ClearField(user.FieldSsoIdentifier, field.TypeString)
	}
	if value, ok := uu.mutation.EmailConfirmed(); ok {
		_spec.SetField(user.FieldEmailConfirmed, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Permissions(); ok {
		_spec.SetField(user.FieldPermissions, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedPermissions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldPermissions, value)
		})
	}
	if value, ok := uu.mutation.Tags(); ok {
		_spec.SetField(user.FieldTags, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldTags, value)
		})
	}
	if value, ok := uu.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
	}
	if uu.mutation.LastLoginCleared() {
		_spec.ClearField(user.FieldLastLogin, field.TypeTime)
	}
	if value, ok := uu.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if uu.mutation.ReportedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReportedTable,
			Columns: []string{user.ReportedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedReportedIDs(); len(nodes) > 0 && !uu.mutation.ReportedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReportedTable,
			Columns: []string{user.ReportedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ReportedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReportedTable,
			Columns: []string{user.ReportedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.ReporterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReporterTable,
			Columns: []string{user.ReporterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedReporterIDs(); len(nodes) > 0 && !uu.mutation.ReporterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReporterTable,
			Columns: []string{user.ReporterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.ReporterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReporterTable,
			Columns: []string{user.ReporterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTokensIDs(); len(nodes) > 0 && !uu.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUUID sets the "uuid" field.
func (uuo *UserUpdateOne) SetUUID(s string) *UserUpdateOne {
	uuo.mutation.SetUUID(s)
	return uuo
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUUID(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUUID(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetSsoIdentifier sets the "ssoIdentifier" field.
func (uuo *UserUpdateOne) SetSsoIdentifier(s string) *UserUpdateOne {
	uuo.mutation.SetSsoIdentifier(s)
	return uuo
}

// SetNillableSsoIdentifier sets the "ssoIdentifier" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSsoIdentifier(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetSsoIdentifier(*s)
	}
	return uuo
}

// ClearSsoIdentifier clears the value of the "ssoIdentifier" field.
func (uuo *UserUpdateOne) ClearSsoIdentifier() *UserUpdateOne {
	uuo.mutation.ClearSsoIdentifier()
	return uuo
}

// SetEmailConfirmed sets the "emailConfirmed" field.
func (uuo *UserUpdateOne) SetEmailConfirmed(b bool) *UserUpdateOne {
	uuo.mutation.SetEmailConfirmed(b)
	return uuo
}

// SetNillableEmailConfirmed sets the "emailConfirmed" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmailConfirmed(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetEmailConfirmed(*b)
	}
	return uuo
}

// SetPermissions sets the "permissions" field.
func (uuo *UserUpdateOne) SetPermissions(s []string) *UserUpdateOne {
	uuo.mutation.SetPermissions(s)
	return uuo
}

// AppendPermissions appends s to the "permissions" field.
func (uuo *UserUpdateOne) AppendPermissions(s []string) *UserUpdateOne {
	uuo.mutation.AppendPermissions(s)
	return uuo
}

// SetTags sets the "tags" field.
func (uuo *UserUpdateOne) SetTags(s []string) *UserUpdateOne {
	uuo.mutation.SetTags(s)
	return uuo
}

// AppendTags appends s to the "tags" field.
func (uuo *UserUpdateOne) AppendTags(s []string) *UserUpdateOne {
	uuo.mutation.AppendTags(s)
	return uuo
}

// SetLastLogin sets the "lastLogin" field.
func (uuo *UserUpdateOne) SetLastLogin(t time.Time) *UserUpdateOne {
	uuo.mutation.SetLastLogin(t)
	return uuo
}

// SetNillableLastLogin sets the "lastLogin" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastLogin(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetLastLogin(*t)
	}
	return uuo
}

// ClearLastLogin clears the value of the "lastLogin" field.
func (uuo *UserUpdateOne) ClearLastLogin() *UserUpdateOne {
	uuo.mutation.ClearLastLogin()
	return uuo
}

// SetCreatedAt sets the "createdAt" field.
func (uuo *UserUpdateOne) SetCreatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetCreatedAt(t)
	return uuo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreatedAt(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetCreatedAt(*t)
	}
	return uuo
}

// AddReportedIDs adds the "reported" edge to the Report entity by IDs.
func (uuo *UserUpdateOne) AddReportedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddReportedIDs(ids...)
	return uuo
}

// AddReported adds the "reported" edges to the Report entity.
func (uuo *UserUpdateOne) AddReported(r ...*Report) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddReportedIDs(ids...)
}

// AddReporterIDs adds the "reporter" edge to the Report entity by IDs.
func (uuo *UserUpdateOne) AddReporterIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddReporterIDs(ids...)
	return uuo
}

// AddReporter adds the "reporter" edges to the Report entity.
func (uuo *UserUpdateOne) AddReporter(r ...*Report) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddReporterIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddGroupIDs(ids...)
	return uuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (uuo *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupIDs(ids...)
}

// AddTokenIDs adds the "tokens" edge to the Token entity by IDs.
func (uuo *UserUpdateOne) AddTokenIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddTokenIDs(ids...)
	return uuo
}

// AddTokens adds the "tokens" edges to the Token entity.
func (uuo *UserUpdateOne) AddTokens(t ...*Token) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTokenIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearReported clears all "reported" edges to the Report entity.
func (uuo *UserUpdateOne) ClearReported() *UserUpdateOne {
	uuo.mutation.ClearReported()
	return uuo
}

// RemoveReportedIDs removes the "reported" edge to Report entities by IDs.
func (uuo *UserUpdateOne) RemoveReportedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveReportedIDs(ids...)
	return uuo
}

// RemoveReported removes "reported" edges to Report entities.
func (uuo *UserUpdateOne) RemoveReported(r ...*Report) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveReportedIDs(ids...)
}

// ClearReporter clears all "reporter" edges to the Report entity.
func (uuo *UserUpdateOne) ClearReporter() *UserUpdateOne {
	uuo.mutation.ClearReporter()
	return uuo
}

// RemoveReporterIDs removes the "reporter" edge to Report entities by IDs.
func (uuo *UserUpdateOne) RemoveReporterIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveReporterIDs(ids...)
	return uuo
}

// RemoveReporter removes "reporter" edges to Report entities.
func (uuo *UserUpdateOne) RemoveReporter(r ...*Report) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveReporterIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uuo *UserUpdateOne) ClearGroups() *UserUpdateOne {
	uuo.mutation.ClearGroups()
	return uuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveGroupIDs(ids...)
	return uuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (uuo *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
}

// ClearTokens clears all "tokens" edges to the Token entity.
func (uuo *UserUpdateOne) ClearTokens() *UserUpdateOne {
	uuo.mutation.ClearTokens()
	return uuo
}

// RemoveTokenIDs removes the "tokens" edge to Token entities by IDs.
func (uuo *UserUpdateOne) RemoveTokenIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveTokenIDs(ids...)
	return uuo
}

// RemoveTokens removes "tokens" edges to Token entities.
func (uuo *UserUpdateOne) RemoveTokens(t ...*Token) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTokenIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UUID(); ok {
		_spec.SetField(user.FieldUUID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.SsoIdentifier(); ok {
		_spec.SetField(user.FieldSsoIdentifier, field.TypeString, value)
	}
	if uuo.mutation.SsoIdentifierCleared() {
		_spec.ClearField(user.FieldSsoIdentifier, field.TypeString)
	}
	if value, ok := uuo.mutation.EmailConfirmed(); ok {
		_spec.SetField(user.FieldEmailConfirmed, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Permissions(); ok {
		_spec.SetField(user.FieldPermissions, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedPermissions(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldPermissions, value)
		})
	}
	if value, ok := uuo.mutation.Tags(); ok {
		_spec.SetField(user.FieldTags, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldTags, value)
		})
	}
	if value, ok := uuo.mutation.LastLogin(); ok {
		_spec.SetField(user.FieldLastLogin, field.TypeTime, value)
	}
	if uuo.mutation.LastLoginCleared() {
		_spec.ClearField(user.FieldLastLogin, field.TypeTime)
	}
	if value, ok := uuo.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeTime, value)
	}
	if uuo.mutation.ReportedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReportedTable,
			Columns: []string{user.ReportedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedReportedIDs(); len(nodes) > 0 && !uuo.mutation.ReportedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReportedTable,
			Columns: []string{user.ReportedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ReportedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReportedTable,
			Columns: []string{user.ReportedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.ReporterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReporterTable,
			Columns: []string{user.ReporterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedReporterIDs(); len(nodes) > 0 && !uuo.mutation.ReporterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReporterTable,
			Columns: []string{user.ReporterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.ReporterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ReporterTable,
			Columns: []string{user.ReporterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(report.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTokensIDs(); len(nodes) > 0 && !uuo.mutation.TokensCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TokensIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TokensTable,
			Columns: []string{user.TokensColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
