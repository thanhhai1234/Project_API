package models

type User struct {
	ID       uint   `json:"ID" gorm:"primary_key"`
	Name     string `json:"Name"`
	Password string `json:"Password" binding:"required"`
	Tasks    []Task `gorm:"foreign_key:UserID"`
}
