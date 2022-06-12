package models

import "time"

type User struct {
	Id        uint
	Name      string
	Email     string `gorm:"unique"`
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}
