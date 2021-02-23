package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanbs9/bankingApp/service"
)

type Customer struct {
	FirstName string `json:"first_name" xml:"first_name"`
	LastName  string `json:"last_name" xml:"last_name"`
	City      string `json:"city" xml:"city"`
}

// Construyo un handler que va a tener el service. De esta manera conecto handler-service
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// Llamo al handler -> service -> func
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		log.Fatal("Error al obtener los clientes | " + err.Error())
		return
	}

	if r.Header.Get("Content-type") == "application/xml" {
		// XML format
		w.Header().Set("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		// Json format
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars:= mux.Vars(r)
	id:= vars["id"]

	// Here we link handler with service
	customer, err := ch.service.GetCustomerById(id)
	if err!= nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(w, err.Error())
	}
	// Case success
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
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
