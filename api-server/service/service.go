package service

import (
	"github.com/kons16/meibun/api-server/model"
	"github.com/kons16/meibun/api-server/repository"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type MeibunApp interface {
	CreateNewUser(name string, email string, password string) error
	FindUserByEmail(email string) (*model.User, error)
	CreateNewToken(userID uint, expiresAt time.Time) (string, error)
	FindUserByToken(token string) (*model.User, error)
	LoginUser(email string, password string) (bool, error)
	LogoutUser(token string) error

	CreateNewBook(sentence string, title string, author string, pages int, userId uint) error
	DeleteBookByBookID(bookID uint, userID uint) error
	GetAllBooksByUserID(userID uint) (*[]model.Books, error)
	Close() error
}

type meibunApp struct {
	repo repository.Repository
}

func NewApp(repo repository.Repository) MeibunApp {
	return &meibunApp{repo: repo}
}

func (app *meibunApp) Close() error {
	return app.repo.Close()
}
