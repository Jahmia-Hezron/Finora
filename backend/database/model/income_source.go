package model

import "gorm.io/gorm"

type IncomeSource struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	UserID      uint   `gorm:"not null;index" json:"user_id"`
	User        User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
}