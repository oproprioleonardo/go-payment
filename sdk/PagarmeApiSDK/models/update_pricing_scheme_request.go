package models

import (
	"encoding/json"
)

// UpdatePricingSchemeRequest represents a UpdatePricingSchemeRequest struct.
// Request for updating a pricing scheme
type UpdatePricingSchemeRequest struct {
	// Scheme type
	SchemeType string `json:"scheme_type"`
	// Price brackets
	PriceBrackets []UpdatePriceBracketRequest `json:"price_brackets"`
	// Price
	Price *int `json:"price,omitempty"`
	// Minimum price
	MinimumPrice *int `json:"minimum_price,omitempty"`
	// percentual value used in pricing_scheme Percent
	Percentage *float64 `json:"percentage,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePricingSchemeRequest.
// It customizes the JSON marshaling process for UpdatePricingSchemeRequest objects.
func (u *UpdatePricingSchemeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdatePricingSchemeRequest object to a map representation for JSON marshaling.
func (u *UpdatePricingSchemeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["scheme_type"] = u.SchemeType
	structMap["price_brackets"] = u.PriceBrackets
	if u.Price != nil {
		structMap["price"] = u.Price
	}
	if u.MinimumPrice != nil {
		structMap["minimum_price"] = u.MinimumPrice
	}
	if u.Percentage != nil {
		structMap["percentage"] = u.Percentage
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePricingSchemeRequest.
// It customizes the JSON unmarshaling process for UpdatePricingSchemeRequest objects.
func (u *UpdatePricingSchemeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SchemeType    string                      `json:"scheme_type"`
		PriceBrackets []UpdatePriceBracketRequest `json:"price_brackets"`
		Price         *int                        `json:"price,omitempty"`
		MinimumPrice  *int                        `json:"minimum_price,omitempty"`
		Percentage    *float64                    `json:"percentage,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.SchemeType = temp.SchemeType
	u.PriceBrackets = temp.PriceBrackets
	u.Price = temp.Price
	u.MinimumPrice = temp.MinimumPrice
	u.Percentage = temp.Percentage
	return nil
}
