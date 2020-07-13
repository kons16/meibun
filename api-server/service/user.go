package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

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

