package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bank application api running")

}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO]Create Customer Endpoint Called")
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
	if customer.Name == "" || customer.Email == "" || customer.Username == "" || customer.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to Hash Password", http.StatusInternalServerError)
		return
	}
	customer.Password = string(hashedPassword)

	collection := DB.Collection("customers")
	var existingCustomer Customer

	err = collection.FindOne(context.Background(), bson.M{"email": customer.Email, "username": customer.Username}).Decode(&existingCustomer)

	if err == nil {
		log.Printf("[WARN]Username or Email Already Exists")
		http.Error(w, "Email or Username already exists", http.StatusBadRequest)
		return
	}
	_, err = collection.InsertOne(context.Background(), customer)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Customer Created Successfully", "id": customer.ID})
	log.Printf("[INFO]New Customer was Created with Customer ID: %s", customer.ID)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO]Get Customer Endpoint Called")
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
		log.Printf("[ERROR]Customer ID Not Found")
		http.Error(w, "Customer Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
	log.Printf("[INFO]Custmer with ID: %s was Fetched", customerID)

}

func GetCustomerSummary(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO]Get Customer Summary Endpoint Called")
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
	log.Printf("[INFO]Customer Summary Fteched For All Customer")
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO]Change Password Endpoint Called")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ChangePasswordRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	collection := DB.Collection(("customers"))
	var customer Customer
	err = collection.FindOne(context.Background(), bson.M{"id": req.CustomerID}).Decode(&customer)

	if err != nil {
		http.Error(w, "Customer Not Found", http.StatusNotFound)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.OldPassword))
	if err != nil {
		log.Printf("[ERROR]Current Password Was Incorrect")
		http.Error(w, "Current Password Incorrect", http.StatusUnauthorized)
		return
	}
	if req.NewPassword == "" {
		http.Error(w, "New Password Cannot Be Empty", http.StatusBadRequest)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(req.OldPassword), []byte(req.NewPassword))
	if err == nil {
		http.Error(w, "New Password Must Be Different", http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)

	_, err = collection.UpdateOne(context.Background(), bson.M{"id": req.CustomerID}, bson.M{"$set": bson.M{"password": string(hashedPassword)}})
	if err != nil {
		log.Printf("[ERROR]Failed To Update Password")
		http.Error(w, "Failed to Update Password", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"message": "Password Updated Successfully"})
	log.Printf("[INFO]Password was Changed Successfully")
}
