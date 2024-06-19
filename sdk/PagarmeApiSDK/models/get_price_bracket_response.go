package models

import (
	"encoding/json"
)

// GetPriceBracketResponse represents a GetPriceBracketResponse struct.
// Response object for getting a price bracket
type GetPriceBracketResponse struct {
	StartQuantity Optional[int] `json:"start_quantity"`
	Price         Optional[int] `json:"price"`
	EndQuantity   Optional[int] `json:"end_quantity"`
	OveragePrice  Optional[int] `json:"overage_price"`
}

// MarshalJSON implements the json.Marshaler interface for GetPriceBracketResponse.
// It customizes the JSON marshaling process for GetPriceBracketResponse objects.
func (g *GetPriceBracketResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetPriceBracketResponse object to a map representation for JSON marshaling.
func (g *GetPriceBracketResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.StartQuantity.IsValueSet() {
		structMap["start_quantity"] = g.StartQuantity.Value()
	}
	if g.Price.IsValueSet() {
		structMap["price"] = g.Price.Value()
	}
	if g.EndQuantity.IsValueSet() {
		structMap["end_quantity"] = g.EndQuantity.Value()
	}
	if g.OveragePrice.IsValueSet() {
		structMap["overage_price"] = g.OveragePrice.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPriceBracketResponse.
// It customizes the JSON unmarshaling process for GetPriceBracketResponse objects.
func (g *GetPriceBracketResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StartQuantity Optional[int] `json:"start_quantity"`
		Price         Optional[int] `json:"price"`
		EndQuantity   Optional[int] `json:"end_quantity"`
		OveragePrice  Optional[int] `json:"overage_price"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.StartQuantity = temp.StartQuantity
	g.Price = temp.Price
	g.EndQuantity = temp.EndQuantity
	g.OveragePrice = temp.OveragePrice
	return nil
}
