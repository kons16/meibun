package repository

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kons16/meibun/api-server/model"
	"time"
)

// name,email,password_hashをもとにusersテーブルに挿入
func (r *repository) CreateNewUser(name string, email string, passwordHash string) error {
	user := &model.User{
		Name: name,
		Email: email,
		PasswordHash: passwordHash,
	}
	// 同じemailで登録したユーザーがいないかどうかチェック
	findedUser, _ := r.FindUserByEmail(email)
	if findedUser != nil {
		return nil
	}

	if dbc := r.db.Create(&user); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

// emailをもとにuserを検索し、idとnameをmodel.Userにマッピング
func (r *repository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	// query := "SELECT * FROM user WHERE email = ? LIMIT 1"
	if dbc := r.db.Where("email = ?", email).First(&user); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &user, nil
}

// emailをもとにuserを検索し、userのハッシュ化パスワードを取得
func (r *repository) FindPasswordHashByEmail(email string) (string, error) {
	var user model.User
	// query := "SELECT password_hash FROM user WHERE email = ?"
	if dbc := r.db.Select("password_hash").Where("email = ?", email).First(&user); dbc.Error != nil {
		return "", dbc.Error
	}
	return user.PasswordHash, nil
}

// userID,token,expiresAtを, user_sessionテーブルに挿入
func (r *repository) CreateNewToken(userID uint, token string, expiresAt time.Time) error {
	userSession := &model.UserSession{
		UserID: userID,
		Token: token,
		ExpiresAt: expiresAt,
	}
	if dbc := r.db.Create(&userSession); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}

// tokenをもとにuserのid,nameを取得し、model.Userにマッピング
func (r *repository) FindUserByToken(token string) (*model.User, error) {
	var user model.User
	dbc := r.db.Raw("SELECT users.id, users.name FROM users JOIN user_sessions ON users.id = user_sessions.user_id WHERE user_sessions.token = ? && user_sessions.expires_at > ?",
		token, time.Now()).Scan(&user)
	if dbc.Error != nil {
		return nil, dbc.Error
	}
	return &user, nil
}

// tokenをもとにuser_sessionの該当行を削除
func (r *repository) DeleteUserSessionByToken(token string) error {
	if dbc := r.db.Where("Token = ?", token).Delete(&model.UserSession{}); dbc.Error != nil {
		return dbc.Error
	}
	return nil
}
