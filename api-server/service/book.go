package service

// 名文を保存するのに必要な情報をrepoのCreateNewBookに渡す
func(app *meibunApp) CreateNewBook(sentence string, title string, author string, pages int, userId uint) error {
	return app.repo.CreateNewBook(sentence, title, author, pages, userId)
}
