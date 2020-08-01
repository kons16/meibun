package model

import (
	"github.com/jinzhu/gorm"
)

type UserHarts struct {
	gorm.Model
	UserID   	uint
	BookID 		uint
}
