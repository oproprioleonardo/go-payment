package models

import (
	"encoding/json"
)

// CreateCheckoutBankTransferRequest represents a CreateCheckoutBankTransferRequest struct.
// Checkout bank transfer payment request
type CreateCheckoutBankTransferRequest struct {
	// Bank
	Bank []string `json:"bank"`
	// Number of retries for processing
	Retries int `json:"retries"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCheckoutBankTransferRequest.
// It customizes the JSON marshaling process for CreateCheckoutBankTransferRequest objects.
func (c *CreateCheckoutBankTransferRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCheckoutBankTransferRequest object to a map representation for JSON marshaling.
func (c *CreateCheckoutBankTransferRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["bank"] = c.Bank
	structMap["retries"] = c.Retries
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCheckoutBankTransferRequest.
// It customizes the JSON unmarshaling process for CreateCheckoutBankTransferRequest objects.
func (c *CreateCheckoutBankTransferRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Bank    []string `json:"bank"`
		Retries int      `json:"retries"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Bank = temp.Bank
	c.Retries = temp.Retries
	return nil
}
