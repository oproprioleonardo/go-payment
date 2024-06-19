package models

import (
	"encoding/json"
	"log"
	"time"
)

// UpdateCurrentCycleEndDateRequest represents a UpdateCurrentCycleEndDateRequest struct.
// Request to update the end date of the current subscription cycle
type UpdateCurrentCycleEndDateRequest struct {
	// Current cycle end date
	EndAt *time.Time `json:"end_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCurrentCycleEndDateRequest.
// It customizes the JSON marshaling process for UpdateCurrentCycleEndDateRequest objects.
func (u *UpdateCurrentCycleEndDateRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateCurrentCycleEndDateRequest object to a map representation for JSON marshaling.
func (u *UpdateCurrentCycleEndDateRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.EndAt != nil {
		structMap["end_at"] = u.EndAt.Format(time.RFC3339)
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCurrentCycleEndDateRequest.
// It customizes the JSON unmarshaling process for UpdateCurrentCycleEndDateRequest objects.
func (u *UpdateCurrentCycleEndDateRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		EndAt *string `json:"end_at,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	if temp.EndAt != nil {
		EndAtVal, err := time.Parse(time.RFC3339, *temp.EndAt)
		if err != nil {
			log.Fatalf("Cannot Parse end_at as % s format.", time.RFC3339)
		}
		u.EndAt = &EndAtVal
	}
	return nil
}
