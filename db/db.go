package db

import (
	"fmt"
	. "Users/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var db *gorm.DB

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	gorm, err := ConnectDB(dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	db = gorm
}

func ConnectDB(dataSourceName string) (*gorm.DB, error){
	gorm, err := gorm.Open("mysql", dataSourceName)

	defer gorm.Close()

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("connect db successed")

	return gorm, nil
}

func GetDB() *gorm.DB {
	return db
}