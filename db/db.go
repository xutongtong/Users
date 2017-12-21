package db

import (
	"fmt"
	. "Users/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var db *gorm.DB

func Init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	db, err = ConnectDB(dataSourceName)

	if err != nil {
		panic(err.Error())
	}
}

func ConnectDB(dataSourceName string) (*gorm.DB, error){
	db, err := gorm.Open("mysql", dataSourceName)

	defer db.Close()

	if err != nil {
		panic("failed to connect database")
	}

	return db, nil
}

func GetDB() *gorm.DB {
	return db
}