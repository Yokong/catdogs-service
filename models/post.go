package models

import "time"

type Post struct {
	Id         int    `xorm:"pk autoincr"`
	Title      string `xorm:"varchar(120)"`
	Author     string `xorm:"varchar(24)"`
	Content    struct{}
	Source     string `xorm:"varchar(10)"`
	Timestamp  int
	CreateTime time.Time `xorm:"created"`
}

func (p *Post) Set() error {
	_, err := Db.Insert(p)
	return err
}

func (p *Post) Get() (has bool, err error) {
	has, err = Db.Get(p)
	return has, err
}
