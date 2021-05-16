package main
//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"strconv"
//)
//
//type User struct {
//	gorm.Model
//	Name string
//	Age int
//	CompanyID int
//	Company Company
//}
//type Company struct {
//	ID int
//	Name string
//}
//
//func main() {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.New(mysql.Config{
//		DSN: dsn,
//		DefaultStringSize: 256,
//	}),&gorm.Config{})
//	if err !=nil {
//		fmt.Println(err)
//	}
//
//
//	db.Exec("drop table users")
//	db.Exec("drop table companies")
//	db.AutoMigrate(&User{},&Company{})
//	for i:=1;i<10;{
//		db.Create(&Company{Name: "abc"+strconv.Itoa(i)})
//		i++
//	}
//	for i:=0 ;i<20;i++{
//		db.Create(&User{Name: "xxx"+strconv.Itoa(i),Age: i*2,CompanyID:(i%9)+1})
//	}
//
//}