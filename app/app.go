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

	// Cableado
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	/*router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers/{id}", getCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)*/

	// Server starting
	log.Println("Conectado puerto 8000")
	http.ListenAndServe(":8000", router)

}
