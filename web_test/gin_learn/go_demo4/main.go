package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	dsn := "root:pharsalia_mariadb@tcp(47.110.240.147:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to mysql failed!")
	}
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		return
	}

	db.Create(&UserInfo{
		ID:     2,
		Name:   "pharsalia",
		Gender: "female",
		Hobby:  "sleep all day",
	})
	//var u UserInfo
	db.Model(&UserInfo{}).
		Where("name = ?", "pharsalia").
		Update("hobby", "play games,sleep")

}
