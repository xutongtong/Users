package model

import (
	"time"
	"github.com/xutongtong/Users/db"
)

type User struct {
	ID        uint       `gorm:"primary_key" json:"id,omitempty"`
	UserName  string     `json:"username" sql:"not null"`
	Mobile    string     `json:"mobile" sql:"not null"`
	Password  string     `json:"password" sql:"not null"`
	Salt      string     `json:"salt" sql:"not null"`
	Source    string     `json:"source" sql:"not null"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

func UserModelInit() {
	db.GetDB().AutoMigrate(&User{})
}

func UserAdd() (user *User) {
	user = &User{
		UserName: "testuser1",
		Mobile: "1351111222",
	}
	db.GetDB().Create(&user)
	return
}

func UserList() (users []*User) {
	db.GetDB().Table("users").Select("user_name, mobile").Scan(&users)
	return
}