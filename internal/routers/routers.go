package routers

import (
	"github.com/pastry-personal5/tt-apiserver/internal/handlers"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(CORSMiddleware())

	expenseTransactionsRoutes := r.Group("/expense_transactions")
	{
		expenseTransactionsRoutes.GET("/", handlers.GetExpenseTransactions)
		expenseTransactionsRoutes.POST("/:id", handlers.UpdateExpenseTransaction)
	}

	expenseTransactionMonthlyAnalysisRoutes := r.Group("/expense_transactions_monthly_analysis")
	{
		expenseTransactionMonthlyAnalysisRoutes.GET("/", handlers.GetExpenseTransactionsMonthlyAnalysis)
		expenseTransactionMonthlyAnalysisRoutes.GET("/count_of_distinct_months", handlers.GetExpenseTransactionsMonthlyAnalysisForCountOfDistinctMonths)
	}

	expenseCategoriesRoutes := r.Group("/expense_categories")
	{
		//expenseCategoriesRoutes.GET("/", handlers.GetExpenseCategories)
		expenseCategoriesRoutes.GET("/count_of_distinct_names", handlers.GetExpenseCategoriesForCountOfDistinctNames)
	}

	return r
}
