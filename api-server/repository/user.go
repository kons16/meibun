package repository

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kons16/meibun/api-server/model"
	"time"
)

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

func (r *repository) FindUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Get(
		&user,
		`SELECT id,name FROM user
			WHERE email = ? LIMIT 1`, email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindPasswordHashByEmail(email string) (string, error) {
	var hash string
	err := r.db.Get(
		&hash,
		`SELECT password_hash FROM user
			WHERE email = ? LIMIT 1`, email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return hash, nil
}

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

func (r *repository) FindUserByToken(token string) (*model.User, error) {
	var user model.User
	err := r.db.Get(
		&user,
		`SELECT id,name FROM user JOIN user_session
			ON user.id = user_session.user_id
				WHERE user_session.token = ? && user_session.expires_at > ?
				LIMIT 1`, token, time.Now(),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &user, nil
}
