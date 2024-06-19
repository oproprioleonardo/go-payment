package models

import (
	"encoding/json"
)

// CreatePricingSchemeRequest represents a CreatePricingSchemeRequest struct.
// Request for creating a pricing scheme
type CreatePricingSchemeRequest struct {
	// Scheme type
	SchemeType string `json:"scheme_type"`
	// Price brackets
	PriceBrackets []CreatePriceBracketRequest `json:"price_brackets,omitempty"`
	// Price
	Price *int `json:"price,omitempty"`
	// Minimum price
	MinimumPrice *int `json:"minimum_price,omitempty"`
	// percentual value used in pricing_scheme Percent
	Percentage *float64 `json:"percentage,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePricingSchemeRequest.
// It customizes the JSON marshaling process for CreatePricingSchemeRequest objects.
func (c *CreatePricingSchemeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePricingSchemeRequest object to a map representation for JSON marshaling.
func (c *CreatePricingSchemeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["scheme_type"] = c.SchemeType
	if c.PriceBrackets != nil {
		structMap["price_brackets"] = c.PriceBrackets
	}
	if c.Price != nil {
		structMap["price"] = c.Price
	}
	if c.MinimumPrice != nil {
		structMap["minimum_price"] = c.MinimumPrice
	}
	if c.Percentage != nil {
		structMap["percentage"] = c.Percentage
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePricingSchemeRequest.
// It customizes the JSON unmarshaling process for CreatePricingSchemeRequest objects.
func (c *CreatePricingSchemeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SchemeType    string                      `json:"scheme_type"`
		PriceBrackets []CreatePriceBracketRequest `json:"price_brackets,omitempty"`
		Price         *int                        `json:"price,omitempty"`
		MinimumPrice  *int                        `json:"minimum_price,omitempty"`
		Percentage    *float64                    `json:"percentage,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.SchemeType = temp.SchemeType
	c.PriceBrackets = temp.PriceBrackets
	c.Price = temp.Price
	c.MinimumPrice = temp.MinimumPrice
	c.Percentage = temp.Percentage
	return nil
}
