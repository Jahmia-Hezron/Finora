package handler

import (
	"finora/api/controller"
	"finora/api/interfaces"
	"finora/database/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BudgetHandler struct {
	BudgetCpntroller interfaces.BudgetInterface
}

func NewBudgetHandler() *BudgetHandler {
	return &BudgetHandler{
		BudgetCpntroller: controller.NewBudgetController(),
	}
}

func (bh *BudgetHandler) CreateBudget(c *gin.Context) {
	var budget model.Budget

	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	createdBudget, err := bh.BudgetCpntroller.CreateBudget(&budget)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create budget: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Budget created successfully",
		"budget": gin.H{
			"id":          createdBudget.ID,
			"name":        createdBudget.Name,
			"description": createdBudget.Description,
			"amount":      createdBudget.Amount,
			"category_id": createdBudget.CategoryID,
			"period":      createdBudget.Period,
			"user_id":     createdBudget.UserID,

		},
	})
}

func (bh *BudgetHandler) GetBudgetByID(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Budget ID not found in context",
		})
		return
	}

	budget, err := bh.BudgetCpntroller.GetBudgetByID(id.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve budget: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"budget": gin.H{
			"id":          budget.ID,
			"name":        budget.Name,
			"description": budget.Description,
			"amount":      budget.Amount,
			"category_id": budget.CategoryID,
			"period":      budget.Period,
			"user_id":     budget.UserID,
		},
	})
}

func (bh *BudgetHandler) GetAllBudgets(c *gin.Context) {
	budgets, err := bh.BudgetCpntroller.GetAllBudgets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve budgets: " + err.Error(),
		})
		return
	}

	var budgetList []gin.H
	for _, budget := range budgets {
		budgetList = append(budgetList, gin.H{
			"id":          budget.ID,
			"name":        budget.Name,
			"description": budget.Description,
			"amount":      budget.Amount,
			"category_id": budget.CategoryID,
			"period":      budget.Period,
			"user_id":     budget.UserID,	
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"budgets": budgetList,
	})
}

func (bh *BudgetHandler) UpdateBudget(c *gin.Context) {
	var budget model.Budget
	if err := c.ShouldBindJSON(&budget); err != nil {
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
			"message": "Budget ID not found in context",
		})
		return
	}

	updatedBudget, err := bh.BudgetCpntroller.UpdateBudget(id.(uint), &budget)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to update budget: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Budget updated successfully",
		"budget": gin.H{
			"id":          updatedBudget.ID,
			"name":        updatedBudget.Name,
			"description": updatedBudget.Description,
			"amount":      updatedBudget.Amount,
			"category_id": updatedBudget.CategoryID,
			"period":      updatedBudget.Period,
			"user_id":     updatedBudget.UserID,
		},
	})
}

func (bh *BudgetHandler) DeleteBudget(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Budget ID not found in context",
		})
		return
	}

	err := bh.BudgetCpntroller.DeleteBudget(id.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete budget: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Budget deleted successfully",
	})
}

func (bh *BudgetHandler) GetBudgetsByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User ID not found in context",
		})
		return
	}

	budgets, err := bh.BudgetCpntroller.GetBudgetsByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve budgets: " + err.Error(),
		})
		return
	}

	var budgetList []gin.H
	for _, budget := range budgets {
		budgetList = append(budgetList, gin.H{
			"id":          budget.ID,
			"name":        budget.Name,
			"description": budget.Description,
			"amount":      budget.Amount,
			"category_id": budget.CategoryID,
			"period":      budget.Period,
			"user_id":     budget.UserID,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"budgets": budgetList,
	})
}

func (bh *BudgetHandler) GetBudgetsByCategoryID(c *gin.Context) {
	categoryID, exists := c.Get("category_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Category ID not found in context",
		})
		return
	}

	budgets, err := bh.BudgetCpntroller.GetBudgetsByCategoryID(categoryID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve budgets: " + err.Error(),
		})
		return
	}

	var budgetList []gin.H
	for _, budget := range budgets {
		budgetList = append(budgetList, gin.H{
			"id":          budget.ID,
			"name":        budget.Name,
			"description": budget.Description,
			"amount":      budget.Amount,
			"category_id": budget.CategoryID,
			"period":      budget.Period,
			"user_id":     budget.UserID,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"budgets": budgetList,
	})
}