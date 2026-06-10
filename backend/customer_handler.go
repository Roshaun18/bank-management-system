package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bank application api running")

}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	customer.ID = fmt.Sprintf("C%d", time.Now().UnixNano())
	if customer.Name == "" || customer.Email == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	collection := DB.Collection("customers")
	var exixtingCustomer Customer

	err = collection.FindOne(context.Background(), bson.M{"email": customer.Email}).Decode(&exixtingCustomer)

	if err == nil {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}
	_, err = collection.InsertOne(context.Background(), customer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer Created Successfully", "id": customer.ID})
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetCustomer called")
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	customerID := r.URL.Query().Get("id")
	if customerID == "" {
		http.Error(w, "Customer ID Required", http.StatusBadRequest)
		return
	}

	collection := DB.Collection("customers")

	var customer Customer
	err := collection.FindOne(context.Background(), map[string]string{"id": customerID}).Decode(&customer)

	if err != nil {
		http.Error(w, "Customer Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)

}

func GetCustomerSummary(w http.ResponseWriter, r *http.Request) {
	customerCollection := DB.Collection("customers")
	accountCollection := DB.Collection("accounts")

	cursor, err := customerCollection.Find(context.Background(), bson.M{})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var customers []Customer

	err = cursor.All(context.Background(), &customers)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result []CustomerSummary

	for _, customer := range customers {
		accountCursor, _ := accountCollection.Find(context.Background(), bson.M{"customer_id": customer.ID})

		var accounts []Account

		accountCursor.All(context.Background(), &accounts)

		var accountInfos []AccountInfo

		for _, account := range accounts {
			accountInfos = append(accountInfos, AccountInfo{AccountID: account.AccountID, Balance: account.Balance})
		}

		result = append(result, CustomerSummary{
			CustomerID: customer.ID,
			Name:       customer.Name,
			Email:      customer.Email,
			Accounts:   accountInfos,
		})
	}

	json.NewEncoder(w).Encode(result)
}
