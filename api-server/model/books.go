package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Sentence 		string
	Title 			string
	Author 			string
	Pages			int
	Harts			int
	UserID   		uint
	Users			[]*User `gorm:"many2many:user_books;"`
}
