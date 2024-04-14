package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"unique_index"`
	Password string
	Fullname string
	Abn string `gorm:"unique_index"`
	PhoneNumber string
	Address string

	Companies []Company
}