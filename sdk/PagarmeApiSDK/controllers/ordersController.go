package controllers

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

type OrdersController struct {
	baseController
}

func NewOrdersController(baseController baseController) *OrdersController {
	ordersController := OrdersController{baseController: baseController}
	return &ordersController
}

// Gets all orders
func (o *OrdersController) GetOrders(
	page *int,
	size *int,
	code *string,
	status *string,
	createdSince *time.Time,
	createdUntil *time.Time,
	customerId *string) (
	https.ApiResponse[models.ListOrderResponse],
	error) {
	req := o.prepareRequest("GET", "/orders")
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
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListOrderResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListOrderResponse]{Response: resp}, err
	}

	var result models.ListOrderResponse
	result, err = utilities.DecodeResults[models.ListOrderResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListOrderResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) UpdateOrderItem(
	orderId string,
	itemId string,
	request *models.UpdateOrderItemRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetOrderItemResponse],
	error) {
	req := o.prepareRequest(
		"PUT",
		fmt.Sprintf("/orders/%s/items/%s", orderId, itemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}

	var result models.GetOrderItemResponse
	result, err = utilities.DecodeResults[models.GetOrderItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) DeleteAllOrderItems(
	orderId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetOrderResponse],
	error) {
	req := o.prepareRequest("DELETE", fmt.Sprintf("/orders/%s/items", orderId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}

	var result models.GetOrderResponse
	result, err = utilities.DecodeResults[models.GetOrderResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) DeleteOrderItem(
	orderId string,
	itemId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetOrderItemResponse],
	error) {
	req := o.prepareRequest(
		"DELETE",
		fmt.Sprintf("/orders/%s/items/%s", orderId, itemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}

	var result models.GetOrderItemResponse
	result, err = utilities.DecodeResults[models.GetOrderItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) CloseOrder(
	id string,
	request *models.UpdateOrderStatusRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetOrderResponse],
	error) {
	req := o.prepareRequest("PATCH", fmt.Sprintf("/orders/%s/closed", id))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}

	var result models.GetOrderResponse
	result, err = utilities.DecodeResults[models.GetOrderResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a new Order
func (o *OrdersController) CreateOrder(
	body *models.CreateOrderRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetOrderResponse],
	error) {
	req := o.prepareRequest("POST", "/orders")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}

	var result models.GetOrderResponse
	result, err = utilities.DecodeResults[models.GetOrderResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) CreateOrderItem(
	orderId string,
	request *models.CreateOrderItemRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetOrderItemResponse],
	error) {
	req := o.prepareRequest("POST", fmt.Sprintf("/orders/%s/items", orderId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}

	var result models.GetOrderItemResponse
	result, err = utilities.DecodeResults[models.GetOrderItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (o *OrdersController) GetOrderItem(
	orderId string,
	itemId string) (
	https.ApiResponse[models.GetOrderItemResponse],
	error) {
	req := o.prepareRequest(
		"GET",
		fmt.Sprintf("/orders/%s/items/%s", orderId, itemId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}

	var result models.GetOrderItemResponse
	result, err = utilities.DecodeResults[models.GetOrderItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetOrderItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the metadata from an order
func (o *OrdersController) UpdateOrderMetadata(
	orderId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetOrderResponse],
	error) {
	req := o.prepareRequest("PATCH", fmt.Sprintf("/Orders/%s/metadata", orderId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}

	var result models.GetOrderResponse
	result, err = utilities.DecodeResults[models.GetOrderResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets an order
func (o *OrdersController) GetOrder(orderId string) (
	https.ApiResponse[models.GetOrderResponse],
	error) {
	req := o.prepareRequest("GET", fmt.Sprintf("/orders/%s", orderId))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}

	var result models.GetOrderResponse
	result, err = utilities.DecodeResults[models.GetOrderResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetOrderResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
