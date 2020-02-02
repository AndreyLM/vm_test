package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").StructTag(`json:"ID"`),
		field.String("user_name").StructTag(`json:"UserName"`),
		field.String("full_name").StructTag(`json:"FullName"`),
		field.String("city").StructTag(`json:"City"`),
		field.Time("birth_date").StructTag(`json:"BirthDate"`),
		field.String("department").StructTag(`json:"Department"`),
		field.String("gender").StructTag(`json:"Gender"`),
		field.Int("experience_years").StructTag(`json:"ExperienceYears"`).Positive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
