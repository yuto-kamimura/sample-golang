package entity

import "time"

type (
	Item struct {
		ID       int
		Name     string
		CreateAt time.Time
	}
)
