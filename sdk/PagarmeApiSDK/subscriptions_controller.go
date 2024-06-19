package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"net/http"
	"pagarmeapisdk/models"
	"time"
)

// SubscriptionsController represents a controller struct.
type SubscriptionsController struct {
	baseController
}

// NewSubscriptionsController creates a new instance of SubscriptionsController.
// It takes a baseController as a parameter and returns a pointer to the SubscriptionsController.
func NewSubscriptionsController(baseController baseController) *SubscriptionsController {
	subscriptionsController := SubscriptionsController{baseController: baseController}
	return &subscriptionsController
}

// RenewSubscription takes context, subscriptionId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetPeriodResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) RenewSubscription(
	ctx context.Context,
	subscriptionId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetPeriodResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/cycles", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetPeriodResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPeriodResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionCard takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
// Updates the credit card from a subscription
func (s *SubscriptionsController) UpdateSubscriptionCard(
	ctx context.Context,
	subscriptionId string,
	request *models.UpdateSubscriptionCardRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/subscriptions/%v/card", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteUsage takes context, subscriptionId, itemId, usageId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetUsageResponse data and
// an error if there was an issue with the request or response.
// Deletes a usage
func (s *SubscriptionsController) DeleteUsage(
	ctx context.Context,
	subscriptionId string,
	itemId string,
	usageId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetUsageResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/subscriptions/%v/items/%v/usages/%v", subscriptionId, itemId, usageId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetUsageResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetUsageResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateDiscount takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetDiscountResponse data and
// an error if there was an issue with the request or response.
// Creates a discount
func (s *SubscriptionsController) CreateDiscount(
	ctx context.Context,
	subscriptionId string,
	request *models.CreateDiscountRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetDiscountResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/discounts", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetDiscountResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetDiscountResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateAnUsage takes context, subscriptionId, itemId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetUsageResponse data and
// an error if there was an issue with the request or response.
// Create Usage
func (s *SubscriptionsController) CreateAnUsage(
	ctx context.Context,
	subscriptionId string,
	itemId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetUsageResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/items/%v/usages", subscriptionId, itemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetUsageResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetUsageResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateCurrentCycleStatus takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with  data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) UpdateCurrentCycleStatus(
	ctx context.Context,
	subscriptionId string,
	request *models.UpdateCurrentCycleStatusRequest,
	idempotencyKey *string) (
	*http.Response,
	error) {
	req := s.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/subscriptions/%v/cycle-status", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	context, err := req.Call()
	if err != nil {
		return context.Response, err
	}
	err = validateResponse(*context.Response)
	if err != nil {
		return context.Response, err
	}
	return context.Response, err
}

// DeleteDiscount takes context, subscriptionId, discountId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetDiscountResponse data and
// an error if there was an issue with the request or response.
// Deletes a discount
func (s *SubscriptionsController) DeleteDiscount(
	ctx context.Context,
	subscriptionId string,
	discountId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetDiscountResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/subscriptions/%v/discounts/%v", subscriptionId, discountId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetDiscountResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetDiscountResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetSubscriptionItems takes context, subscriptionId, page, size, name, code, status, description, createdSince, createdUntil as parameters and
// returns an models.ApiResponse with models.ListSubscriptionItemsResponse data and
// an error if there was an issue with the request or response.
// Get Subscription Items
func (s *SubscriptionsController) GetSubscriptionItems(
	ctx context.Context,
	subscriptionId string,
	page *int,
	size *int,
	name *string,
	code *string,
	status *string,
	description *string,
	createdSince *string,
	createdUntil *string) (
	models.ApiResponse[models.ListSubscriptionItemsResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/items", subscriptionId),
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

	var result models.ListSubscriptionItemsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListSubscriptionItemsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionPaymentMethod takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
// Updates the payment method from a subscription
func (s *SubscriptionsController) UpdateSubscriptionPaymentMethod(
	ctx context.Context,
	subscriptionId string,
	request *models.UpdateSubscriptionPaymentMethodRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/subscriptions/%v/payment-method", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetSubscriptionItem takes context, subscriptionId, itemId as parameters and
// returns an models.ApiResponse with models.GetSubscriptionItemResponse data and
// an error if there was an issue with the request or response.
// Get Subscription Item
func (s *SubscriptionsController) GetSubscriptionItem(
	ctx context.Context,
	subscriptionId string,
	itemId string) (
	models.ApiResponse[models.GetSubscriptionItemResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/items/%v", subscriptionId, itemId),
	)
	req.Authenticate(true)

	var result models.GetSubscriptionItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetSubscriptions takes context, page, size, code, billingType, customerId, planId, cardId, status, nextBillingSince, nextBillingUntil, createdSince, createdUntil as parameters and
// returns an models.ApiResponse with models.ListSubscriptionsResponse data and
// an error if there was an issue with the request or response.
// Gets all subscriptions
func (s *SubscriptionsController) GetSubscriptions(
	ctx context.Context,
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
	models.ApiResponse[models.ListSubscriptionsResponse],
	error) {
	req := s.prepareRequest(ctx, "GET", "/subscriptions")
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
	var result models.ListSubscriptionsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListSubscriptionsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CancelSubscription takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
// Cancels a subscription
func (s *SubscriptionsController) CancelSubscription(
	ctx context.Context,
	subscriptionId string,
	request *models.CreateCancelSubscriptionRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/subscriptions/%v", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	if request != nil {
		req.Json(*request)
	}

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateIncrement takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetIncrementResponse data and
// an error if there was an issue with the request or response.
// Creates a increment
func (s *SubscriptionsController) CreateIncrement(
	ctx context.Context,
	subscriptionId string,
	request *models.CreateIncrementRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetIncrementResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/increments", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetIncrementResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetIncrementResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateUsage takes context, subscriptionId, itemId, body, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetUsageResponse data and
// an error if there was an issue with the request or response.
// Creates a usage
func (s *SubscriptionsController) CreateUsage(
	ctx context.Context,
	subscriptionId string,
	itemId string,
	body *models.CreateUsageRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetUsageResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/items/%v/usages", subscriptionId, itemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)

	var result models.GetUsageResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetUsageResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetDiscountById takes context, subscriptionId, discountId as parameters and
// returns an models.ApiResponse with models.GetDiscountResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) GetDiscountById(
	ctx context.Context,
	subscriptionId string,
	discountId string) (
	models.ApiResponse[models.GetDiscountResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/discounts/%v", subscriptionId, discountId),
	)
	req.Authenticate(true)

	var result models.GetDiscountResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetDiscountResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateSubscription takes context, body, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
// Creates a new subscription
func (s *SubscriptionsController) CreateSubscription(
	ctx context.Context,
	body *models.CreateSubscriptionRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(ctx, "POST", "/subscriptions")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)
	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetIncrementById takes context, subscriptionId, incrementId as parameters and
// returns an models.ApiResponse with models.GetIncrementResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) GetIncrementById(
	ctx context.Context,
	subscriptionId string,
	incrementId string) (
	models.ApiResponse[models.GetIncrementResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/increments/%v", subscriptionId, incrementId),
	)
	req.Authenticate(true)

	var result models.GetIncrementResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetIncrementResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionAffiliationId takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) UpdateSubscriptionAffiliationId(
	ctx context.Context,
	subscriptionId string,
	request *models.UpdateSubscriptionAffiliationIdRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/subscriptions/%v/gateway-affiliation-id", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionMetadata takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
// Updates the metadata from a subscription
func (s *SubscriptionsController) UpdateSubscriptionMetadata(
	ctx context.Context,
	subscriptionId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/Subscriptions/%v/metadata", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteIncrement takes context, subscriptionId, incrementId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetIncrementResponse data and
// an error if there was an issue with the request or response.
// Deletes a increment
func (s *SubscriptionsController) DeleteIncrement(
	ctx context.Context,
	subscriptionId string,
	incrementId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetIncrementResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/subscriptions/%v/increments/%v", subscriptionId, incrementId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetIncrementResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetIncrementResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetSubscriptionCycles takes context, subscriptionId, page, size as parameters and
// returns an models.ApiResponse with models.ListCyclesResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) GetSubscriptionCycles(
	ctx context.Context,
	subscriptionId string,
	page string,
	size string) (
	models.ApiResponse[models.ListCyclesResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/cycles", subscriptionId),
	)
	req.Authenticate(true)
	req.QueryParam("page", page)
	req.QueryParam("size", size)

	var result models.ListCyclesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListCyclesResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetDiscounts takes context, subscriptionId, page, size as parameters and
// returns an models.ApiResponse with models.ListDiscountsResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) GetDiscounts(
	ctx context.Context,
	subscriptionId string,
	page int,
	size int) (
	models.ApiResponse[models.ListDiscountsResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/discounts/", subscriptionId),
	)
	req.Authenticate(true)
	req.QueryParam("page", page)
	req.QueryParam("size", size)

	var result models.ListDiscountsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListDiscountsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionBillingDate takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
// Updates the billing date from a subscription
func (s *SubscriptionsController) UpdateSubscriptionBillingDate(
	ctx context.Context,
	subscriptionId string,
	request *models.UpdateSubscriptionBillingDateRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/subscriptions/%v/billing-date", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeleteSubscriptionItem takes context, subscriptionId, subscriptionItemId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionItemResponse data and
// an error if there was an issue with the request or response.
// Deletes a subscription item
func (s *SubscriptionsController) DeleteSubscriptionItem(
	ctx context.Context,
	subscriptionId string,
	subscriptionItemId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionItemResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/subscriptions/%v/items/%v", subscriptionId, subscriptionItemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetSubscriptionItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetIncrements takes context, subscriptionId, page, size as parameters and
// returns an models.ApiResponse with models.ListIncrementsResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) GetIncrements(
	ctx context.Context,
	subscriptionId string,
	page *int,
	size *int) (
	models.ApiResponse[models.ListIncrementsResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/increments/", subscriptionId),
	)
	req.Authenticate(true)
	if page != nil {
		req.QueryParam("page", *page)
	}
	if size != nil {
		req.QueryParam("size", *size)
	}

	var result models.ListIncrementsResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListIncrementsResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionDueDays takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
// Updates the boleto due days from a subscription
func (s *SubscriptionsController) UpdateSubscriptionDueDays(
	ctx context.Context,
	subscriptionId string,
	request *models.UpdateSubscriptionDueDaysRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/subscriptions/%v/boleto-due-days", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionStartAt takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
// Updates the start at date from a subscription
func (s *SubscriptionsController) UpdateSubscriptionStartAt(
	ctx context.Context,
	subscriptionId string,
	request *models.UpdateSubscriptionStartAtRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/subscriptions/%v/start-at", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionItem takes context, subscriptionId, itemId, body, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionItemResponse data and
// an error if there was an issue with the request or response.
// Updates a subscription item
func (s *SubscriptionsController) UpdateSubscriptionItem(
	ctx context.Context,
	subscriptionId string,
	itemId string,
	body *models.UpdateSubscriptionItemRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionItemResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/subscriptions/%v/items/%v", subscriptionId, itemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)

	var result models.GetSubscriptionItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreateSubscriptionItem takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionItemResponse data and
// an error if there was an issue with the request or response.
// Creates a new Subscription item
func (s *SubscriptionsController) CreateSubscriptionItem(
	ctx context.Context,
	subscriptionId string,
	request *models.CreateSubscriptionItemRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionItemResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"POST",
		fmt.Sprintf("/subscriptions/%v/items", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetSubscriptionItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetSubscription takes context, subscriptionId as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
// Gets a subscription
func (s *SubscriptionsController) GetSubscription(
	ctx context.Context,
	subscriptionId string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v", subscriptionId),
	)
	req.Authenticate(true)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetUsages takes context, subscriptionId, itemId, page, size, code, group, usedSince, usedUntil as parameters and
// returns an models.ApiResponse with models.ListUsagesResponse data and
// an error if there was an issue with the request or response.
// Lists all usages from a subscription item
func (s *SubscriptionsController) GetUsages(
	ctx context.Context,
	subscriptionId string,
	itemId string,
	page *int,
	size *int,
	code *string,
	group *string,
	usedSince *time.Time,
	usedUntil *time.Time) (
	models.ApiResponse[models.ListUsagesResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/items/%v/usages", subscriptionId, itemId),
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

	var result models.ListUsagesResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListUsagesResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateLatestPeriodEndAt takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) UpdateLatestPeriodEndAt(
	ctx context.Context,
	subscriptionId string,
	request *models.UpdateCurrentCycleEndDateRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/subscriptions/%v/periods/latest/end-at", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSubscriptionMiniumPrice takes context, subscriptionId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
// Atualização do valor mínimo da assinatura
func (s *SubscriptionsController) UpdateSubscriptionMiniumPrice(
	ctx context.Context,
	subscriptionId string,
	request *models.UpdateSubscriptionMinimumPriceRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"PATCH",
		fmt.Sprintf("/subscriptions/%v/minimum_price", subscriptionId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetSubscriptionCycleById takes context, subscriptionId, cycleId as parameters and
// returns an models.ApiResponse with models.GetPeriodResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) GetSubscriptionCycleById(
	ctx context.Context,
	subscriptionId string,
	cycleId string) (
	models.ApiResponse[models.GetPeriodResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/cycles/%v", subscriptionId, cycleId),
	)
	req.Authenticate(true)

	var result models.GetPeriodResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPeriodResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetUsageReport takes context, subscriptionId, periodId as parameters and
// returns an models.ApiResponse with models.GetUsageReportResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) GetUsageReport(
	ctx context.Context,
	subscriptionId string,
	periodId string) (
	models.ApiResponse[models.GetUsageReportResponse],
	error) {
	req := s.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/subscriptions/%v/periods/%v/usages/report", subscriptionId, periodId),
	)
	req.Authenticate(true)

	var result models.GetUsageReportResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetUsageReportResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdateSplitSubscription takes context, id, request as parameters and
// returns an models.ApiResponse with models.GetSubscriptionResponse data and
// an error if there was an issue with the request or response.
func (s *SubscriptionsController) UpdateSplitSubscription(
	ctx context.Context,
	id string,
	request *models.UpdateSubscriptionSplitRequest) (
	models.ApiResponse[models.GetSubscriptionResponse],
	error) {
	req := s.prepareRequest(ctx, "PATCH", fmt.Sprintf("/subscriptions/%v/split", id))
	req.Authenticate(true)
	req.Json(request)

	var result models.GetSubscriptionResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetSubscriptionResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
