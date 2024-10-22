package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ItemProgress holds the schema definition for the ItemProgress entity.
type ItemProgress struct {
	ent.Schema
}

// Fields of the ItemProgress.
func (ItemProgress) Fields() []ent.Field {
	return []ent.Field{
		field.Time("next_review_date"),
		field.Int("interval_days").Default(1),
		field.Float("ease_factor").Default(2.5),
		field.Int("streak").Default(0),
	}
}

// Edges of the ItemProgress.
func (ItemProgress) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item", Item.Type).
			Ref("progress").
			Unique(),
	}
}
