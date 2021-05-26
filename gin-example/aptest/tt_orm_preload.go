package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

type User struct {
	gorm.Model
	Username string
	Orders []Order
	CompanyID int
	Company Company
}
type Order struct {
	gorm.Model
	UserID int
	Price float64
}
type Company struct {
	ID int
	Name string
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println(err)


	}
	db.Exec("drop table orders")
	db.Exec("drop table users")
	db.AutoMigrate(&User{}, &Order{})
	initm(db)
	us := []*User{}
	db.Preload("Orders").Preload("Company").Find(&us)
	for _,n :=range us{
		fmt.Println(n.Orders)
		fmt.Println(n.Company)
	}


}

func initm(db *gorm.DB){
	db.Create(&Company{Name: "com1"})
	db.Create(&Company{Name: "com2"})
	for i:=1;i<5;i++ {
		user := &User{Username: "abc"+strconv.Itoa(i),CompanyID: i%2+1}
		db.Create(&user)
	}
	for i:=0;i<23;i++{
		db.Create(&Order{
			UserID: i&3+1,
			Price: 123.2+float64(i),
		})
	}
}