package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kons16/meibun/api-server/model"
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
		fmt.Println(dbc.Error)
		return dbc.Error
	}

	bookHart := &model.BookHart{
		BookID: book.ID,
		Hart: 0,
	}
	if dbc := r.db.Create(&bookHart); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return dbc.Error
	}
	return nil
}

// 指定されたbooksレコードをcookieから得たuserIDをもとに削除
func (r *repository) DeleteBookByBookID(bookID uint, userID uint) error {
	if dbc := r.db.Where("id = ? AND user_id = ?", bookID, userID).Delete(&model.Book{}); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return dbc.Error
	}
	return nil
}

// userIDに紐づくbookレコードを全て取得
func (r *repository) GetAllBooksByUserID(userID uint) (*[]model.FrontBook, error) {
	var books []model.Book
	var user model.User

	r.db.First(&user, userID)
	if dbc := r.db.Model(&user).Related(&books); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}

	frontBook := make([]model.FrontBook, len(books))

	for i, _ := range books {
		frontBook[i].ID = books[i].ID
		frontBook[i].UserID = books[i].UserID
		frontBook[i].CreatedAt = books[i].CreatedAt
		frontBook[i].UpdatedAt = books[i].UpdatedAt
		frontBook[i].Sentence = books[i].Sentence
		frontBook[i].Title = books[i].Title
		frontBook[i].Author = books[i].Author
		frontBook[i].Pages = books[i].Pages
		var bookHart model.BookHart
		r.db.Model(books[i]).Related(&bookHart)
		frontBook[i].Harts = bookHart.Hart
	}

	return &frontBook, nil
}

// booksにハートを押したときusers_hartsテーブルにレコードを追加し、books該当レコードのハート数を1上げる
func (r * repository) MakeHart(bookID uint, userID uint) (int, error) {
	var book model.Book
	var userBookHart model.UsersHarts
	var bookHart model.BookHart

	// bookIDから該当するbook_hart_idを取得し、users_hartsからbook_hartに紐づくユーザーを取得する
	r.db.First(&book, bookID)
	if dbc := r.db.Model(&book).Related(&bookHart); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return 0, dbc.Error
	}
	r.db.Raw("SELECT * FROM users_book_harts WHERE user_id = ? AND book_hart_id = ?", userID, bookHart.ID).Scan(&userBookHart)

	// book_hartのhartを1増やす。ただしusers_hartに挿入するレコードが入っていないときのみ(初回のみ)
	if userBookHart.UserID != userID {
		if dbc := r.db.Exec("UPDATE book_harts SET hart = ? WHERE book_id = ?", bookHart.Hart+1, bookID); dbc.Error != nil {
			fmt.Println(dbc.Error)
			return 0, dbc.Error
		}
	}

	// users_book_hartsに重複がなければINSERTする
	sql := "INSERT INTO users_book_harts(user_id, book_hart_id) SELECT ?, ? WHERE NOT EXISTS (SELECT user_id FROM users_book_harts WHERE user_id = ? AND book_hart_id = ?)"
	if dbc := r.db.Exec(sql, userID, bookHart.ID, userID, bookHart.ID); dbc.Error != nil {
		return 0, dbc.Error
	}

	return bookHart.Hart+1, nil
}

// bookの編集情報を受取り、更新する
func (r *repository) UpdateBook(updateData *model.Book) error {
	var book model.Book

	r.db.First(&book, updateData.ID)
	dbc := r.db.Model(&book).Updates(model.Book{
			Author: updateData.Author, Sentence: updateData.Sentence, Title: updateData.Title,
			Pages: updateData.Pages, UserID: updateData.UserID})
	if dbc.Error != nil {
		fmt.Println(dbc.Error)
		return dbc.Error
	}
	return nil
}

// users_book_hartsから該当レコードを削除し、book_hartの該当hartを-1する
func (r *repository) RemoveMyHart(bookID uint, userID uint) (*[]model.FrontBook, error) {
	var book model.Book
	var bookHart model.BookHart

	// bookIDから該当するbook_hart_idを取得し、users_book_hartsからuser_idとbook_hart_idレコードを削除する
	r.db.First(&book, bookID)
	if dbc := r.db.Model(&book).Related(&bookHart); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}
	if dbc := r.db.Exec("DELETE FROM users_book_harts WHERE user_id = ? AND book_hart_id = ?", userID, bookHart.ID); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}

	// book_hartの該当レコードのhartを-1する
	if dbc := r.db.Exec("UPDATE book_harts SET hart = ? WHERE book_id = ?", bookHart.Hart-1, bookID); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}

	// remove後のユーザーに紐づくハートbookを全件返す
	return r.GetMyHart(userID)
}

// userがハートしたbook全件を取得
func (r *repository) GetMyHart(userID uint) (*[]model.FrontBook, error) {
	var user model.User

	r.db.First(&user, userID).Related(&user.UsersHarts, "UsersHarts")
	frontBooks := make([]model.FrontBook, len(user.UsersHarts))
	for i, v := range user.UsersHarts {
		var book model.Book
		r.db.First(&book, v.BookID)
		frontBooks[i].ID = book.ID
		frontBooks[i].Sentence = book.Sentence
		frontBooks[i].Title = book.Title
		frontBooks[i].Author = book.Author
		frontBooks[i].Pages = book.Pages
		frontBooks[i].Harts = v.Hart
	}

	return &frontBooks, nil
}
