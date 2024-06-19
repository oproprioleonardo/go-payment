package models

import (
	"encoding/json"
)

// CreateCashPaymentRequest represents a CreateCashPaymentRequest struct.
type CreateCashPaymentRequest struct {
	// Description
	Description string `json:"description"`
	// Indicates whether cash collection will be confirmed in the act of creation
	Confirm bool `json:"confirm"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCashPaymentRequest.
// It customizes the JSON marshaling process for CreateCashPaymentRequest objects.
func (c *CreateCashPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCashPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateCashPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["description"] = c.Description
	structMap["confirm"] = c.Confirm
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCashPaymentRequest.
// It customizes the JSON unmarshaling process for CreateCashPaymentRequest objects.
func (c *CreateCashPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Description string `json:"description"`
		Confirm     bool   `json:"confirm"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Description = temp.Description
	c.Confirm = temp.Confirm
	return nil
}
