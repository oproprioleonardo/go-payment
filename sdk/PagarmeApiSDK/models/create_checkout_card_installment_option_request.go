package models

import (
	"encoding/json"
)

// CreateCheckoutCardInstallmentOptionRequest represents a CreateCheckoutCardInstallmentOptionRequest struct.
// Options for card installment
type CreateCheckoutCardInstallmentOptionRequest struct {
	// Installment quantity
	Number int `json:"number"`
	// Total amount
	Total int `json:"total"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCheckoutCardInstallmentOptionRequest.
// It customizes the JSON marshaling process for CreateCheckoutCardInstallmentOptionRequest objects.
func (c *CreateCheckoutCardInstallmentOptionRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCheckoutCardInstallmentOptionRequest object to a map representation for JSON marshaling.
func (c *CreateCheckoutCardInstallmentOptionRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["number"] = c.Number
	structMap["total"] = c.Total
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCheckoutCardInstallmentOptionRequest.
// It customizes the JSON unmarshaling process for CreateCheckoutCardInstallmentOptionRequest objects.
func (c *CreateCheckoutCardInstallmentOptionRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Number int `json:"number"`
		Total  int `json:"total"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Number = temp.Number
	c.Total = temp.Total
	return nil
}
