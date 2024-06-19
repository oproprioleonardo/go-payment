package models

import (
	"encoding/json"
)

// CreateFineRequest represents a CreateFineRequest struct.
// Fine Request
type CreateFineRequest struct {
	// Days
	Days int `json:"days"`
	// Type
	Type string `json:"type"`
	// Amount
	Amount int `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for CreateFineRequest.
// It customizes the JSON marshaling process for CreateFineRequest objects.
func (c *CreateFineRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateFineRequest object to a map representation for JSON marshaling.
func (c *CreateFineRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["days"] = c.Days
	structMap["type"] = c.Type
	structMap["amount"] = c.Amount
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateFineRequest.
// It customizes the JSON unmarshaling process for CreateFineRequest objects.
func (c *CreateFineRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Days   int    `json:"days"`
		Type   string `json:"type"`
		Amount int    `json:"amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Days = temp.Days
	c.Type = temp.Type
	c.Amount = temp.Amount
	return nil
}
