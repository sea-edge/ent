package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type HogeAdministrator struct {
	ent.Schema
}

func (HogeAdministrator) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").MaxLen(255).Optional(),
		field.String("last_name").MaxLen(255).Optional(),
		field.String("email").MaxLen(255).Unique().NotEmpty(),
		field.Bool("is_active").Default(false),
	}
}

func (HogeAdministrator) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ULIDMixin{},
		TimeMixin{},
	}
}

func (HogeAdministrator) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hoge", Hoge.Type).
			Ref("hoge_administrators").
			Unique(),
	}
}
