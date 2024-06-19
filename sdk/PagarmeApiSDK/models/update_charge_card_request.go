package models

import (
	"encoding/json"
)

// UpdateChargeCardRequest represents a UpdateChargeCardRequest struct.
// Request for updating card data
type UpdateChargeCardRequest struct {
	// Indicates if the subscriptions using this card must also be updated
	UpdateSubscription bool `json:"update_subscription"`
	// Card id
	CardId string `json:"card_id"`
	// Card data
	Card CreateCardRequest `json:"card"`
	// Indicates a recurrence
	Recurrence bool `json:"recurrence"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateChargeCardRequest.
// It customizes the JSON marshaling process for UpdateChargeCardRequest objects.
func (u *UpdateChargeCardRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateChargeCardRequest object to a map representation for JSON marshaling.
func (u *UpdateChargeCardRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["update_subscription"] = u.UpdateSubscription
	structMap["card_id"] = u.CardId
	structMap["card"] = u.Card
	structMap["recurrence"] = u.Recurrence
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateChargeCardRequest.
// It customizes the JSON unmarshaling process for UpdateChargeCardRequest objects.
func (u *UpdateChargeCardRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		UpdateSubscription bool              `json:"update_subscription"`
		CardId             string            `json:"card_id"`
		Card               CreateCardRequest `json:"card"`
		Recurrence         bool              `json:"recurrence"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.UpdateSubscription = temp.UpdateSubscription
	u.CardId = temp.CardId
	u.Card = temp.Card
	u.Recurrence = temp.Recurrence
	return nil
}
