package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserInfo 在数据库中创建表 user_infos
type UserInfo struct {
	ID     int
	Name   string
	Age    int
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:MySql123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// 创建表 自动迁移（将结构体与数据表进行关联）。
	// 表中添加新字段时，自动导入数据库。
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		return
	}

	////插入数据
	//user := &UserInfo{
	//	ID:     1,
	//	Name:   "John Doe",
	//	Age:    20,
	//	Gender: "male",
	//	Hobby:  "coding",
	//}
	//db.Create(user)

	////批量插入数据
	//u1 := UserInfo{Name: "Jack", Age: 21, Gender: "male", Hobby: "sing"}
	//u2 := UserInfo{Name: "Mack", Age: 32, Gender: "male", Hobby: "dance"}
	//u3 := UserInfo{Name: "Tome", Age: 18, Gender: "female", Hobby: "code"}
	//u4 := UserInfo{Name: "Sari", Age: 26, Gender: "female", Hobby: "game"}
	//userInfos := []UserInfo{u1, u2, u3, u4}
	//db.Create(&userInfos)

	//查找、更新数据
	var u UserInfo
	db.First(&u, 1)
	db.Model(&u).Update("age", 22)
	fmt.Printf("u:%#v\n", u)
	//u:main.UserInfo{ID:1, Name:"John Doe", Age:22, Gender:"male", Hobby:"coding"}

	////删除数据
	//db.Delete(&u)
}
