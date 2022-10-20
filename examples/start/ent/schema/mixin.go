package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/examples/start/ent/util/ulid"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// TimeMixin implements a mixin declares "created_at" and "updated_at" fields.
type TimeMixin struct {
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// ULIDMixin implements a mixin overrides "id" fields to use pulid.
type ULIDMixin struct {
	mixin.Schema
}

func (ULIDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(ulid.ID("")).
			DefaultFunc(func() ulid.ID { return ulid.New() }),
	}
}
