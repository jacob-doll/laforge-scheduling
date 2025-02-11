// Code generated by ent, DO NOT EDIT.

package buildcommit

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Revision applies equality check predicate on the "revision" field. It's identical to RevisionEQ.
func Revision(v int) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevision), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.BuildCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.BuildCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// RevisionEQ applies the EQ predicate on the "revision" field.
func RevisionEQ(v int) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevision), v))
	})
}

// RevisionNEQ applies the NEQ predicate on the "revision" field.
func RevisionNEQ(v int) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRevision), v))
	})
}

// RevisionIn applies the In predicate on the "revision" field.
func RevisionIn(vs ...int) predicate.BuildCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRevision), v...))
	})
}

// RevisionNotIn applies the NotIn predicate on the "revision" field.
func RevisionNotIn(vs ...int) predicate.BuildCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRevision), v...))
	})
}

// RevisionGT applies the GT predicate on the "revision" field.
func RevisionGT(v int) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRevision), v))
	})
}

// RevisionGTE applies the GTE predicate on the "revision" field.
func RevisionGTE(v int) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRevision), v))
	})
}

// RevisionLT applies the LT predicate on the "revision" field.
func RevisionLT(v int) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRevision), v))
	})
}

// RevisionLTE applies the LTE predicate on the "revision" field.
func RevisionLTE(v int) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRevision), v))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.BuildCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.BuildCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BuildCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BuildCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// HasBuildCommitToBuild applies the HasEdge predicate on the "BuildCommitToBuild" edge.
func HasBuildCommitToBuild() predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildCommitToBuildTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BuildCommitToBuildTable, BuildCommitToBuildColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBuildCommitToBuildWith applies the HasEdge predicate on the "BuildCommitToBuild" edge with a given conditions (other predicates).
func HasBuildCommitToBuildWith(preds ...predicate.Build) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildCommitToBuildInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, BuildCommitToBuildTable, BuildCommitToBuildColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBuildCommitToServerTask applies the HasEdge predicate on the "BuildCommitToServerTask" edge.
func HasBuildCommitToServerTask() predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildCommitToServerTaskTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BuildCommitToServerTaskTable, BuildCommitToServerTaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBuildCommitToServerTaskWith applies the HasEdge predicate on the "BuildCommitToServerTask" edge with a given conditions (other predicates).
func HasBuildCommitToServerTaskWith(preds ...predicate.ServerTask) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildCommitToServerTaskInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BuildCommitToServerTaskTable, BuildCommitToServerTaskColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBuildCommitToPlanDiffs applies the HasEdge predicate on the "BuildCommitToPlanDiffs" edge.
func HasBuildCommitToPlanDiffs() predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildCommitToPlanDiffsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BuildCommitToPlanDiffsTable, BuildCommitToPlanDiffsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBuildCommitToPlanDiffsWith applies the HasEdge predicate on the "BuildCommitToPlanDiffs" edge with a given conditions (other predicates).
func HasBuildCommitToPlanDiffsWith(preds ...predicate.PlanDiff) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildCommitToPlanDiffsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, BuildCommitToPlanDiffsTable, BuildCommitToPlanDiffsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BuildCommit) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BuildCommit) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
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
func Not(p predicate.BuildCommit) predicate.BuildCommit {
	return predicate.BuildCommit(func(s *sql.Selector) {
		p(s.Not())
	})
}
