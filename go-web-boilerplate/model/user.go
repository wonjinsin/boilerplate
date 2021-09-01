package model

import (
	"github.com/google/uuid"
)

// User ...
type User struct {
	UID   string `json:"uid" gorm:"primaryKey"`
	Email string `json:"email"`
	Nick  string `json:"nick"`
}

// ValidateNewUser ...
func (u *User) ValidateNewUser() bool {
	if _, err := uuid.Parse(u.UID); err != nil {
		return false
	}

	return true
}

// ValidateUpdateUser ...
func (u *User) ValidateUpdateUser() bool {
	return true
}
