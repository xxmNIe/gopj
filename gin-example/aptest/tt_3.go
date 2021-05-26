package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)




func main() {
	r := gin.Default()

	r.GET("/someYMAL", func(c *gin.Context) {
		c.YAML(http.StatusOK,gin.H{
			"name":"zhangsan" ,
		})
	})

	//返回protobuff
	r.GET("/someproto", func(c *gin.Context) {
		label := "Label"

		reps :=[]int64{int64(1),int64(2)}

		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		c.ProtoBuf(http.StatusOK,data)
	})

	r.Run(":8080")
}

