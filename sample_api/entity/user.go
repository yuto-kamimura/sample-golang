package entity

import (
	"sampleApi/db"
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
	dbInstance := db.GetDB()
	if err := dbInstance.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func FindUser(username string) (User, error) {
	dbInstance := db.GetDB()
	var target User
	if err := dbInstance.Where("username = ?", username).First(&target).Error; err != nil {
		return User{}, err
	}
	return target, nil
}

func (u *User) CompareLoginPassword(target User) error {
	if err := bcrypt.CompareHashAndPassword([]byte(target.Password), []byte(u.Password)); err != nil {
		return err
	}

	return nil
}
