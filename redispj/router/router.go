package router

import (
	"dbpj/tool"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Res struct {
	Msg  string
	Err  string
	Code int
}

func Set(c *gin.Context) {
	res := Res{Code: 200}
	defer handle(c, &res)
	type pair struct {
		Key    string `json:"key" form:"key"`
		Values string `json:"value" form:"value"`
	}
	var p pair
	if err := c.ShouldBind(&p); err != nil {
		log.Println("Set err: ", err)
		res.Code = http.StatusBadRequest
		res.Err = err.Error()
		return
	}
	redis := tool.Red
	if err := redis.SET(p.Key, p.Values); err != nil {
		res.Code = http.StatusInternalServerError
		res.Err = err.Error()
		return
	}
}

func Get(c *gin.Context) {
	res := Res{Code: 200}
	defer handle(c, &res)

	type pair struct {
		Key    string `json:"key" form:"key"`
		Values string `json:"value" form:"value"`
	}
	var p pair
	if err := c.ShouldBind(&p); err != nil {
		log.Println("Get err: ", err)
		res.Code = http.StatusBadRequest
		res.Err = err.Error()
		return
	}
	redis := tool.Red
	value, err := redis.GET(p.Key)
	if err != nil {
		res.Code = http.StatusBadRequest
		res.Err = err.Error()
	}
	res.Msg = value
}

func handle(c *gin.Context, res *Res) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  res.Msg,
		"err":  res.Err,
		"code": res.Code,
	})
}
