package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name			string
	Email			string
	PasswordHash	string
	UserSession		UserSession
	Books			[]*Book
	UsersHarts		[]BookHart `gorm:"many2many:users_book_harts"`
}
