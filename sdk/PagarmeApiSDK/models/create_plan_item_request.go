package models

import (
	"encoding/json"
)

// CreatePlanItemRequest represents a CreatePlanItemRequest struct.
// Request for creating a plan item
type CreatePlanItemRequest struct {
	// Item name
	Name string `json:"name"`
	// Item's pricing scheme
	PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
	// Item's id
	Id string `json:"id"`
	// Item's description
	Description string `json:"description"`
	// Number of cycles where the item will be charged
	Cycles *int `json:"cycles,omitempty"`
	// Quantity
	Quantity *int `json:"quantity,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePlanItemRequest.
// It customizes the JSON marshaling process for CreatePlanItemRequest objects.
func (c *CreatePlanItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePlanItemRequest object to a map representation for JSON marshaling.
func (c *CreatePlanItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = c.Name
	structMap["pricing_scheme"] = c.PricingScheme
	structMap["id"] = c.Id
	structMap["description"] = c.Description
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.Quantity != nil {
		structMap["quantity"] = c.Quantity
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePlanItemRequest.
// It customizes the JSON unmarshaling process for CreatePlanItemRequest objects.
func (c *CreatePlanItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name          string                     `json:"name"`
		PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
		Id            string                     `json:"id"`
		Description   string                     `json:"description"`
		Cycles        *int                       `json:"cycles,omitempty"`
		Quantity      *int                       `json:"quantity,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Name = temp.Name
	c.PricingScheme = temp.PricingScheme
	c.Id = temp.Id
	c.Description = temp.Description
	c.Cycles = temp.Cycles
	c.Quantity = temp.Quantity
	return nil
}
