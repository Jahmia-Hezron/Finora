package interfaces

import "finora/database/model"

type IncomeInterface interface {
	CreateIncome(income *model.Income) (*model.Income, error)
	GetIncomeByID(id uint) (*model.Income, error)
	GetAllIncomes() ([]model.Income, error)
	UpdateIncome(id uint, income *model.Income) (*model.Income, error)
	DeleteIncome(id uint) error
	GetIncomesByUserID(userID uint) ([]model.Income, error)
	GetIncomesBySourceID(sourceID uint) ([]model.Income, error)
}