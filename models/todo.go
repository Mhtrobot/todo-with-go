package models

import (
	_ "gorm.io/gorm"
	"time"
)

type User struct {
	ID		uint `gorm:"primaryKey"`
	Username	string `gorm:"unique"`
	HashedPassword	string
	CreatedAt	time.Time
	UpdatedAt	time.Time 
}

type Todo struct {
	ID		uint `gorm:"primaryKey"`
	Title		string
	Description	string
	Completed	bool `gorm:"default:false"`
	User		User `gorm:"foreignKey:UserID"`
	UserID		uint
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
