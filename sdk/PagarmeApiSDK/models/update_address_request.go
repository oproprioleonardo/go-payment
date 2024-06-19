package models

import (
	"encoding/json"
)

// UpdateAddressRequest represents a UpdateAddressRequest struct.
// Request for updating an address
type UpdateAddressRequest struct {
	// Number
	Number string `json:"number"`
	// Complement
	Complement string `json:"complement"`
	// Metadata
	Metadata map[string]string `json:"metadata"`
	// Line 2 for address
	Line2 string `json:"line_2"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateAddressRequest.
// It customizes the JSON marshaling process for UpdateAddressRequest objects.
func (u *UpdateAddressRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateAddressRequest object to a map representation for JSON marshaling.
func (u *UpdateAddressRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["number"] = u.Number
	structMap["complement"] = u.Complement
	structMap["metadata"] = u.Metadata
	structMap["line_2"] = u.Line2
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateAddressRequest.
// It customizes the JSON unmarshaling process for UpdateAddressRequest objects.
func (u *UpdateAddressRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Number     string            `json:"number"`
		Complement string            `json:"complement"`
		Metadata   map[string]string `json:"metadata"`
		Line2      string            `json:"line_2"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Number = temp.Number
	u.Complement = temp.Complement
	u.Metadata = temp.Metadata
	u.Line2 = temp.Line2
	return nil
}
