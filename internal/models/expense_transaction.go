package models

import "gorm.io/gorm"

type ExpenseTransaction struct {
	gorm.Model
	Amount              int    `json:"amount"`
	Category0           string `json:"category0"`
	Category1           string `json:"category1"`
	Currency            string `json:"currency"`
	Memo0               string `json:"memo0"`
	Memo1               string `json:"memo1"`
	SourceAccount       string `json:"source_account"`
	TargetAccount       string `json:"target_account"`
	TransactionDatetime string `json:"transaction_datetime"`
}

type ExpenseTransactionResponse struct {
	Data  []ExpenseTransaction `json:"data"`
	Total int64                `json:"total"`
}

type ExpenseTransactionMonthlyAnalysis struct {
	gorm.Model
	UserIdentifier string `json:"user_identifier"`
	Month          string `json:"month"`
	Category0      string `json:"category0"`
	TotalSum       string `json:"total_sum"`
}

type ExpenseTransactionMonthlyAnalysisResponse struct {
	Data  []ExpenseTransactionMonthlyAnalysis `json:"data"`
	Total int64                               `json:"total"`
}

type ExpenseTransactionMonthlyAnalysisForCountOfDistinctMonthsResponse struct {
	UserIdentifier string `json:"user_identifier"`
	Total          int    `json:"total"`
}

type ExpenseCategories struct {
	gorm.Model
	UserIdentifier string `json:"user_identifier"`
	Name           string `json:"name"`
	UUID           string `json:"uuid"`
}

type ExpenseCategoriesForCountOfDistinctNamesResponse struct {
	UserIdentifier string `json:"user_identifier"`
	Total          int    `json:"total"`
}
