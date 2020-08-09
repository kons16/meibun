package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kons16/meibun/api-server/model"
	"time"
)

// booksテーブルに新しく名文を追加
func (r *repository) CreateNewBook(sentence string, title string, author string, pages int, userId uint) error {
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

// 指定されたbooksレコードをcookieから得たuserIDをもとに削除
func (r *repository) DeleteBookByBookID(bookID uint, userID uint) error {
	if dbc := r.db.Where("id = ? AND user_id = ?", bookID, userID).Delete(&model.Books{}); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

// userIDに紐づくbookレコードを全て取得
func (r *repository) GetAllBooksByUserID(userID uint) (*[]model.Books, error) {
	var books []model.Books
	var user model.User
	r.db.First(&user, userID)
	if dbc := r.db.Model(&user).Related(&books); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &books, nil
}

// booksにハートを押したときuser_hartsテーブルにレコードを追加し、books該当レコードのハート数を1上げる
func (r * repository) MakeHart(bookID uint, userID uint) (int, error) {
	// user_hartsに重複がなければINSERTする
	now := time.Now()
	sql := "INSERT INTO user_harts(user_id, book_id, created_at, updated_at) SELECT ?, ?, ?, ? WHERE NOT EXISTS (SELECT user_id FROM user_harts WHERE user_id = ? AND book_id = ?)"
	if dbc := r.db.Exec(sql, userID, bookID, now, now, userID, bookID); dbc.Error != nil {
		return 0, dbc.Error
	}

	// booksのhartを1増やす
	var book model.Books
	r.db.Raw("SELECT harts FROM books WHERE id = ?", bookID).Scan(&book)
	if dbc := r.db.Exec("UPDATE books SET harts = ? WHERE id = ?", book.Harts+1, bookID); dbc.Error != nil {
		return 0, dbc.Error
	}

	return book.Harts+1, nil
}
