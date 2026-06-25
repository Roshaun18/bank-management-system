package main

import (
	"log"
	"net/http"
)

func main() {
	err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", HomeHandler)
	log.Println("Server Started on: 8080")
	http.HandleFunc("/customer", CreateCustomer)
	http.HandleFunc("/get-customer", GetCustomer)
	http.HandleFunc("/account", CreateAccount)
	http.HandleFunc("/deposit", Deposit)
	http.HandleFunc("/withdraw", Withdraw)
	http.HandleFunc("/balance", GetBalance)
	http.HandleFunc("/transactions", GetTransactions)
	http.HandleFunc("/transfer", TransferMoney)
	http.HandleFunc("/dashboard", GetDashboard)
	http.HandleFunc("/get-accounts", GetAccount)
	http.HandleFunc("/customer-summary", GetCustomerSummary)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/customer/accounts", GetCustomerAccounts)
	http.HandleFunc("/change-password", ChangePassword)
	http.HandleFunc("/generate-otp", GenerateOtp)
	http.HandleFunc("/verify-otp", VerifyOtp)
	http.HandleFunc("/reset-password", ResetPassword)
	http.HandleFunc("/beneficiary", AddBeneficiary)
	http.HandleFunc("/beneficiaries", GetBeneficiaries)
	http.HandleFunc("/beneficiary/delete", DeleteBeneficiary)
	handler := enableCORS(http.DefaultServeMux)
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}

}
