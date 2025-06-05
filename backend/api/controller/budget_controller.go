package controller

import (
	"finora/api/interfaces"
	"finora/database"
	"finora/database/model"
)

type BudgetController struct{}

func NewBudgetController() interfaces.BudgetInterface {
	return &BudgetController{}
}

func (bc *BudgetController) CreateBudget(budget *model.Budget) (*model.Budget, error) {
	if err := database.DB.Create(budget).Error; err != nil {
		return nil, err
	}
	return budget, nil
}

func (bc *BudgetController) GetBudgetByID(id uint) (*model.Budget, error) {
	var budget model.Budget
	if err := database.DB.First(&budget, id).Error; err != nil {
		return nil, err
	}
	return &budget, nil
}

func (bc *BudgetController) GetAllBudgets() ([]model.Budget, error) {
	var budgets []model.Budget
	if err := database.DB.Find(&budgets).Error; err != nil {
		return nil, err
	}
	return budgets, nil
}

func (bc *BudgetController) UpdateBudget(id uint, budget *model.Budget) (*model.Budget, error) {
	var existingBudget model.Budget
	if err := database.DB.First(&existingBudget, id).Error; err != nil {
		return nil, err
	}

	existingBudget.Name = budget.Name
	existingBudget.Amount = budget.Amount
	existingBudget.Description = budget.Description
	existingBudget.Period = budget.Period

	if err := database.DB.Save(&existingBudget).Error; err != nil {
		return nil, err
	}
	return &existingBudget, nil
}

func (bc *BudgetController) DeleteBudget(id uint) error {
	var budget model.Budget
	if err := database.DB.First(&budget, id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&budget).Error; err != nil {
		return err
	}
	return nil
}

func (bc *BudgetController) GetBudgetsByUserID(userID uint) ([]model.Budget, error) {
	var budgets []model.Budget
	if err := database.DB.Preload("User").Where("user_id = ?", userID).Find(&budgets).Error; err != nil {
		return nil, err
	}
	return budgets, nil
}

func (bc *BudgetController) GetBudgetsByCategoryID(categoryID uint) ([]model.Budget, error) {
	var budgets []model.Budget
	if err := database.DB.Preload("ExpenseCategory").Where("category_id = ?", categoryID).Find(&budgets).Error; err != nil {
		return nil, err
	}
	return budgets, nil
}