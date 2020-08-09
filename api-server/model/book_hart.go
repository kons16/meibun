package model

import (
	"github.com/jinzhu/gorm"
)

type BookHart struct {
	gorm.Model
	BookID   		uint
	Hart			int
}
