package main

import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key")
		if err !=nil {
			c.SetCookie("key","haha",62,"/","localhost",false,true)

		}
		c.JSON(200,gin.H{
			"cookie":cookie,
		})
	})

	r.Run(":8080")
}


