package service

import "github.com/kons16/meibun/api-server/model"

// 名文を保存するのに必要な情報をrepoのCreateNewBookに渡し,errorを返す
func (app *meibunApp) CreateNewBook(sentence string, title string, author string, pages int, userId uint) error {
	return app.repo.CreateNewBook(sentence, title, author, pages, userId)
}

// userIDに紐づくbooksを返す
func (app *meibunApp) GetAllBooksByUserID(userID uint) (*[]model.Books, error) {
	return app.repo.GetAllBooksByUserID(userID)
}
