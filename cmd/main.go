package main

import (
	transaction "RoadToTribal2.0"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"strconv"
	"time"
)

func getTransaction(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "ID")
	if ID == "" {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	trans := &transaction.Transaction{}
	err, trans := transaction.GetTransactionById(ID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Contact with ID: %s not found", ID), 404)
		return
	}
	json.NewEncoder(w).Encode(trans)
	w.WriteHeader(200)
}

func getAllTransactions(w http.ResponseWriter, r *http.Request) {
	err, transactions := transaction.GetAllTransaction()
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	json.NewEncoder(w).Encode(transactions)
	w.WriteHeader(200)
}

func addTransaction(w http.ResponseWriter, r *http.Request) {
	existingtransaction := &transaction.Transaction{}
	var trans transaction.Transaction
	json.NewDecoder(r.Body).Decode(&trans)
	trans.CreatedAt = time.Now()
	err, existingtransaction := transaction.GetTransactionById(strconv.Itoa(existingtransaction.ID))
	if err == nil {
		http.Error(w, fmt.Sprintf("Contact with ID: %s already exist", strconv.Itoa(trans.ID)), 400)
		return
	}
	err, _ = transaction.Addtransction(&trans)
	if err != nil {
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	w.Write([]byte("Transaction created successfully"))
	w.WriteHeader(201)
}

func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/transactions", func(r chi.Router) {
		r.Get("/", getAllTransactions) //GET /transactions
		r.Get("/{ID}", getTransaction) //GET /transactions/0147344454
		r.Post("/", addTransaction)    //POST /transactions
		/*r.Put("/{ID}", updateTransaction)    //PUT /transactions/0147344454
		r.Delete("/{ID}", deleteTransaction) //DELETE /transactions/0147344454*/
	})
	return r
}

func main() {
	r := registerRoutes()
	log.Fatal(http.ListenAndServe(":8080", r))
}
