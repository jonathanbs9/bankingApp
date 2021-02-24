package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonathanbs9/bankingApp/domain"
	"github.com/jonathanbs9/bankingApp/service"
)

// Start func => Create mux, handlers and server.
func Start() {
	// We create a new server with mux
	router := mux.NewRouter()

	// Wired (Cableado)
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)


	// Server starting
	log.Println("Conected on port 8000")
	http.ListenAndServe(":8000", router)

}
