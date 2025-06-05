package routes

import (
	"finora/api/handler"
	"finora/api/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userHandler := handler.NewUserHandler()
	incomeSourceHandler := handler.NewIncomeSourceHandler()
	incomeHandler := handler.NewIncomeHandler()
	expenseCategoryHandler := handler.NewExpenseCategoryHandler()
	expenseHandler := handler.NewExpenseHandler()
	budgetHandler := handler.NewBudgetHandler()

	// ------------------- User Routes -------------------
	userRoutes := router.Group("/api/v1/user")
	{
		userRoutes.POST("/register", userHandler.CreateUser)
		userRoutes.POST("/login", userHandler.LogInUser)
		userRoutes.POST("/logout", middleware.AuthMiddleware(), userHandler.LogOutUser)

		userRoutes.GET("/", middleware.AuthMiddleware(), userHandler.GetAllUsers)
		userRoutes.GET("/:id", middleware.AuthMiddleware(), userHandler.GetUserByID)
		userRoutes.PUT("/:id", middleware.AuthMiddleware(), userHandler.UpdateUser)
		userRoutes.PUT("/:id/reset-password", middleware.AuthMiddleware(), userHandler.ResetPassword)
		userRoutes.DELETE("/:id", middleware.AuthMiddleware(), userHandler.DeleteUser)
	}

	// ------------------- Income Source Routes -------------------
	incomeSourceRoutes := router.Group("/api/v1/income-sources")
	incomeSourceRoutes.Use(middleware.AuthMiddleware())
	{
		incomeSourceRoutes.GET("/", incomeSourceHandler.GetAllIncomeSources)
		incomeSourceRoutes.POST("/create", incomeSourceHandler.CreateIncomeSource)
		incomeSourceRoutes.GET("/:id", incomeSourceHandler.GetIncomeSourceByID)
		incomeSourceRoutes.GET("/user/:userID", incomeSourceHandler.GetIncomeSourcesByUserID)
		incomeSourceRoutes.PUT("/:id", incomeSourceHandler.UpdateIncomeSource)
		incomeSourceRoutes.DELETE("/:id", incomeSourceHandler.DeleteIncomeSource)
	}

	// ------------------- Income Routes -------------------
	incomeRoutes := router.Group("/api/v1/incomes")
	incomeRoutes.Use(middleware.AuthMiddleware())
	{
		incomeRoutes.POST("/create", incomeHandler.CreateIncome)
		incomeRoutes.GET("/", incomeHandler.GetAllIncomes)
		incomeRoutes.GET("/:id", incomeHandler.GetIncomeByID)
		incomeRoutes.GET("/user/:userID", incomeHandler.GetIncomesByUserID)
		incomeRoutes.GET("/source/:sourceID", incomeHandler.GetIncomesBySourceID)
		incomeRoutes.PUT("/:id", incomeHandler.UpdateIncome)
		incomeRoutes.DELETE("/:id", incomeHandler.DeleteIncome)
	}

	// ------------------- Expense Category Routes -------------------
	expenseCategoryRoutes := router.Group("/api/v1/expense-categories")
	expenseCategoryRoutes.Use(middleware.AuthMiddleware())
	{
		expenseCategoryRoutes.POST("/create", expenseCategoryHandler.CreateExpenseCategory)
		expenseCategoryRoutes.GET("/", expenseCategoryHandler.GetAllExpenseCategories)
		expenseCategoryRoutes.GET("/:id", expenseCategoryHandler.GetExpenseCategoryByID)
		expenseCategoryRoutes.GET("/user/:userID", expenseCategoryHandler.GetExpenseCategoriesByUserID)
		expenseCategoryRoutes.PUT("/:id", expenseCategoryHandler.UpdateExpenseCategory)
		expenseCategoryRoutes.DELETE("/:id", expenseCategoryHandler.DeleteExpenseCategory)
	}

	// ------------------- Expense Routes -------------------
	expenseRoutes := router.Group("/api/v1/expenses")
	expenseRoutes.Use(middleware.AuthMiddleware())
	{
		expenseRoutes.POST("/create", expenseHandler.CreateExpense)
		expenseRoutes.GET("/", expenseHandler.GetAllExpenses)
		expenseRoutes.GET("/:id", expenseHandler.GetExpenseByID)
		expenseRoutes.GET("/user/:userID", expenseHandler.GetExpensesByUserID)
		expenseRoutes.GET("/category/:categoryID", expenseHandler.GetExpensesByCategoryID)
		expenseRoutes.PUT("/:id", expenseHandler.UpdateExpense)
		expenseRoutes.DELETE("/:id", expenseHandler.DeleteExpense)
	}

	// ------------------- Budget Routes -------------------
	budgetRoutes := router.Group("/api/v1/budgets")
	budgetRoutes.Use(middleware.AuthMiddleware())
	{
		budgetRoutes.POST("/create", budgetHandler.CreateBudget)
		budgetRoutes.GET("/", budgetHandler.GetAllBudgets)
		budgetRoutes.GET("/:id", budgetHandler.GetBudgetByID)
		budgetRoutes.GET("/user/:userID", budgetHandler.GetBudgetsByUserID)
		budgetRoutes.GET("/category/:categoryID", budgetHandler.GetBudgetsByCategoryID)
		budgetRoutes.PUT("/:id", budgetHandler.UpdateBudget)
		budgetRoutes.DELETE("/:id", budgetHandler.DeleteBudget)
	}
}
