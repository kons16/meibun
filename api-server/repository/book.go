package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kons16/meibun/api-server/model"
	"time"
)

// booksテーブルに新しく名文を追加
func (r *repository) CreateNewBook(sentence string, title string, author string, pages int, userId uint) error {
	book := &model.Book{
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
	if dbc := r.db.Where("id = ? AND user_id = ?", bookID, userID).Delete(&model.Book{}); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

// userIDに紐づくbookレコードを全て取得
func (r *repository) GetAllBooksByUserID(userID uint) (*[]model.Book, error) {
	var books []model.Book
	var user model.User
	r.db.First(&user, userID)
	if dbc := r.db.Model(&user).Related(&books); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &books, nil
}

// booksにハートを押したときusers_hartsテーブルにレコードを追加し、books該当レコードのハート数を1上げる
func (r * repository) MakeHart(bookID uint, userID uint) (int, error) {
	// booksのhartを1増やす。ただしusers_hartに挿入するレコードが入っていないときのみ(初回のみ)
	var book model.Book
	var userHart model.UsersHarts
	r.db.Raw("SELECT * FROM users_harts WHERE user_id = ? AND book_id = ?", userID, bookID).Scan(&userHart)

	if userHart.UserID != userID {
		r.db.Raw("SELECT harts FROM books WHERE id = ?", bookID).Scan(&book)
		if dbc := r.db.Exec("UPDATE books SET harts = ? WHERE id = ?", book.Harts+1, bookID); dbc.Error != nil {
			return 0, dbc.Error
		}
	}

	// users_hartsに重複がなければINSERTする
	now := time.Now()
	sql := "INSERT INTO users_harts(user_id, book_id, created_at, updated_at) SELECT ?, ?, ?, ? WHERE NOT EXISTS (SELECT user_id FROM users_harts WHERE user_id = ? AND book_id = ?)"
	if dbc := r.db.Exec(sql, userID, bookID, now, now, userID, bookID); dbc.Error != nil {
		return 0, dbc.Error
	}

	return book.Harts+1, nil
}

// userがハートしたbook全件を取得
func (r *repository) GetMyHart(userID uint) (*[]model.Book, error) {
	var books []model.Book
	var user model.User

	r.db.First(&user, userID)
	if dbc := r.db.Model(&books).Related(&user, "Users"); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}

	fmt.Println(books)
	return &books, nil
}
