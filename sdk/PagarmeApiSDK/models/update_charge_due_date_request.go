package models

import (
	"encoding/json"
	"log"
	"time"
)

// UpdateChargeDueDateRequest represents a UpdateChargeDueDateRequest struct.
// Request for updating a charge due date
type UpdateChargeDueDateRequest struct {
	// The charge's new due date
	DueAt *time.Time `json:"due_at,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateChargeDueDateRequest.
// It customizes the JSON marshaling process for UpdateChargeDueDateRequest objects.
func (u *UpdateChargeDueDateRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateChargeDueDateRequest object to a map representation for JSON marshaling.
func (u *UpdateChargeDueDateRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.DueAt != nil {
		structMap["due_at"] = u.DueAt.Format(time.RFC3339)
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateChargeDueDateRequest.
// It customizes the JSON unmarshaling process for UpdateChargeDueDateRequest objects.
func (u *UpdateChargeDueDateRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		DueAt *string `json:"due_at,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	if temp.DueAt != nil {
		DueAtVal, err := time.Parse(time.RFC3339, *temp.DueAt)
		if err != nil {
			log.Fatalf("Cannot Parse due_at as % s format.", time.RFC3339)
		}
		u.DueAt = &DueAtVal
	}
	return nil
}
