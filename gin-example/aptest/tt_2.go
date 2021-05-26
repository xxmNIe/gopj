package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Login struct {
	User string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	PAssword string `form:"password" json:"password" uri:"password"`
}


func main() {
	r := gin.Default()
	//分组路由
	v1 := r.Group("/v1")
	{
		v1.GET("/login",login)
		v1.GET("/submit",)
	}
	v2 := r.Group("v2")
	{
		v2.POST("/login")
		v2.POST("/submit",submit)
	}


	//数据绑定
	r.POST("loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json);err !=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"msessage" : strings.Join([]string{"err : ",err.Error()," \n"},""),
			})
			return
		}
		if json.User != "root" || json.PAssword != "admin" {
			c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"message":"ok"})
	})

	r.Run(":8080")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name","jack")
	c.String(http.StatusOK,fmt.Sprintf("hello %s\n",name))

}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name","lili")
	c.String(http.StatusOK,"hello "+name+"\n")
}