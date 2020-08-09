package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kons16/meibun/api-server/model"
	"time"
)

// booksテーブルに新しくレコードを追加。book_hartはhartを0でレコード追加。
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

	bookHart := &model.BookHart{
		BookID: book.ID,
		Hart: 0,
	}
	if dbc := r.db.Create(&bookHart); dbc.Error != nil {
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
	var book model.Book
	var userHart model.UsersHarts
	var bookHart model.BookHart

	// bookIDから該当するbook_hart_idを取得し、users_hartsからbook_hartに紐づくユーザーを取得する
	r.db.First(&book, bookID)
	if dbc := r.db.Model(&book).Related(&bookHart); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return 0, dbc.Error
	}
	r.db.Raw("SELECT * FROM users_harts WHERE user_id = ? AND book_hart_id = ?", userID, bookHart.ID).Scan(&userHart)

	// book_hartのhartを1増やす。ただしusers_hartに挿入するレコードが入っていないときのみ(初回のみ)
	if userHart.UserID != userID {
		if dbc := r.db.Exec("UPDATE book_harts SET hart = ? WHERE book_id = ?", bookHart.Hart+1, bookID); dbc.Error != nil {
			return 0, dbc.Error
		}
	}

	// users_hartsに重複がなければINSERTする
	now := time.Now()
	sql := "INSERT INTO users_harts(user_id, book_hart_id, created_at, updated_at) SELECT ?, ?, ?, ? WHERE NOT EXISTS (SELECT user_id FROM users_harts WHERE user_id = ? AND book_hart_id = ?)"
	if dbc := r.db.Exec(sql, userID, bookHart.ID, now, now, userID, bookHart.ID); dbc.Error != nil {
		return 0, dbc.Error
	}

	return bookHart.Hart+1, nil
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
