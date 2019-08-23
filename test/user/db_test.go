package user

import (
	"catdogs-service/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func TestConnDb(t *testing.T) {
	url := "catdogs:catdogs1024:@(118.24.146.34:3306)/catdogs"
	e, err := xorm.NewEngine("mysql", url)
	if err != nil {
		t.Error(err)
	}
	user := models.User{
		Email: "18836617@qq.com",
	}
	_, err = e.Get(&user)
	if err != nil {
		t.Error(err)
	}
	// fmt.Println(has)
	e.SetMaxIdleConns(100)
}
