package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pastry-personal5/tt-apiserver/internal/config"
	"github.com/pastry-personal5/tt-apiserver/internal/models"
	"gorm.io/gorm"
)

func Paginate(paramPage string, paramPageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(paramPage)
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(paramPageSize)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func GetExpenseTransactions(c *gin.Context) {
	var expense_transactions []models.ExpenseTransaction
	paramPage := c.DefaultQuery("page", "1")
	paramPageSize := c.DefaultQuery("page_size", "100")
	config.DB.Unscoped().Scopes(Paginate(paramPage, paramPageSize)).Order("transaction_datetime asc").Find(&expense_transactions)
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.JSON(http.StatusOK, expense_transactions)
}
