package handler

import (
	"finora/api/controller"
	"finora/api/interfaces"
	"finora/database/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IncomeSourceHandler struct {
	IncomeSourceController interfaces.IncomeSourceInterface
}

func NewIncomeSourceHandler() *IncomeSourceHandler {
	return &IncomeSourceHandler{
		IncomeSourceController: controller.NewIncomeSourceController(),
	}
}

func (ish *IncomeSourceHandler) CreateIncomeSource(c *gin.Context) {
	var incomeSource model.IncomeSource

	if err := c.ShouldBindJSON(&incomeSource); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid input: " + err.Error(),
		})
		return
	}

	createdIncomeSource, err := ish.IncomeSourceController.CreateIncomeSource(&incomeSource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create income source: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Income source created successfully",
		"income_source": gin.H{
			"id":          createdIncomeSource.ID,
			"name":        createdIncomeSource.Name,
			"description": createdIncomeSource.Description,
			"user_id":     createdIncomeSource.UserID,
		},
	})
}

func (ish *IncomeSourceHandler) GetIncomeSourceByID(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Income source ID not found in context",
		})
		return
	}

	incomeSource, err := ish.IncomeSourceController.GetIncomeSourceByID(id.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Income source not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"income_source": gin.H{
			"id":          incomeSource.ID,
			"name":        incomeSource.Name,
			"description": incomeSource.Description,
			"user_id":     incomeSource.UserID,
		},
	})
}

func (ish *IncomeSourceHandler) GetAllIncomeSources(c *gin.Context) {
	incomeSources, err := ish.IncomeSourceController.GetAllIncomeSources()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve income sources: " + err.Error(),
		})
		return
	}

	var response []gin.H
	for _, source := range incomeSources {
		response = append(response, gin.H{
			"id":          source.ID,
			"name":        source.Name,
			"description": source.Description,
			"user_id":     source.UserID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          http.StatusOK,
		"income_sources": response,
	})
}

func (ish *IncomeSourceHandler) UpdateIncomeSource(c *gin.Context) {
	var incomeSource model.IncomeSource
	if err := c.ShouldBindJSON(&incomeSource); err != nil {
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
			"message": "Income source ID not found in context",
		})
		return
	}

	updatedIncomeSource, err := ish.IncomeSourceController.UpdateIncomeSource(id.(uint), &incomeSource)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to update income source: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Income source updated successfully",
		"income_source": gin.H{
			"id":          updatedIncomeSource.ID,
			"name":        updatedIncomeSource.Name,
			"description": updatedIncomeSource.Description,
			"user_id":     updatedIncomeSource.UserID,
		},
	})
}

func (ish *IncomeSourceHandler) DeleteIncomeSource(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Income source ID not found in context",
		})
		return
	}

	if err := ish.IncomeSourceController.DeleteIncomeSource(id.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to delete income source: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Income source deleted successfully",
	})
}

func (ish *IncomeSourceHandler) GetIncomeSourcesByUserID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User ID not found in context",
		})
		return
	}

	incomeSources, err := ish.IncomeSourceController.GetIncomeSourcesByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to retrieve income sources: " + err.Error(),
		})
		return
	}

	var response []gin.H
	for _, source := range incomeSources {
		response = append(response, gin.H{
			"id":          source.ID,
			"name":        source.Name,
			"description": source.Description,
			"user_id":     source.UserID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          http.StatusOK,
		"income_sources": response,
	})
}
