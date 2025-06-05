package interfaces

import "finora/database/model"

type BudgetInterface interface {
	CreateBudget(budget *model.Budget) (*model.Budget, error)
	GetBudgetByID(id uint) (*model.Budget, error)
	GetAllBudgets() ([]model.Budget, error)
	UpdateBudget(id uint, budget *model.Budget) (*model.Budget, error)
	DeleteBudget(id uint) error
	GetBudgetsByUserID(userID uint) ([]model.Budget, error)
	GetBudgetsByCategoryID(categoryID uint) ([]model.Budget, error)
}