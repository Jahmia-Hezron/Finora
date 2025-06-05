package handler

import (
	"finora/api/controller"
	"finora/api/interfaces"
	"finora/database/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IncomeHandler struct {
	IncomeController interfaces.IncomeInterface
}

func NewIncomeHandler() *IncomeHandler {
	return &IncomeHandler{
		IncomeController: controller.NewIncomeController(),
	}
}

func (ih *IncomeHandler) CreateIncome(c *gin.Context) {
	var income model.Income

	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	createdIncome, err := ih.IncomeController.CreateIncome(&income)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create income: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Income created successfully",
		"income": gin.H{
			"id":          createdIncome.ID,
			"amount":      createdIncome.Amount,
			"description": createdIncome.Description,
			"income_source_id": createdIncome.SourceID,
			"user_id":     createdIncome.UserID,
			"date":        createdIncome.Date,
		},
	})
}

func (ih *IncomeHandler) GetIncomeByID(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Income ID not found in context",
		})
		return
	}

	income, err := ih.IncomeController.GetIncomeByID(id.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve income: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"income": gin.H{
			"id":          income.ID,
			"amount":      income.Amount,
			"description": income.Description,
			"income_source_id": income.SourceID,
			"user_id":     income.UserID,
			"date":        income.Date,
		},
	})
}

func (ih *IncomeHandler) GetAllIncomes(c *gin.Context) {
	incomes, err := ih.IncomeController.GetAllIncomes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve incomes: " + err.Error(),
		})
		return
	}

	var incomeList []gin.H
	for _, income := range incomes {
		incomeList = append(incomeList, gin.H{
			"id":          income.ID,
			"amount":      income.Amount,
			"description": income.Description,
			"income_source_id": income.SourceID,
			"user_id":     income.UserID,
			"date":        income.Date,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"incomes": incomeList,
	})
}

func (ih *IncomeHandler) UpdateIncome(c *gin.Context) {
	var income model.Income
	if err := c.ShouldBindJSON(&income); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Income ID not found in context",
		})
		return
	}

	updatedIncome, err := ih.IncomeController.UpdateIncome(id.(uint), &income)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to update income: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status" :  http.StatusOK,
		"message": "Income updated successfully",
		"income": gin.H{
			"id":          updatedIncome.ID,
			"amount":      updatedIncome.Amount,
			"description": updatedIncome.Description,
			"income_source_id": updatedIncome.SourceID,
			"user_id":     updatedIncome.UserID,
			"date":        updatedIncome.Date,
		},
	})
}

func (ih *IncomeHandler) DeleteIncome(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Income ID not found in context",
		})
		return
	}

	if err := ih.IncomeController.DeleteIncome(id.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete income: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Income deleted successfully",
	})
}

func (ih *IncomeHandler) GetIncomesByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User ID not found in context",
		})
		return
	}

	incomes, err := ih.IncomeController.GetIncomesByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve incomes: " + err.Error(),
		})
		return
	}

	var incomeList []gin.H
	for _, income := range incomes {
		incomeList = append(incomeList, gin.H{
			"id":          income.ID,
			"amount":      income.Amount,
			"description": income.Description,
			"income_source_id": income.SourceID,
			"user_id":     income.UserID,
			"date":        income.Date,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"incomes": incomeList,
	})
}

func (ih *IncomeHandler) GetIncomesBySourceID(c *gin.Context) {
	sourceID, exists := c.Get("source_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Source ID not found in context",
		})
		return
	}

	incomes, err := ih.IncomeController.GetIncomesBySourceID(sourceID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve incomes: " + err.Error(),
		})
		return
	}

	var incomeList []gin.H
	for _, income := range incomes {
		incomeList = append(incomeList, gin.H{
			"id":          income.ID,
			"amount":      income.Amount,
			"description": income.Description,
			"income_source_id": income.SourceID,
			"user_id":     income.UserID,
			"date":        income.Date,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"incomes": incomeList,
	})
}