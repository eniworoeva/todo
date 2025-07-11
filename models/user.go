package models

import "time"

type User struct {
	ID        string     `gorm:"primaryKey"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	ToDoList  []ToDoList `gorm:"foreignKey:UserID"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt time.Time  `json:"deleted_at"`
}
