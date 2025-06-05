package model

import (
	"time"

	"gorm.io/gorm"
)

type Income struct {
	gorm.Model
	Amount      float64      `json:"amount" gorm:"not null"`
	Description string       `json:"description" gorm:"not null"`
	Date        time.Time    `json:"date" gorm:"not null"`
	SourceID    uint         `json:"source_id" gorm:"not null;index"`
	Source      IncomeSource `gorm:"foreignKey:SourceID;references: ID"`

	UserID      uint         `json:"user_id" gorm:"not null;index"`
	User        User         `gorm:"foreignKey:UserID; references: ID"`
}