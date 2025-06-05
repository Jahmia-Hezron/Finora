package interfaces

import "finora/database/model"

type ExpenseInterface interface {
	CreateExpense(expense *model.Expense) (*model.Expense, error)
	GetExpenseByID(id uint) (*model.Expense, error)
	GetAllExpenses() ([]model.Expense, error)
	UpdateExpense(id uint, expense *model.Expense) (*model.Expense, error)
	DeleteExpense(id uint) error
	GetExpensesByUserID(userID uint) ([]model.Expense, error)
	GetExpensesByCategoryID(categoryID uint) ([]model.Expense, error)
}