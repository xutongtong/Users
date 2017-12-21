package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/account")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	r.GET("/hello", func(context *gin.Context) {
		_, err := db.Query("SELECT * FROM users")
		if err != nil {
			panic(err.Error())
		}
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":9000")
	//apis.Init()
}