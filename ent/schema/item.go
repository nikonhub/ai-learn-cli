package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("description").Optional(),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("topic", Topic.Type).
			Ref("items").
			Unique(),
		edge.To("generated_problems", GeneratedProblem.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
		edge.To("progress", ItemProgress.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
	}
}
