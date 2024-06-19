package models

import (
	"encoding/json"
)

// UpdatePlanItemRequest represents a UpdatePlanItemRequest struct.
// Request for updating a plan item
type UpdatePlanItemRequest struct {
	// Item name
	Name string `json:"name"`
	// Description
	Description string `json:"description"`
	// Item status
	Status string `json:"status"`
	// Pricing scheme
	PricingScheme UpdatePricingSchemeRequest `json:"pricing_scheme"`
	// Quantity
	Quantity *int `json:"quantity,omitempty"`
	// Number of cycles that the item will be charged
	Cycles *int `json:"cycles,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePlanItemRequest.
// It customizes the JSON marshaling process for UpdatePlanItemRequest objects.
func (u *UpdatePlanItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdatePlanItemRequest object to a map representation for JSON marshaling.
func (u *UpdatePlanItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = u.Name
	structMap["description"] = u.Description
	structMap["status"] = u.Status
	structMap["pricing_scheme"] = u.PricingScheme
	if u.Quantity != nil {
		structMap["quantity"] = u.Quantity
	}
	if u.Cycles != nil {
		structMap["cycles"] = u.Cycles
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePlanItemRequest.
// It customizes the JSON unmarshaling process for UpdatePlanItemRequest objects.
func (u *UpdatePlanItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name          string                     `json:"name"`
		Description   string                     `json:"description"`
		Status        string                     `json:"status"`
		PricingScheme UpdatePricingSchemeRequest `json:"pricing_scheme"`
		Quantity      *int                       `json:"quantity,omitempty"`
		Cycles        *int                       `json:"cycles,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Name = temp.Name
	u.Description = temp.Description
	u.Status = temp.Status
	u.PricingScheme = temp.PricingScheme
	u.Quantity = temp.Quantity
	u.Cycles = temp.Cycles
	return nil
}
