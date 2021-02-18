package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Finding holds the schema definition for the Finding entity.
type Finding struct {
	ent.Schema
}

// Fields of the Finding.
func (Finding) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			StructTag(`hcl:"name,attr"`),
		field.String("description").
			StructTag(`hcl:"description,optional"`),
		field.Enum("severity").Values("ZeroSeverity", "LowSeverity", "MediumSeverity", "HighSeverity", "CriticalSeverity", "NullSeverity").
			StructTag(`hcl:"severity,attr"`),
		field.Enum("difficulty").Values("ZeroDifficulty", "NoviceDifficulty", "AdvancedDifficulty", "ExpertDifficulty", "NullDifficulty").
			StructTag(`hcl:"difficulty,attr"`),
		field.JSON("tags", map[string]string{}).
			StructTag(`hcl:"tags,optional"`),
	}
}

// Edges of the Finding.
func (Finding) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("FindingToUser", User.Type).
			StructTag(`hcl:"maintainer,block"`),
		edge.To("FindingToTag", Tag.Type),
		edge.To("FindingToHost", Host.Type),
		edge.To("FindingToScript", Script.Type),
	}
}
