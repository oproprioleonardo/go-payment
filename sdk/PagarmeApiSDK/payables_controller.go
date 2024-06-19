package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

// PayablesController represents a controller struct.
type PayablesController struct {
	baseController
}

// NewPayablesController creates a new instance of PayablesController.
// It takes a baseController as a parameter and returns a pointer to the PayablesController.
func NewPayablesController(baseController baseController) *PayablesController {
	payablesController := PayablesController{baseController: baseController}
	return &payablesController
}

// GetPayables takes context, mType, splitId, bulkAnticipationId, installment, status, recipientId, amount, chargeId, paymentDateUntil, paymentDateSince, updatedUntil, updatedSince, createdUntil, createdSince, liquidationArrangementId, page, size, gatewayId as parameters and
// returns an models.ApiResponse with models.ListPayablesResponse data and
// an error if there was an issue with the request or response.
func (p *PayablesController) GetPayables(
	ctx context.Context,
	mType *string,
	splitId *string,
	bulkAnticipationId *string,
	installment *int,
	status *string,
	recipientId *string,
	amount *int,
	chargeId *string,
	paymentDateUntil *string,
	paymentDateSince *time.Time,
	updatedUntil *time.Time,
	updatedSince *time.Time,
	createdUntil *time.Time,
	createdSince *time.Time,
	liquidationArrangementId *string,
	page *int,
	size *int,
	gatewayId *int64) (
	models.ApiResponse[models.ListPayablesResponse],
	error) {
	req := p.prepareRequest(ctx, "GET", "/payables")
	req.Authenticate(true)
	if mType != nil {
		req.QueryParam("type", *mType)
	}
	if splitId != nil {
		req.QueryParam("split_id", *splitId)
	}
	if bulkAnticipationId != nil {
		req.QueryParam("bulk_anticipation_id", *bulkAnticipationId)
	}
	if installment != nil {
		req.QueryParam("installment", *installment)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if recipientId != nil {
		req.QueryParam("recipient_id", *recipientId)
	}
	if amount != nil {
		req.QueryParam("amount", *amount)
	}
	if chargeId != nil {
		req.QueryParam("charge_id", *chargeId)
	}
	if paymentDateUntil != nil {
		req.QueryParam("payment_date_until", *paymentDateUntil)
	}
	if paymentDateSince != nil {
		req.QueryParam("payment_date_since", paymentDateSince.Format(time.RFC3339))
	}
	if updatedUntil != nil {
		req.QueryParam("updated_until", updatedUntil.Format(time.RFC3339))
	}
	if updatedSince != nil {
		req.QueryParam("updated_since", updatedSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if liquidationArrangementId != nil {
		req.QueryParam("liquidation_arrangement_id", *liquidationArrangementId)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if gatewayId != nil {
		req.QueryParam("gateway_id", *gatewayId)
	}
	var result models.ListPayablesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListPayablesResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetPayableById takes context, id as parameters and
// returns an models.ApiResponse with models.GetPayableResponse data and
// an error if there was an issue with the request or response.
func (p *PayablesController) GetPayableById(
	ctx context.Context,
	id int64) (
	models.ApiResponse[models.GetPayableResponse],
	error) {
	req := p.prepareRequest(ctx, "GET", fmt.Sprintf("/payables/%v", id))
	req.Authenticate(true)

	var result models.GetPayableResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPayableResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
