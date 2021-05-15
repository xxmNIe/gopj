package main
//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"log"
//	"net/http"
//)
//
//func main() {
//
//	r :=gin.Default()
//	//限制上传文件最大值
//	r.MaxMultipartMemory = 8<<20
//
//	r.GET("/", func(context *gin.Context) {
//		context.String(http.StatusOK,"hello world")
//	})
//
//	r.GET("/user/:name/*ac", func(context *gin.Context) {
//
//		name:= context.Param("name")
//		ac :=context.Param("ac")
//		context.String(http.StatusOK,name +" is "+ac)
//	})
//
//	r.GET("/welcome", func(context *gin.Context) {
//		//?name= 这种情况会取得name是空的
//		name:=context.DefaultQuery("name","abc")
//		context.String(http.StatusOK,name )
//	})
//	//表单提交
//	r.POST("/form", func(c *gin.Context) {
//		 ptype:= c.DefaultPostForm("type","alert")
//
//		 username :=c.PostForm("username")
//		 password :=c.PostForm("password")
//
//		 //多选框
//		 hobbys := c.PostFormArray("hobby")
//		 c.String(http.StatusOK,fmt.Sprintf("type: %s ;usenrma: %s ;password: %s;habbys: %v",ptype,username,password,hobbys))
//	})
//	// 表单上传
//	r.POST("/upload", func(c *gin.Context) {
//
//		file,_ := c.FormFile("file")
//		//save
//		log.Println(file.Filename)
//		c.SaveUploadedFile(file,file.Filename)
//		c.String(http.StatusOK,"ok")
//	})
//
//	r.POST("/mupload", func(c *gin.Context) {
//
//		form, err := c.MultipartForm()
//		if err != nil {
//			c.String(http.StatusBadRequest,"recv err")
//		}
//		files := form.File["file"]
//
//		for _,file :=range files {
//			if err := c.SaveUploadedFile(file,file.Filename);err !=nil {
//				c.String(http.StatusBadRequest,"save err")
//				return
//			}
//		}
//
//		c.String(http.StatusOK,"ok")
//	})
//
//
//	r.Run(":8080")
//}