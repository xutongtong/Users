package db

import (
	"fmt"
	"../config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() {
	c := config.GetConfig()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		c.GetString("mysql.user"),
		c.GetString("mysql.password"),
		c.GetString("mysql.host"),
		c.GetString("mysql.port"),
		c.GetString("mysql.dbname"))

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