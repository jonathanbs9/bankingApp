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

/*func greet(w http.ResponseWriter, r *http.Request) {
	message := "Hola Mundo bankApp"

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(message)
}*/

// Construyo un handler que va a tener el service. De esta manera conecto handler-service
type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// llamo al handler -> service -> func
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		log.Fatal("Error al obtener los clientes | " + err.Error())
		return
	}
	// HardCoded customers
	/*customers := []Customer{
		{
			FirstName: "Jonathan",
			LastName:  "Brull Schroeder",
			City:      "Mar del Plata",
		},
		{
			FirstName: "Hernan",
			LastName:  "Cattaneo",
			City:      "Buenos Aires",
		},
		{
			FirstName: "Armin",
			LastName:  "Van Buuren",
			City:      "Jarbeus",
		},
	}*/

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

	customer, err := ch.service.GetCustomerById(id)
	if err!= nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(w, err.Error())
	}
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
