package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GinFileMiddleware holds the schema definition for the GinFileMiddleware entity.
type GinFileMiddleware struct {
	ent.Schema
}

// Fields of the GinFileMiddleware.
func (GinFileMiddleware) Fields() []ent.Field {
	return []ent.Field{
		field.String("url_path"),
		field.String("file_path"),
		field.Bool("accessed").
			Default(false),
	}
}

// Edges of the GinFileMiddleware.
func (GinFileMiddleware) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("GinFileMiddlewareToProvisionedHost", ProvisionedHost.Type).
			Unique(),
		edge.To("GinFileMiddlewareToProvisioningStep", ProvisioningStep.Type).
			Unique(),
	}
}
