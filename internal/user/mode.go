package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"uniqueIndex;not null"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func NewUser(email string, name string, password string) *User {
	return &User{Email: email, Name: name, Password: password}
}
