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
	var transactions []models.ExpenseTransaction
	var response models.ExpenseTransactionResponse
	paramPage := c.DefaultQuery("page", "1")
	paramPageSize := c.DefaultQuery("page_size", "100")
	config.DB.Unscoped().Scopes(Paginate(paramPage, paramPageSize)).Order("transaction_datetime asc").Find(&transactions)
	response.Data = transactions

	var total int64
	config.DB.Table("expense_transactions").Count(&total)
	response.Total = total

	c.IndentedJSON(http.StatusOK, response)
}

func UpdateExpenseTransaction(c *gin.Context) {
	transaction_id := c.Param("transaction_id")
	transaction_id_as_int, _ := strconv.Atoi(transaction_id)
	var t models.ExpenseTransaction
	db := config.DB
	if err := db.Unscoped().First(&t, transaction_id_as_int).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Expense transaction not found"})
		return
	}

	var input models.ExpenseTransaction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	t.Amount = input.Amount
	t.Category0 = input.Category0
	t.Category1 = input.Category1
	t.Currency = input.Currency
	t.Memo0 = input.Memo0
	t.Memo1 = input.Memo1
	t.SourceAccount = input.SourceAccount
	t.TargetAccount = input.TargetAccount
	t.TransactionDatetime = input.TransactionDatetime

	db.Save(&t)

	c.JSON(http.StatusOK, t)
}
