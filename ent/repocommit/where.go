// Code generated by entc, DO NOT EDIT.

package repocommit

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Revision applies equality check predicate on the "revision" field. It's identical to RevisionEQ.
func Revision(v int) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevision), v))
	})
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// PgpSignature applies equality check predicate on the "pgp_signature" field. It's identical to PgpSignatureEQ.
func PgpSignature(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPgpSignature), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// TreeHash applies equality check predicate on the "tree_hash" field. It's identical to TreeHashEQ.
func TreeHash(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTreeHash), v))
	})
}

// RevisionEQ applies the EQ predicate on the "revision" field.
func RevisionEQ(v int) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevision), v))
	})
}

// RevisionNEQ applies the NEQ predicate on the "revision" field.
func RevisionNEQ(v int) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRevision), v))
	})
}

// RevisionIn applies the In predicate on the "revision" field.
func RevisionIn(vs ...int) predicate.RepoCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRevision), v...))
	})
}

// RevisionNotIn applies the NotIn predicate on the "revision" field.
func RevisionNotIn(vs ...int) predicate.RepoCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRevision), v...))
	})
}

// RevisionGT applies the GT predicate on the "revision" field.
func RevisionGT(v int) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRevision), v))
	})
}

// RevisionGTE applies the GTE predicate on the "revision" field.
func RevisionGTE(v int) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRevision), v))
	})
}

// RevisionLT applies the LT predicate on the "revision" field.
func RevisionLT(v int) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRevision), v))
	})
}

// RevisionLTE applies the LTE predicate on the "revision" field.
func RevisionLTE(v int) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRevision), v))
	})
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHash), v))
	})
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.RepoCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHash), v...))
	})
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.RepoCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHash), v...))
	})
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHash), v))
	})
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHash), v))
	})
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHash), v))
	})
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHash), v))
	})
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHash), v))
	})
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHash), v))
	})
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHash), v))
	})
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHash), v))
	})
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHash), v))
	})
}

// PgpSignatureEQ applies the EQ predicate on the "pgp_signature" field.
func PgpSignatureEQ(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPgpSignature), v))
	})
}

// PgpSignatureNEQ applies the NEQ predicate on the "pgp_signature" field.
func PgpSignatureNEQ(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPgpSignature), v))
	})
}

// PgpSignatureIn applies the In predicate on the "pgp_signature" field.
func PgpSignatureIn(vs ...string) predicate.RepoCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPgpSignature), v...))
	})
}

// PgpSignatureNotIn applies the NotIn predicate on the "pgp_signature" field.
func PgpSignatureNotIn(vs ...string) predicate.RepoCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPgpSignature), v...))
	})
}

// PgpSignatureGT applies the GT predicate on the "pgp_signature" field.
func PgpSignatureGT(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPgpSignature), v))
	})
}

// PgpSignatureGTE applies the GTE predicate on the "pgp_signature" field.
func PgpSignatureGTE(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPgpSignature), v))
	})
}

// PgpSignatureLT applies the LT predicate on the "pgp_signature" field.
func PgpSignatureLT(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPgpSignature), v))
	})
}

// PgpSignatureLTE applies the LTE predicate on the "pgp_signature" field.
func PgpSignatureLTE(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPgpSignature), v))
	})
}

// PgpSignatureContains applies the Contains predicate on the "pgp_signature" field.
func PgpSignatureContains(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPgpSignature), v))
	})
}

// PgpSignatureHasPrefix applies the HasPrefix predicate on the "pgp_signature" field.
func PgpSignatureHasPrefix(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPgpSignature), v))
	})
}

// PgpSignatureHasSuffix applies the HasSuffix predicate on the "pgp_signature" field.
func PgpSignatureHasSuffix(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPgpSignature), v))
	})
}

// PgpSignatureEqualFold applies the EqualFold predicate on the "pgp_signature" field.
func PgpSignatureEqualFold(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPgpSignature), v))
	})
}

// PgpSignatureContainsFold applies the ContainsFold predicate on the "pgp_signature" field.
func PgpSignatureContainsFold(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPgpSignature), v))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.RepoCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.RepoCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// TreeHashEQ applies the EQ predicate on the "tree_hash" field.
func TreeHashEQ(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTreeHash), v))
	})
}

// TreeHashNEQ applies the NEQ predicate on the "tree_hash" field.
func TreeHashNEQ(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTreeHash), v))
	})
}

// TreeHashIn applies the In predicate on the "tree_hash" field.
func TreeHashIn(vs ...string) predicate.RepoCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTreeHash), v...))
	})
}

// TreeHashNotIn applies the NotIn predicate on the "tree_hash" field.
func TreeHashNotIn(vs ...string) predicate.RepoCommit {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.RepoCommit(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTreeHash), v...))
	})
}

// TreeHashGT applies the GT predicate on the "tree_hash" field.
func TreeHashGT(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTreeHash), v))
	})
}

// TreeHashGTE applies the GTE predicate on the "tree_hash" field.
func TreeHashGTE(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTreeHash), v))
	})
}

// TreeHashLT applies the LT predicate on the "tree_hash" field.
func TreeHashLT(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTreeHash), v))
	})
}

// TreeHashLTE applies the LTE predicate on the "tree_hash" field.
func TreeHashLTE(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTreeHash), v))
	})
}

// TreeHashContains applies the Contains predicate on the "tree_hash" field.
func TreeHashContains(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTreeHash), v))
	})
}

// TreeHashHasPrefix applies the HasPrefix predicate on the "tree_hash" field.
func TreeHashHasPrefix(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTreeHash), v))
	})
}

// TreeHashHasSuffix applies the HasSuffix predicate on the "tree_hash" field.
func TreeHashHasSuffix(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTreeHash), v))
	})
}

// TreeHashEqualFold applies the EqualFold predicate on the "tree_hash" field.
func TreeHashEqualFold(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTreeHash), v))
	})
}

// TreeHashContainsFold applies the ContainsFold predicate on the "tree_hash" field.
func TreeHashContainsFold(v string) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTreeHash), v))
	})
}

// HasRepoCommitToRepository applies the HasEdge predicate on the "RepoCommitToRepository" edge.
func HasRepoCommitToRepository() predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepoCommitToRepositoryTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepoCommitToRepositoryTable, RepoCommitToRepositoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepoCommitToRepositoryWith applies the HasEdge predicate on the "RepoCommitToRepository" edge with a given conditions (other predicates).
func HasRepoCommitToRepositoryWith(preds ...predicate.Repository) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RepoCommitToRepositoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RepoCommitToRepositoryTable, RepoCommitToRepositoryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RepoCommit) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RepoCommit) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
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
func Not(p predicate.RepoCommit) predicate.RepoCommit {
	return predicate.RepoCommit(func(s *sql.Selector) {
		p(s.Not())
	})
}
