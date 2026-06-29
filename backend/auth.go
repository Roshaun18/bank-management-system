package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("[INFO] Login endpoint called")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Printf("[ERROR] Failed to decode login request: %v", err)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if req.Username == "admin" && req.Password == "admin123" {
		json.NewEncoder(w).Encode(map[string]string{
			"message": "login Successful",
			"role":    "admin",
		})
		return
	}

	collection := DB.Collection("customers")
	var customer Customer

	err = collection.FindOne(context.Background(), bson.M{"username": req.Username}).Decode(&customer)

	if err != nil {
		log.Printf("[ERROR]Invalid Username or Password")
		http.Error(w, "Invalid Username or Password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password))

	if err != nil {
		log.Printf("[ERROR]Invalid Username or Password")
		http.Error(w, "Invalid Username or Password", http.StatusUnauthorized)
		return
	}

	accountCollection := DB.Collection("accounts")

	var account Account

	err = accountCollection.FindOne(context.Background(), bson.M{"customer_id": customer.ID}).Decode(&account)

	if err != nil {
		http.Error(w, "Acount Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Login Successful",
		"role":        "customer",
		"customer_id": customer.ID,
	})

	log.Printf("[INFO]User %s logged in ", req.Username)

}
