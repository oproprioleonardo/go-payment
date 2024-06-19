package models

import (
	"encoding/json"
)

// UpdateOrderStatusRequest represents a UpdateOrderStatusRequest struct.
type UpdateOrderStatusRequest struct {
	// Order status
	Status string `json:"status"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateOrderStatusRequest.
// It customizes the JSON marshaling process for UpdateOrderStatusRequest objects.
func (u *UpdateOrderStatusRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateOrderStatusRequest object to a map representation for JSON marshaling.
func (u *UpdateOrderStatusRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["status"] = u.Status
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateOrderStatusRequest.
// It customizes the JSON unmarshaling process for UpdateOrderStatusRequest objects.
func (u *UpdateOrderStatusRequest) UnmarshalJSON(input []byte) error {
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
