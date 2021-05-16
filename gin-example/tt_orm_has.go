package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)


type Userho struct {
	gorm.Model
	CreditCard CreditCard
}

type CreditCard struct {
	gorm.Model
	Number string
	UserhoID int
}


// has money

type Userhm struct {
	gorm.Model
	CreditCardm []CreditCardm
}

type CreditCardm struct {
	gorm.Model
	Number string
	UserhmID int
}

//m 2 m
type Userm2m struct {
	gorm.Model
	//Name string
	Languages []Language `gorm:"many2many:user_languages"`
}

type Language struct {
	gorm.Model
	Name string 	`gorm:"many2many:user_languages"`
}



func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
				DSN: dsn,
				DefaultStringSize: 256,
			}),&gorm.Config{})
	if err !=nil {
		fmt.Println(err)
	}

	// has one
	//db.AutoMigrate(&Userho{},&CreditCard{})
	// has money
	//db.AutoMigrate(&Userhm{},&CreditCardm{})
	//many2many
	db.AutoMigrate(&Userm2m{},&Language{})

	//many2mnyinit(db)
	var userm2m Userm2m
	userm2m.ID=1
	//us := []*Userm2m{}
	ls := []Language{}

	//
	db.Model(&userm2m).Association("Languages").Find(&ls)
	//添加关联的记录
	//db.Model(&userm2m).Association("Languages").Append([]Language{{Name: "yyddddddddddddddddddddd"}})
	//替换
	//db.Model(&userm2m).Association("Languages").Replace([]Language{{Name: "agc"}})
	//假删除 要删除的关联的逐渐不能为空
	e :=db.Model(&userm2m).Association("Languages").Delete(&Language{Model:gorm.Model{ID: 25}})
	fmt.Println(e)

	db.Model(&userm2m).Association("Languages").Find(&ls)
	fmt.Println(ls)
}



func many2mnyinit(db *gorm.DB){
	for i:=1;i<10;i++ {
		db.Save(&Userm2m{
			Languages: []Language{
				{Name: "abv" +strconv.Itoa(i)},
				{Name: "adddd"+strconv.Itoa(i*2)},
			}},
		)
	}
}