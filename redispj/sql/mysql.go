package sql

import (
	"database/sql"
	"fmt"

	"log"
	"time"

	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/manager"
	"github.com/didi/gendry/scanner"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func GetMysql() (*sql.DB, error) {
	if db == nil {
		Init("bg", "root", "123456", "127.0.0.1", 3306)
	}
	if db == nil {
		log.Printf("mysql init err:%v\n", err)
		return nil, err
	}
	scanner.SetTagName("json")
	return db, err
}

func Init(dbname, username, password, host string, port int) {
	db, err = manager.New(dbname, username, password, host).
		Set(manager.SetCharset("utf8"),
			manager.SetAllowAllFiles(true),
			manager.SetTimeout(1*time.Second),
			manager.SetReadTimeout(1*time.Second)).
		Port(port).Open(true)
	if err != nil {
		log.Printf("DB open err: %s,%s,%s,%s, port:%d  err:%v\n", dbname, username, password, host, port, err)
	}
}

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `josn:"age"`
	Flag int    `json:"flag"`
}

func GetUser(name string) {
	where := map[string]interface{}{
		"name": name,
	}

	table := "user"

	selectFiled := []string{"id", "name", "age", "flag"}

	cond, vals, err2 := builder.BuildSelect(table, where, selectFiled)
	if err2 != nil {
		log.Println("sql build err:", err)
		return
	}
	fmt.Println(cond)
	fmt.Println(vals...)
	rows, err2 := db.Query(cond, vals...)
	if err2 != nil {
		log.Println("query err: ", err)
	}
	var users []user
	err2 = scanner.Scan(rows, &users)
	if err2 != nil {
		log.Println("scan err: ", err)
	}
	fmt.Println(users)
}
