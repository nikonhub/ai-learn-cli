package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GeneratedProblem holds the schema definition for the GeneratedProblem entity.
type GeneratedProblem struct {
	ent.Schema
}

// Fields of the GeneratedProblem.
func (GeneratedProblem) Fields() []ent.Field {
	return []ent.Field{
		field.String("problem_text").NotEmpty(),
		field.Time("generated_on").Default(time.Now),
	}
}

// Edges of the GeneratedProblem.
func (GeneratedProblem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item", Item.Type).
			Ref("generated_problems").
			Unique(),
	}
}
