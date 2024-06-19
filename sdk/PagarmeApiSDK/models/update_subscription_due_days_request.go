package models

import (
	"encoding/json"
)

// UpdateSubscriptionDueDaysRequest represents a UpdateSubscriptionDueDaysRequest struct.
type UpdateSubscriptionDueDaysRequest struct {
	BoletoDueDays int `json:"boleto_due_days"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionDueDaysRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionDueDaysRequest objects.
func (u *UpdateSubscriptionDueDaysRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionDueDaysRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionDueDaysRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["boleto_due_days"] = u.BoletoDueDays
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionDueDaysRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionDueDaysRequest objects.
func (u *UpdateSubscriptionDueDaysRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		BoletoDueDays int `json:"boleto_due_days"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.BoletoDueDays = temp.BoletoDueDays
	return nil
}
