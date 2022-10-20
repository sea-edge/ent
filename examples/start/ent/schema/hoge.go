package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Hoge struct {
	ent.Schema
}

// Fields of the Hoge.
func (Hoge) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50).NotEmpty(),
	}
}

func (Hoge) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ULIDMixin{},
		TimeMixin{},
	}
}

func (Hoge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hoge_administrators", HogeAdministrator.Type).StorageKey(edge.Column("hoge_id")),
	}
}
