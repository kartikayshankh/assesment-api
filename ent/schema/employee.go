package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Employee holds the schema definition for the Employee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Employee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("ID").
			Unique(),
		field.String("name"),
		field.String("position"),
		field.Float("salary"),
	}
}

// Edges of the Employee.
func (Employee) Edges() []ent.Edge {
	return nil
}
