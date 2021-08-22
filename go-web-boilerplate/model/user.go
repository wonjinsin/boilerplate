package model

// User ...
type User struct {
	UID   string `json:"uid" gorm:"primaryKey"`
	Email string `json:"email"`
	Nick  string `json:"nick"`
}
