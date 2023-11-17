package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type (
	User struct {
		ID        int
		Username  string
		Password  string
		Email     string
		UpdatedAt time.Time
	}

	Users []Users
)

func (u *User) EncriptPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

func (u *User) CreateUser() error {

	return nil
}
