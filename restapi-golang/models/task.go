package models

import "time"

type Task struct {
	ID        uint   `json:"ID" gorm:"primary_key"`
	Title     string `json:"Title"`
	Completed string `json:"Completed"`
	CreatedAt time.Time
	UserID    uint
}
