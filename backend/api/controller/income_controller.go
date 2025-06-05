package controller

import (
	"finora/api/interfaces"
	"finora/database"
	"finora/database/model"
)

type IncomeController struct{}

func NewIncomeController() interfaces.IncomeInterface {
	return &IncomeController{}
}


func (ic *IncomeController) CreateIncome(income *model.Income) (*model.Income, error) {
	if err := database.DB.Create(income).Error; err != nil {
		return nil, err
	}
	return income, nil
}

func (ic *IncomeController) GetIncomeByID(id uint) (*model.Income, error) {
	var income model.Income
	if err := database.DB.First(&income, id).Error; err != nil {
		return nil, err
	}
	return &income, nil
}

func (ic *IncomeController) GetAllIncomes() ([]model.Income, error) {
	var incomes []model.Income
	if err := database.DB.Find(&incomes).Error; err != nil {
		return nil, err
	}
	return incomes, nil
}

func (ic *IncomeController) UpdateIncome(id uint, income *model.Income) (*model.Income, error) {
	var existingIncome model.Income
	if err := database.DB.First(&existingIncome, id).Error; err != nil {
		return nil, err
	}

	existingIncome.Amount = income.Amount
	existingIncome.Date = income.Date
	existingIncome.Description = income.Description

	if err := database.DB.Save(&existingIncome).Error; err != nil {
		return nil, err
	}
	return &existingIncome, nil
}

func (ic *IncomeController) DeleteIncome(id uint) error {
	var income model.Income
	if err := database.DB.First(&income, id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&income).Error; err != nil {
		return err
	}
	return nil
}

func (ic *IncomeController) GetIncomesByUserID(userID uint) ([]model.Income, error) {
	var incomes []model.Income
	if err := database.DB.Preload("User").Where("user_id = ?", userID).Find(&incomes).Error; err != nil {
		return nil, err
	}
	return incomes, nil
}

func (ic *IncomeController) GetIncomesBySourceID(sourceID uint) ([]model.Income, error) {
	var incomes []model.Income
	if err := database.DB.Preload("IncomeSource").Where("source_id = ?", sourceID).Find(&incomes).Error; err != nil {
		return nil, err
	}
	return incomes, nil
}
