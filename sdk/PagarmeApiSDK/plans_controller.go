package pagarmeapisdk

import (
	"context"
	"fmt"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

// PlansController represents a controller struct.
type PlansController struct {
	baseController
}

// NewPlansController creates a new instance of PlansController.
// It takes a baseController as a parameter and returns a pointer to the PlansController.
func NewPlansController(baseController baseController) *PlansController {
	plansController := PlansController{baseController: baseController}
	return &plansController
}

// GetPlan takes context, planId as parameters and
// returns an models.ApiResponse with models.GetPlanResponse data and
// an error if there was an issue with the request or response.
// Gets a plan
func (p *PlansController) GetPlan(
	ctx context.Context,
	planId string) (
	models.ApiResponse[models.GetPlanResponse],
	error) {
	req := p.prepareRequest(ctx, "GET", fmt.Sprintf("/plans/%v", planId))
	req.Authenticate(true)

	var result models.GetPlanResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPlanResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeletePlan takes context, planId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetPlanResponse data and
// an error if there was an issue with the request or response.
// Deletes a plan
func (p *PlansController) DeletePlan(
	ctx context.Context,
	planId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetPlanResponse],
	error) {
	req := p.prepareRequest(ctx, "DELETE", fmt.Sprintf("/plans/%v", planId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetPlanResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPlanResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdatePlanMetadata takes context, planId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetPlanResponse data and
// an error if there was an issue with the request or response.
// Updates the metadata from a plan
func (p *PlansController) UpdatePlanMetadata(
	ctx context.Context,
	planId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetPlanResponse],
	error) {
	req := p.prepareRequest(ctx, "PATCH", fmt.Sprintf("/Plans/%v/metadata", planId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetPlanResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPlanResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdatePlanItem takes context, planId, planItemId, body, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetPlanItemResponse data and
// an error if there was an issue with the request or response.
// Updates a plan item
func (p *PlansController) UpdatePlanItem(
	ctx context.Context,
	planId string,
	planItemId string,
	body *models.UpdatePlanItemRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetPlanItemResponse],
	error) {
	req := p.prepareRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/plans/%v/items/%v", planId, planItemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)

	var result models.GetPlanItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPlanItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreatePlanItem takes context, planId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetPlanItemResponse data and
// an error if there was an issue with the request or response.
// Adds a new item to a plan
func (p *PlansController) CreatePlanItem(
	ctx context.Context,
	planId string,
	request *models.CreatePlanItemRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetPlanItemResponse],
	error) {
	req := p.prepareRequest(ctx, "POST", fmt.Sprintf("/plans/%v/items", planId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetPlanItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPlanItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetPlanItem takes context, planId, planItemId as parameters and
// returns an models.ApiResponse with models.GetPlanItemResponse data and
// an error if there was an issue with the request or response.
// Gets a plan item
func (p *PlansController) GetPlanItem(
	ctx context.Context,
	planId string,
	planItemId string) (
	models.ApiResponse[models.GetPlanItemResponse],
	error) {
	req := p.prepareRequest(
		ctx,
		"GET",
		fmt.Sprintf("/plans/%v/items/%v", planId, planItemId),
	)
	req.Authenticate(true)

	var result models.GetPlanItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPlanItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// CreatePlan takes context, body, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetPlanResponse data and
// an error if there was an issue with the request or response.
// Creates a new plan
func (p *PlansController) CreatePlan(
	ctx context.Context,
	body *models.CreatePlanRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetPlanResponse],
	error) {
	req := p.prepareRequest(ctx, "POST", "/plans")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)
	var result models.GetPlanResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPlanResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// DeletePlanItem takes context, planId, planItemId, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetPlanItemResponse data and
// an error if there was an issue with the request or response.
// Removes an item from a plan
func (p *PlansController) DeletePlanItem(
	ctx context.Context,
	planId string,
	planItemId string,
	idempotencyKey *string) (
	models.ApiResponse[models.GetPlanItemResponse],
	error) {
	req := p.prepareRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/plans/%v/items/%v", planId, planItemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	var result models.GetPlanItemResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPlanItemResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// GetPlans takes context, page, size, name, status, billingType, createdSince, createdUntil as parameters and
// returns an models.ApiResponse with models.ListPlansResponse data and
// an error if there was an issue with the request or response.
// Gets all plans
func (p *PlansController) GetPlans(
	ctx context.Context,
	page *int,
	size *int,
	name *string,
	status *string,
	billingType *string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	models.ApiResponse[models.ListPlansResponse],
	error) {
	req := p.prepareRequest(ctx, "GET", "/plans")
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
	if status != nil {
		req.QueryParam("status", *status)
	}
	if billingType != nil {
		req.QueryParam("billing_type", *billingType)
	}
	if createdSince != nil {
		req.QueryParam("created_since", createdSince.Format(time.RFC3339))
	}
	if createdUntil != nil {
		req.QueryParam("created_until", createdUntil.Format(time.RFC3339))
	}
	var result models.ListPlansResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.ListPlansResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}

// UpdatePlan takes context, planId, request, idempotencyKey as parameters and
// returns an models.ApiResponse with models.GetPlanResponse data and
// an error if there was an issue with the request or response.
// Updates a plan
func (p *PlansController) UpdatePlan(
	ctx context.Context,
	planId string,
	request *models.UpdatePlanRequest,
	idempotencyKey *string) (
	models.ApiResponse[models.GetPlanResponse],
	error) {
	req := p.prepareRequest(ctx, "PUT", fmt.Sprintf("/plans/%v", planId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	var result models.GetPlanResponse
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}
	err = validateResponse(*resp)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	result, err = utilities.DecodeResults[models.GetPlanResponse](decoder)
	if err != nil {
		return models.NewApiResponse(result, resp), err
	}

	return models.NewApiResponse(result, resp), err
}
