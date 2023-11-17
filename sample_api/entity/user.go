package entity

import "time"

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

func (u *User) EncriptPassword() {

}

func (u *User) CreateUser() error {

	return nil
}
