package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jonathanbs9/bankingApp/dto"
	"github.com/jonathanbs9/bankingApp/logger"
	"github.com/jonathanbs9/bankingApp/service"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	logger.Info("AccountHanlders => customerid " + customerId)

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, errApp := h.service.NewAccount(request)
		if errApp != nil {
			writeResponse(w, errApp.Code, errApp.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandler) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	// Get account_id and customer_id from the url
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	// Decode incoming request
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		// Build the request object
		request.AccountId = accountId
		request.CustomerId = customerId

		// Make the transaction
		account, appError := h.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
