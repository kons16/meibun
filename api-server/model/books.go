package model

import (
	"github.com/jinzhu/gorm"
)

type Books struct {
	gorm.Model
	Sentence 		string
	Title 			string
	Author 			string
	Pages			int
	Harts			int
	UserID   		uint
}
