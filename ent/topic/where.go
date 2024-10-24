// Code generated by ent, DO NOT EDIT.

package topic

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/nikonhub/ai-learn-cli/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Topic {
	return predicate.Topic(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Topic {
	return predicate.Topic(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Topic {
	return predicate.Topic(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Topic {
	return predicate.Topic(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Topic {
	return predicate.Topic(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Topic {
	return predicate.Topic(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Topic {
	return predicate.Topic(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldName, v))
}

// Instructions applies equality check predicate on the "instructions" field. It's identical to InstructionsEQ.
func Instructions(v string) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldInstructions, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Topic {
	return predicate.Topic(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Topic {
	return predicate.Topic(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Topic {
	return predicate.Topic(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Topic {
	return predicate.Topic(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Topic {
	return predicate.Topic(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Topic {
	return predicate.Topic(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Topic {
	return predicate.Topic(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Topic {
	return predicate.Topic(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Topic {
	return predicate.Topic(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Topic {
	return predicate.Topic(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Topic {
	return predicate.Topic(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Topic {
	return predicate.Topic(sql.FieldContainsFold(FieldName, v))
}

// InstructionsEQ applies the EQ predicate on the "instructions" field.
func InstructionsEQ(v string) predicate.Topic {
	return predicate.Topic(sql.FieldEQ(FieldInstructions, v))
}

// InstructionsNEQ applies the NEQ predicate on the "instructions" field.
func InstructionsNEQ(v string) predicate.Topic {
	return predicate.Topic(sql.FieldNEQ(FieldInstructions, v))
}

// InstructionsIn applies the In predicate on the "instructions" field.
func InstructionsIn(vs ...string) predicate.Topic {
	return predicate.Topic(sql.FieldIn(FieldInstructions, vs...))
}

// InstructionsNotIn applies the NotIn predicate on the "instructions" field.
func InstructionsNotIn(vs ...string) predicate.Topic {
	return predicate.Topic(sql.FieldNotIn(FieldInstructions, vs...))
}

// InstructionsGT applies the GT predicate on the "instructions" field.
func InstructionsGT(v string) predicate.Topic {
	return predicate.Topic(sql.FieldGT(FieldInstructions, v))
}

// InstructionsGTE applies the GTE predicate on the "instructions" field.
func InstructionsGTE(v string) predicate.Topic {
	return predicate.Topic(sql.FieldGTE(FieldInstructions, v))
}

// InstructionsLT applies the LT predicate on the "instructions" field.
func InstructionsLT(v string) predicate.Topic {
	return predicate.Topic(sql.FieldLT(FieldInstructions, v))
}

// InstructionsLTE applies the LTE predicate on the "instructions" field.
func InstructionsLTE(v string) predicate.Topic {
	return predicate.Topic(sql.FieldLTE(FieldInstructions, v))
}

// InstructionsContains applies the Contains predicate on the "instructions" field.
func InstructionsContains(v string) predicate.Topic {
	return predicate.Topic(sql.FieldContains(FieldInstructions, v))
}

// InstructionsHasPrefix applies the HasPrefix predicate on the "instructions" field.
func InstructionsHasPrefix(v string) predicate.Topic {
	return predicate.Topic(sql.FieldHasPrefix(FieldInstructions, v))
}

// InstructionsHasSuffix applies the HasSuffix predicate on the "instructions" field.
func InstructionsHasSuffix(v string) predicate.Topic {
	return predicate.Topic(sql.FieldHasSuffix(FieldInstructions, v))
}

// InstructionsEqualFold applies the EqualFold predicate on the "instructions" field.
func InstructionsEqualFold(v string) predicate.Topic {
	return predicate.Topic(sql.FieldEqualFold(FieldInstructions, v))
}

// InstructionsContainsFold applies the ContainsFold predicate on the "instructions" field.
func InstructionsContainsFold(v string) predicate.Topic {
	return predicate.Topic(sql.FieldContainsFold(FieldInstructions, v))
}

// HasItems applies the HasEdge predicate on the "items" edge.
func HasItems() predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemsWith applies the HasEdge predicate on the "items" edge with a given conditions (other predicates).
func HasItemsWith(preds ...predicate.Item) predicate.Topic {
	return predicate.Topic(func(s *sql.Selector) {
		step := newItemsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Topic) predicate.Topic {
	return predicate.Topic(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Topic) predicate.Topic {
	return predicate.Topic(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Topic) predicate.Topic {
	return predicate.Topic(sql.NotPredicates(p))
}
