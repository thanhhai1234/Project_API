package models

import "time"

type Task struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Completed string `json:"completed"`
	CreatedAt time.Time
}
