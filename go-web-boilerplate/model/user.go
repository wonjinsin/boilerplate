package model

// User ...
type User struct {
	UID   string `json:"uid" gorm:"primaryKey"`
	Email string `json:"email"`
	Nick  string `json:"nick"`
}

// ValidateNewUser ...
func (u *User) ValidateNewUser() bool {
	return u.Email != "" && u.Nick != ""
}

// ValidateUpdateUser ...
func (u *User) ValidateUpdateUser() bool {
	return true
}

// UpdateUser ...
func (u *User) UpdateUser(user *User) *User {
	u = user
	return u
}
