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
	vars :=mux.Vars(r)
	customerId := vars["customer_id"]
	logger.Info("AccountHanlders => customerid "+customerId)

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




