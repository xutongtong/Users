package main

import (
	"Users/db"
	"Users/apis"
)

func main() {
	apis.Init()
	defer db.CloseDB()
}