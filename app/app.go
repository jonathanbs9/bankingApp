package app

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jonathanbs9/bankingApp/logger"
	"log"
	"net/http"
	"os"
	"time"

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

	// We create a DB client
	dbClient := getDbClient()

	// Wired (Cableado)
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	// Define handlers
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// Define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	// Getting ENV variables
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	fmt.Println("address=" + address+ "\nport="+port)

	// Server starting
	http.ListenAndServe(fmt.Sprintf("%s:%s",address, port), router)
	logger.Info("Connected on port: " + address)

}

func getDbClient() *sqlx.DB{
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName:= os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",dbUser, dbPass, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos => " + err.Error())
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetConnMaxIdleTime(10)
	client.SetMaxOpenConns(10)

	return client
}

func sanityCheck(){
	if os.Getenv("SERVER_ADDRESS") == "" || os.Getenv("SERVER_PORT") == ""{
		log.Fatal("Entorno no ha sido definido")
	}
}
