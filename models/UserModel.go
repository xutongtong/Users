package models

import (
	"time"
)

type User struct {
	ID         int64
	Name       string
	Mobile     string
	Password   string
	Salt       string
	Source     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}