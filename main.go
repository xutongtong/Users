package main

import (
	//"Users/apis"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	//"fmt"
	//"time"
	//"os/user"
	//"time"
	"fmt"
	//"time"
	//"github.com/jinzhu/gorm"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
	"math/rand"
)

type Users struct {
	ID int64
	Name string
	CountryCode string
	Mobile int64
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}
//
//type Result struct {
//	Name string
//	Password string
//}

func main() {
	uuidTime, _, err := uuid.GetTime()
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println(uuid.Must(uuid.NewRandom()))
	//fmt.Println(uint64(os.Getegid()))

	db, err := gorm.Open("mysql", "account:123456789aABC@tcp(drds4ln637rupoos.drds.aliyuncs.com:3306)/account")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	//
	//var id int64
	//db.Exec("select seq.nextval").Scan(id)
	//
	//fmt.Println(id)
	////
	////
	id, _ := uuidTime.UnixTime()
	name := fmt.Sprintf("xutt%d", rand.Intn(10000))
	password  := rand.Intn(9999999)
	mobile := 18680663925 + rand.Int63n(1000000)
	user := Users{ID: id, Name:name,Password:string(password),CountryCode:"86",Mobile:mobile}
	db.Create(&user)

	//if db.NewRecord(user) {
	//	fmt.Println("true")
	//} else {
	//	panic("error")
	//}

	//fmt.Println(user)
	//
	//db.First(user2, 1)
	//
	//fmt.Println(user2)
	//apis.Init()
	//user2 := Users{}
	//db.First(&user2)
	//
	//fmt.Println(user.Name)
	//
	//var result Result
	//db.Raw("SELECT name, password FROM users WHERE name = ?", "xutongtong").Scan(&result)
	//
	//fmt.Println(result)
}