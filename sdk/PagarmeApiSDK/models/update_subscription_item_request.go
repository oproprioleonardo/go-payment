package models

import (
	"encoding/json"
)

// UpdateSubscriptionItemRequest represents a UpdateSubscriptionItemRequest struct.
// Request for updating a subscription item
type UpdateSubscriptionItemRequest struct {
	// Description
	Description string `json:"description"`
	// Status
	Status string `json:"status"`
	// Pricing scheme
	PricingScheme UpdatePricingSchemeRequest `json:"pricing_scheme"`
	// Item name
	Name string `json:"name"`
	// Number of cycles that the item will be charged
	Cycles *int `json:"cycles,omitempty"`
	// Quantity
	Quantity *int `json:"quantity,omitempty"`
	// Minimum price
	MinimumPrice *int `json:"minimum_price,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionItemRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionItemRequest objects.
func (u *UpdateSubscriptionItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionItemRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["description"] = u.Description
	structMap["status"] = u.Status
	structMap["pricing_scheme"] = u.PricingScheme
	structMap["name"] = u.Name
	if u.Cycles != nil {
		structMap["cycles"] = u.Cycles
	}
	if u.Quantity != nil {
		structMap["quantity"] = u.Quantity
	}
	if u.MinimumPrice != nil {
		structMap["minimum_price"] = u.MinimumPrice
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionItemRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionItemRequest objects.
func (u *UpdateSubscriptionItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Description   string                     `json:"description"`
		Status        string                     `json:"status"`
		PricingScheme UpdatePricingSchemeRequest `json:"pricing_scheme"`
		Name          string                     `json:"name"`
		Cycles        *int                       `json:"cycles,omitempty"`
		Quantity      *int                       `json:"quantity,omitempty"`
		MinimumPrice  *int                       `json:"minimum_price,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Description = temp.Description
	u.Status = temp.Status
	u.PricingScheme = temp.PricingScheme
	u.Name = temp.Name
	u.Cycles = temp.Cycles
	u.Quantity = temp.Quantity
	u.MinimumPrice = temp.MinimumPrice
	return nil
}
