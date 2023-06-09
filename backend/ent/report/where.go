// Code generated by ent, DO NOT EDIT.

package report

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/bo-mathadventure/admin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldID, id))
}

// ReportedUserComment applies equality check predicate on the "reportedUserComment" field. It's identical to ReportedUserCommentEQ.
func ReportedUserComment(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldReportedUserComment, v))
}

// RoomUrl applies equality check predicate on the "roomUrl" field. It's identical to RoomUrlEQ.
func RoomUrl(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldRoomUrl, v))
}

// Hide applies equality check predicate on the "hide" field. It's identical to HideEQ.
func Hide(v bool) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldHide, v))
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldCreatedAt, v))
}

// ReportedUserCommentEQ applies the EQ predicate on the "reportedUserComment" field.
func ReportedUserCommentEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldReportedUserComment, v))
}

// ReportedUserCommentNEQ applies the NEQ predicate on the "reportedUserComment" field.
func ReportedUserCommentNEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldReportedUserComment, v))
}

// ReportedUserCommentIn applies the In predicate on the "reportedUserComment" field.
func ReportedUserCommentIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldReportedUserComment, vs...))
}

// ReportedUserCommentNotIn applies the NotIn predicate on the "reportedUserComment" field.
func ReportedUserCommentNotIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldReportedUserComment, vs...))
}

// ReportedUserCommentGT applies the GT predicate on the "reportedUserComment" field.
func ReportedUserCommentGT(v string) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldReportedUserComment, v))
}

// ReportedUserCommentGTE applies the GTE predicate on the "reportedUserComment" field.
func ReportedUserCommentGTE(v string) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldReportedUserComment, v))
}

// ReportedUserCommentLT applies the LT predicate on the "reportedUserComment" field.
func ReportedUserCommentLT(v string) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldReportedUserComment, v))
}

// ReportedUserCommentLTE applies the LTE predicate on the "reportedUserComment" field.
func ReportedUserCommentLTE(v string) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldReportedUserComment, v))
}

// ReportedUserCommentContains applies the Contains predicate on the "reportedUserComment" field.
func ReportedUserCommentContains(v string) predicate.Report {
	return predicate.Report(sql.FieldContains(FieldReportedUserComment, v))
}

// ReportedUserCommentHasPrefix applies the HasPrefix predicate on the "reportedUserComment" field.
func ReportedUserCommentHasPrefix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasPrefix(FieldReportedUserComment, v))
}

// ReportedUserCommentHasSuffix applies the HasSuffix predicate on the "reportedUserComment" field.
func ReportedUserCommentHasSuffix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasSuffix(FieldReportedUserComment, v))
}

// ReportedUserCommentEqualFold applies the EqualFold predicate on the "reportedUserComment" field.
func ReportedUserCommentEqualFold(v string) predicate.Report {
	return predicate.Report(sql.FieldEqualFold(FieldReportedUserComment, v))
}

// ReportedUserCommentContainsFold applies the ContainsFold predicate on the "reportedUserComment" field.
func ReportedUserCommentContainsFold(v string) predicate.Report {
	return predicate.Report(sql.FieldContainsFold(FieldReportedUserComment, v))
}

// RoomUrlEQ applies the EQ predicate on the "roomUrl" field.
func RoomUrlEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldRoomUrl, v))
}

// RoomUrlNEQ applies the NEQ predicate on the "roomUrl" field.
func RoomUrlNEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldRoomUrl, v))
}

// RoomUrlIn applies the In predicate on the "roomUrl" field.
func RoomUrlIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldRoomUrl, vs...))
}

// RoomUrlNotIn applies the NotIn predicate on the "roomUrl" field.
func RoomUrlNotIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldRoomUrl, vs...))
}

// RoomUrlGT applies the GT predicate on the "roomUrl" field.
func RoomUrlGT(v string) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldRoomUrl, v))
}

// RoomUrlGTE applies the GTE predicate on the "roomUrl" field.
func RoomUrlGTE(v string) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldRoomUrl, v))
}

// RoomUrlLT applies the LT predicate on the "roomUrl" field.
func RoomUrlLT(v string) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldRoomUrl, v))
}

// RoomUrlLTE applies the LTE predicate on the "roomUrl" field.
func RoomUrlLTE(v string) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldRoomUrl, v))
}

// RoomUrlContains applies the Contains predicate on the "roomUrl" field.
func RoomUrlContains(v string) predicate.Report {
	return predicate.Report(sql.FieldContains(FieldRoomUrl, v))
}

// RoomUrlHasPrefix applies the HasPrefix predicate on the "roomUrl" field.
func RoomUrlHasPrefix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasPrefix(FieldRoomUrl, v))
}

// RoomUrlHasSuffix applies the HasSuffix predicate on the "roomUrl" field.
func RoomUrlHasSuffix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasSuffix(FieldRoomUrl, v))
}

// RoomUrlEqualFold applies the EqualFold predicate on the "roomUrl" field.
func RoomUrlEqualFold(v string) predicate.Report {
	return predicate.Report(sql.FieldEqualFold(FieldRoomUrl, v))
}

// RoomUrlContainsFold applies the ContainsFold predicate on the "roomUrl" field.
func RoomUrlContainsFold(v string) predicate.Report {
	return predicate.Report(sql.FieldContainsFold(FieldRoomUrl, v))
}

// HideEQ applies the EQ predicate on the "hide" field.
func HideEQ(v bool) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldHide, v))
}

// HideNEQ applies the NEQ predicate on the "hide" field.
func HideNEQ(v bool) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldHide, v))
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldCreatedAt, v))
}

// HasReportedUser applies the HasEdge predicate on the "reportedUser" edge.
func HasReportedUser() predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReportedUserTable, ReportedUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReportedUserWith applies the HasEdge predicate on the "reportedUser" edge with a given conditions (other predicates).
func HasReportedUserWith(preds ...predicate.User) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		step := newReportedUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReporterUser applies the HasEdge predicate on the "reporterUser" edge.
func HasReporterUser() predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReporterUserTable, ReporterUserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReporterUserWith applies the HasEdge predicate on the "reporterUser" edge with a given conditions (other predicates).
func HasReporterUserWith(preds ...predicate.User) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		step := newReporterUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Report) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Report) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
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
func Not(p predicate.Report) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		p(s.Not())
	})
}
