package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
)

// TransactionsController represents a controller struct.
type TransactionsController struct {
	baseController
}

// NewTransactionsController creates a new instance of TransactionsController.
// It takes a baseController as a parameter and returns a pointer to the TransactionsController.
func NewTransactionsController(baseController baseController) *TransactionsController {
	transactionsController := TransactionsController{baseController: baseController}
	return &transactionsController
}

// GetTransaction takes context, transactionId as parameters and
// returns an models.ApiResponse with models.GetTransactionResponseInterface data and
// an error if there was an issue with the request or response.
func (t *TransactionsController) GetTransaction(
	ctx context.Context,
	transactionId string) (
	models.ApiResponse[models.GetTransactionResponseInterface],
	error) {
	req := t.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/transactions/%v", transactionId),
	)
	req.Authenticate(true)

	var result models.GetTransactionResponseInterface
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[*models.GetTransactionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
