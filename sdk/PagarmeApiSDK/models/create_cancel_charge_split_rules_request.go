package models

import (
	"encoding/json"
)

// CreateCancelChargeSplitRulesRequest represents a CreateCancelChargeSplitRulesRequest struct.
// Creates a refund with split rules
type CreateCancelChargeSplitRulesRequest struct {
	// The split rule gateway id
	Id string `json:"id"`
	// The split rule amount
	Amount int `json:"Amount"`
	// The amount type (flat ou percentage)
	Type string `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCancelChargeSplitRulesRequest.
// It customizes the JSON marshaling process for CreateCancelChargeSplitRulesRequest objects.
func (c *CreateCancelChargeSplitRulesRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCancelChargeSplitRulesRequest object to a map representation for JSON marshaling.
func (c *CreateCancelChargeSplitRulesRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["id"] = c.Id
	structMap["Amount"] = c.Amount
	structMap["type"] = c.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCancelChargeSplitRulesRequest.
// It customizes the JSON unmarshaling process for CreateCancelChargeSplitRulesRequest objects.
func (c *CreateCancelChargeSplitRulesRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id     string `json:"id"`
		Amount int    `json:"Amount"`
		Type   string `json:"type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Id = temp.Id
	c.Amount = temp.Amount
	c.Type = temp.Type
	return nil
}
