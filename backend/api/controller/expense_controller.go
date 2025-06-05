package controller

import (
	"finora/api/interfaces"
	"finora/database"
	"finora/database/model"
)

type ExpenseController struct{}

func NewExpenseController() interfaces.ExpenseInterface {
	return &ExpenseController{}
}

func (ec *ExpenseController) CreateExpense(expense *model.Expense) (*model.Expense, error) {
	if err := database.DB.Create(expense).Error; err != nil {
		return nil, err
	}
	return expense, nil
}

func (ec *ExpenseController) GetExpenseByID(id uint) (*model.Expense, error) {
	var expense model.Expense
	if err := database.DB.First(&expense, id).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

func (ec *ExpenseController) GetAllExpenses() ([]model.Expense, error) {
	var expenses []model.Expense
	if err := database.DB.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func (ec *ExpenseController) UpdateExpense(id uint, expense *model.Expense) (*model.Expense, error) {
	var existingExpense model.Expense
	if err := database.DB.First(&existingExpense, id).Error; err != nil {
		return nil, err
	}

	existingExpense.Amount = expense.Amount
	existingExpense.Date = expense.Date
	existingExpense.Description = expense.Description

	if err := database.DB.Save(&existingExpense).Error; err != nil {
		return nil, err
	}
	return &existingExpense, nil
}

func (ec *ExpenseController) DeleteExpense(id uint) error {
	var expense model.Expense
	if err := database.DB.First(&expense, id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&expense).Error; err != nil {
		return err
	}
	return nil
}

func (ec *ExpenseController) GetExpensesByUserID(userID uint) ([]model.Expense, error) {
	var expenses []model.Expense
	if err := database.DB.Preload("User").Where("user_id = ?", userID).Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}

func (ec *ExpenseController) GetExpensesByCategoryID(categoryID uint) ([]model.Expense, error) {
	var expenses []model.Expense
	if err := database.DB.Preload("ExpenseCategory").Where("category_id = ?", categoryID).Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}