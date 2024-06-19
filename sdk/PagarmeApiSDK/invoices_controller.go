package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

// InvoicesController represents a controller struct.
type InvoicesController struct {
	baseController
}

// NewInvoicesController creates a new instance of InvoicesController.
// It takes a baseController as a parameter and returns a pointer to the InvoicesController.
func NewInvoicesController(baseController baseController) *InvoicesController {
	invoicesController := InvoicesController{baseController: baseController}
	return &invoicesController
}

// UpdateInvoiceMetadata takes context, invoiceId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetInvoiceResponse data and
// an error if there was an issue with the request or response.
// Updates the metadata from an invoice
func (i *InvoicesController) UpdateInvoiceMetadata(
	ctx context.Context,
	invoiceId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/invoices/%v/metadata", invoiceId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetInvoiceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetPartialInvoice takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.GetInvoiceResponse data and
// an error if there was an issue with the request or response.
func (i *InvoicesController) GetPartialInvoice(
	ctx context.Context,
	subscriptionId string) (
	models.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/partial-invoice", subscriptionId),
	)
	req.Authenticate(true)

	var result models.GetInvoiceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CancelInvoice takes context, invoiceId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetInvoiceResponse data and
// an error if there was an issue with the request or response.
// Cancels an invoice
func (i *InvoicesController) CancelInvoice(
	ctx context.Context,
	invoiceId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest(ctx, "DELETE", fmt.Sprintf("/invoices/%v", invoiceId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetInvoiceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateInvoice takes context, subscriptionId, cycleId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetInvoiceResponse data and
// an error if there was an issue with the request or response.
// Create an Invoice
func (i *InvoicesController) CreateInvoice(
	ctx context.Context,
	subscriptionId string,
	cycleId string,
	request *models.CreateInvoiceRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/cycles/%v/pay", subscriptionId, cycleId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	if request != nil {
		req.Json(*request)
	}

	var result models.GetInvoiceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetInvoices takes context, page, size, code, customerId, subscriptionId, createdSince, createdUntil, status, dueSince, dueUntil, customerDocument as parameters and
// returns an models.ApiResponse with models.ListInvoicesResponse data and
// an error if there was an issue with the request or response.
// Gets all invoices
func (i *InvoicesController) GetInvoices(
	ctx context.Context,
	page *int,
	size *int,
	code *string,
	customerId *string,
	subscriptionId *string,
	createdSince *time.Time,
	createdUntil *time.Time,
	status *string,
	dueSince *time.Time,
	dueUntil *time.Time,
	customerDocument *string) (
	models.ApiResponse[models.ListInvoicesResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", "/invoices")
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
	if customerId != nil {
		req.QueryParam("customer_id", *customerId)
	}
	if subscriptionId != nil {
		req.QueryParam("subscription_id", *subscriptionId)
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if dueSince != nil {
		req.QueryParam("due_since", dueSince.Format(time.RFC3339))
	}
	if dueUntil != nil {
		req.QueryParam("due_until", dueUntil.Format(time.RFC3339))
	}
	if customerDocument != nil {
		req.QueryParam("customer_document", *customerDocument)
	}
	var result models.ListInvoicesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListInvoicesResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetInvoice takes context, invoiceId as parameters and
// returns an models.ApiResponse with models.GetInvoiceResponse data and
// an error if there was an issue with the request or response.
// Gets an invoice
func (i *InvoicesController) GetInvoice(
	ctx context.Context,
	invoiceId string) (
	models.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest(ctx, "GET", fmt.Sprintf("/invoices/%v", invoiceId))
	req.Authenticate(true)

	var result models.GetInvoiceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateInvoiceStatus takes context, invoiceId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetInvoiceResponse data and
// an error if there was an issue with the request or response.
// Updates the status from an invoice
func (i *InvoicesController) UpdateInvoiceStatus(
	ctx context.Context,
	invoiceId string,
	request *models.UpdateInvoiceStatusRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/invoices/%v/status", invoiceId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetInvoiceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
