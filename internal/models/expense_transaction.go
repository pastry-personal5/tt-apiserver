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
	TransactionID       int    `json:"transaction_id" gorm:"primaryKey"`
}
