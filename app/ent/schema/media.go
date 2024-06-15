package schema

import (
	"api/lib/db"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Annotations of the Media.
func (Media) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("media"),
		entsql.WithComments(true),
		schema.Comment("媒体表"),
	}
}

// Mixin of the Media.
func (Media) Mixin() []ent.Mixin {
	return []ent.Mixin{
		db.CommonMixin{},
	}
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.String("media_id").Default("").Comment("媒体ID"),
		field.String("file_name").Default("").Comment("文件名称"),
		field.Int64("file_id").Optional().Default(0).Comment("文件ID"),
	}
}

// Indexes of the Media.
func (Media) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("media_id").Unique().StorageKey("media_id_key"),
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("file", File.Type).Field("file_id").Ref("medias").Unique(),
	}
}
