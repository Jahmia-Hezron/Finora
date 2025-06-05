package handler

import (
	"finora/api/controller"
	"finora/api/interfaces"
	"finora/database/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpenseCateforyHandler struct {
	ExpenseCategoryController interfaces.ExpenseCategoryInterface
}

func NewExpenseCategoryHandler() *ExpenseCateforyHandler {
	return &ExpenseCateforyHandler{
		ExpenseCategoryController: controller.NewExpenseCategoryController(),
	}
}

func (ech *ExpenseCateforyHandler) CreateExpenseCategory(c *gin.Context) {
	var expenseCategory model.ExpenseCategory

	if err := c.ShouldBindJSON(&expenseCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	createdExpenseCategory, err := ech.ExpenseCategoryController.CreateExpenseCategory(&expenseCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create expense category: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Expense category created successfully",
		"expense_category": gin.H{
			"id":          createdExpenseCategory.ID,
			"name":        createdExpenseCategory.Name,
			"description": createdExpenseCategory.Description,
			"user_id":     createdExpenseCategory.UserID,
		},
	})
}

func (ech *ExpenseCateforyHandler) GetExpenseCategoryByID(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Expense category ID not found in context",
		})
		return
	}

	expenseCategory, err := ech.ExpenseCategoryController.GetExpenseCategoryByID(id.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve expense category: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Expense category retrieved successfully",
		"expense_category": gin.H{
			"id":          expenseCategory.ID,
			"name":        expenseCategory.Name,
			"description": expenseCategory.Description,
			"user_id":     expenseCategory.UserID,
		},
	})
}

func (ech *ExpenseCateforyHandler) GetAllExpenseCategories(c *gin.Context) {
	expenseCategories, err := ech.ExpenseCategoryController.GetAllExpenseCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve expense categories: " + err.Error(),
		})
		return
	}

	var response []gin.H
	for _, category := range expenseCategories {
		response = append(response, gin.H{
			"id":          category.ID,
			"name":        category.Name,
			"description": category.Description,
			"user_id":     category.UserID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Expense categories retrieved successfully",
		"expense_categories": response,
	})
}

func (ech *ExpenseCateforyHandler) UpdateExpenseCategory(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Expense category ID not found in context",
		})
		return
	}

	var expenseCategory model.ExpenseCategory
	if err := c.ShouldBindJSON(&expenseCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	updatedExpenseCategory, err := ech.ExpenseCategoryController.UpdateExpenseCategory(id.(uint), &expenseCategory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to update expense category: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Expense category updated successfully",
		"expense_category": gin.H{
			"id":          updatedExpenseCategory.ID,
			"name":        updatedExpenseCategory.Name,
			"description": updatedExpenseCategory.Description,
			"user_id":     updatedExpenseCategory.UserID,
		},
	})
}

func (ech *ExpenseCateforyHandler) DeleteExpenseCategory(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Expense category ID not found in context",
		})
		return
	}

	err := ech.ExpenseCategoryController.DeleteExpenseCategory(id.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete expense category: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Expense category deleted successfully",
	})
}

func (ech *ExpenseCateforyHandler) GetExpenseCategoriesByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User ID not found in context",
		})
		return
	}

	expenseCategories, err := ech.ExpenseCategoryController.GetExpenseCategoriesByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve expense categories for user: " + err.Error(),
		})
		return
	}

	var response []gin.H
	for _, category := range expenseCategories {
		response = append(response, gin.H{
			"id":          category.ID,
			"name":        category.Name,
			"description": category.Description,
			"user_id":     category.UserID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Expense categories for user retrieved successfully",
		"expense_categories": response,
	})
}

