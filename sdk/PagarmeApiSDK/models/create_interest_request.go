package models

import (
	"encoding/json"
)

// CreateInterestRequest represents a CreateInterestRequest struct.
// Interest Request
type CreateInterestRequest struct {
	// Days
	Days int `json:"days"`
	// Type
	Type string `json:"type"`
	// Amount
	Amount int `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInterestRequest.
// It customizes the JSON marshaling process for CreateInterestRequest objects.
func (c *CreateInterestRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateInterestRequest object to a map representation for JSON marshaling.
func (c *CreateInterestRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["days"] = c.Days
	structMap["type"] = c.Type
	structMap["amount"] = c.Amount
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInterestRequest.
// It customizes the JSON unmarshaling process for CreateInterestRequest objects.
func (c *CreateInterestRequest) UnmarshalJSON(input []byte) error {
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
