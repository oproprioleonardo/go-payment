package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
)

// CustomersController represents a controller struct.
type CustomersController struct {
	baseController
}

// NewCustomersController creates a new instance of CustomersController.
// It takes a baseController as a parameter and returns a pointer to the CustomersController.
func NewCustomersController(baseController baseController) *CustomersController {
	customersController := CustomersController{baseController: baseController}
	return &customersController
}

// UpdateCard takes context, customerId, cardId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetCardResponse data and
// an error if there was an issue with the request or response.
// Updates a card
func (c *CustomersController) UpdateCard(
	ctx context.Context,
	customerId string,
	cardId string,
	request *models.UpdateCardRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetCardResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/customers/%v/cards/%v", customerId, cardId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetCardResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetCardResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateAddress takes context, customerId, addressId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetAddressResponse data and
// an error if there was an issue with the request or response.
// Updates an address
func (c *CustomersController) UpdateAddress(
	ctx context.Context,
	customerId string,
	addressId string,
	request *models.UpdateAddressRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetAddressResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/customers/%v/addresses/%v", customerId, addressId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetAddressResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetAddressResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteAccessToken takes context, customerId, tokenId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetAccessTokenResponse data and
// an error if there was an issue with the request or response.
// Delete a customer's access token
func (c *CustomersController) DeleteAccessToken(
	ctx context.Context,
	customerId string,
	tokenId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetAccessTokenResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/customers/%v/access-tokens/%v", customerId, tokenId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetAccessTokenResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetAccessTokenResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateCustomer takes context, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetCustomerResponse data and
// an error if there was an issue with the request or response.
// Creates a new customer
func (c *CustomersController) CreateCustomer(
	ctx context.Context,
	request *models.CreateCustomerRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetCustomerResponse],
	error) {
	req := c.prepareRequest(ctx, "POST", "/customers")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)
	var result models.GetCustomerResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetCustomerResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateAddress takes context, customerId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetAddressResponse data and
// an error if there was an issue with the request or response.
// Creates a new address for a customer
func (c *CustomersController) CreateAddress(
	ctx context.Context,
	customerId string,
	request *models.CreateAddressRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetAddressResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/customers/%v/addresses", customerId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetAddressResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetAddressResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteAccessTokens takes context, customerId as parameters and
// returns an models.ApiResponse with models.ListAccessTokensResponse data and
// an error if there was an issue with the request or response.
// Delete a Customer's access tokens
func (c *CustomersController) DeleteAccessTokens(
	ctx context.Context,
	customerId string) (
	models.ApiResponse[models.ListAccessTokensResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%v/access-tokens/", customerId),
	)
	req.Authenticate(true)

	var result models.ListAccessTokensResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListAccessTokensResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAddress takes context, customerId, addressId as parameters and
// returns an models.ApiResponse with models.GetAddressResponse data and
// an error if there was an issue with the request or response.
// Get a customer's address
func (c *CustomersController) GetAddress(
	ctx context.Context,
	customerId string,
	addressId string) (
	models.ApiResponse[models.GetAddressResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%v/addresses/%v", customerId, addressId),
	)
	req.Authenticate(true)

	var result models.GetAddressResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetAddressResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteAddress takes context, customerId, addressId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetAddressResponse data and
// an error if there was an issue with the request or response.
// Delete a Customer's address
func (c *CustomersController) DeleteAddress(
	ctx context.Context,
	customerId string,
	addressId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetAddressResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/customers/%v/addresses/%v", customerId, addressId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetAddressResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetAddressResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateCard takes context, customerId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetCardResponse data and
// an error if there was an issue with the request or response.
// Creates a new card for a customer
func (c *CustomersController) CreateCard(
	ctx context.Context,
	customerId string,
	request *models.CreateCardRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetCardResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/customers/%v/cards", customerId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetCardResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetCardResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetCustomers takes context, name, document, page, size, email, code as parameters and
// returns an models.ApiResponse with models.ListCustomersResponse data and
// an error if there was an issue with the request or response.
// Get all Customers
func (c *CustomersController) GetCustomers(
	ctx context.Context,
	name *string,
	document *string,
	page *int,
	size *int,
	email *string,
	code *string) (
	models.ApiResponse[models.ListCustomersResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", "/customers")
	req.Authenticate(true)
	if name != nil {
		req.QueryParam("name", *name)
	}
	if document != nil {
		req.QueryParam("document", *document)
	}
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if email != nil {
		req.QueryParam("email", *email)
	}
	if code != nil {
		req.QueryParam("Code", *code)
	}
	var result models.ListCustomersResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListCustomersResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateCustomer takes context, customerId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetCustomerResponse data and
// an error if there was an issue with the request or response.
// Updates a customer
func (c *CustomersController) UpdateCustomer(
	ctx context.Context,
	customerId string,
	request *models.UpdateCustomerRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetCustomerResponse],
	error) {
	req := c.prepareRequest(ctx, "PUT", fmt.Sprintf("/customers/%v", customerId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetCustomerResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetCustomerResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateAccessToken takes context, customerId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetAccessTokenResponse data and
// an error if there was an issue with the request or response.
// Creates a access token for a customer
func (c *CustomersController) CreateAccessToken(
	ctx context.Context,
	customerId string,
	request *models.CreateAccessTokenRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetAccessTokenResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/customers/%v/access-tokens", customerId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetAccessTokenResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetAccessTokenResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAccessTokens takes context, customerId, page, size as parameters and
// returns an models.ApiResponse with models.ListAccessTokensResponse data and
// an error if there was an issue with the request or response.
// Get all access tokens from a customer
func (c *CustomersController) GetAccessTokens(
	ctx context.Context,
	customerId string,
	page *int,
	size *int) (
	models.ApiResponse[models.ListAccessTokensResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%v/access-tokens", customerId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}

	var result models.ListAccessTokensResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListAccessTokensResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetCards takes context, customerId, page, size as parameters and
// returns an models.ApiResponse with models.ListCardsResponse data and
// an error if there was an issue with the request or response.
// Get all cards from a customer
func (c *CustomersController) GetCards(
	ctx context.Context,
	customerId string,
	page *int,
	size *int) (
	models.ApiResponse[models.ListCardsResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%v/cards", customerId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}

	var result models.ListCardsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListCardsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// RenewCard takes context, customerId, cardId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetCardResponse data and
// an error if there was an issue with the request or response.
// Renew a card
func (c *CustomersController) RenewCard(
	ctx context.Context,
	customerId string,
	cardId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetCardResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/customers/%v/cards/%v/renew", customerId, cardId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetCardResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetCardResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAccessToken takes context, customerId, tokenId as parameters and
// returns an models.ApiResponse with models.GetAccessTokenResponse data and
// an error if there was an issue with the request or response.
// Get a Customer's access token
func (c *CustomersController) GetAccessToken(
	ctx context.Context,
	customerId string,
	tokenId string) (
	models.ApiResponse[models.GetAccessTokenResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%v/access-tokens/%v", customerId, tokenId),
	)
	req.Authenticate(true)

	var result models.GetAccessTokenResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetAccessTokenResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateCustomerMetadata takes context, customerId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetCustomerResponse data and
// an error if there was an issue with the request or response.
// Updates the metadata a customer
func (c *CustomersController) UpdateCustomerMetadata(
	ctx context.Context,
	customerId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetCustomerResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/Customers/%v/metadata", customerId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetCustomerResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetCustomerResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteCard takes context, customerId, cardId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetCardResponse data and
// an error if there was an issue with the request or response.
// Delete a customer's card
func (c *CustomersController) DeleteCard(
	ctx context.Context,
	customerId string,
	cardId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetCardResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/customers/%v/cards/%v", customerId, cardId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetCardResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetCardResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetAddresses takes context, customerId, page, size as parameters and
// returns an models.ApiResponse with models.ListAddressesResponse data and
// an error if there was an issue with the request or response.
// Gets all adressess from a customer
func (c *CustomersController) GetAddresses(
	ctx context.Context,
	customerId string,
	page *int,
	size *int) (
	models.ApiResponse[models.ListAddressesResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%v/addresses", customerId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}

	var result models.ListAddressesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListAddressesResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetCustomer takes context, customerId as parameters and
// returns an models.ApiResponse with models.GetCustomerResponse data and
// an error if there was an issue with the request or response.
// Get a customer
func (c *CustomersController) GetCustomer(
	ctx context.Context,
	customerId string) (
	models.ApiResponse[models.GetCustomerResponse],
	error) {
	req := c.prepareRequest(ctx, "GET", fmt.Sprintf("/customers/%v", customerId))
	req.Authenticate(true)

	var result models.GetCustomerResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetCustomerResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetCard takes context, customerId, cardId as parameters and
// returns an models.ApiResponse with models.GetCardResponse data and
// an error if there was an issue with the request or response.
// Get a customer's card
func (c *CustomersController) GetCard(
	ctx context.Context,
	customerId string,
	cardId string) (
	models.ApiResponse[models.GetCardResponse],
	error) {
	req := c.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/customers/%v/cards/%v", customerId, cardId),
	)
	req.Authenticate(true)

	var result models.GetCardResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetCardResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
