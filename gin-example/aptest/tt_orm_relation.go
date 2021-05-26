package main
//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//
////belong  to
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
//
//
//// has money
//
//
//
//
//func main() {
//	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.New(mysql.Config{
//				DSN: dsn,
//				DefaultStringSize: 256,
//			}),&gorm.Config{})
//	if err !=nil {
//		fmt.Println(err)
//	}
//
//	//db.AutoMigrate(&User{},&Company{})
//
//
//	us :=[]*User{}
//	//事先预加载 全部
//	//db.Preload(clause.Associations).Limit(5).Find(&us)
//	db.Preload("companies").Limit(5).Find(&us)
//	for _,v := range us {
//		fmt.Println(v.Name," : ",v.Company.Name)
//	}
//
//}