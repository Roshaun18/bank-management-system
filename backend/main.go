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
	handler := enableCORS(http.DefaultServeMux)
	err = http.ListenAndServe(":8080", handler)
	if err != nil {
		log.Fatal(err)
	}

}
