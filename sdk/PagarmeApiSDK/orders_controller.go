package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

// OrdersController represents a controller struct.
type OrdersController struct {
	baseController
}

// NewOrdersController creates a new instance of OrdersController.
// It takes a baseController as a parameter and returns a pointer to the OrdersController.
func NewOrdersController(baseController baseController) *OrdersController {
	ordersController := OrdersController{baseController: baseController}
	return &ordersController
}

// GetOrders takes context, page, size, code, status, createdSince, createdUntil, customerId as parameters and
// returns an models.ApiResponse with models.ListOrderResponse data and
// an error if there was an issue with the request or response.
// Gets all orders
func (o *OrdersController) GetOrders(
	ctx context.Context,
	page *int,
	size *int,
	code *string,
	status *string,
	createdSince *time.Time,
	createdUntil *time.Time,
	customerId *string) (
	models.ApiResponse[models.ListOrderResponse],
	error) {
	req := o.prepareRequest(ctx, "GET", "/orders")
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
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}
	if customerId != nil {
		req.QueryParam("customer_id", *customerId)
	}
	var result models.ListOrderResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListOrderResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateOrderItem takes context, orderId, itemId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetOrderItemResponse data and
// an error if there was an issue with the request or response.
func (o *OrdersController) UpdateOrderItem(
	ctx context.Context,
	orderId string,
	itemId string,
	request *models.UpdateOrderItemRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetOrderItemResponse],
	error) {
	req := o.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/orders/%v/items/%v", orderId, itemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetOrderItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetOrderItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteAllOrderItems takes context, orderId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetOrderResponse data and
// an error if there was an issue with the request or response.
func (o *OrdersController) DeleteAllOrderItems(
	ctx context.Context,
	orderId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetOrderResponse],
	error) {
	req := o.prepareRequest(ctx, "DELETE", fmt.Sprintf("/orders/%v/items", orderId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetOrderResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetOrderResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteOrderItem takes context, orderId, itemId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetOrderItemResponse data and
// an error if there was an issue with the request or response.
func (o *OrdersController) DeleteOrderItem(
	ctx context.Context,
	orderId string,
	itemId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetOrderItemResponse],
	error) {
	req := o.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/orders/%v/items/%v", orderId, itemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetOrderItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetOrderItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CloseOrder takes context, id, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetOrderResponse data and
// an error if there was an issue with the request or response.
func (o *OrdersController) CloseOrder(
	ctx context.Context,
	id string,
	request *models.UpdateOrderStatusRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetOrderResponse],
	error) {
	req := o.prepareRequest(ctx, "PATCH", fmt.Sprintf("/orders/%v/closed", id))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetOrderResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetOrderResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateOrder takes context, body, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetOrderResponse data and
// an error if there was an issue with the request or response.
// Creates a new Order
func (o *OrdersController) CreateOrder(
	ctx context.Context,
	body *models.CreateOrderRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetOrderResponse],
	error) {
	req := o.prepareRequest(ctx, "POST", "/orders")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)
	var result models.GetOrderResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetOrderResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateOrderItem takes context, orderId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetOrderItemResponse data and
// an error if there was an issue with the request or response.
func (o *OrdersController) CreateOrderItem(
	ctx context.Context,
	orderId string,
	request *models.CreateOrderItemRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetOrderItemResponse],
	error) {
	req := o.prepareRequest(ctx, "POST", fmt.Sprintf("/orders/%v/items", orderId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetOrderItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetOrderItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetOrderItem takes context, orderId, itemId as parameters and
// returns an models.ApiResponse with models.GetOrderItemResponse data and
// an error if there was an issue with the request or response.
func (o *OrdersController) GetOrderItem(
	ctx context.Context,
	orderId string,
	itemId string) (
	models.ApiResponse[models.GetOrderItemResponse],
	error) {
	req := o.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/orders/%v/items/%v", orderId, itemId),
	)
	req.Authenticate(true)

	var result models.GetOrderItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetOrderItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateOrderMetadata takes context, orderId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetOrderResponse data and
// an error if there was an issue with the request or response.
// Updates the metadata from an order
func (o *OrdersController) UpdateOrderMetadata(
	ctx context.Context,
	orderId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetOrderResponse],
	error) {
	req := o.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/Orders/%v/metadata", orderId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetOrderResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetOrderResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetOrder takes context, orderId as parameters and
// returns an models.ApiResponse with models.GetOrderResponse data and
// an error if there was an issue with the request or response.
// Gets an order
func (o *OrdersController) GetOrder(
	ctx context.Context,
	orderId string) (
	models.ApiResponse[models.GetOrderResponse],
	error) {
	req := o.prepareRequest(ctx, "GET", fmt.Sprintf("/orders/%v", orderId))
	req.Authenticate(true)

	var result models.GetOrderResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetOrderResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
