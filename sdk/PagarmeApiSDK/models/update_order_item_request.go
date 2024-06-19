package models

import (
	"encoding/json"
)

// UpdateOrderItemRequest represents a UpdateOrderItemRequest struct.
// Update Order item Request
type UpdateOrderItemRequest struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Category    string `json:"category"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateOrderItemRequest.
// It customizes the JSON marshaling process for UpdateOrderItemRequest objects.
func (u *UpdateOrderItemRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateOrderItemRequest object to a map representation for JSON marshaling.
func (u *UpdateOrderItemRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["amount"] = u.Amount
	structMap["description"] = u.Description
	structMap["quantity"] = u.Quantity
	structMap["category"] = u.Category
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateOrderItemRequest.
// It customizes the JSON unmarshaling process for UpdateOrderItemRequest objects.
func (u *UpdateOrderItemRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Amount      int    `json:"amount"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
		Category    string `json:"category"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Amount = temp.Amount
	u.Description = temp.Description
	u.Quantity = temp.Quantity
	u.Category = temp.Category
	return nil
}
