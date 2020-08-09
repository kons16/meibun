package model

import (
	"github.com/jinzhu/gorm"
)

type UsersHarts struct {
	gorm.Model
	UserID   	uint
	BookID 		uint
}
