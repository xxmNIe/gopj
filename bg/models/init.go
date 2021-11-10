package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	)

var  UserOrm orm.Ormer
func init () {
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:123456@tcp(127.0.0.1:3306)/bg?charset=utf8")
}



