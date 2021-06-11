package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// UserExtend holds the schema definition for the UserExtend entity.
type UserExtend struct {
	ent.Schema
}

// 用户实体的注解
func (UserExtend) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "plat_user_extend"},
	}
}

// Fields of the UserExtend.
func (UserExtend) Fields() []ent.Field {
	return []ent.Field{
		field.Uint("plat_user_id"),
		field.String("real_name").Default(""),
		field.String("id_number").Default(""),
		field.Int8("age").Default(0),
		field.Int8("sex").Default(1),
		field.Int("birth").Default(0),
		field.String("icon").Default(""),
		field.String("nick_name").Default(""),
		field.String("imei").Default(""),
		field.String("oaid").Default(""),
		field.String("device_id").Default(""),
		field.String("system_name").Default(""),
		field.String("system_version").Default(""),
		field.String("adid").Default(""),
		field.String("game_id").Default(""),
		field.Int8("third_platform_id").Default(0),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "timestamp",
		}),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp",
		}),
	}
}

// Edges of the UserExtend.
func (UserExtend) Edges() []ent.Edge {
	return nil
}
