package models

import (
	configs "catdogs-service/configs/common"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/mgo.v2"
)

var Db *xorm.Engine
var Mdb *mgo.Session

func InitModel() {
	var err error

	// 初始化mysql
	Db, err = xorm.NewEngine("mysql", configs.C.DbAddr)
	if err != nil {
		fmt.Println(configs.C.DbAddr)
		panic(err)
	}
	Db.SetMaxIdleConns(configs.C.IdleConns)
	Db.SetMaxOpenConns(configs.C.OpenConns)

	initTables()
}

func InitMongoDb() {
	// 初始化mongodb
	var err error
	Mdb, err = mgo.Dial(configs.C.MongoDbAddr)
	if err != nil {
		panic(err)
	}
}

func initTables() {
	Db.Sync2(new(User))
	Db.Sync2(new(VerifyCode))
	Db.Sync2(new(Post))
}
