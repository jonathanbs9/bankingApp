package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/jonathanbs9/bankingApp/errs"
	"github.com/jonathanbs9/bankingApp/logger"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status ) values (?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error al crear nueva cuenta => " + err.Error())

		return nil, errs.NewUnexpectedError("Error inesperado de la Base de datos")
	}

	// obtengo el ultimo id insertado
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error al obtener el último id insertado de la cuenta => " + err.Error())
		return nil, errs.NewUnexpectedError("Error inesperado de la Base de datos")
	}

	// Caso exitoso
	// Paso el int a string
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb (dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}