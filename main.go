package main

import (
	//"Users/apis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Users struct {
	gorm.Model
	ID int64 `gorm:"primary_key"`
	Name string
	Password string
}

func main() {
	db, err := gorm.Open("mysql", "account:123456789aABC$@tcp(drds4ln637rupoos.drds.aliyuncs.com:3306)/account?charset=utf8bmp&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.AutoMigrate(&Users{})

	db.Create(Users{Name:"xutt",Password:"123456"})
	//apis.Init()
}