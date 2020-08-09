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
	UserID   		uint
	BookHart		BookHart
}
