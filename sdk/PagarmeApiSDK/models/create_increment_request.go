package models

import (
	"encoding/json"
)

// CreateIncrementRequest represents a CreateIncrementRequest struct.
// Request for creating a new increment
type CreateIncrementRequest struct {
	// The increment value
	Value float64 `json:"value"`
	// Increment type. Can be either flat or percentage.
	IncrementType string `json:"increment_type"`
	// The item where the increment will be applied
	ItemId string `json:"item_id"`
	// Number of cycles that the increment will be applied
	Cycles *int `json:"cycles,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateIncrementRequest.
// It customizes the JSON marshaling process for CreateIncrementRequest objects.
func (c *CreateIncrementRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateIncrementRequest object to a map representation for JSON marshaling.
func (c *CreateIncrementRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["value"] = c.Value
	structMap["increment_type"] = c.IncrementType
	structMap["item_id"] = c.ItemId
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.Description != nil {
		structMap["description"] = c.Description
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateIncrementRequest.
// It customizes the JSON unmarshaling process for CreateIncrementRequest objects.
func (c *CreateIncrementRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Value         float64 `json:"value"`
		IncrementType string  `json:"increment_type"`
		ItemId        string  `json:"item_id"`
		Cycles        *int    `json:"cycles,omitempty"`
		Description   *string `json:"description,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Value = temp.Value
	c.IncrementType = temp.IncrementType
	c.ItemId = temp.ItemId
	c.Cycles = temp.Cycles
	c.Description = temp.Description
	return nil
}
