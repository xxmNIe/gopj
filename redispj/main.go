package main

import (
	"dbpj/router"
	"dbpj/sql"
	_ "dbpj/tool"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.Use(gin.Recovery(), gin.Logger())

	engine.GET("/set", router.Set)
	engine.GET("/get", router.Get)
	//pprof.Register(engine)
	_, err := sql.GetMysql()
	if err != nil {
		fmt.Println(err)
	}
	sql.GetUser("xx")
	log.Fatal(engine.Run(":9999"))

}
