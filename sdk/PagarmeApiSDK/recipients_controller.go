package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

// RecipientsController represents a controller struct.
type RecipientsController struct {
	baseController
}

// NewRecipientsController creates a new instance of RecipientsController.
// It takes a baseController as a parameter and returns a pointer to the RecipientsController.
func NewRecipientsController(baseController baseController) *RecipientsController {
	recipientsController := RecipientsController{baseController: baseController}
	return &recipientsController
}

// UpdateRecipient takes context, recipientId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetRecipientResponse data and
// an error if there was an issue with the request or response.
// Updates a recipient
func (r *RecipientsController) UpdateRecipient(
	ctx context.Context,
	recipientId string,
	request *models.UpdateRecipientRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "PUT", fmt.Sprintf("/recipients/%v", recipientId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetRecipientResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateAnticipation takes context, recipientId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetAnticipationResponse data and
// an error if there was an issue with the request or response.
// Creates an anticipation
func (r *RecipientsController) CreateAnticipation(
	ctx context.Context,
	recipientId string,
	request *models.CreateAnticipationRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetAnticipationResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/recipients/%v/anticipations", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetAnticipationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetAnticipationResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAnticipationLimits takes context, recipientId, timeframe, paymentDate as parameters and
// returns an models.ApiResponse with models.GetAnticipationLimitResponse data and
// an error if there was an issue with the request or response.
// Gets the anticipation limits for a recipient
func (r *RecipientsController) GetAnticipationLimits(
	ctx context.Context,
	recipientId string,
	timeframe string,
	paymentDate time.Time) (
	models.ApiResponse[models.GetAnticipationLimitResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%v/anticipation_limits", recipientId),
	)
	req.Authenticate(true)
	req.QueryParam("timeframe", timeframe)
	req.QueryParam("payment_date", paymentDate.Format(time.RFC3339))

	var result models.GetAnticipationLimitResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetAnticipationLimitResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetRecipients takes context, page, size as parameters and
// returns an models.ApiResponse with models.ListRecipientResponse data and
// an error if there was an issue with the request or response.
// Retrieves paginated recipients information
func (r *RecipientsController) GetRecipients(
	ctx context.Context,
	page *int,
	size *int) (
	models.ApiResponse[models.ListRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", "/recipients")
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	var result models.ListRecipientResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListRecipientResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetWithdrawById takes context, recipientId, withdrawalId as parameters and
// returns an models.ApiResponse with models.GetWithdrawResponse data and
// an error if there was an issue with the request or response.
func (r *RecipientsController) GetWithdrawById(
	ctx context.Context,
	recipientId string,
	withdrawalId string) (
	models.ApiResponse[models.GetWithdrawResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%v/withdrawals/%v", recipientId, withdrawalId),
	)
	req.Authenticate(true)

	var result models.GetWithdrawResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetWithdrawResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateRecipientDefaultBankAccount takes context, recipientId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetRecipientResponse data and
// an error if there was an issue with the request or response.
// Updates the default bank account from a recipient
func (r *RecipientsController) UpdateRecipientDefaultBankAccount(
	ctx context.Context,
	recipientId string,
	request *models.UpdateRecipientBankAccountRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/recipients/%v/default-bank-account", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetRecipientResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateRecipientMetadata takes context, recipientId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetRecipientResponse data and
// an error if there was an issue with the request or response.
// Updates recipient metadata
func (r *RecipientsController) UpdateRecipientMetadata(
	ctx context.Context,
	recipientId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/recipients/%v/metadata", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetRecipientResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetTransfers takes context, recipientId, page, size, status, createdSince, createdUntil as parameters and
// returns an models.ApiResponse with models.ListTransferResponse data and
// an error if there was an issue with the request or response.
// Gets a paginated list of transfers for the recipient
func (r *RecipientsController) GetTransfers(
	ctx context.Context,
	recipientId string,
	page *int,
	size *int,
	status *string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	models.ApiResponse[models.ListTransferResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%v/transfers", recipientId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}

	var result models.ListTransferResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListTransferResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetTransfer takes context, recipientId, transferId as parameters and
// returns an models.ApiResponse with models.GetTransferResponse data and
// an error if there was an issue with the request or response.
// Gets a transfer
func (r *RecipientsController) GetTransfer(
	ctx context.Context,
	recipientId string,
	transferId string) (
	models.ApiResponse[models.GetTransferResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%v/transfers/%v", recipientId, transferId),
	)
	req.Authenticate(true)

	var result models.GetTransferResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetTransferResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateWithdraw takes context, recipientId, request as parameters and
// returns an models.ApiResponse with models.GetWithdrawResponse data and
// an error if there was an issue with the request or response.
func (r *RecipientsController) CreateWithdraw(
	ctx context.Context,
	recipientId string,
	request *models.CreateWithdrawRequest) (
	models.ApiResponse[models.GetWithdrawResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/recipients/%v/withdrawals", recipientId),
	)
	req.Authenticate(true)
	req.Json(request)

	var result models.GetWithdrawResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetWithdrawResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateAutomaticAnticipationSettings takes context, recipientId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetRecipientResponse data and
// an error if there was an issue with the request or response.
// Updates recipient metadata
func (r *RecipientsController) UpdateAutomaticAnticipationSettings(
	ctx context.Context,
	recipientId string,
	request *models.UpdateAutomaticAnticipationSettingsRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/recipients/%v/automatic-anticipation-settings", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetRecipientResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAnticipation takes context, recipientId, anticipationId as parameters and
// returns an models.ApiResponse with models.GetAnticipationResponse data and
// an error if there was an issue with the request or response.
// Gets an anticipation
func (r *RecipientsController) GetAnticipation(
	ctx context.Context,
	recipientId string,
	anticipationId string) (
	models.ApiResponse[models.GetAnticipationResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%v/anticipations/%v", recipientId, anticipationId),
	)
	req.Authenticate(true)

	var result models.GetAnticipationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetAnticipationResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateRecipientTransferSettings takes context, recipientId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetRecipientResponse data and
// an error if there was an issue with the request or response.
func (r *RecipientsController) UpdateRecipientTransferSettings(
	ctx context.Context,
	recipientId string,
	request *models.UpdateTransferSettingsRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/recipients/%v/transfer-settings", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetRecipientResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAnticipations takes context, recipientId, page, size, status, timeframe, paymentDateSince, paymentDateUntil, createdSince, createdUntil as parameters and
// returns an models.ApiResponse with models.ListAnticipationResponse data and
// an error if there was an issue with the request or response.
// Retrieves a paginated list of anticipations from a recipient
func (r *RecipientsController) GetAnticipations(
	ctx context.Context,
	recipientId string,
	page *int,
	size *int,
	status *string,
	timeframe *string,
	paymentDateSince *time.Time,
	paymentDateUntil *time.Time,
	createdSince *time.Time,
	createdUntil *time.Time) (
	models.ApiResponse[models.ListAnticipationResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%v/anticipations", recipientId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if timeframe != nil {
		req.QueryParam("timeframe", *timeframe)
	}
	if paymentDateSince != nil {
		req.QueryParam("payment_date_since", paymentDateSince.Format(time.RFC3339))
	}
	if paymentDateUntil != nil {
		req.QueryParam("payment_date_until", paymentDateUntil.Format(time.RFC3339))
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}

	var result models.ListAnticipationResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListAnticipationResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetRecipient takes context, recipientId as parameters and
// returns an models.ApiResponse with models.GetRecipientResponse data and
// an error if there was an issue with the request or response.
// Retrieves recipient information
func (r *RecipientsController) GetRecipient(
	ctx context.Context,
	recipientId string) (
	models.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", fmt.Sprintf("/recipients/%v", recipientId))
	req.Authenticate(true)

	var result models.GetRecipientResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetBalance takes context, recipientId as parameters and
// returns an models.ApiResponse with models.GetBalanceResponse data and
// an error if there was an issue with the request or response.
// Get balance information for a recipient
func (r *RecipientsController) GetBalance(
	ctx context.Context,
	recipientId string) (
	models.ApiResponse[models.GetBalanceResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%v/balance", recipientId),
	)
	req.Authenticate(true)

	var result models.GetBalanceResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetBalanceResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetWithdrawals takes context, recipientId, page, size, status, createdSince, createdUntil as parameters and
// returns an models.ApiResponse with models.ListWithdrawals data and
// an error if there was an issue with the request or response.
// Gets a paginated list of transfers for the recipient
func (r *RecipientsController) GetWithdrawals(
	ctx context.Context,
	recipientId string,
	page *int,
	size *int,
	status *string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	models.ApiResponse[models.ListWithdrawals],
	error) {
	req := r.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/recipients/%v/withdrawals", recipientId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}

	var result models.ListWithdrawals
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListWithdrawals](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateTransfer takes context, recipientId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetTransferResponse data and
// an error if there was an issue with the request or response.
// Creates a transfer for a recipient
func (r *RecipientsController) CreateTransfer(
	ctx context.Context,
	recipientId string,
	request *models.CreateTransferRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetTransferResponse],
	error) {
	req := r.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/recipients/%v/transfers", recipientId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetTransferResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetTransferResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateRecipient takes context, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetRecipientResponse data and
// an error if there was an issue with the request or response.
// Creates a new recipient
func (r *RecipientsController) CreateRecipient(
	ctx context.Context,
	request *models.CreateRecipientRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "POST", "/recipients")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)
	var result models.GetRecipientResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetRecipientByCode takes context, code as parameters and
// returns an models.ApiResponse with models.GetRecipientResponse data and
// an error if there was an issue with the request or response.
// Retrieves recipient information
func (r *RecipientsController) GetRecipientByCode(
	ctx context.Context,
	code string) (
	models.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", fmt.Sprintf("/recipients/%v", code))
	req.Authenticate(true)

	var result models.GetRecipientResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetDefaultRecipient takes context as parameters and
// returns an models.ApiResponse with models.GetRecipientResponse data and
// an error if there was an issue with the request or response.
func (r *RecipientsController) GetDefaultRecipient(ctx context.Context) (
	models.ApiResponse[models.GetRecipientResponse],
	error) {
	req := r.prepareRequest(ctx, "GET", "/recipients/default")
	req.Authenticate(true)
	var result models.GetRecipientResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetRecipientResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
