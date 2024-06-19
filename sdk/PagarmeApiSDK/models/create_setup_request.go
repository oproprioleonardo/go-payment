package models

import (
	"encoding/json"
)

// CreateSetupRequest represents a CreateSetupRequest struct.
// Request for creating a Setup for a subscription. The setup is an order that will be created at the subscription creation.
type CreateSetupRequest struct {
	// Setup amount
	Amount int `json:"amount"`
	// Description
	Description string `json:"description"`
	// Payment data
	Payment CreatePaymentRequest `json:"payment"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSetupRequest.
// It customizes the JSON marshaling process for CreateSetupRequest objects.
func (c *CreateSetupRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSetupRequest object to a map representation for JSON marshaling.
func (c *CreateSetupRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	structMap["description"] = c.Description
	structMap["payment"] = c.Payment
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSetupRequest.
// It customizes the JSON unmarshaling process for CreateSetupRequest objects.
func (c *CreateSetupRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount      int                  `json:"amount"`
		Description string               `json:"description"`
		Payment     CreatePaymentRequest `json:"payment"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Description = temp.Description
	c.Payment = temp.Payment
	return nil
}
