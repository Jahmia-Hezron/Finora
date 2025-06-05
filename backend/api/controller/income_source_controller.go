package controller

import (
	"finora/api/interfaces"
	"finora/database"
	"finora/database/model"
)

type IncomeSourceController struct{}

func NewIncomeSourceController() interfaces.IncomeSourceInterface {
	return &IncomeSourceController{}
}



func (isc *IncomeSourceController) CreateIncomeSource(incomeSource *model.IncomeSource) (*model.IncomeSource, error) {
	if err := database.DB.Create(incomeSource).Error; err != nil {
		return nil, err
	}
	return incomeSource, nil
}

func (isc *IncomeSourceController) GetIncomeSourceByID(id uint) (*model.IncomeSource, error) {
	var incomeSource model.IncomeSource
	if err := database.DB.First(&incomeSource, id).Error; err != nil {
		return nil, err
	}
	return &incomeSource, nil
}

func (isc *IncomeSourceController) GetAllIncomeSources() ([]model.IncomeSource, error) {
	var incomeSources []model.IncomeSource
	if err := database.DB.Find(&incomeSources).Error; err != nil {
		return nil, err
	}
	return incomeSources, nil
}

func (isc *IncomeSourceController) UpdateIncomeSource(id uint, incomeSource *model.IncomeSource) (*model.IncomeSource, error) {
	var existingIncomeSource model.IncomeSource
	if err := database.DB.First(&existingIncomeSource, id).Error; err != nil {
		return nil, err
	}

	existingIncomeSource.Name = incomeSource.Name
	existingIncomeSource.Description = incomeSource.Description

	if err := database.DB.Save(&existingIncomeSource).Error; err != nil {
		return nil, err
	}
	return &existingIncomeSource, nil
}

func (isc *IncomeSourceController) DeleteIncomeSource(id uint) error {
	var incomeSource model.IncomeSource
	if err := database.DB.First(&incomeSource, id).Error; err != nil {
		return err
	}

	if err := database.DB.Delete(&incomeSource).Error; err != nil {
		return err
	}
	return nil
}

func (isc *IncomeSourceController) GetIncomeSourcesByUserID(userID uint) ([]model.IncomeSource, error) {
	var incomeSources []model.IncomeSource
	if err := database.DB.Preload("User").Where("user_id = ?", userID).Find(&incomeSources).Error; err != nil {
		return nil, err
	}
	return incomeSources, nil
}
