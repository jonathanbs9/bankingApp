package domain

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jonathanbs9/bankingApp/errs"
	"github.com/jonathanbs9/bankingApp/logger"
	"log"
	"os"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows  *sql.Rows
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		query := "select customer_id, first_name, last_name, date_birth, city, zipcode, status from customers"
		err = d.client.Select(&customers, query)
	} else {
		query := "select customer_id, first_name, last_name, date_birth, city, zipcode, status from customers where status = ?"
		// 1 destination, 2 query, 3 argument
		err = d.client.Select(&customers, query, status)
	}

	if err != nil {
		logger.Error("No se pueden obtener resultados (GetAllCustomers) de la BD => " +err.Error())
		return nil, errs.NewUnexpectedError("Error inesperado en la base de datos")
	}
	return customers, nil
}

// Func GetCustomerById
func (d CustomerRepositoryDb) GetCustomerById(id string) (*Customer, *errs.AppError) {
	// Query
	customerSql := "select customer_id, first_name, last_name, city, zipcode, date_birth, status from customers where customer_id = ?"
	var c Customer
	// Hacemos una llamada a la base de datos, obtenemos en c un customer. Si no trae nada devuelve un error
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Cliente no encontrado => "+ sql.ErrNoRows.Error())
		} else {
			logger.Error("Error al buscar un cliente => " + err.Error())
			return nil, errs.NewUnexpectedError("Error inesperado en la base de datos")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
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

	return CustomerRepositoryDb{
		client: client,
	}
}
