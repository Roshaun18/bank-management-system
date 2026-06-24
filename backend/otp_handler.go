package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func GenerateOtp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req GenerateOtpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid Request Body", http.StatusBadRequest)
		return
	}

	customerCollection := DB.Collection("customers")

	var customer Customer

	err = customerCollection.FindOne(context.Background(), bson.M{"username": req.Username}).Decode(&customer)

	if err != nil {
		http.Error(w, "Customer Not FOund", http.StatusNotFound)
		return
	}

	otp := fmt.Sprintf("%06d", rand.Intn(900000)+100000)

	log.Println("Generated OTP:", otp)

	otpCollection := DB.Collection("otp_store")

	_, err = otpCollection.UpdateOne(context.Background(), bson.M{"username": req.Username}, bson.M{"$set": bson.M{"usernmae": req.Username, "otp": otp, "expires_at": time.Now().Add(5 * time.Minute)}}, options.UpdateOne().SetUpsert(true))

	if err != nil {
		http.Error(w, "Failed to Generate OTP", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "OTP Sent Successfully",
	})

}

func VerifyOtp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req VerifyOtpRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	otpCollection := DB.Collection("otp_store")

	var otpRecord OtpRecord

	err = otpCollection.FindOne(context.Background(), bson.M{"username": req.Username}).Decode(&otpRecord)

	if err != nil {
		http.Error(w, "OTP Not Found", http.StatusNotFound)
		return
	}

	if time.Now().After(otpRecord.ExpiresAt) {
		http.Error(w, "OTP Expired", http.StatusUnauthorized)
		return
	}

	if otpRecord.Otp != req.Otp {
		http.Error(w, "Invalid OTP", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "OTP Verified Successfully",
	})
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ResetPasswordRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Failed to Hash Password", http.StatusInternalServerError)
		return
	}

	customerCollection := DB.Collection("customers")

	_, err = customerCollection.UpdateOne(context.Background(), bson.M{"username": req.Username}, bson.M{"$set": bson.M{"password": string(hashedPassword)}})

	if err != nil {
		http.Error(w, "Failed To Reset Password", http.StatusInternalServerError)
		return
	}

	otpCollection := DB.Collection("otp_score")

	otpCollection.DeleteOne(context.Background(), bson.M{"username": req.Username})

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Password Reset Successfully",
	})
}
