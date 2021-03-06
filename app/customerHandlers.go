package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanbs9/bankingApp/service"
)

// Construyo un handler que va a tener el service. De esta manera conecto handler-service
type CustomerHandlers struct {
	service service.CustomerService
}

// getAllCustomers func =>
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	// Llamo al handler -> service -> func
	customers, err := ch.service.GetAllCustomer(status)
	if err!= nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		// Case success
		writeResponse(w, http.StatusOK, customers)
	}
}

// getCustomer func =>
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	id:= vars["id"]

	// Here we link handler with service
	customer, err := ch.service.GetCustomerById(id)

	if err!= nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		// Case success
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}){
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil{
		panic(err)
	}
}

func getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, vars["id"])
	fmt.Println(vars["id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	message := "Customer created"
	json.NewEncoder(w).Encode(message)
}
