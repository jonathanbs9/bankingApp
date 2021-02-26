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

func (d AccountRepositoryDb) Findby(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount FROM accounts WHERE account_id = ?"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error al recuperar la info de la cuenta:  " + err.Error())
		return nil, errs.NewUnexpectedError("Error inesperado en la BD (fetchAccount)")
	}
	return &account, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}

// Save Account
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

// Save Transaction = Make entry transaction table  AND update the balance in account table
func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// Start transaction  DB block
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error al iniciar la nueva transacción ... | " + err.Error())
		return nil, errs.NewUnexpectedError("Error inesperado en la BD (TX)")
	}

	// Insert bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?,?,?,?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// Updating account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE  accounts SET amount = amount - ? WHERE account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE  accounts SET amount = amount + ? WHERE account_id = ?`, t.Amount, t.AccountId)
	}

	// In case error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error al guardar la transacción... | " + err.Error())
		return nil, errs.NewUnexpectedError("Error inesperado en la BD (TX-2)")
	}

	// commit the transaction when all is ok
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error al comitear transacciones para la cuenta del banco | " + err.Error())
		return nil, errs.NewUnexpectedError("Error inesperado en la BD (TX-commit)")
	}

	// Getting Last Transaction ID from transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error al obtener el ultimo ID de transacción | " + err.Error())
		return nil, errs.NewUnexpectedError("Error insesperado en la BD (Tx-ID)")
	}

	// Get the latest account information from the accounts table
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// Update the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	sqlGetAccount := "SELECT account_id, customer_id, opening_date, account_type, amount from accounts where account_id = ?"
	var account Account
	err := d.client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}


