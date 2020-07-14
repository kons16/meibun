package model

import "time"

type User struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
	Email string `db:"email"`
	PasswordHash string `db:"password_hash"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
