package models

import (
	"encoding/json"
)

// UpdateSubscriptionCardRequest represents a UpdateSubscriptionCardRequest struct.
// Request for updating the card from a subscription
type UpdateSubscriptionCardRequest struct {
	// Credit card data
	Card CreateCardRequest `json:"card"`
	// Credit card id
	CardId string `json:"card_id"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionCardRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionCardRequest objects.
func (u *UpdateSubscriptionCardRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionCardRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionCardRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["card"] = u.Card
	structMap["card_id"] = u.CardId
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionCardRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionCardRequest objects.
func (u *UpdateSubscriptionCardRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Card   CreateCardRequest `json:"card"`
		CardId string            `json:"card_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Card = temp.Card
	u.CardId = temp.CardId
	return nil
}
