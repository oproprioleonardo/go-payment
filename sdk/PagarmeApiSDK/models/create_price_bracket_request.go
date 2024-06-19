package models

import (
	"encoding/json"
)

// CreatePriceBracketRequest represents a CreatePriceBracketRequest struct.
// Request for creating a price bracket
type CreatePriceBracketRequest struct {
	// Start quantity
	StartQuantity int `json:"start_quantity"`
	// Price
	Price int `json:"price"`
	// End quantity
	EndQuantity *int `json:"end_quantity,omitempty"`
	// Overage price
	OveragePrice *int `json:"overage_price,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePriceBracketRequest.
// It customizes the JSON marshaling process for CreatePriceBracketRequest objects.
func (c *CreatePriceBracketRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePriceBracketRequest object to a map representation for JSON marshaling.
func (c *CreatePriceBracketRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["start_quantity"] = c.StartQuantity
	structMap["price"] = c.Price
	if c.EndQuantity != nil {
		structMap["end_quantity"] = c.EndQuantity
	}
	if c.OveragePrice != nil {
		structMap["overage_price"] = c.OveragePrice
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePriceBracketRequest.
// It customizes the JSON unmarshaling process for CreatePriceBracketRequest objects.
func (c *CreatePriceBracketRequest) UnmarshalJSON(input []byte) error {
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

	c.StartQuantity = temp.StartQuantity
	c.Price = temp.Price
	c.EndQuantity = temp.EndQuantity
	c.OveragePrice = temp.OveragePrice
	return nil
}
