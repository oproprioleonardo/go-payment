package pagarmeapisdk

import (
	"github.com/apimatic/go-core-runtime/https"
	"net/http"
	"pagarmeapisdk/errors"
)

// callBuilderFactory is an interface that defines a method to get a CallBuilderFactory.
// It allows objects to get a reference to a CallBuilderFactory for creating API call.
type callBuilderFactory interface {
	GetCallBuilder() https.CallBuilderFactory
}

// baseController represents a controller used as a base for other controllers.
// It encapsulates common functionality required by controllers for making API call.
type baseController struct {
	callBuilder    callBuilderFactory
	prepareRequest https.CallBuilderFactory
}

// NewBaseController creates a new instance of baseController.
// It takes a callBuilderFactory as a parameter and returns a pointer to the baseController.
func NewBaseController(cb callBuilderFactory) *baseController {
	baseController := baseController{callBuilder: cb}
	baseController.prepareRequest = baseController.callBuilder.GetCallBuilder()
	return &baseController
}

// validateResponse is a function used to validate the HTTP response.
// It takes an http.Response object as a parameter and returns an error, if any.
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
