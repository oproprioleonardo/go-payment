package models

import (
	"encoding/json"
)

// CreateSubscriptionItemRequest represents a CreateSubscriptionItemRequest struct.
// Request for creating a new subscription item
type CreateSubscriptionItemRequest struct {
	// Item description
	Description string `json:"description"`
	// Pricing scheme
	PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
	// Item id
	Id string `json:"id"`
	// Plan item id
	PlanItemId string `json:"plan_item_id"`
	// Discounts for the item
	Discounts []CreateDiscountRequest `json:"discounts"`
	// Item name
	Name string `json:"name"`
	// Number of cycles which the item will be charged
	Cycles *int `json:"cycles,omitempty"`
	// Quantity of items
	Quantity *int `json:"quantity,omitempty"`
	// Minimum price
	MinimumPrice *int `json:"minimum_price,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscriptionItemRequest.
// It customizes the JSON marshaling process for CreateSubscriptionItemRequest objects.
func (c *CreateSubscriptionItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscriptionItemRequest object to a map representation for JSON marshaling.
func (c *CreateSubscriptionItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["description"] = c.Description
	structMap["pricing_scheme"] = c.PricingScheme
	structMap["id"] = c.Id
	structMap["plan_item_id"] = c.PlanItemId
	structMap["discounts"] = c.Discounts
	structMap["name"] = c.Name
	if c.Cycles != nil {
		structMap["cycles"] = c.Cycles
	}
	if c.Quantity != nil {
		structMap["quantity"] = c.Quantity
	}
	if c.MinimumPrice != nil {
		structMap["minimum_price"] = c.MinimumPrice
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscriptionItemRequest.
// It customizes the JSON unmarshaling process for CreateSubscriptionItemRequest objects.
func (c *CreateSubscriptionItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Description   string                     `json:"description"`
		PricingScheme CreatePricingSchemeRequest `json:"pricing_scheme"`
		Id            string                     `json:"id"`
		PlanItemId    string                     `json:"plan_item_id"`
		Discounts     []CreateDiscountRequest    `json:"discounts"`
		Name          string                     `json:"name"`
		Cycles        *int                       `json:"cycles,omitempty"`
		Quantity      *int                       `json:"quantity,omitempty"`
		MinimumPrice  *int                       `json:"minimum_price,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Description = temp.Description
	c.PricingScheme = temp.PricingScheme
	c.Id = temp.Id
	c.PlanItemId = temp.PlanItemId
	c.Discounts = temp.Discounts
	c.Name = temp.Name
	c.Cycles = temp.Cycles
	c.Quantity = temp.Quantity
	c.MinimumPrice = temp.MinimumPrice
	return nil
}
