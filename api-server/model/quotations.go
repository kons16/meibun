package model

import (
	"github.com/jinzhu/gorm"
)

type Quotations struct {
	gorm.Model
	Sentence 		string
	PersonName 		string
	Harts			int
	UserID   		uint
}
