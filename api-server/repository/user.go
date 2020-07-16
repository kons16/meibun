package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kons16/meibun/api-server/model"
	"time"
)

// name,email,password_hashをもとにusersテーブルに挿入
func (r *repository) CreateNewUser(name string, email string, passwordHash string) error {
	id, err := r.generateID()
	if err != nil {
		return err
	}
	now := time.Now()
	user := &model.User{
		ID: id,
		Name: name,
		Email: email,
		PasswordHash: passwordHash,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := r.dbMap.Insert(user); err != nil {
		return err
	}
	return nil
}

// emailをもとにuserを検索し、idとnameをmodel.Userにマッピング
func (r *repository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	query := "SELECT id,name FROM user WHERE email = ?"
	if err := r.dbMap.SelectOne(&user, query, email); err != nil {
		return nil, err
	}
	return &user, nil
}

// emailをもとにuserを検索し、userのハッシュ化パスワードを取得
func (r *repository) FindPasswordHashByEmail(email string) (string, error) {
	var hash string
	query := "SELECT password_hash FROM user WHERE email = ?"
	if err := r.dbMap.SelectOne(hash, query, email); err != nil {
		return "", err
	}
	return hash, nil
}

// userID,token,expiresAtを, user_sessionテーブルに挿入
func (r *repository) CreateNewToken(userID uint64, token string, expiresAt time.Time) error {
	now := time.Now()
	userSession := &model.UserSession{
		UserId: userID,
		Token: token,
		ExpiresAt: expiresAt,
		CreatedAt: now,
		UpdatedAt: now,
	}
	if err := r.dbMap.Insert(userSession); err != nil {
		return err
	}
	return nil
}

// tokenをもとにuserのid,nameを取得し、model.Userにマッピング
func (r *repository) FindUserByToken(token string) (*model.User, error) {
	var user model.User
	query := "SELECT id,name FROM user JOIN user_session ON user.id = user_session.user_id WHERE user_session.token = ? && user_session.expires_at > ?"
	if err := r.dbMap.SelectOne(&user, query, token, time.Now()); err != nil {
		return nil, err
	}
	return &user, nil
}
