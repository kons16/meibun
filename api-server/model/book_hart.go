package model

import (
	"github.com/jinzhu/gorm"
)

type BookHart struct {
	gorm.Model
	BookID   		uint
	Hart			int
	UsersHarts		[]User `gorm:"many2many:users_book_harts"`
}
