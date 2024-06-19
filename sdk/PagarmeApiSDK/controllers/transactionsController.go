package controllers

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
)

type TransactionsController struct {
	baseController
}

func NewTransactionsController(baseController baseController) *TransactionsController {
	transactionsController := TransactionsController{baseController: baseController}
	return &transactionsController
}

// TODO: type endpoint description here
func (t *TransactionsController) GetTransaction(transactionId string) (
	https.ApiResponse[models.GetTransactionResponseInterface],
	error) {
	req := t.prepareRequest("GET", fmt.Sprintf("/transactions/%s", transactionId))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetTransactionResponseInterface]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetTransactionResponseInterface]{Response: resp}, err
	}

	var result models.GetTransactionResponseInterface
	result, err = utilities.DecodeResults[*models.GetTransactionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetTransactionResponseInterface]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
