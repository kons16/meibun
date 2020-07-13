package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Repository interface {
	CreateNewUser(name string, email string, passwordHash string) error
	// FindUserByEmail(email string) (*model.User, error)
	CreateNewToken(userID uint64, token string, expiresAt time.Time) error
	// FindUserByToken(token string) (*model.User, error)
	// FindPasswordHashByEmail(email string) (string, error)
	Close() error
}

type repository struct {
	db *sqlx.DB
}

func New(dsn string) (Repository, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"password",
		"127.0.0.1",
		"3306",
		"db_name",
	)
	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("Open mysql failed: %v", err)
	}
	return &repository{db: db}, nil
}

func (r *repository) generateID() (uint64, error) {
	var id uint64
	err := r.db.Get(&id, "SELECT UUID_SHORT()")
	return id, err
}

func (r *repository) Close() error {
	return r.db.Close()
}
