package model

import "time"

// bookをフロントへ返すときはhartも同時に返したため、bookとは別に構造体FrontBookを定義する
type FrontBook struct {
	ID				uint
	UserID			uint
	CreatedAt		time.Time
	UpdatedAt		time.Time
	Sentence 		string
	Title 			string
	Author 			string
	Pages			int
	Harts  			int
}
