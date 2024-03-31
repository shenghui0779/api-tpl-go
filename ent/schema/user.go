package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Table("user"),
		entsql.WithComments(true),
		schema.Comment("用户表"),
	}
}

// Mixin of the User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone").MaxLen(16).Comment("手机号"),
		field.String("nickname").MaxLen(32).Default("").Comment("昵称"),
		field.String("password").MaxLen(32).Default("").Comment("密码"),
		field.String("salt").MaxLen(16).Default("").Comment("加密盐"),
		field.Time("login_at").Optional().Comment("登录时间"),
		field.String("login_token").MaxLen(32).Default("").Comment("登录Token"),
	}
}

// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("phone").Unique().StorageKey("uniq_phone"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{}
}
