package models

import (
	"encoding/json"
)

// UpdateCurrentCycleStatusRequest represents a UpdateCurrentCycleStatusRequest struct.
type UpdateCurrentCycleStatusRequest struct {
	// Status
	Status string `json:"status"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCurrentCycleStatusRequest.
// It customizes the JSON marshaling process for UpdateCurrentCycleStatusRequest objects.
func (u *UpdateCurrentCycleStatusRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateCurrentCycleStatusRequest object to a map representation for JSON marshaling.
func (u *UpdateCurrentCycleStatusRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["status"] = u.Status
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCurrentCycleStatusRequest.
// It customizes the JSON unmarshaling process for UpdateCurrentCycleStatusRequest objects.
func (u *UpdateCurrentCycleStatusRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Status string `json:"status"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Status = temp.Status
	return nil
}
