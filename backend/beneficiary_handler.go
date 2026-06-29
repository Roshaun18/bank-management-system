package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func AddBeneficiary(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO]Add Beneficiary Endpoint Called")
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var beneficiary Beneficiary

	err := json.NewDecoder(r.Body).Decode(&beneficiary)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if beneficiary.CustomerID == "" || beneficiary.AccountID == "" || beneficiary.Beneficiary == "" {
		http.Error(w, "All Fields Are Required", http.StatusBadRequest)
		return
	}

	collection := DB.Collection("beneficiaries")

	var existing Beneficiary

	err = collection.FindOne(context.Background(), bson.M{"customer_id": beneficiary.CustomerID, "account_id": beneficiary.AccountID}).Decode(&existing)

	if err == nil {
		http.Error(w, "Beneficiary Already Exixits", http.StatusBadRequest)
		return
	}

	_, err = collection.InsertOne(context.Background(), beneficiary)

	if err != nil {
		http.Error(w, "Failed To Add Beneficiary", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Beneficiary Added Successfully",
	})
	log.Printf("[INFO]New Beneficiary with User Name: %s & Account ID: %s was Created", beneficiary.Beneficiary, beneficiary.AccountID)
}

func GetBeneficiaries(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO]Get Beneficiary Endpoint Called")
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	customerID := r.URL.Query().Get("customer_id")

	if customerID == "" {
		http.Error(w, "Customer ID Required", http.StatusBadRequest)
		return
	}

	collection := DB.Collection("beneficiaries")

	cursor, err := collection.Find(context.Background(), bson.M{"customer_id": customerID})

	if err != nil {
		log.Printf("[ERROR]Failed To Fetch the Beneficiaries")
		http.Error(w, "Failed To Fetch Beneficiaries", http.StatusInternalServerError)
		return
	}

	var beneficiaries []Beneficiary

	err = cursor.All(context.Background(), &beneficiaries)

	if err != nil {
		http.Error(w, "Failed To Decode Data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(beneficiaries)
	log.Printf("[INFO]Beneficiaries added in the Customer ID; %s was Fetched", customerID)
}

func DeleteBeneficiary(w http.ResponseWriter, r *http.Request) {
	log.Printf("[INFO]Delete Beneficiary Endpoint Called")
	if r.Method != http.MethodDelete {
		http.Error(w, "method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Beneficiary

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	collection := DB.Collection("beneficiaries")

	result, err := collection.DeleteOne(context.Background(), bson.M{"customer_id": req.CustomerID, "account_id": req.AccountID})
	if err != nil {
		log.Printf("[ERROR]Failed to Delete the Beneficiary")
		http.Error(w, "Failed To Delete Beneficiary", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Beneficiary Not Found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Benificiary Deleted Successfully",
	})
	log.Printf("[INFO]Beneficiary Was Deleted Successfully")

}
