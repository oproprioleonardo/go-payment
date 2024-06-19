package controllers

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

type BalanceOperationsController struct {
	baseController
}

func NewBalanceOperationsController(baseController baseController) *BalanceOperationsController {
	balanceOperationsController := BalanceOperationsController{baseController: baseController}
	return &balanceOperationsController
}

// TODO: type endpoint description here
func (b *BalanceOperationsController) GetBalanceOperations(
	status *string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	https.ApiResponse[models.ListBalanceOperationResponse],
	error) {
	req := b.prepareRequest("GET", "/balance/operations")
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
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListBalanceOperationResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListBalanceOperationResponse]{Response: resp}, err
	}

	var result models.ListBalanceOperationResponse
	result, err = utilities.DecodeResults[models.ListBalanceOperationResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListBalanceOperationResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (b *BalanceOperationsController) GetBalanceOperationById(id int64) (
	https.ApiResponse[models.GetBalanceOperationResponse],
	error) {
	req := b.prepareRequest("GET", fmt.Sprintf("/balance/operations/%s", id))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetBalanceOperationResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetBalanceOperationResponse]{Response: resp}, err
	}

	var result models.GetBalanceOperationResponse
	result, err = utilities.DecodeResults[models.GetBalanceOperationResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetBalanceOperationResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
