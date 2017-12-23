package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	. "Users/config"
)

type DBConfig struct {
	driver   string
	host     string
	port     string
	user     string
	password string
	name     string
	charset  string
}

var db *gorm.DB

func init() {
	// Read config
	dbConfig := DBConfig{driver:DB_DRIVER,host:DB_HOST,port:DB_PORT,user:DB_USER,password:DB_PASSWORD,name:DB_NAME}

	orm, err := Connect(dbConfig)

	if err != nil {
		panic(err.Error())
	}

	db = orm
}

func Connect(config DBConfig) (*gorm.DB, error){
	source := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.user, config.password, config.host, config.port, config.name)

	gorm, err := gorm.Open(config.driver, source)

	if err != nil {
		panic(err.Error())
	}

	return gorm, nil
}

func GetDB() (*gorm.DB){
	return db
}

func CloseDB() {
	db.Close()
}