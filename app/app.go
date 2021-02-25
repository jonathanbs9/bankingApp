package app

import (
	"fmt"
	"github.com/jonathanbs9/bankingApp/logger"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jonathanbs9/bankingApp/domain"
	"github.com/jonathanbs9/bankingApp/service"
)

// Start func => Create mux, handlers and server.
func Start() {
	// Sanity Check
	sanityCheck()
	// We create a new server with mux
	router := mux.NewRouter()

	// Wired (Cableado)
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// Getting ENV variables
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	fmt.Println("address=" + address+ "\nport="+port)



	// Server starting
	http.ListenAndServe(fmt.Sprintf("%s:%s",address, port), router)
	logger.Info("Connected on port: " + address)

}

func sanityCheck(){
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == ""{
		log.Fatal("Entorno no ha sido definido")
	}
}
