package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jonathanbs9/bankingApp/errs"
	"github.com/jonathanbs9/bankingApp/logger"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	/*c, err := sql.Open("mysql", "root:@tcp(locahost:3306)/banking")
	if err != nil {
		log.Println("Error al conectar a la base de datos  => " + err.Error())
		return nil, errs.NewUnexpectedError("Error inesperado en la base de datos")
	}
	c.SetConnMaxLifetime(time.Second * 5)
	c.SetMaxOpenConns(10)
	c.SetMaxIdleConns(10)*/
	var rows  *sql.Rows
	var err error

	if status == ""{
		query := "select customer_id, first_name, last_name, date_birth, city, zipcode, status from customers"
		rows, err = d.client.Query(query)
	} else {
		query := "select customer_id, first_name, last_name, date_birth, city, zipcode, status from customers where status = ?"
		rows, err = d.client.Query(query, status)
	}

	//rows, err = d.client.Query(query)
	if err != nil {
		logger.Error("No se pueden obtener resultados (GetAllCustomers) de la BD => \n " + err.Error())
		return nil, errs.NewUnexpectedError("Error inesperado en la base de datos")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
		if err != nil {
			logger.Error("Error al scanear customers => " + err.Error())
			return nil, errs.NewUnexpectedError("Error inesperado en la base de datos")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) GetCustomerById(id string) (*Customer, *errs.AppError) {
	// Hacemos una llamada a la base de datos
	customerSql := "select customer_id, first_name, last_name, city, zipcode, date_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer

	err:= row.Scan(&c.Id, &c.FirstName, &c.LastName, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil{
		if err == sql.ErrNoRows{
			return nil, errs.NewNotFoundError("Cliente no encontrado")
		} else {
			logger.Error("Error al buscar un cliente => " + err.Error())
			return nil , errs.NewUnexpectedError("Error inesperado en la base de datos")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:@tcp(localhost:3306)/banking")
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
