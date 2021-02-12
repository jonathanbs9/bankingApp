package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start func => Create mux, handlers and server.
func Start() {
	// We create a new server with mux
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id}", getCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	// Server starting
	log.Println("Conectado puerto 8000")
	http.ListenAndServe(":8000", router)

}
