package controllers

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

type InvoicesController struct {
	baseController
}

func NewInvoicesController(baseController baseController) *InvoicesController {
	invoicesController := InvoicesController{baseController: baseController}
	return &invoicesController
}

// Updates the metadata from an invoice
func (i *InvoicesController) UpdateInvoiceMetadata(
	invoiceId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest(
		"PATCH",
		fmt.Sprintf("/invoices/%s/metadata", invoiceId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	var result models.GetInvoiceResponse
	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (i *InvoicesController) GetPartialInvoice(subscriptionId string) (
	https.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/partial-invoice", subscriptionId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	var result models.GetInvoiceResponse
	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Cancels an invoice
func (i *InvoicesController) CancelInvoice(
	invoiceId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest("DELETE", fmt.Sprintf("/invoices/%s", invoiceId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	var result models.GetInvoiceResponse
	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Create an Invoice
func (i *InvoicesController) CreateInvoice(
	subscriptionId string,
	cycleId string,
	request *models.CreateInvoiceRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest(
		"POST",
		fmt.Sprintf("/subscriptions/%s/cycles/%s/pay", subscriptionId, cycleId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	if request != nil {
		req.Json(*request)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	var result models.GetInvoiceResponse
	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets all invoices
func (i *InvoicesController) GetInvoices(
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
	https.ApiResponse[models.ListInvoicesResponse],
	error) {
	req := i.prepareRequest("GET", "/invoices")
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
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListInvoicesResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListInvoicesResponse]{Response: resp}, err
	}

	var result models.ListInvoicesResponse
	result, err = utilities.DecodeResults[models.ListInvoicesResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListInvoicesResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets an invoice
func (i *InvoicesController) GetInvoice(invoiceId string) (
	https.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest("GET", fmt.Sprintf("/invoices/%s", invoiceId))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	var result models.GetInvoiceResponse
	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the status from an invoice
func (i *InvoicesController) UpdateInvoiceStatus(
	invoiceId string,
	request *models.UpdateInvoiceStatusRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetInvoiceResponse],
	error) {
	req := i.prepareRequest("PATCH", fmt.Sprintf("/invoices/%s/status", invoiceId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	var result models.GetInvoiceResponse
	result, err = utilities.DecodeResults[models.GetInvoiceResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetInvoiceResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
