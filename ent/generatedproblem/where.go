// Code generated by ent, DO NOT EDIT.

package generatedproblem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/nikonhub/ai-learn-cli/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldLTE(FieldID, id))
}

// ProblemText applies equality check predicate on the "problem_text" field. It's identical to ProblemTextEQ.
func ProblemText(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldEQ(FieldProblemText, v))
}

// GeneratedOn applies equality check predicate on the "generated_on" field. It's identical to GeneratedOnEQ.
func GeneratedOn(v time.Time) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldEQ(FieldGeneratedOn, v))
}

// ProblemTextEQ applies the EQ predicate on the "problem_text" field.
func ProblemTextEQ(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldEQ(FieldProblemText, v))
}

// ProblemTextNEQ applies the NEQ predicate on the "problem_text" field.
func ProblemTextNEQ(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldNEQ(FieldProblemText, v))
}

// ProblemTextIn applies the In predicate on the "problem_text" field.
func ProblemTextIn(vs ...string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldIn(FieldProblemText, vs...))
}

// ProblemTextNotIn applies the NotIn predicate on the "problem_text" field.
func ProblemTextNotIn(vs ...string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldNotIn(FieldProblemText, vs...))
}

// ProblemTextGT applies the GT predicate on the "problem_text" field.
func ProblemTextGT(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldGT(FieldProblemText, v))
}

// ProblemTextGTE applies the GTE predicate on the "problem_text" field.
func ProblemTextGTE(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldGTE(FieldProblemText, v))
}

// ProblemTextLT applies the LT predicate on the "problem_text" field.
func ProblemTextLT(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldLT(FieldProblemText, v))
}

// ProblemTextLTE applies the LTE predicate on the "problem_text" field.
func ProblemTextLTE(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldLTE(FieldProblemText, v))
}

// ProblemTextContains applies the Contains predicate on the "problem_text" field.
func ProblemTextContains(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldContains(FieldProblemText, v))
}

// ProblemTextHasPrefix applies the HasPrefix predicate on the "problem_text" field.
func ProblemTextHasPrefix(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldHasPrefix(FieldProblemText, v))
}

// ProblemTextHasSuffix applies the HasSuffix predicate on the "problem_text" field.
func ProblemTextHasSuffix(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldHasSuffix(FieldProblemText, v))
}

// ProblemTextEqualFold applies the EqualFold predicate on the "problem_text" field.
func ProblemTextEqualFold(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldEqualFold(FieldProblemText, v))
}

// ProblemTextContainsFold applies the ContainsFold predicate on the "problem_text" field.
func ProblemTextContainsFold(v string) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldContainsFold(FieldProblemText, v))
}

// GeneratedOnEQ applies the EQ predicate on the "generated_on" field.
func GeneratedOnEQ(v time.Time) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldEQ(FieldGeneratedOn, v))
}

// GeneratedOnNEQ applies the NEQ predicate on the "generated_on" field.
func GeneratedOnNEQ(v time.Time) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldNEQ(FieldGeneratedOn, v))
}

// GeneratedOnIn applies the In predicate on the "generated_on" field.
func GeneratedOnIn(vs ...time.Time) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldIn(FieldGeneratedOn, vs...))
}

// GeneratedOnNotIn applies the NotIn predicate on the "generated_on" field.
func GeneratedOnNotIn(vs ...time.Time) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldNotIn(FieldGeneratedOn, vs...))
}

// GeneratedOnGT applies the GT predicate on the "generated_on" field.
func GeneratedOnGT(v time.Time) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldGT(FieldGeneratedOn, v))
}

// GeneratedOnGTE applies the GTE predicate on the "generated_on" field.
func GeneratedOnGTE(v time.Time) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldGTE(FieldGeneratedOn, v))
}

// GeneratedOnLT applies the LT predicate on the "generated_on" field.
func GeneratedOnLT(v time.Time) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldLT(FieldGeneratedOn, v))
}

// GeneratedOnLTE applies the LTE predicate on the "generated_on" field.
func GeneratedOnLTE(v time.Time) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.FieldLTE(FieldGeneratedOn, v))
}

// HasItem applies the HasEdge predicate on the "item" edge.
func HasItem() predicate.GeneratedProblem {
	return predicate.GeneratedProblem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ItemTable, ItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemWith applies the HasEdge predicate on the "item" edge with a given conditions (other predicates).
func HasItemWith(preds ...predicate.Item) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(func(s *sql.Selector) {
		step := newItemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GeneratedProblem) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GeneratedProblem) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GeneratedProblem) predicate.GeneratedProblem {
	return predicate.GeneratedProblem(sql.NotPredicates(p))
}
