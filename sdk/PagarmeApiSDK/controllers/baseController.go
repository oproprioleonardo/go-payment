/*
Package pagarmeapisdk

This file was automatically generated by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package controllers

import (
	"github.com/apimatic/go-core-runtime/https"
	"net/http"
	"pagarmeapisdk/errors"
)

type callBuilderFactory interface {
	GetCallBuilder() https.CallBuilderFactory
}

type baseController struct {
	callBuilder    callBuilderFactory
	prepareRequest https.CallBuilderFactory
}

func NewBaseController(cb callBuilderFactory) *baseController {
	baseController := baseController{callBuilder: cb}
	baseController.prepareRequest = baseController.callBuilder.GetCallBuilder()
	return &baseController
}

func validateResponse(response http.Response) error {
	if response.StatusCode == 400 {
		return errors.NewError(400, "Invalid request")
	}
	if response.StatusCode == 401 {
		return errors.NewError(401, "Invalid API key")
	}
	if response.StatusCode == 404 {
		return errors.NewError(404, "An informed resource was not found")
	}
	if response.StatusCode == 412 {
		return errors.NewError(412, "Business validation error")
	}
	if response.StatusCode == 422 {
		return errors.NewError(422, "Contract validation error")
	}
	if response.StatusCode == 500 {
		return errors.NewError(500, "Internal server error")
	}
	return nil
}
