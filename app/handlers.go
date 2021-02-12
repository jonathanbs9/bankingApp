package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	FirstName string `json:"first_name" xml:"first_name"`
	LastName  string `json:"last_name" xml:"last_name"`
	City      string `json:"city" xml:"city"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	message := "Hola Mundo bankApp"

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers := []Customer{
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
		{
			FirstName: "Popof",
			LastName:  "Popote",
			City:      "Paris",
		}, {
			FirstName: "Joseph",
			LastName:  "Capriati",
			City:      "Napoli",
		},
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
