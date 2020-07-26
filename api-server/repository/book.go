package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kons16/meibun/api-server/model"
)

// tokenからuser情報を取得し、booksテーブルに新しく名文を追加
func(r *repository) CreateNewBook(sentence string, title string, author string, pages int, userId uint) error {
	book := &model.Books{
		Sentence: sentence,
		Title: title,
		Author: author,
		Pages: pages,
		UserID: userId,
	}
	if dbc := r.db.Create(&book); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}
