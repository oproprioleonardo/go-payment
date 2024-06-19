package models

import (
	"encoding/json"
)

// UpdatePriceBracketRequest represents a UpdatePriceBracketRequest struct.
// Request for updating a price bracket
type UpdatePriceBracketRequest struct {
	// Start quantity of the bracket
	StartQuantity int `json:"start_quantity"`
	// Price
	Price int `json:"price"`
	// End quantity of the bracket
	EndQuantity *int `json:"end_quantity,omitempty"`
	// Overage price
	OveragePrice *int `json:"overage_price,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePriceBracketRequest.
// It customizes the JSON marshaling process for UpdatePriceBracketRequest objects.
func (u *UpdatePriceBracketRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdatePriceBracketRequest object to a map representation for JSON marshaling.
func (u *UpdatePriceBracketRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["start_quantity"] = u.StartQuantity
	structMap["price"] = u.Price
	if u.EndQuantity != nil {
		structMap["end_quantity"] = u.EndQuantity
	}
	if u.OveragePrice != nil {
		structMap["overage_price"] = u.OveragePrice
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePriceBracketRequest.
// It customizes the JSON unmarshaling process for UpdatePriceBracketRequest objects.
func (u *UpdatePriceBracketRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StartQuantity int  `json:"start_quantity"`
		Price         int  `json:"price"`
		EndQuantity   *int `json:"end_quantity,omitempty"`
		OveragePrice  *int `json:"overage_price,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.StartQuantity = temp.StartQuantity
	u.Price = temp.Price
	u.EndQuantity = temp.EndQuantity
	u.OveragePrice = temp.OveragePrice
	return nil
}
