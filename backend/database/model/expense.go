package model

import (
	"time"

	"gorm.io/gorm"
)

type Expense struct {
	gorm.Model
	Amount      float64         `json:"amount" gorm:"not null"`
	Description string          `json:"description" gorm:"not null"`
	Date time.Time 				`json:"date" gorm:"not null"`

	CategoryID  uint            `json:"category_id" gorm:"not null;index"`
	Category ExpenseCategory 	`gorm:"foreignKey:CategoryID;references:ID"`

	UserID      uint            `json:"user_id" gorm:"not null;index"`
	User        User            `gorm:"foreignKey:UserID;references:ID"`
}