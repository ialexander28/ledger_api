package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Transactions []Transaction

var Accounts []Account

var is_equal bool = true

type Transaction struct {
	Id         string `json:"Id"`
	Amount     string `json:"Amount"`
	Direction  string `json:"Direction"`
	Account_id string `json:"Account_id"`
}

type Account struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	Direction string `json:"Direction"`
	Balance   string `json:"Balance"`
}

func getAll_Tx(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Transactions)
}

func get_Tx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, transaction := range Transactions {
		if transaction.Id == key {
			json.NewEncoder(w).Encode(transaction)
		}
	}
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	var transaction Transaction
	var account Account

	_ = json.NewDecoder(r.Body).Decode(&transaction)
	//transaction.Id = vars["id"]
	Transactions = append(Transactions, transaction)
	json.NewEncoder(w).Encode(transaction)

	_ = json.NewDecoder(r.Body).Decode(&account.Balance)
	//account.Id = vars["id"]
	Accounts = append(Accounts, account)
	json.NewEncoder(w).Encode(account)

	if account.Id == transaction.Id {
		fmt.Println(is_equal)
		Accounts = append(Accounts, account)
		json.NewEncoder(w).Encode(account)
	} else {
		fmt.Println(!is_equal)
	}
	if is_equal {
		fmt.Println("Ending balance is: ", (account.Balance + transaction.Amount))
	} else {
		fmt.Println("Ending balance is: negative") //(account.Balance + transaction.Amount))
		Accounts = append(Accounts, account)

	}
}

func getAllAccounts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Accounts)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, account := range Accounts {
		if account.Id == key {
			json.NewEncoder(w).Encode(account)
		}
	}
}

func createAccount(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	var account Account
	_ = json.NewDecoder(r.Body).Decode(&account)
	//account.Id = vars["id"]
	Accounts = append(Accounts, account)
	json.NewEncoder(w).Encode(account)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "To create a test transaction use the following curl input: curl -d {} -X POST http://localhost:3000/transaction/")
}

func apiRequests() {
	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/", homePage)
	route.HandleFunc("/transaction", getAll_Tx)
	route.HandleFunc("/createTransaction", createTransaction).Methods("POST")
	route.HandleFunc("/transaction/{id}", get_Tx)
	route.HandleFunc("/account", getAllAccounts)
	route.HandleFunc("/account/{id}", getAccount)
	route.HandleFunc("/createAccount", createAccount).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", route))
}

func main() {
	Transactions = []Transaction{

		{Id: "5",
			Amount:     "TEST",
			Direction:  "TEST",
			Account_id: "TEST"},
	}
	Accounts = []Account{
		{Id: "0x48f2f18febd5799f",
			Name:      "bridge",
			Direction: "debit",
			Balance:   "0"},
	}
	apiRequests()

}
