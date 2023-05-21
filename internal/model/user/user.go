package user

import (
	"time"
)

type User struct {
	Username        string
	Email           string
	Password        string
	PasswordConfirm string
	Role            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
