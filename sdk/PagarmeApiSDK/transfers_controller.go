package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
)

// TransfersController represents a controller struct.
type TransfersController struct {
	baseController
}

// NewTransfersController creates a new instance of TransfersController.
// It takes a baseController as a parameter and returns a pointer to the TransfersController.
func NewTransfersController(baseController baseController) *TransfersController {
	transfersController := TransfersController{baseController: baseController}
	return &transfersController
}

// GetTransferById takes context, transferId as parameters and
// returns an models.ApiResponse with models.GetTransfer data and
// an error if there was an issue with the request or response.
func (t *TransfersController) GetTransferById(
	ctx context.Context,
	transferId string) (
	models.ApiResponse[models.GetTransfer],
	error) {
	req := t.prepareRequest(ctx, "GET", fmt.Sprintf("/transfers/%v", transferId))
	req.Authenticate(true)

	var result models.GetTransfer
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetTransfer](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateTransfer takes context, request as parameters and
// returns an models.ApiResponse with models.GetTransfer data and
// an error if there was an issue with the request or response.
func (t *TransfersController) CreateTransfer(
	ctx context.Context,
	request *models.CreateTransfer) (
	models.ApiResponse[models.GetTransfer],
	error) {
	req := t.prepareRequest(ctx, "POST", "/transfers/recipients")
	req.Authenticate(true)
	req.Json(request)
	var result models.GetTransfer
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetTransfer](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetTransfers takes context as parameters and
// returns an models.ApiResponse with models.ListTransfers data and
// an error if there was an issue with the request or response.
// Gets all transfers
func (t *TransfersController) GetTransfers(ctx context.Context) (
	models.ApiResponse[models.ListTransfers],
	error) {
	req := t.prepareRequest(ctx, "GET", "/transfers")
	req.Authenticate(true)
	var result models.ListTransfers
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListTransfers](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
