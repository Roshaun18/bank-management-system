package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var account Account

	err := json.NewDecoder(r.Body).Decode(&account)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	customerCollection := DB.Collection("customers")

	var customer Customer
	err = customerCollection.FindOne(context.Background(), bson.M{"id": account.CustomerID}).Decode(&customer)

	if err != nil {
		http.Error(w, "Customer Not Found", http.StatusNotFound)
		return
	}

	accountCollection := DB.Collection("accounts")

	account.AccountID = fmt.Sprintf("A%d", time.Now().UnixNano())

	if account.Balance < 0 {
		http.Error(w, "Initial balance cannot be negative", http.StatusBadRequest)
		return
	}

	_, err = accountCollection.InsertOne(context.Background(), account)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Account Created Successfully", "account_id": account.AccountID})
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetCustomer Accounts called")
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	accountID := r.URL.Query().Get("account_id")

	if accountID == "" {
		http.Error(w, "Account ID is required", http.StatusBadRequest)
		return
	}

	accountCollection := DB.Collection("accounts")

	var account Account

	err := accountCollection.FindOne(context.Background(), bson.M{"account_id": accountID}).Decode(&account)

	if err != nil {
		http.Error(w, "Account Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func GetCustomerAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	customerID := r.URL.Query().Get("customer_id")

	if customerID == "" {
		http.Error(w, "Customer ID is required", http.StatusBadRequest)
		return
	}

	accountCollection := DB.Collection("accounts")

	cursor, err := accountCollection.Find(context.Background(), bson.M{"customer_id": customerID})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var accounts []Account
	err = cursor.All(context.Background(), &accounts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accounts)
}
