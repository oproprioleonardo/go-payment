package models

import (
	"encoding/json"
	"log"
	"time"
)

// UpdateSubscriptionStartAtRequest represents a UpdateSubscriptionStartAtRequest struct.
// Request for updating the start date from a subscription
type UpdateSubscriptionStartAtRequest struct {
	// The date when the subscription periods will start
	StartAt time.Time `json:"start_at"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscriptionStartAtRequest.
// It customizes the JSON marshaling process for UpdateSubscriptionStartAtRequest objects.
func (u *UpdateSubscriptionStartAtRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscriptionStartAtRequest object to a map representation for JSON marshaling.
func (u *UpdateSubscriptionStartAtRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["start_at"] = u.StartAt.Format(time.RFC3339)
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscriptionStartAtRequest.
// It customizes the JSON unmarshaling process for UpdateSubscriptionStartAtRequest objects.
func (u *UpdateSubscriptionStartAtRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StartAt string `json:"start_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	StartAtVal, err := time.Parse(time.RFC3339, temp.StartAt)
	if err != nil {
		log.Fatalf("Cannot Parse start_at as % s format.", time.RFC3339)
	}
	u.StartAt = StartAtVal
	return nil
}
