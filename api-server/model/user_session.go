package model

import "time"

type UserSession struct {
	UserId   	uint64 `db:"user_id"`
	Token 		string `db:"token"`
	ExpiresAt 	time.Time `db:"expires_at"`
	CreatedAt 	time.Time `db:"created_at"`
	UpdatedAt 	time.Time `db:"updated_at"`
}