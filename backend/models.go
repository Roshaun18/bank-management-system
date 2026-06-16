package main

import (
	"time"
)

type Customer struct {
	ID       string `json:"id" bson:"id"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type Account struct {
	AccountID  string  `json:"account_id" bson:"account_id"`
	CustomerID string  `json:"customer_id" bson:"customer_id"`
	Balance    float64 `json:"balance" bson:"balance"`
}

type Transaction struct {
	ID          string    `json:"id" bson:"id"`
	AccountID   string    `json:"account_id" bson:"account_id"`
	ToAccountID string    `json:"to_account_id,omitempty" bson:"to_account_id,omitempty"`
	Amount      float64   `json:"amount" bson:"amount"`
	Type        string    `json:"type" bson:"type"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}

type TransactionRequest struct {
	AccountID string  `json:"account_id"`
	Amount    float64 `json:"amount"`
}

type TransferRequest struct {
	FromAccount string  `json:"from_account"`
	ToAccount   string  `json:"to_account"`
	Amount      float64 `json:"amount"`
}

type DashboardResponse struct {
	AccountID        string  `json:"account_id"`
	Balance          float64 `json:"balance"`
	TransactionCount int     `json:"transaction_count"`
	TotalDeposit     float64 `json:"total_deposit"`
	TotalWithdraw    float64 `json:"total_withdraw"`
	TotalTransfer    float64 `json:"total_transfer"`
}

type AccountInfo struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}

type CustomerSummary struct {
	CustomerID string        `json:"customer_id"`
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Accounts   []AccountInfo `json:"accounts"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	CustomerID  string `json:"customer_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
