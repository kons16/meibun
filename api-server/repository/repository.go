package repository

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kons16/meibun/api-server/model"
	"os"
	"time"
)

type Repository interface {
	CreateNewUser(name string, email string, passwordHash string) error
	FindUserByEmail(email string) (*model.User, error)
	CreateNewToken(userID uint64, token string, expiresAt time.Time) error
	FindUserByToken(token string) (*model.User, error)
	FindPasswordHashByEmail(email string) (string, error)
	Close() error
}

type repository struct {
	dbMap *gorp.DbMap
}

func New(dsn string) (Repository, error) {
	if err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV"))); err != nil {
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
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("Open mysql failed: %v", err)
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

	dbmap.AddTableWithName(model.User{}, "users")
	dbmap.AddTableWithName(model.UserSession{}, "user_session")

	return &repository{dbMap: dbmap}, nil
}

func (r *repository) generateID() (uint64, error) {
	var id uint64
	_, err := r.dbMap.Select(&id, "SELECT UUID_SHORT()")
	return id, err
}

func (r *repository) Close() error {
	return r.dbMap.Db.Close()
}
