package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// 用户实体的注解
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "plat_user"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_name").Default(""),
		field.String("email").Default(""),
		field.String("tel_num").Default(""),
		field.String("password").Default(""),
		field.String("password_str").Default(""),
		field.Int8("reg_type").Default(0),
		field.Time("register_time").SchemaType(map[string]string{
			dialect.MySQL: "timestamp",
		}),
		field.Int("register_ip").Default(0),
		field.Int8("tel_status").Default(2),
		field.Int8("status").Default(1),
		field.Time("created_at").Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "timestamp",
		}),
		field.Time("updated_at").SchemaType(map[string]string{
			dialect.MySQL: "timestamp",
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
