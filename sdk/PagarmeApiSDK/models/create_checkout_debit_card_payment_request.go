package models

import (
	"encoding/json"
)

// CreateCheckoutDebitCardPaymentRequest represents a CreateCheckoutDebitCardPaymentRequest struct.
// Checkout credit card payment request
type CreateCheckoutDebitCardPaymentRequest struct {
	// Card invoice text descriptor
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
	// Creates payment authentication
	Authentication CreatePaymentAuthenticationRequest `json:"authentication"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCheckoutDebitCardPaymentRequest.
// It customizes the JSON marshaling process for CreateCheckoutDebitCardPaymentRequest objects.
func (c *CreateCheckoutDebitCardPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCheckoutDebitCardPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateCheckoutDebitCardPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.StatementDescriptor != nil {
		structMap["statement_descriptor"] = c.StatementDescriptor
	}
	structMap["authentication"] = c.Authentication
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCheckoutDebitCardPaymentRequest.
// It customizes the JSON unmarshaling process for CreateCheckoutDebitCardPaymentRequest objects.
func (c *CreateCheckoutDebitCardPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StatementDescriptor *string                            `json:"statement_descriptor,omitempty"`
		Authentication      CreatePaymentAuthenticationRequest `json:"authentication"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.StatementDescriptor = temp.StatementDescriptor
	c.Authentication = temp.Authentication
	return nil
}
