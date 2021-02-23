package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	c, err := sql.Open("mysql", "root:@tcp(locahost:3306)/banking")
	if err != nil {
		log.Fatal("Error al conectar a la base de datos => " + err.Error())
	}

	c.SetConnMaxLifetime(time.Second * 5)
	c.SetMaxOpenConns(10)
	c.SetMaxIdleConns(10)

	query := "select customer_id, first_name, last_name, date_birth, city, zipcode, status from customers"

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

func (d CustomerRepositoryDb) GetCustomerById(id string) (*Customer, error) {
	// Hacemos una llamada a la base de datos
	customerSql := "select customer_id, first_name, last_name, city, zipcode, date_birth, status from customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer

	err:= row.Scan(&c.Id, &c.FirstName, &c.LastName, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil{
		log.Println("Error al buscar un cliente => " + err.Error())
		return nil, err
	}
	// Implementar si no encuentra el customer porque no existe en BD (404). No debería ser un error. Sino devolver vacio

	return &c, nil
}
