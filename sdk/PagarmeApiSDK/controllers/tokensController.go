package controllers

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
)

type TokensController struct {
	baseController
}

func NewTokensController(baseController baseController) *TokensController {
	tokensController := TokensController{baseController: baseController}
	return &tokensController
}

// TODO: type endpoint description here
func (t *TokensController) CreateToken(
	publicKey string,
	request *models.CreateTokenRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetTokenResponse],
	error) {
	req := t.prepareRequest("POST", fmt.Sprintf("/tokens?appId=%s", publicKey))
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetTokenResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetTokenResponse]{Response: resp}, err
	}

	var result models.GetTokenResponse
	result, err = utilities.DecodeResults[models.GetTokenResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetTokenResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets a token from its id
func (t *TokensController) GetToken(
	id string,
	publicKey string) (
	https.ApiResponse[models.GetTokenResponse],
	error) {
	req := t.prepareRequest(
		"GET",
		fmt.Sprintf("/tokens/%s?appId=%s", id, publicKey),
	)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetTokenResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetTokenResponse]{Response: resp}, err
	}

	var result models.GetTokenResponse
	result, err = utilities.DecodeResults[models.GetTokenResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetTokenResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
