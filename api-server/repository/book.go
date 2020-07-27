package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kons16/meibun/api-server/model"
)

// booksテーブルに新しく名文を追加
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

// userIDに紐づくbookレコードを全て取得
func(r *repository) GetAllBooksByUserID(userID uint) (*model.Books, error) {
	var books model.Books
	if dbc := r.db.Model(r.db.First(&model.User{}, userID)).Related(&books); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &books, nil
}
