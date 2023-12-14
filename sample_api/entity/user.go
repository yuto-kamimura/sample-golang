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
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
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

func (u *User) VerifyPassword(target User) error {
	if err := bcrypt.CompareHashAndPassword([]byte(target.Password), []byte(u.Password)); err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateUser() error {
	dbInstance := db.GetDB()
	if err := dbInstance.Save(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) DeleteUser() error {
	dbInstance := db.GetDB()
	if err := dbInstance.Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *Users) FindAllUsers() error {
	dbInstance := db.GetDB()
	if err := dbInstance.Find(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FindUserByID() error {
	dbInstance := db.GetDB()
	if err := dbInstance.Where("id = ?", u.ID).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FindUserByEmail() error {
	dbInstance := db.GetDB()
	if err := dbInstance.Where("email = ?", u.Email).First(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) UpdatePassword() error {
	dbInstance := db.GetDB()
	if err := dbInstance.Model(&u).Update("password", u.Password).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateEmail() error {
	dbInstance := db.GetDB()
	if err := dbInstance.Model(&u).Update("email", u.Email).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUsername() error {
	dbInstance := db.GetDB()
	if err := dbInstance.Model(&u).Update("username", u.Username).Error; err != nil {
		return err
	}
	return nil
}
