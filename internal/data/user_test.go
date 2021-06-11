package data

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"user/internal/data/ent"
	"user/internal/data/ent/user"
)

func TestUser(t *testing.T) {

	client, err := ent.Open(
		"mysql",
		"work:123456@tcp(192.168.20.222:3306)/dl_member?parseTime=true",
	)
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	all, err := client.User.Query().Where(user.UserNameEQ("30000000068")).All(ctx)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(all)
}
