package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

// ChargesController represents a controller struct.
type ChargesController struct {
	baseController
}

// NewChargesController creates a new instance of ChargesController.
// It takes a baseController as a parameter and returns a pointer to the ChargesController.
func NewChargesController(baseController baseController) *ChargesController {
	chargesController := ChargesController{baseController: baseController}
	return &chargesController
}

// UpdateChargeMetadata takes context, chargeId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetChargeResponse data and
// an error if there was an issue with the request or response.
// Updates the metadata from a charge
func (c *ChargesController) UpdateChargeMetadata(
	ctx context.Context,
	chargeId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/Charges/%v/metadata", chargeId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetChargeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateChargePaymentMethod takes context, chargeId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetChargeResponse data and
// an error if there was an issue with the request or response.
// Updates a charge's payment method
func (c *ChargesController) UpdateChargePaymentMethod(
	ctx context.Context,
	chargeId string,
	request *models.UpdateChargePaymentMethodRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/charges/%v/payment-method", chargeId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetChargeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetChargeTransactions takes context, chargeId, page, size as parameters and
// returns an models.ApiResponse with models.ListChargeTransactionsResponse data and
// an error if there was an issue with the request or response.
func (c *ChargesController) GetChargeTransactions(
	ctx context.Context,
	chargeId string,
	page *int,
	size *int) (
	models.ApiResponse[models.ListChargeTransactionsResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/charges/%v/transactions", chargeId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}

	var result models.ListChargeTransactionsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListChargeTransactionsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateChargeDueDate takes context, chargeId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetChargeResponse data and
// an error if there was an issue with the request or response.
// Updates the due date from a charge
func (c *ChargesController) UpdateChargeDueDate(
	ctx context.Context,
	chargeId string,
	request *models.UpdateChargeDueDateRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/Charges/%v/due-date", chargeId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetChargeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetCharges takes context, page, size, code, status, paymentMethod, customerId, orderId, createdSince, createdUntil as parameters and
// returns an models.ApiResponse with models.ListChargesResponse data and
// an error if there was an issue with the request or response.
// Lists all charges
func (c *ChargesController) GetCharges(
	ctx context.Context,
	page *int,
	size *int,
	code *string,
	status *string,
	paymentMethod *string,
	customerId *string,
	orderId *string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	models.ApiResponse[models.ListChargesResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", "/charges")
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if code != nil {
		req.QueryParam("code", *code)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if paymentMethod != nil {
		req.QueryParam("payment_method", *paymentMethod)
	}
	if customerId != nil {
		req.QueryParam("customer_id", *customerId)
	}
	if orderId != nil {
		req.QueryParam("order_id", *orderId)
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}
	var result models.ListChargesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListChargesResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CaptureCharge takes context, chargeId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetChargeResponse data and
// an error if there was an issue with the request or response.
// Captures a charge
func (c *ChargesController) CaptureCharge(
	ctx context.Context,
	chargeId string,
	request *models.CreateCaptureChargeRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/charges/%v/capture", chargeId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	if request != nil {
		req.Json(*request)
	}

	var result models.GetChargeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateChargeCard takes context, chargeId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetChargeResponse data and
// an error if there was an issue with the request or response.
// Updates the card from a charge
func (c *ChargesController) UpdateChargeCard(
	ctx context.Context,
	chargeId string,
	request *models.UpdateChargeCardRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "PATCH", fmt.Sprintf("/charges/%v/card", chargeId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetChargeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetCharge takes context, chargeId as parameters and
// returns an models.ApiResponse with models.GetChargeResponse data and
// an error if there was an issue with the request or response.
// Get a charge from its id
func (c *ChargesController) GetCharge(
	ctx context.Context,
	chargeId string) (
	models.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", fmt.Sprintf("/charges/%v", chargeId))
	req.Authenticate(true)

	var result models.GetChargeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetChargesSummary takes context, status, createdSince, createdUntil as parameters and
// returns an models.ApiResponse with models.GetChargesSummaryResponse data and
// an error if there was an issue with the request or response.
func (c *ChargesController) GetChargesSummary(
	ctx context.Context,
	status string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	models.ApiResponse[models.GetChargesSummaryResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", "/charges/summary")
	req.Authenticate(true)
	req.QueryParam("status", status)
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}
	var result models.GetChargesSummaryResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargesSummaryResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// RetryCharge takes context, chargeId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetChargeResponse data and
// an error if there was an issue with the request or response.
// Retries a charge
func (c *ChargesController) RetryCharge(
	ctx context.Context,
	chargeId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "POST", fmt.Sprintf("/charges/%v/retry", chargeId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetChargeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CancelCharge takes context, chargeId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetChargeResponse data and
// an error if there was an issue with the request or response.
// Cancel a charge
func (c *ChargesController) CancelCharge(
	ctx context.Context,
	chargeId string,
	request *models.CreateCancelChargeRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "DELETE", fmt.Sprintf("/charges/%v", chargeId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	if request != nil {
		req.Json(*request)
	}

	var result models.GetChargeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateCharge takes context, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetChargeResponse data and
// an error if there was an issue with the request or response.
// Creates a new charge
func (c *ChargesController) CreateCharge(
	ctx context.Context,
	request *models.CreateChargeRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(ctx, "POST", "/Charges")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)
	var result models.GetChargeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// ConfirmPayment takes context, chargeId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetChargeResponse data and
// an error if there was an issue with the request or response.
func (c *ChargesController) ConfirmPayment(
	ctx context.Context,
	chargeId string,
	request *models.CreateConfirmPaymentRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetChargeResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/charges/%v/confirm-payment", chargeId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	if request != nil {
		req.Json(*request)
	}

	var result models.GetChargeResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetChargeResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
