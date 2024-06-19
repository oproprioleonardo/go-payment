package models

import (
	"encoding/json"
)

// CreateBankTransferPaymentRequest represents a CreateBankTransferPaymentRequest struct.
// Request for creating a bank transfer payment
type CreateBankTransferPaymentRequest struct {
	// Bank
	Bank string `json:"bank"`
	// Number of retries
	Retries int `json:"retries"`
}

// MarshalJSON implements the json.Marshaler interface for CreateBankTransferPaymentRequest.
// It customizes the JSON marshaling process for CreateBankTransferPaymentRequest objects.
func (c *CreateBankTransferPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateBankTransferPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateBankTransferPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["bank"] = c.Bank
	structMap["retries"] = c.Retries
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateBankTransferPaymentRequest.
// It customizes the JSON unmarshaling process for CreateBankTransferPaymentRequest objects.
func (c *CreateBankTransferPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Bank    string `json:"bank"`
		Retries int    `json:"retries"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Bank = temp.Bank
	c.Retries = temp.Retries
	return nil
}
