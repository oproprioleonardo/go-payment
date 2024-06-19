package models

import (
	"encoding/json"
)

// GetPricingSchemeResponse represents a GetPricingSchemeResponse struct.
// Response object for getting a pricing scheme
type GetPricingSchemeResponse struct {
	Price         Optional[int]                       `json:"price"`
	SchemeType    Optional[string]                    `json:"scheme_type"`
	PriceBrackets Optional[[]GetPriceBracketResponse] `json:"price_brackets"`
	MinimumPrice  Optional[int]                       `json:"minimum_price"`
	// percentual value used in pricing_scheme Percent
	Percentage Optional[float64] `json:"percentage"`
}

// MarshalJSON implements the json.Marshaler interface for GetPricingSchemeResponse.
// It customizes the JSON marshaling process for GetPricingSchemeResponse objects.
func (g *GetPricingSchemeResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetPricingSchemeResponse object to a map representation for JSON marshaling.
func (g *GetPricingSchemeResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Price.IsValueSet() {
		structMap["price"] = g.Price.Value()
	}
	if g.SchemeType.IsValueSet() {
		structMap["scheme_type"] = g.SchemeType.Value()
	}
	if g.PriceBrackets.IsValueSet() {
		structMap["price_brackets"] = g.PriceBrackets.Value()
	}
	if g.MinimumPrice.IsValueSet() {
		structMap["minimum_price"] = g.MinimumPrice.Value()
	}
	if g.Percentage.IsValueSet() {
		structMap["percentage"] = g.Percentage.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPricingSchemeResponse.
// It customizes the JSON unmarshaling process for GetPricingSchemeResponse objects.
func (g *GetPricingSchemeResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Price         Optional[int]                       `json:"price"`
		SchemeType    Optional[string]                    `json:"scheme_type"`
		PriceBrackets Optional[[]GetPriceBracketResponse] `json:"price_brackets"`
		MinimumPrice  Optional[int]                       `json:"minimum_price"`
		Percentage    Optional[float64]                   `json:"percentage"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Price = temp.Price
	g.SchemeType = temp.SchemeType
	g.PriceBrackets = temp.PriceBrackets
	g.MinimumPrice = temp.MinimumPrice
	g.Percentage = temp.Percentage
	return nil
}
