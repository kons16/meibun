package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserSession struct {
	gorm.Model
	UserID   	uint
	Token 		string
	ExpiresAt 	time.Time
}
