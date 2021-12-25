package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID                int32          `json:"id"`
	HashedPassword    sql.NullString `json:"hashed_password"`
	FullName          sql.NullString `json:"full_name"`
	Email             string         `json:"email"`
	PasswordChangedAt sql.NullTime   `json:"password_changed_at"`
	CreatedAt         time.Time      `json:"created_at"`
}
