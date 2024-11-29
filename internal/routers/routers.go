package routers

import (
	"github.com/pastry-personal5/tt-apiserver/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	expenseTransactionRoutes := r.Group("/expense_transactions")
	{
		expenseTransactionRoutes.GET("/", handlers.GetExpenseTransactions)
	}

	return r
}
