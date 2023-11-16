package entity

import (
	"sampleApi/db"
	"time"
)

type (
	Item struct {
		ID          int       `gorm:"primarykey;autoIncrement"`
		Name        string    `gorm:"not null"`
		Description string    `gorm:"not null"`
		CreatedAt   time.Time `gorm:"not null"`
	}
	Items []Item
)

// 商品を登録する
func (u *Items) CreateItems() error {
	dbIsntance := db.GetDB()
	if err := dbIsntance.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

// 商品IDで商品を取得する
func (u *Items) SelectItemsByItemID(id []int) error {
	dbInstance := db.GetDB()
	if err := dbInstance.Where("id IN ?", id).Find(&u).Error; err != nil {
		return err
	}
	return nil
}
