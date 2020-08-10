package service

import "github.com/kons16/meibun/api-server/model"

// 名文を保存するのに必要な情報をrepoのCreateNewBookに渡し,errorを返す
func (app *meibunApp) CreateNewBook(sentence string, title string, author string, pages int, userId uint) error {
	return app.repo.CreateNewBook(sentence, title, author, pages, userId)
}

// userIDに紐づくbooksを取得
func (app *meibunApp) GetAllBooksByUserID(userID uint) (*[]model.FrontBook, error) {
	return app.repo.GetAllBooksByUserID(userID)
}

// bookIDをもとに該当bookレコードを削除する
func (app *meibunApp) DeleteBookByBookID(bookID uint, userID uint) error {
	return app.repo.DeleteBookByBookID(bookID, userID)
}

// userがハートしたbookを保存させ、保存後のbookのハート数を取得
func (app *meibunApp) MakeHart(bookID uint, userID uint) (int, error) {
	return app.repo.MakeHart(bookID, userID)
}

// userがハートしたbookを全件取得
func (app *meibunApp) GetMyHart(userID uint) (*[]model.FrontBook, error) {
	return app.repo.GetMyHart(userID)
}
