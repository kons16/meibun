package model

import (
	"github.com/jinzhu/gorm"
)

type BookHarts struct {
	gorm.Model
	BookID   		uint
	Harts			int
}
