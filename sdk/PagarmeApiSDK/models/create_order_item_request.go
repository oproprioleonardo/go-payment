package models

import (
	"encoding/json"
)

// CreateOrderItemRequest represents a CreateOrderItemRequest struct.
// Request for creating an order item
type CreateOrderItemRequest struct {
	// Amount
	Amount int `json:"amount"`
	// Description
	Description string `json:"description"`
	// Quantity
	Quantity int `json:"quantity"`
	// Category
	Category string `json:"category"`
	// The item code passed by the client
	Code *string `json:"code,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrderItemRequest.
// It customizes the JSON marshaling process for CreateOrderItemRequest objects.
func (c *CreateOrderItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateOrderItemRequest object to a map representation for JSON marshaling.
func (c *CreateOrderItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = c.Amount
	structMap["description"] = c.Description
	structMap["quantity"] = c.Quantity
	structMap["category"] = c.Category
	if c.Code != nil {
		structMap["code"] = c.Code
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrderItemRequest.
// It customizes the JSON unmarshaling process for CreateOrderItemRequest objects.
func (c *CreateOrderItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount      int     `json:"amount"`
		Description string  `json:"description"`
		Quantity    int     `json:"quantity"`
		Category    string  `json:"category"`
		Code        *string `json:"code,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Amount = temp.Amount
	c.Description = temp.Description
	c.Quantity = temp.Quantity
	c.Category = temp.Category
	c.Code = temp.Code
	return nil
}
