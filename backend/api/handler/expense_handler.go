package handler

import (
	"finora/api/controller"
	"finora/api/interfaces"
	"finora/database/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpenseHandler struct {
	ExpenseController interfaces.ExpenseInterface
}

func NewExpenseHandler() *ExpenseHandler {
	return &ExpenseHandler{
		ExpenseController: controller.NewExpenseController(),
	}
}

func (eh *ExpenseHandler) CreateExpense(c *gin.Context) {
	var expense model.Expense

	if err := c.ShouldBindJSON(&expense); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	createdExpense, err := eh.ExpenseController.CreateExpense(&expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create expense: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Expense created successfully",
		"expense": gin.H{
			"id":          createdExpense.ID,
			"amount":      createdExpense.Amount,
			"description": createdExpense.Description,
			"user_id":     createdExpense.UserID,
			"date":        createdExpense.Date,
			"category_id": createdExpense.CategoryID,
		},
	})
}

func (eh *ExpenseHandler) GetExpenseByID(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Expense ID not found in context",
		})
		return
	}

	expense, err := eh.ExpenseController.GetExpenseByID(id.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve expense: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Expense retrieved successfully",
		"expense": expense,
	})
}

func (eh *ExpenseHandler) GetAllExpenses(c *gin.Context) {
	expenses, err := eh.ExpenseController.GetAllExpenses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve expenses: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Expenses retrieved successfully",
		"expenses": expenses,
	})
}

func (eh *ExpenseHandler) UpdateExpense(c *gin.Context) {
	var expense model.Expense
	if err := c.ShouldBindJSON(&expense); err != nil {
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
			"message": "Expense ID not found in context",
		})
		return
	}

	updatedExpense, err := eh.ExpenseController.UpdateExpense(id.(uint), &expense)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to update expense: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Expense updated successfully",
		"expense": updatedExpense,
	})
}

func (eh *ExpenseHandler) DeleteExpense(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Expense ID not found in context",
		})
		return
	}

	if err := eh.ExpenseController.DeleteExpense(id.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete expense: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Expense deleted successfully",
	})
}

func (eh *ExpenseHandler) GetExpensesByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User ID not found in context",
		})
		return
	}

	expenses, err := eh.ExpenseController.GetExpensesByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve expenses: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Expenses retrieved successfully",
		"expenses": expenses,
	})
}

func (eh *ExpenseHandler) GetExpensesByCategoryID(c *gin.Context) {
	categoryID, exists := c.Get("category_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Category ID not found in context",
		})
		return
	}

	expenses, err := eh.ExpenseController.GetExpensesByCategoryID(categoryID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve expenses: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"message":  "Expenses retrieved successfully",
		"expenses": expenses,
	})
}
