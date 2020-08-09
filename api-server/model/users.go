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
	Books			[]*Book `gorm:"many2many:user_books;"`
}
