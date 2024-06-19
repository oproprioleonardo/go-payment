package controllers

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"net/http"
	"pagarmeapisdk/models"
	"time"
)

type SubscriptionsController struct {
	baseController
}

func NewSubscriptionsController(baseController baseController) *SubscriptionsController {
	subscriptionsController := SubscriptionsController{baseController: baseController}
	return &subscriptionsController
}

// TODO: type endpoint description here
func (s *SubscriptionsController) RenewSubscription(
	subscriptionId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetPeriodResponse],
	error) {
	req := s.prepareRequest(
		"POST",
		fmt.Sprintf("/subscriptions/%s/cycles", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPeriodResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPeriodResponse]{Response: resp}, err
	}

	var result models.GetPeriodResponse
	result, err = utilities.DecodeResults[models.GetPeriodResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPeriodResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the credit card from a subscription
func (s *SubscriptionsController) UpdateSubscriptionCard(
	subscriptionId string,
	request *models.UpdateSubscriptionCardRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"PATCH",
		fmt.Sprintf("/subscriptions/%s/card", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Deletes a usage
func (s *SubscriptionsController) DeleteUsage(
	subscriptionId string,
	itemId string,
	usageId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetUsageResponse],
	error) {
	req := s.prepareRequest(
		"DELETE",
		fmt.Sprintf("/subscriptions/%s/items/%s/usages/%s", subscriptionId, itemId, usageId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetUsageResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetUsageResponse]{Response: resp}, err
	}

	var result models.GetUsageResponse
	result, err = utilities.DecodeResults[models.GetUsageResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetUsageResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a discount
func (s *SubscriptionsController) CreateDiscount(
	subscriptionId string,
	request *models.CreateDiscountRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetDiscountResponse],
	error) {
	req := s.prepareRequest(
		"POST",
		fmt.Sprintf("/subscriptions/%s/discounts", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetDiscountResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetDiscountResponse]{Response: resp}, err
	}

	var result models.GetDiscountResponse
	result, err = utilities.DecodeResults[models.GetDiscountResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetDiscountResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Create Usage
func (s *SubscriptionsController) CreateAnUsage(
	subscriptionId string,
	itemId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetUsageResponse],
	error) {
	req := s.prepareRequest(
		"POST",
		fmt.Sprintf("/subscriptions/%s/items/%s/usages", subscriptionId, itemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetUsageResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetUsageResponse]{Response: resp}, err
	}

	var result models.GetUsageResponse
	result, err = utilities.DecodeResults[models.GetUsageResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetUsageResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) UpdateCurrentCycleStatus(
	subscriptionId string,
	request *models.UpdateCurrentCycleStatusRequest,
	idempotencyKey *string) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		"PATCH",
		fmt.Sprintf("/subscriptions/%s/cycle-status", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	context, err := req.Call()
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// Deletes a discount
func (s *SubscriptionsController) DeleteDiscount(
	subscriptionId string,
	discountId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetDiscountResponse],
	error) {
	req := s.prepareRequest(
		"DELETE",
		fmt.Sprintf("/subscriptions/%s/discounts/%s", subscriptionId, discountId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetDiscountResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetDiscountResponse]{Response: resp}, err
	}

	var result models.GetDiscountResponse
	result, err = utilities.DecodeResults[models.GetDiscountResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetDiscountResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get Subscription Items
func (s *SubscriptionsController) GetSubscriptionItems(
	subscriptionId string,
	page *int,
	size *int,
	name *string,
	code *string,
	status *string,
	description *string,
	createdSince *string,
	createdUntil *string) (
	https.ApiResponse[models.ListSubscriptionItemsResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/items", subscriptionId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}
	if name != nil {
		req.QueryParam("name", *name)
	}
	if code != nil {
		req.QueryParam("code", *code)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if description != nil {
		req.QueryParam("description", *description)
	}
	if createdSince != nil {
		req.QueryParam("created_since", *createdSince)
	}
	if createdUntil != nil {
		req.QueryParam("created_until", *createdUntil)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListSubscriptionItemsResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListSubscriptionItemsResponse]{Response: resp}, err
	}

	var result models.ListSubscriptionItemsResponse
	result, err = utilities.DecodeResults[models.ListSubscriptionItemsResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListSubscriptionItemsResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the payment method from a subscription
func (s *SubscriptionsController) UpdateSubscriptionPaymentMethod(
	subscriptionId string,
	request *models.UpdateSubscriptionPaymentMethodRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"PATCH",
		fmt.Sprintf("/subscriptions/%s/payment-method", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Get Subscription Item
func (s *SubscriptionsController) GetSubscriptionItem(
	subscriptionId string,
	itemId string) (
	https.ApiResponse[models.GetSubscriptionItemResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/items/%s", subscriptionId, itemId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionItemResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets all subscriptions
func (s *SubscriptionsController) GetSubscriptions(
	page *int,
	size *int,
	code *string,
	billingType *string,
	customerId *string,
	planId *string,
	cardId *string,
	status *string,
	nextBillingSince *time.Time,
	nextBillingUntil *time.Time,
	createdSince *time.Time,
	createdUntil *time.Time) (
	https.ApiResponse[models.ListSubscriptionsResponse],
	error) {
	req := s.prepareRequest("GET", "/subscriptions")
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
	if billingType != nil {
		req.QueryParam("billing_type", *billingType)
	}
	if customerId != nil {
		req.QueryParam("customer_id", *customerId)
	}
	if planId != nil {
		req.QueryParam("plan_id", *planId)
	}
	if cardId != nil {
		req.QueryParam("card_id", *cardId)
	}
	if status != nil {
		req.QueryParam("status", *status)
	}
	if nextBillingSince != nil {
		req.QueryParam("next_billing_since", nextBillingSince.Format(time.RFC3339))
	}
	if nextBillingUntil != nil {
		req.QueryParam("next_billing_until", nextBillingUntil.Format(time.RFC3339))
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListSubscriptionsResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListSubscriptionsResponse]{Response: resp}, err
	}

	var result models.ListSubscriptionsResponse
	result, err = utilities.DecodeResults[models.ListSubscriptionsResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListSubscriptionsResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Cancels a subscription
func (s *SubscriptionsController) CancelSubscription(
	subscriptionId string,
	request *models.CreateCancelSubscriptionRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"DELETE",
		fmt.Sprintf("/subscriptions/%s", subscriptionId),
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
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a increment
func (s *SubscriptionsController) CreateIncrement(
	subscriptionId string,
	request *models.CreateIncrementRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetIncrementResponse],
	error) {
	req := s.prepareRequest(
		"POST",
		fmt.Sprintf("/subscriptions/%s/increments", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetIncrementResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetIncrementResponse]{Response: resp}, err
	}

	var result models.GetIncrementResponse
	result, err = utilities.DecodeResults[models.GetIncrementResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetIncrementResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a usage
func (s *SubscriptionsController) CreateUsage(
	subscriptionId string,
	itemId string,
	body *models.CreateUsageRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetUsageResponse],
	error) {
	req := s.prepareRequest(
		"POST",
		fmt.Sprintf("/subscriptions/%s/items/%s/usages", subscriptionId, itemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetUsageResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetUsageResponse]{Response: resp}, err
	}

	var result models.GetUsageResponse
	result, err = utilities.DecodeResults[models.GetUsageResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetUsageResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetDiscountById(
	subscriptionId string,
	discountId string) (
	https.ApiResponse[models.GetDiscountResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/discounts/%s", subscriptionId, discountId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetDiscountResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetDiscountResponse]{Response: resp}, err
	}

	var result models.GetDiscountResponse
	result, err = utilities.DecodeResults[models.GetDiscountResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetDiscountResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a new subscription
func (s *SubscriptionsController) CreateSubscription(
	body *models.CreateSubscriptionRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest("POST", "/subscriptions")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetIncrementById(
	subscriptionId string,
	incrementId string) (
	https.ApiResponse[models.GetIncrementResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/increments/%s", subscriptionId, incrementId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetIncrementResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetIncrementResponse]{Response: resp}, err
	}

	var result models.GetIncrementResponse
	result, err = utilities.DecodeResults[models.GetIncrementResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetIncrementResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) UpdateSubscriptionAffiliationId(
	subscriptionId string,
	request *models.UpdateSubscriptionAffiliationIdRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"PATCH",
		fmt.Sprintf("/subscriptions/%s/gateway-affiliation-id", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the metadata from a subscription
func (s *SubscriptionsController) UpdateSubscriptionMetadata(
	subscriptionId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"PATCH",
		fmt.Sprintf("/Subscriptions/%s/metadata", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Deletes a increment
func (s *SubscriptionsController) DeleteIncrement(
	subscriptionId string,
	incrementId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetIncrementResponse],
	error) {
	req := s.prepareRequest(
		"DELETE",
		fmt.Sprintf("/subscriptions/%s/increments/%s", subscriptionId, incrementId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetIncrementResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetIncrementResponse]{Response: resp}, err
	}

	var result models.GetIncrementResponse
	result, err = utilities.DecodeResults[models.GetIncrementResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetIncrementResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetSubscriptionCycles(
	subscriptionId string,
	page string,
	size string) (
	https.ApiResponse[models.ListCyclesResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/cycles", subscriptionId),
	)
	req.Authenticate(true)
	req.QueryParam("page", page)
	req.QueryParam("size", size)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListCyclesResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListCyclesResponse]{Response: resp}, err
	}

	var result models.ListCyclesResponse
	result, err = utilities.DecodeResults[models.ListCyclesResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListCyclesResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetDiscounts(
	subscriptionId string,
	page int,
	size int) (
	https.ApiResponse[models.ListDiscountsResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/discounts/", subscriptionId),
	)
	req.Authenticate(true)
	req.QueryParam("page", page)
	req.QueryParam("size", size)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListDiscountsResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListDiscountsResponse]{Response: resp}, err
	}

	var result models.ListDiscountsResponse
	result, err = utilities.DecodeResults[models.ListDiscountsResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListDiscountsResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the billing date from a subscription
func (s *SubscriptionsController) UpdateSubscriptionBillingDate(
	subscriptionId string,
	request *models.UpdateSubscriptionBillingDateRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"PATCH",
		fmt.Sprintf("/subscriptions/%s/billing-date", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Deletes a subscription item
func (s *SubscriptionsController) DeleteSubscriptionItem(
	subscriptionId string,
	subscriptionItemId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionItemResponse],
	error) {
	req := s.prepareRequest(
		"DELETE",
		fmt.Sprintf("/subscriptions/%s/items/%s", subscriptionId, subscriptionItemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionItemResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetIncrements(
	subscriptionId string,
	page *int,
	size *int) (
	https.ApiResponse[models.ListIncrementsResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/increments/", subscriptionId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListIncrementsResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListIncrementsResponse]{Response: resp}, err
	}

	var result models.ListIncrementsResponse
	result, err = utilities.DecodeResults[models.ListIncrementsResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListIncrementsResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the boleto due days from a subscription
func (s *SubscriptionsController) UpdateSubscriptionDueDays(
	subscriptionId string,
	request *models.UpdateSubscriptionDueDaysRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"PATCH",
		fmt.Sprintf("/subscriptions/%s/boleto-due-days", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the start at date from a subscription
func (s *SubscriptionsController) UpdateSubscriptionStartAt(
	subscriptionId string,
	request *models.UpdateSubscriptionStartAtRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"PATCH",
		fmt.Sprintf("/subscriptions/%s/start-at", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates a subscription item
func (s *SubscriptionsController) UpdateSubscriptionItem(
	subscriptionId string,
	itemId string,
	body *models.UpdateSubscriptionItemRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionItemResponse],
	error) {
	req := s.prepareRequest(
		"PUT",
		fmt.Sprintf("/subscriptions/%s/items/%s", subscriptionId, itemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionItemResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a new Subscription item
func (s *SubscriptionsController) CreateSubscriptionItem(
	subscriptionId string,
	request *models.CreateSubscriptionItemRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionItemResponse],
	error) {
	req := s.prepareRequest(
		"POST",
		fmt.Sprintf("/subscriptions/%s/items", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionItemResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets a subscription
func (s *SubscriptionsController) GetSubscription(subscriptionId string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s", subscriptionId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Lists all usages from a subscription item
func (s *SubscriptionsController) GetUsages(
	subscriptionId string,
	itemId string,
	page *int,
	size *int,
	code *string,
	group *string,
	usedSince *time.Time,
	usedUntil *time.Time) (
	https.ApiResponse[models.ListUsagesResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/items/%s/usages", subscriptionId, itemId),
	)
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
	if group != nil {
		req.QueryParam("group", *group)
	}
	if usedSince != nil {
		req.QueryParam("used_since", usedSince.Format(time.RFC3339))
	}
	if usedUntil != nil {
		req.QueryParam("used_until", usedUntil.Format(time.RFC3339))
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListUsagesResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListUsagesResponse]{Response: resp}, err
	}

	var result models.ListUsagesResponse
	result, err = utilities.DecodeResults[models.ListUsagesResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListUsagesResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) UpdateLatestPeriodEndAt(
	subscriptionId string,
	request *models.UpdateCurrentCycleEndDateRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"PATCH",
		fmt.Sprintf("/subscriptions/%s/periods/latest/end-at", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Atualização do valor mínimo da assinatura
func (s *SubscriptionsController) UpdateSubscriptionMiniumPrice(
	subscriptionId string,
	request *models.UpdateSubscriptionMinimumPriceRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		"PATCH",
		fmt.Sprintf("/subscriptions/%s/minimum_price", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetSubscriptionCycleById(
	subscriptionId string,
	cycleId string) (
	https.ApiResponse[models.GetPeriodResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/cycles/%s", subscriptionId, cycleId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPeriodResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPeriodResponse]{Response: resp}, err
	}

	var result models.GetPeriodResponse
	result, err = utilities.DecodeResults[models.GetPeriodResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPeriodResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) GetUsageReport(
	subscriptionId string,
	periodId string) (
	https.ApiResponse[models.GetUsageReportResponse],
	error) {
	req := s.prepareRequest(
		"GET",
		fmt.Sprintf("/subscriptions/%s/periods/%s/usages/report", subscriptionId, periodId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetUsageReportResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetUsageReportResponse]{Response: resp}, err
	}

	var result models.GetUsageReportResponse
	result, err = utilities.DecodeResults[models.GetUsageReportResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetUsageReportResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// TODO: type endpoint description here
func (s *SubscriptionsController) UpdateSplitSubscription(
	id string,
	request *models.UpdateSubscriptionSplitRequest) (
	https.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest("PATCH", fmt.Sprintf("/subscriptions/%s/split", id))
	req.Authenticate(true)
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	var result models.GetSubscriptionResponse
	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetSubscriptionResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
