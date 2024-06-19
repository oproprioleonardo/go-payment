package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
)

// TokensController represents a controller struct.
type TokensController struct {
	baseController
}

// NewTokensController creates a new instance of TokensController.
// It takes a baseController as a parameter and returns a pointer to the TokensController.
func NewTokensController(baseController baseController) *TokensController {
	tokensController := TokensController{baseController: baseController}
	return &tokensController
}

// CreateToken takes context, publicKey, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetTokenResponse data and
// an error if there was an issue with the request or response.
func (t *TokensController) CreateToken(
	ctx context.Context,
	publicKey string,
	request *models.CreateTokenRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetTokenResponse],
	error) {
	req := t.prepareRequest(ctx, "POST", fmt.Sprintf("/tokens?appId=%v", publicKey))
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetTokenResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetTokenResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetToken takes context, id, publicKey as parameters and
// returns an models.ApiResponse with models.GetTokenResponse data and
// an error if there was an issue with the request or response.
// Gets a token from its id
func (t *TokensController) GetToken(
	ctx context.Context,
	id string,
	publicKey string) (
	models.ApiResponse[models.GetTokenResponse],
	error) {
	req := t.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/tokens/%v?appId=%v", id, publicKey),
	)

	var result models.GetTokenResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetTokenResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
