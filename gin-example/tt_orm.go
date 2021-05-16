package main
//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
//type User struct {
//	gorm.Model
//	Name string
//	Age int64
//	AbC int64
//}
//type Som struct {
//	Name string
//	Age int64
//}
//
//func main() {
//
//	//1
//	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
//	//_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	//if err !=nil {
//	//	fmt.Println(err)
//	//}
//
//	//2
//	db, err := gorm.Open(mysql.New(mysql.Config{
//				DSN: dsn,
//				DefaultStringSize: 256,
//			}),&gorm.Config{})
//	if err !=nil {
//		fmt.Println(err)
//	}
//
//	db.AutoMigrate(&User{})
//
//	//user :=User{Name: "xxxx",Age: 123}
//	//result :=db.Create(&user)
//	//
//	//fmt.Println(user.ID)
//	//fmt.Println(result.RowsAffected)
//
//
//	//map创建  可以是多个
//	//res :=db.Model(&User{}).Create([]map[string]interface{}{
//	//			{"Name":"wyf","Age":23},
//	//			{"Name":"wyf2","Age":213},
//	//		})
//	//fmt.Println(res.RowsAffected)
//
//	u :=User{}
//	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
//	//first := db.First(&u)
//	// SELECT * FROM users WHERE id = 10;
//	//db.First(&u,10)
//	// SELECT * FROM users WHERE id IN (1,2,3);
//	//db.First(&u,[]int{10,11,12})
//	//db.Take(&u)
//	db.Last(&u)
//	//fmt.Println(first.Row())
//	fmt.Println(u)
//
//
	/*
	select
	 */
//
//	result := map[string]interface{}{}
//	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
//	db.Model(&User{}).First(&result)
//	fmt.Println(result)
//
//	result1 := map[string]interface{}{}
//	result2 := map[string]interface{}{}
//	/*
//	不行
//	 SELECT * FROM `users` ORDER BY `users`. LIMIT 1
//	*/
////	db.Table("users").First(&result1)
//	db.Table("users").Take(&result2)
//	fmt.Println(result1)
//	fmt.Println(result2)
//
//
//	// SELECT * FROM users;
//	users :=[]*User{}
//	db.Find(&users)
//	fmt.Println(users)
//
////条件
//	fmt.Println("-------条件-----")
//	cu :=User{}
//	//sql
//	db.Where("id =?",2).First(&cu)
//	fmt.Println(cu)
//	//struct
//	cu2 :=User{}
//	//where 查找不到时会err，最终会放返回frist的结果
//	db.Where(map[string]interface{}{"id":"2","name":"xxxx"}).First(&cu2)
//	fmt.Println(cu2)
//
//	users2 :=[]*User{}
//	// SELECT * FROM users WHERE name = "xxxx" AND age = 0;
//	db.Where(&User{Name: "xxxx"},"name","age").Find(&users2)
//
//
//
//
//	//not
//	users3 :=[]*User{}
//	user3 :=User{}
//	// SELECT * FROM users WHERE NOT name = "xxxx" ORDER BY id LIMIT 1;
//	//db.Not("name = ?","xxxx").Find(users3)
//	db.Not("name = ?","xxxx").First(&user3)
//	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");
//	db.Not(map[string]interface{}{"name":[]string{"xxxx","xagg"}}).Find(&users3)
//
////or
//	users4 := []*User{}
//	// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';
//	//可以传入struct  ，mao
//	db.Where("name = ?","abc").Or("name = ?","xxxx").Find(&users4)
//
//}