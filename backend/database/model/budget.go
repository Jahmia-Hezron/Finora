package model

import "gorm.io/gorm"

type Budget struct {
	gorm.Model
	Name		string                `json:"name" gorm:"not null"`
	Amount      float64               `json:"amount" gorm:"not null"`
	Description string                `json:"description" gorm:"not null"`
	Period     string                `json:"period" gorm:"not null"`
	CategoryID  uint                  `json:"category_id" gorm:"not null;index"`
	Category    ExpenseCategory `gorm:"foreignKey:CategoryID;references:ID"`
	UserID      uint                  `json:"user_id" gorm:"not null;index"`
	User        User            `gorm:"foreignKey:UserID;references:ID"`
}