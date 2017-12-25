package main

import (
	"github.com/xutongtong/Users/router"
	"github.com/xutongtong/Users/db"
	"github.com/xutongtong/Users/model"
)

func main() {
	db.Init()
	model.Init()

	router.Init()
}