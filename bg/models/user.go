package models

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id int			`json:"id"`
	Name string		`json:"name"`
	Age int			`json:"age"`
	Flag int		`json:"flag"`
}

func (u *User) Update(db orm.Ormer) error{
	str_1 := "update user "
	str_2 := "values("
	flag := false
	parme := []interface{}{}
	if u.Name != ""{
		if flag != true {
			flag = true
			str_1 += "set ( "
		}
		str_1 +=" name"
		str_2 +="?"
		parme = append(parme, u.Name)
	}
	if u.Age != 0 {
		if flag != true {
			flag = true
			str_1 += "set ( "
		}
		str_1 +=",age "
		str_2 +=",?"
		parme = append(parme, u.Age)
	}
	if flag == false {
		return errors.New("no parme need update")
	}
	str_1 += ")"
	str_2 += ")"
	str_1 += str_2
	str_1 = str_1+"where id = ? and flag = ?"
	fmt.Println(str_1)
	_, err := db.Raw(str_1, parme, 1, u.Flag).Exec()
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
