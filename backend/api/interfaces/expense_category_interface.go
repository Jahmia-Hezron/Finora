package interfaces

import "finora/database/model"

type ExpenseCategoryInterface interface {
	CreateExpenseCategory(category *model.ExpenseCategory) (*model.ExpenseCategory, error)
	GetExpenseCategoryByID(id uint) (*model.ExpenseCategory, error)
	GetAllExpenseCategories() ([]model.ExpenseCategory, error)
	UpdateExpenseCategory(id uint, category *model.ExpenseCategory) (*model.ExpenseCategory, error)
	DeleteExpenseCategory(id uint) error
	GetExpenseCategoriesByUserID(userID uint) ([]model.ExpenseCategory, error)
}