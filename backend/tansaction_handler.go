package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func Deposit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var req TransactionRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, "Amount must be greater than zero", http.StatusBadRequest)
		return
	}

	accountCollection := DB.Collection("accounts")

	var account Account

	err = accountCollection.FindOne(context.Background(), bson.M{"account_id": req.AccountID}).Decode(&account)

	if err != nil {
		http.Error(w, "Account Not Found", http.StatusNotFound)
		return
	}

	newBalance := account.Balance + req.Amount

	_, err = accountCollection.UpdateOne(context.Background(), bson.M{"account_id": req.AccountID}, bson.M{"$set": bson.M{"balance": newBalance}})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transactionCollection := DB.Collection("transactions")

	transaction := Transaction{
		ID:        fmt.Sprintf("T%d", time.Now().UnixNano()),
		AccountID: req.AccountID,
		Amount:    req.Amount,
		Type:      "DEPOSIT",
	}

	_, err = transactionCollection.InsertOne(context.Background(), transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":     "Deposit Successful",
		"new_balance": newBalance,
	})
}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method NOt Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TransactionRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, "Amount must be greater than zero", http.StatusBadRequest)
		return
	}

	accountCollection := DB.Collection("accounts")
	var account Account

	err = accountCollection.FindOne(context.Background(), bson.M{"account_id": req.AccountID}).Decode(&account)

	if err != nil {
		http.Error(w, "Account Not Found", http.StatusNotFound)
		return
	}

	if req.Amount > account.Balance {
		http.Error(w, "Insufficient Balance", http.StatusBadRequest)
		return
	}

	newBalance := account.Balance - req.Amount

	_, err = accountCollection.UpdateOne(context.Background(), bson.M{"account_id": req.AccountID}, bson.M{"$set": bson.M{"balance": newBalance}})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transactionCollection := DB.Collection("transactions")

	transaction := Transaction{
		ID:        fmt.Sprintf("T%d", time.Now().UnixNano()),
		AccountID: req.AccountID,
		Amount:    req.Amount,
		Type:      "WITHDRAW",
	}

	_, err = transactionCollection.InsertOne(context.Background(), transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":     "Withdraw Successful",
		"new_balance": newBalance,
	})
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	accountID := r.URL.Query().Get("account_id")

	if accountID == "" {
		http.Error(w, "Account ID Required", http.StatusBadRequest)
		return
	}

	collection := DB.Collection("accounts")

	var account Account

	err := collection.FindOne(context.Background(), bson.M{"account_id": accountID}).Decode(&account)

	if err != nil {
		http.Error(w, "Account Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"account_id": account.AccountID,
		"balance":    account.Balance,
	})
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	accountID := r.URL.Query().Get("account_id")

	if accountID == "" {
		http.Error(w, "Account ID Required", http.StatusBadRequest)
		return

	}

	collection := DB.Collection("transactions")

	var transactions []Transaction

	cursor, err := collection.Find(context.Background(), bson.M{"account_id": accountID})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &transactions)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(transactions)
}

func TransferMoney(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req TransferRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, "Amount must be greater than zero", http.StatusBadRequest)
		return
	}

	if req.FromAccount == req.ToAccount {
		http.Error(w, "Cannnot transfer to same account", http.StatusBadRequest)
		return
	}

	collection := DB.Collection("accounts")

	var sender, receiver Account

	err = collection.FindOne(context.Background(), bson.M{"account_id": req.FromAccount}).Decode(&sender)

	if err != nil {
		http.Error(w, "Sender Account Not FOund", http.StatusNotFound)
		return
	}

	err = collection.FindOne(context.Background(), bson.M{"account_id": req.ToAccount}).Decode(&receiver)

	if err != nil {
		http.Error(w, "Reveiver Account Not Found", http.StatusNotFound)
		return
	}

	if sender.Balance < req.Amount {
		http.Error(w, "Insuficient Balance", http.StatusNotFound)
		return
	}

	newSenderBalance := sender.Balance - req.Amount
	newReceiverBalance := receiver.Balance + req.Amount

	_, err = collection.UpdateOne(context.Background(), bson.M{"account_id": req.FromAccount}, bson.M{"$set": bson.M{"balance": newSenderBalance}})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"account_id": req.ToAccount}, bson.M{"$set": bson.M{"Balance": newReceiverBalance}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transactionCollection := DB.Collection("transactions")

	transaction := Transaction{
		ID:          fmt.Sprintf("T%d", time.Now().UnixNano()),
		AccountID:   req.FromAccount,
		ToAccountID: req.ToAccount,
		Amount:      req.Amount,
		Type:        "TRANSFER",
	}

	_, err = transactionCollection.InsertOne(context.Background(), transaction)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Transfer Successful",
		"from":    req.FromAccount,
		"to":      req.ToAccount,
		"amount":  req.Amount,
	})

}

func GetDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	accountID := r.URL.Query().Get("account_id")

	if accountID == "" {
		http.Error(w, "Account ID Required", http.StatusBadRequest)
		return
	}

	accountCollection := DB.Collection("accounts")

	var account Account

	err := accountCollection.FindOne(context.Background(), bson.M{"account_id": accountID}).Decode(&account)

	if err != nil {
		http.Error(w, "Account Not Found", http.StatusNotFound)
		return
	}

	transactionCollection := DB.Collection("transactions")

	var transactions []Transaction

	cursor, err := transactionCollection.Find(context.Background(), bson.M{"account_id": accountID})

	err = cursor.All(context.Background(), &transactions)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	var totalDeposit float64
	var totalWithdraw float64
	var totalTransfer float64

	for _, txn := range transactions {
		switch txn.Type {
		case "DEPOSIT":
			totalDeposit += txn.Amount
		case "WITHDRAW":
			totalWithdraw += txn.Amount
		case "TRANSFER":
			totalTransfer += txn.Amount

		}
	}

	response := DashboardResponse{
		AccountID:        accountID,
		Balance:          account.Balance,
		TransactionCount: len(transactions),
		TotalDeposit:     totalDeposit,
		TotalWithdraw:    totalWithdraw,
		TotalTransfer:    totalTransfer,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
