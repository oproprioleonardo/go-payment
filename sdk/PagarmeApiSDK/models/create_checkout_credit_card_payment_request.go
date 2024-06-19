package models

import (
	"encoding/json"
)

// CreateCheckoutCreditCardPaymentRequest represents a CreateCheckoutCreditCardPaymentRequest struct.
// Checkout card payment request
type CreateCheckoutCreditCardPaymentRequest struct {
	// Card invoice text descriptor
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// Payment installment options
	Installments []CreateCheckoutCardInstallmentOptionRequest `json:"installments,omitempty"`
	// Creates payment authentication
	Authentication *CreatePaymentAuthenticationRequest `json:"authentication,omitempty"`
	// Authorize and capture?
	Capture *bool `json:"capture,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCheckoutCreditCardPaymentRequest.
// It customizes the JSON marshaling process for CreateCheckoutCreditCardPaymentRequest objects.
func (c *CreateCheckoutCreditCardPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCheckoutCreditCardPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateCheckoutCreditCardPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.StatementDescriptor != nil {
		structMap["statement_descriptor"] = c.StatementDescriptor
	}
	if c.Installments != nil {
		structMap["installments"] = c.Installments
	}
	if c.Authentication != nil {
		structMap["authentication"] = c.Authentication
	}
	if c.Capture != nil {
		structMap["capture"] = c.Capture
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCheckoutCreditCardPaymentRequest.
// It customizes the JSON unmarshaling process for CreateCheckoutCreditCardPaymentRequest objects.
func (c *CreateCheckoutCreditCardPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor *string                                      `json:"statement_descriptor,omitempty"`
		Installments        []CreateCheckoutCardInstallmentOptionRequest `json:"installments,omitempty"`
		Authentication      *CreatePaymentAuthenticationRequest          `json:"authentication,omitempty"`
		Capture             *bool                                        `json:"capture,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.StatementDescriptor = temp.StatementDescriptor
	c.Installments = temp.Installments
	c.Authentication = temp.Authentication
	c.Capture = temp.Capture
	return nil
}
