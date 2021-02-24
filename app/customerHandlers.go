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
		w.Header().Add("Content-type", "application/xml")
		w.WriteHeader(http.StatusOK)
		xml.NewEncoder(w).Encode(customers)
	} else {
		// Json format
		writeResponse(w, http.StatusOK, customers)
	}
}

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
