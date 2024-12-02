package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pastry-personal5/tt-apiserver/internal/models"
	"github.com/pastry-personal5/tt-apiserver/internal/services"
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
	tableName := "expense_transactions"
	services.DB.Table(tableName).Unscoped().Scopes(Paginate(paramPage, paramPageSize)).Order("transaction_datetime asc").Find(&transactions)
	response.Data = transactions

	var total int64
	services.DB.Table("expense_transactions").Count(&total)
	response.Total = total

	c.IndentedJSON(http.StatusOK, response)
}

func UpdateExpenseTransaction(c *gin.Context) {
	id := c.Param("id")
	id_as_int, _ := strconv.Atoi(id)
	var t models.ExpenseTransaction
	db := services.DB
	if err := db.Unscoped().First(&t, id_as_int).Error; err != nil {
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

func GetExpenseTransactionsMonthlyAnalysis(c *gin.Context) {
	var t []models.ExpenseTransactionMonthlyAnalysis
	var response models.ExpenseTransactionMonthlyAnalysisResponse
	paramPage := c.DefaultQuery("page", "1")
	paramPageSize := c.DefaultQuery("page_size", "100")
	userIdentifier := c.DefaultQuery("user_identifier", "user1")
	tableName := "expense_transactions_monthly_analysis"
	services.DB.Table(tableName).Unscoped().Scopes(Paginate(paramPage, paramPageSize)).Order("month, category0 asc").Find(&t, "user_identifier = ?", userIdentifier)
	response.Data = t

	var total int64
	services.DB.Table("expense_transactions_monthly_analysis").Count(&total)
	response.Total = total

	c.IndentedJSON(http.StatusOK, response)
}

func GetExpenseTransactionsMonthlyAnalysisForCountOfDistinctMonths(c *gin.Context) {
	var t []models.ExpenseTransactionMonthlyAnalysis
	tableName := "expense_transactions_monthly_analysis"
	userIdentifier := c.DefaultQuery("user_identifier", "user1")
	services.DB.Table(tableName).Unscoped().Distinct("month").Find(&t, "user_identifier = ?", userIdentifier)
	var response models.ExpenseTransactionMonthlyAnalysisForCountOfDistinctMonthsResponse
	response.UserIdentifier = userIdentifier
	response.Total = len(t)
	c.IndentedJSON(http.StatusOK, response)
}

func GetExpenseCategoriesForCountOfDistinctNames(c *gin.Context) {
	var t []models.ExpenseCategories
	tableName := "expense_categories"
	userIdentifier := c.DefaultQuery("user_identifier", "user1")
	services.DB.Table(tableName).Unscoped().Distinct("name").Find(&t, "user_identifier = ?", userIdentifier)
	var response models.ExpenseCategoriesForCountOfDistinctNamesResponse
	response.UserIdentifier = userIdentifier
	response.Total = len(t)
	c.IndentedJSON(http.StatusOK, response)
}
