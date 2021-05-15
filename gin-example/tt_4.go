package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)


/**
最重要的 content.Next()
         先执行之后的操作(主体请求啥的)，最后执行中间件之后的操作
 */

func main() {
	r := gin.Default()

	r.Use(mid())

	r.GET("/mid", func(c *gin.Context) {
		get, _ := c.Get("mid")
		fmt.Println("msg is ",get)
		c.JSON(200,gin.H{
			"msg":get,
		})
	})

	r.Run(":8080")
}


//中间件
func mid() gin.HandlerFunc{
	return func(c *gin.Context) {
		t := time.Now()

		fmt.Println("中间件开始执行")

		c.Set("mid","哈哈我")
		c.Next()
		fmt.Println("中间件执行完毕")
		t2 := time.Since(t)
		fmt.Println("用时: ",t2)

	}
}