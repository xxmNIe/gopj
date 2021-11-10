package controllers

import (
	"bg/models"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

type Userreq struct {
	Id int			`json:"id"`
	Name string		`json:"name"`
	Age int			`json:"age"`
	flag int		`json:"flag"`
}

type user_red struct {
	ErrCode int
	Msg string
	Data interface{}
}

func (c *UserController)Post(){
	ret := user_red{}
	user := models.User{}
	var errcode = 200
	var errMsg string
	db := orm.NewOrm()
	defer func() {
		ret.ErrCode = errcode
		ret.Msg = errMsg
		c.Data["json"] = ret
		c.ServeJSON()
	}()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody,&user);err != nil {
		fmt.Println(err)
		fmt.Println(user)
		errcode = 401
		errMsg = "json parm error"
		return
	}
	if user.Id ==0 && (user.Name == "" && user.Age == 0) {
		errcode = 1099
		errMsg = "parme err"
		return
	}
	if err := checkuser(db, &user, 5);err != nil {
		errcode = 666
		errMsg = "id err"
	}

	if err := checkuser(db,&user,5);err != nil {
		fmt.Println(err)
		errcode = 1000
		errMsg = "id exist"
		return
	}
	if err := user.Update(db);err != nil {
		errcode = 3444
		errMsg =  err.Error()
		return
	}
	return
}

func checkuser(o orm.Ormer,user *models.User,maxtry int) error  {
	for i := 0; i<maxtry ; i++ {
		begin, err := o.Begin()
		if err != nil && i ==maxtry{
			return errors.New("chceck error")
		}else if err !=nil {
			continue
		}
		var p []orm.Params
		values, err := begin.Raw("select count(*) from user where name = ?", user.Name).Values(&p)
		if err !=nil && i ==maxtry{
			//log.Logger.Println("exec  select error",err)
			begin.Rollback()
			return errors.New("exec  select error"+err.Error())
		}
		if values == 1 && p[0]["count(*)"] == 1 {
			return errors.New("id already exist")
		}
		begin.Commit()
		break

	}
	return nil
}
