package schema

import (
	"time"

	"ariga.io/atlas/sql/postgres"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Mixin definition

// CommonMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type CommonMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (CommonMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.Postgres: postgres.TypeBigSerial,
		}).Comment("自增ID"),
		field.Time("created_at").
			Immutable().
			Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).Comment("更新时间"),
	}
}
