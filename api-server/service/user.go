package service

import (
	"errors"
	"github.com/kons16/meibun/api-server/model"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

// name, email, passwordをもとにpasswordはハッシュ化しusersテーブルに登録する
func (app *meibunApp) CreateNewUser(name string, email string, password string) error {
	if name == "" {
		return errors.New("empty user name")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return app.repo.CreateNewUser(name, email, string(passwordHash))
}

// token文字列を作成する
func generateToken() string {
	table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_@"
	l := len(table)
	ret := make([]byte, 128)
	src := make([]byte, 128)
	rand.Read(src)
	for i := 0; i < 128; i++ {
		ret[i] = table[int(src[i])%l]
	}
	return string(ret)
}

// tokenを生成し、user_sessionに挿入する
func (app *meibunApp) CreateNewToken(userID uint64, expiresAt time.Time) (string, error) {
	token := generateToken()
	if err := app.repo.CreateNewToken(userID, token, expiresAt); err != nil {
		return "", err
	}
	return token, nil
}

// emailからユーザー情報を取得する
func (app *meibunApp) FindUserByEmail(email string) (*model.User, error) {
	return app.repo.FindUserByEmail(email)
}

// tokenからからユーザー情報を取得する
func (app *meibunApp) FindUserByToken(token string) (*model.User, error) {
	return app.repo.FindUserByToken(token)
}

// passwordがusersテーブルのハッシュ値と正しければtrueを返す
func (app *meibunApp) LoginUser(email string, password string) (bool, error) {
	passwordHash, err := app.repo.FindPasswordHashByEmail(email)
	if err != nil {
		return false, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
