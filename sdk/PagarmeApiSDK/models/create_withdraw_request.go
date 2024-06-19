package models

import (
	"encoding/json"
)

// CreateWithdrawRequest represents a CreateWithdrawRequest struct.
type CreateWithdrawRequest struct {
	Amount   int               `json:"amount"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateWithdrawRequest.
// It customizes the JSON marshaling process for CreateWithdrawRequest objects.
func (c *CreateWithdrawRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateWithdrawRequest object to a map representation for JSON marshaling.
func (c *CreateWithdrawRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	if c.Metadata != nil {
		structMap["metadata"] = c.Metadata
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateWithdrawRequest.
// It customizes the JSON unmarshaling process for CreateWithdrawRequest objects.
func (c *CreateWithdrawRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount   int               `json:"amount"`
		Metadata map[string]string `json:"metadata,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Metadata = temp.Metadata
	return nil
}
