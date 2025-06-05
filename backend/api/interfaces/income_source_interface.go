package interfaces

import "finora/database/model"

type IncomeSourceInterface interface {
	CreateIncomeSource(incomeSource *model.IncomeSource) (*model.IncomeSource, error)
	GetIncomeSourceByID(id uint) (*model.IncomeSource, error)
	GetAllIncomeSources() ([]model.IncomeSource, error)
	UpdateIncomeSource(id uint, incomeSource *model.IncomeSource) (*model.IncomeSource, error)
	DeleteIncomeSource(id uint) error
	GetIncomeSourcesByUserID(userID uint) ([]model.IncomeSource, error)
}