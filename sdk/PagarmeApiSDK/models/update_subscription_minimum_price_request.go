package models

import (
	"encoding/json"
)

// UpdateSubscriptionMinimumPriceRequest represents a UpdateSubscriptionMinimumPriceRequest struct.
// Atualização do valor mínimo da assinatura
type UpdateSubscriptionMinimumPriceRequest struct {
	// Valor mínimo da assinatura
	MinimumPrice *int `json:"minimum_price,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionMinimumPriceRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionMinimumPriceRequest objects.
func (u *UpdateSubscriptionMinimumPriceRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionMinimumPriceRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionMinimumPriceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.MinimumPrice != nil {
		structMap["minimum_price"] = u.MinimumPrice
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionMinimumPriceRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionMinimumPriceRequest objects.
func (u *UpdateSubscriptionMinimumPriceRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		MinimumPrice *int `json:"minimum_price,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.MinimumPrice = temp.MinimumPrice
	return nil
}
