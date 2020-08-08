package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kons16/meibun/api-server/model"
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
	userHart := &model.UserHarts{
		UserID: userID,
		BookID: bookID,
	}
	if dbc := r.db.Create(&userHart); dbc.Error != nil {
		return 0, dbc.Error
	}

	// 現状のhart数を取得し、+1でUPDATEする
	var nowHart int
	r.db.Raw("SELECT harts FROM books WHERE id = ?", bookID).Scan(&nowHart)
	if dbc := r.db.Exec("UPDATE books SET harts = ? WHERE id = ?", nowHart+1, bookID); dbc.Error != nil {
		return 0, dbc.Error
	}
	
	return nowHart+1, nil
}
