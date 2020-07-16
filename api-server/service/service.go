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
	CreateNewToken(userID uint64, expiresAt time.Time) (string, error)
	FindUserByToken(token string) (*model.User, error)
	LoginUser(email string, password string) (bool, error)
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
