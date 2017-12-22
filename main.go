package main

import (
	//"Users/apis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Users struct {
	gorm.Model
	ID int64 `gorm:"primary_key"`
	Name string
	CountryCode string
	Mobile int64
	Password string
}

func main() {
	db, err := gorm.Open("mysql", "account:123456789aABC@tcp(drds4ln637rupoos.drds.aliyuncs.com:3306)/account")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	//db.AutoMigrate(&Users{})

	user := Users{Name:"xutt",Password:"123456",CountryCode:"86",Mobile:18680663925}
	db.NewRecord(user)

	fmt.Println(user)
	//apis.Init()
}