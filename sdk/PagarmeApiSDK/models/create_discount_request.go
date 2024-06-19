package models

import (
	"encoding/json"
)

// CreateDiscountRequest represents a CreateDiscountRequest struct.
// Request for creating a new discount
type CreateDiscountRequest struct {
	// The discount value
	Value float64 `json:"value"`
	// Discount type. Can be either flat or percentage.
	DiscountType string `json:"discount_type"`
	// The item where the discount will be applied
	ItemId string `json:"item_id"`
	// Number of cycles that the discount will be applied
	Cycles *int `json:"cycles,omitempty"`
	// Description
	Description *string `json:"description,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateDiscountRequest.
// It customizes the JSON marshaling process for CreateDiscountRequest objects.
func (c *CreateDiscountRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateDiscountRequest object to a map representation for JSON marshaling.
func (c *CreateDiscountRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["value"] = c.Value
	structMap["discount_type"] = c.DiscountType
	structMap["item_id"] = c.ItemId
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.Description != nil {
		structMap["description"] = c.Description
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateDiscountRequest.
// It customizes the JSON unmarshaling process for CreateDiscountRequest objects.
func (c *CreateDiscountRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Value        float64 `json:"value"`
		DiscountType string  `json:"discount_type"`
		ItemId       string  `json:"item_id"`
		Cycles       *int    `json:"cycles,omitempty"`
		Description  *string `json:"description,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Value = temp.Value
	c.DiscountType = temp.DiscountType
	c.ItemId = temp.ItemId
	c.Cycles = temp.Cycles
	c.Description = temp.Description
	return nil
}
