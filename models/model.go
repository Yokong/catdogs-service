package models

import (
	configs "catdogs-service/configs/common"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Db *xorm.Engine

func InitModel() {
	var err error
	Db, err = xorm.NewEngine("mysql", configs.DbAddr)
	if err != nil {
		panic(err)
	}
	Db.SetMaxIdleConns(configs.IdleConns)
	Db.SetMaxOpenConns(configs.OpenConns)

	initTables()
}

func initTables() {
	Db.Sync2(new(User))
	Db.Sync2(new(VerifyCode))
}
