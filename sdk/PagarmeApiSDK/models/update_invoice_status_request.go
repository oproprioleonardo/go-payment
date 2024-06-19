package models

import (
	"encoding/json"
)

// UpdateInvoiceStatusRequest represents a UpdateInvoiceStatusRequest struct.
// Invoice Update Status Request
type UpdateInvoiceStatusRequest struct {
	// Status
	Status string `json:"status"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateInvoiceStatusRequest.
// It customizes the JSON marshaling process for UpdateInvoiceStatusRequest objects.
func (u *UpdateInvoiceStatusRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateInvoiceStatusRequest object to a map representation for JSON marshaling.
func (u *UpdateInvoiceStatusRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["status"] = u.Status
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateInvoiceStatusRequest.
// It customizes the JSON unmarshaling process for UpdateInvoiceStatusRequest objects.
func (u *UpdateInvoiceStatusRequest) UnmarshalJSON(input []byte) error {
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
