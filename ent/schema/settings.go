// schema/settings.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		// Key field to store the setting's name, unique to avoid duplicate keys
		field.String("key").
			NotEmpty().
			Unique(),

		// Value field to store the corresponding setting's value
		field.String("value").
			Optional(),
	}
}

// Edges of the Settings.
func (Settings) Edges() []ent.Edge {
	return nil
}
