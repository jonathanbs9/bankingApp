package domain

import (
	"database/sql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d *CustomerRepositoryDb) FindAll() ([]Customer, error) {
	c, err := sql.Open("mysql", "root:@tcp(locahost:3306)/banking")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos => " + err.Error())
	}

	c.SetConnMaxLifetime(time.Second * 5)
	c.SetMaxOpenConns(10)
	c.SetMaxIdleConns(10)

	query := "select customer_id, first_name, last_name, city, zip_code, date_birth, status from customers"

	rows, err := d.client.Query(query)
	if err != nil {
		log.Fatal("No se pueden obtener resultados de la BD => " + err.Error())
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.FirstName, &c.LastName, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
		if err != nil {
			log.Fatal("Error al scanear customers => " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDb {
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

// Database adapter 11:00
