package entity

import "time"

type (
	Item struct {
		ID       int       `gorm:"primarykey;autoIncrement"`
		Name     string    `gorm:"not null"`
		CreateAt time.Time `gorm:"not null"`
	}
)
