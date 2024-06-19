package controllers

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

type PayablesController struct {
	baseController
}

func NewPayablesController(baseController baseController) *PayablesController {
	payablesController := PayablesController{baseController: baseController}
	return &payablesController
}

// TODO: type endpoint description here
func (p *PayablesController) GetPayables(
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
	https.ApiResponse[models.ListPayablesResponse],
	error) {
	req := p.prepareRequest("GET", "/payables")
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
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListPayablesResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListPayablesResponse]{Response: resp}, err
	}

	var result models.ListPayablesResponse
	result, err = utilities.DecodeResults[models.ListPayablesResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListPayablesResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (p *PayablesController) GetPayableById(id int64) (
	https.ApiResponse[models.GetPayableResponse],
	error) {
	req := p.prepareRequest("GET", fmt.Sprintf("/payables/%s", id))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPayableResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPayableResponse]{Response: resp}, err
	}

	var result models.GetPayableResponse
	result, err = utilities.DecodeResults[models.GetPayableResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPayableResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
