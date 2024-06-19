package models

import (
	"encoding/json"
)

// CreateTransferRequest represents a CreateTransferRequest struct.
// Request for creating a transfer
type CreateTransferRequest struct {
	// Transfer amount
	Amount int `json:"amount"`
	// Metadata
	Metadata map[string]string `json:"metadata"`
}

// MarshalJSON implements the json.Marshaler interface for CreateTransferRequest.
// It customizes the JSON marshaling process for CreateTransferRequest objects.
func (c *CreateTransferRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateTransferRequest object to a map representation for JSON marshaling.
func (c *CreateTransferRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	structMap["metadata"] = c.Metadata
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateTransferRequest.
// It customizes the JSON unmarshaling process for CreateTransferRequest objects.
func (c *CreateTransferRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount   int               `json:"amount"`
		Metadata map[string]string `json:"metadata"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Metadata = temp.Metadata
	return nil
}
