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

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Annotations of the File.
func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("File"),
		entsql.WithComments(true),
		schema.Comment("文件表"),
	}
}

// Mixin of the File.
func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		db.CommonMixin{},
	}
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("fingerprint").Default("").Comment("文件指纹"),
		field.Int64("size").Default(0).Comment("文件大小(b)"),
		field.String("format").Default("").Comment("文件格式"),
		field.Int("width").Default(0).Comment("图片宽(X)"),
		field.Int("height").Default(0).Comment("图片高(X)"),
		field.Int("orientation").Default(0).Comment("图片方向"),
		field.Int("duration").Default(0).Comment("视频时长(秒)"),
	}
}

// Indexes of the File.
func (File) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("fingerprint").Unique().StorageKey("File_fingerprint_key"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("medias", Media.Type),
	}
}
