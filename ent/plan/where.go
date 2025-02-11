// Code generated by ent, DO NOT EDIT.

package plan

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// StepNumber applies equality check predicate on the "step_number" field. It's identical to StepNumberEQ.
func StepNumber(v int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStepNumber), v))
	})
}

// BuildID applies equality check predicate on the "build_id" field. It's identical to BuildIDEQ.
func BuildID(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBuildID), v))
	})
}

// StepNumberEQ applies the EQ predicate on the "step_number" field.
func StepNumberEQ(v int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStepNumber), v))
	})
}

// StepNumberNEQ applies the NEQ predicate on the "step_number" field.
func StepNumberNEQ(v int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStepNumber), v))
	})
}

// StepNumberIn applies the In predicate on the "step_number" field.
func StepNumberIn(vs ...int) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStepNumber), v...))
	})
}

// StepNumberNotIn applies the NotIn predicate on the "step_number" field.
func StepNumberNotIn(vs ...int) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStepNumber), v...))
	})
}

// StepNumberGT applies the GT predicate on the "step_number" field.
func StepNumberGT(v int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStepNumber), v))
	})
}

// StepNumberGTE applies the GTE predicate on the "step_number" field.
func StepNumberGTE(v int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStepNumber), v))
	})
}

// StepNumberLT applies the LT predicate on the "step_number" field.
func StepNumberLT(v int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStepNumber), v))
	})
}

// StepNumberLTE applies the LTE predicate on the "step_number" field.
func StepNumberLTE(v int) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStepNumber), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// BuildIDEQ applies the EQ predicate on the "build_id" field.
func BuildIDEQ(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBuildID), v))
	})
}

// BuildIDNEQ applies the NEQ predicate on the "build_id" field.
func BuildIDNEQ(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBuildID), v))
	})
}

// BuildIDIn applies the In predicate on the "build_id" field.
func BuildIDIn(vs ...string) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBuildID), v...))
	})
}

// BuildIDNotIn applies the NotIn predicate on the "build_id" field.
func BuildIDNotIn(vs ...string) predicate.Plan {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBuildID), v...))
	})
}

// BuildIDGT applies the GT predicate on the "build_id" field.
func BuildIDGT(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBuildID), v))
	})
}

// BuildIDGTE applies the GTE predicate on the "build_id" field.
func BuildIDGTE(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBuildID), v))
	})
}

// BuildIDLT applies the LT predicate on the "build_id" field.
func BuildIDLT(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBuildID), v))
	})
}

// BuildIDLTE applies the LTE predicate on the "build_id" field.
func BuildIDLTE(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBuildID), v))
	})
}

// BuildIDContains applies the Contains predicate on the "build_id" field.
func BuildIDContains(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBuildID), v))
	})
}

// BuildIDHasPrefix applies the HasPrefix predicate on the "build_id" field.
func BuildIDHasPrefix(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBuildID), v))
	})
}

// BuildIDHasSuffix applies the HasSuffix predicate on the "build_id" field.
func BuildIDHasSuffix(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBuildID), v))
	})
}

// BuildIDEqualFold applies the EqualFold predicate on the "build_id" field.
func BuildIDEqualFold(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBuildID), v))
	})
}

// BuildIDContainsFold applies the ContainsFold predicate on the "build_id" field.
func BuildIDContainsFold(v string) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBuildID), v))
	})
}

// HasPrevPlan applies the HasEdge predicate on the "PrevPlan" edge.
func HasPrevPlan() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrevPlanTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PrevPlanTable, PrevPlanPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrevPlanWith applies the HasEdge predicate on the "PrevPlan" edge with a given conditions (other predicates).
func HasPrevPlanWith(preds ...predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PrevPlanTable, PrevPlanPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNextPlan applies the HasEdge predicate on the "NextPlan" edge.
func HasNextPlan() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NextPlanTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, NextPlanTable, NextPlanPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNextPlanWith applies the HasEdge predicate on the "NextPlan" edge with a given conditions (other predicates).
func HasNextPlanWith(preds ...predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, NextPlanTable, NextPlanPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlanToBuild applies the HasEdge predicate on the "PlanToBuild" edge.
func HasPlanToBuild() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToBuildTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PlanToBuildTable, PlanToBuildColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanToBuildWith applies the HasEdge predicate on the "PlanToBuild" edge with a given conditions (other predicates).
func HasPlanToBuildWith(preds ...predicate.Build) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToBuildInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PlanToBuildTable, PlanToBuildColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlanToTeam applies the HasEdge predicate on the "PlanToTeam" edge.
func HasPlanToTeam() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToTeamTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PlanToTeamTable, PlanToTeamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanToTeamWith applies the HasEdge predicate on the "PlanToTeam" edge with a given conditions (other predicates).
func HasPlanToTeamWith(preds ...predicate.Team) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToTeamInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PlanToTeamTable, PlanToTeamColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlanToProvisionedNetwork applies the HasEdge predicate on the "PlanToProvisionedNetwork" edge.
func HasPlanToProvisionedNetwork() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToProvisionedNetworkTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PlanToProvisionedNetworkTable, PlanToProvisionedNetworkColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanToProvisionedNetworkWith applies the HasEdge predicate on the "PlanToProvisionedNetwork" edge with a given conditions (other predicates).
func HasPlanToProvisionedNetworkWith(preds ...predicate.ProvisionedNetwork) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToProvisionedNetworkInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PlanToProvisionedNetworkTable, PlanToProvisionedNetworkColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlanToProvisionedHost applies the HasEdge predicate on the "PlanToProvisionedHost" edge.
func HasPlanToProvisionedHost() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToProvisionedHostTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PlanToProvisionedHostTable, PlanToProvisionedHostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanToProvisionedHostWith applies the HasEdge predicate on the "PlanToProvisionedHost" edge with a given conditions (other predicates).
func HasPlanToProvisionedHostWith(preds ...predicate.ProvisionedHost) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToProvisionedHostInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PlanToProvisionedHostTable, PlanToProvisionedHostColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlanToProvisioningStep applies the HasEdge predicate on the "PlanToProvisioningStep" edge.
func HasPlanToProvisioningStep() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToProvisioningStepTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PlanToProvisioningStepTable, PlanToProvisioningStepColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanToProvisioningStepWith applies the HasEdge predicate on the "PlanToProvisioningStep" edge with a given conditions (other predicates).
func HasPlanToProvisioningStepWith(preds ...predicate.ProvisioningStep) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToProvisioningStepInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PlanToProvisioningStepTable, PlanToProvisioningStepColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlanToStatus applies the HasEdge predicate on the "PlanToStatus" edge.
func HasPlanToStatus() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToStatusTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PlanToStatusTable, PlanToStatusColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanToStatusWith applies the HasEdge predicate on the "PlanToStatus" edge with a given conditions (other predicates).
func HasPlanToStatusWith(preds ...predicate.Status) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToStatusInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PlanToStatusTable, PlanToStatusColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPlanToPlanDiffs applies the HasEdge predicate on the "PlanToPlanDiffs" edge.
func HasPlanToPlanDiffs() predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToPlanDiffsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, PlanToPlanDiffsTable, PlanToPlanDiffsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlanToPlanDiffsWith applies the HasEdge predicate on the "PlanToPlanDiffs" edge with a given conditions (other predicates).
func HasPlanToPlanDiffsWith(preds ...predicate.PlanDiff) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PlanToPlanDiffsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, PlanToPlanDiffsTable, PlanToPlanDiffsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
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
func Not(p predicate.Plan) predicate.Plan {
	return predicate.Plan(func(s *sql.Selector) {
		p(s.Not())
	})
}
