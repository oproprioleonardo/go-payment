package controllers

import (
	"fmt"
	"github.com/apimatic/go-core-runtime/https"
	"github.com/apimatic/go-core-runtime/utilities"
	"pagarmeapisdk/models"
	"time"
)

type PlansController struct {
	baseController
}

func NewPlansController(baseController baseController) *PlansController {
	plansController := PlansController{baseController: baseController}
	return &plansController
}

// Gets a plan
func (p *PlansController) GetPlan(planId string) (
	https.ApiResponse[models.GetPlanResponse],
	error) {
	req := p.prepareRequest("GET", fmt.Sprintf("/plans/%s", planId))
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}

	var result models.GetPlanResponse
	result, err = utilities.DecodeResults[models.GetPlanResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Deletes a plan
func (p *PlansController) DeletePlan(
	planId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetPlanResponse],
	error) {
	req := p.prepareRequest("DELETE", fmt.Sprintf("/plans/%s", planId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}

	var result models.GetPlanResponse
	result, err = utilities.DecodeResults[models.GetPlanResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates the metadata from a plan
func (p *PlansController) UpdatePlanMetadata(
	planId string,
	request *models.UpdateMetadataRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetPlanResponse],
	error) {
	req := p.prepareRequest("PATCH", fmt.Sprintf("/Plans/%s/metadata", planId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}

	var result models.GetPlanResponse
	result, err = utilities.DecodeResults[models.GetPlanResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates a plan item
func (p *PlansController) UpdatePlanItem(
	planId string,
	planItemId string,
	body *models.UpdatePlanItemRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetPlanItemResponse],
	error) {
	req := p.prepareRequest(
		"PUT",
		fmt.Sprintf("/plans/%s/items/%s", planId, planItemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}

	var result models.GetPlanItemResponse
	result, err = utilities.DecodeResults[models.GetPlanItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Adds a new item to a plan
func (p *PlansController) CreatePlanItem(
	planId string,
	request *models.CreatePlanItemRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetPlanItemResponse],
	error) {
	req := p.prepareRequest("POST", fmt.Sprintf("/plans/%s/items", planId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}

	var result models.GetPlanItemResponse
	result, err = utilities.DecodeResults[models.GetPlanItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets a plan item
func (p *PlansController) GetPlanItem(
	planId string,
	planItemId string) (
	https.ApiResponse[models.GetPlanItemResponse],
	error) {
	req := p.prepareRequest(
		"GET",
		fmt.Sprintf("/plans/%s/items/%s", planId, planItemId),
	)
	req.Authenticate(true)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}

	var result models.GetPlanItemResponse
	result, err = utilities.DecodeResults[models.GetPlanItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Creates a new plan
func (p *PlansController) CreatePlan(
	body *models.CreatePlanRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetPlanResponse],
	error) {
	req := p.prepareRequest("POST", "/plans")
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(body)
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}

	var result models.GetPlanResponse
	result, err = utilities.DecodeResults[models.GetPlanResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Removes an item from a plan
func (p *PlansController) DeletePlanItem(
	planId string,
	planItemId string,
	idempotencyKey *string) (
	https.ApiResponse[models.GetPlanItemResponse],
	error) {
	req := p.prepareRequest(
		"DELETE",
		fmt.Sprintf("/plans/%s/items/%s", planId, planItemId),
	)
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}

	var result models.GetPlanItemResponse
	result, err = utilities.DecodeResults[models.GetPlanItemResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPlanItemResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Gets all plans
func (p *PlansController) GetPlans(
	page *int,
	size *int,
	name *string,
	status *string,
	billingType *string,
	createdSince *time.Time,
	createdUntil *time.Time) (
	https.ApiResponse[models.ListPlansResponse],
	error) {
	req := p.prepareRequest("GET", "/plans")
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
	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.ListPlansResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.ListPlansResponse]{Response: resp}, err
	}

	var result models.ListPlansResponse
	result, err = utilities.DecodeResults[models.ListPlansResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.ListPlansResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}

// Updates a plan
func (p *PlansController) UpdatePlan(
	planId string,
	request *models.UpdatePlanRequest,
	idempotencyKey *string) (
	https.ApiResponse[models.GetPlanResponse],
	error) {
	req := p.prepareRequest("PUT", fmt.Sprintf("/plans/%s", planId))
	req.Authenticate(true)
	if idempotencyKey != nil {
		req.Header("idempotency-key", *idempotencyKey)
	}
	req.Json(request)

	decoder, resp, err := req.CallAsJson()
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}
	err = validateResponse(*resp)
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}

	var result models.GetPlanResponse
	result, err = utilities.DecodeResults[models.GetPlanResponse](decoder)
	if err != nil {
		return https.ApiResponse[models.GetPlanResponse]{Response: resp}, err
	}

	return https.NewApiResponse(result, resp), err
}
