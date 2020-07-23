package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/kons16/meibun/api-server/model"
	"os"
	"time"
)

type Repository interface {
	CreateNewUser(name string, email string, passwordHash string) error
	FindUserByEmail(email string) (*model.User, error)
	CreateNewToken(userID uint, token string, expiresAt time.Time) error
	FindUserByToken(token string) (*model.User, error)
	FindPasswordHashByEmail(email string) (string, error)
	Close() error
}

type repository struct {
	db *gorm.DB
}

func New(dsn string) (Repository, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_IP_ADDRESS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("Open mysql failed: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&model.User{}, &model.UserSession{})

	return &repository{db: db}, nil
}

func (r *repository) Close() error {
	return r.db.Close()
}
