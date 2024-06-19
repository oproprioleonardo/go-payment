package models

import (
	"encoding/json"
)

// CreateConfirmPaymentRequest represents a CreateConfirmPaymentRequest struct.
type CreateConfirmPaymentRequest struct {
	// Description
	Description string `json:"description"`
	// Amount
	Amount *int `json:"Amount,omitempty"`
	// Code reference
	Code string `json:"Code"`
}

// MarshalJSON implements the json.Marshaler interface for CreateConfirmPaymentRequest.
// It customizes the JSON marshaling process for CreateConfirmPaymentRequest objects.
func (c *CreateConfirmPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateConfirmPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateConfirmPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["description"] = c.Description
	if c.Amount != nil {
		structMap["Amount"] = c.Amount
	}
	structMap["Code"] = c.Code
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateConfirmPaymentRequest.
// It customizes the JSON unmarshaling process for CreateConfirmPaymentRequest objects.
func (c *CreateConfirmPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Description string `json:"description"`
		Amount      *int   `json:"Amount,omitempty"`
		Code        string `json:"Code"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Description = temp.Description
	c.Amount = temp.Amount
	c.Code = temp.Code
	return nil
}
