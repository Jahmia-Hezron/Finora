package controller

import (
	"finora/api/interfaces"
	"finora/database"
	"finora/database/model"
)

type ExpenseCategoryController struct{}

func NewExpenseCategoryController() interfaces.ExpenseCategoryInterface {
	return &ExpenseCategoryController{}
}

func (ecc *ExpenseCategoryController) CreateExpenseCategory(category *model.ExpenseCategory) (*model.ExpenseCategory, error) {
	if err := database.DB.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (ecc *ExpenseCategoryController) GetExpenseCategoryByID(id uint) (*model.ExpenseCategory, error) {
	var category model.ExpenseCategory
	if err := database.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (ecc *ExpenseCategoryController) GetAllExpenseCategories() ([]model.ExpenseCategory, error) {
	var categories []model.ExpenseCategory
	if err := database.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (ecc *ExpenseCategoryController) UpdateExpenseCategory(id uint, category *model.ExpenseCategory) (*model.ExpenseCategory, error) {
	var existingCategory model.ExpenseCategory
	if err := database.DB.First(&existingCategory, id).Error; err != nil {
		return nil, err
	}

	existingCategory.Name = category.Name
	existingCategory.Description = category.Description

	if err := database.DB.Save(&existingCategory).Error; err != nil {
		return nil, err
	}
	return &existingCategory, nil
}

func (ecc *ExpenseCategoryController) DeleteExpenseCategory(id uint) error {
	var category model.ExpenseCategory
	if err := database.DB.First(&category, id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

func (ecc *ExpenseCategoryController) GetExpenseCategoriesByUserID(userID uint) ([]model.ExpenseCategory, error) {
	var categories []model.ExpenseCategory
	if err := database.DB.Preload("User").Where("user_id = ?", userID).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (ecc *ExpenseCategoryController) GetExpenseCategoriesByExpenseID(expenseID uint) ([]model.ExpenseCategory, error) {
	var categories []model.ExpenseCategory
	if err := database.DB.Preload("Expenses").Where("expense_id = ?", expenseID).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}