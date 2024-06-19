package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

// BalanceOperationsController represents a controller struct.
type BalanceOperationsController struct {
	baseController
}

// NewBalanceOperationsController creates a new instance of BalanceOperationsController.
// It takes a baseController as a parameter and returns a pointer to the BalanceOperationsController.
func NewBalanceOperationsController(baseController baseController) *BalanceOperationsController {
	balanceOperationsController := BalanceOperationsController{baseController: baseController}
	return &balanceOperationsController
}

// GetBalanceOperations takes context, status, createdSince, createdUntil, recipientId as parameters and
// returns an models.ApiResponse with models.ListBalanceOperationResponse data and
// an error if there was an issue with the request or response.
func (b *BalanceOperationsController) GetBalanceOperations(
	ctx context.Context,
	status *string,
	createdSince *time.Time,
	createdUntil *time.Time,
	recipientId *string) (
	models.ApiResponse[models.ListBalanceOperationResponse],
	error) {
	req := b.prepareRequest(ctx, "GET", "/balance/operations")
	req.Authenticate(true)
	if status != nil {
		req.QueryParam("status", *status)
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}
	if recipientId != nil {
		req.QueryParam("recipient_id", *recipientId)
	}
	var result models.ListBalanceOperationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListBalanceOperationResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBalanceOperationById takes context, id as parameters and
// returns an models.ApiResponse with models.GetBalanceOperationResponse data and
// an error if there was an issue with the request or response.
func (b *BalanceOperationsController) GetBalanceOperationById(
	ctx context.Context,
	id int64) (
	models.ApiResponse[models.GetBalanceOperationResponse],
	error) {
	req := b.prepareRequest(ctx, "GET", fmt.Sprintf("/balance/operations/%v", id))
	req.Authenticate(true)

	var result models.GetBalanceOperationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetBalanceOperationResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
