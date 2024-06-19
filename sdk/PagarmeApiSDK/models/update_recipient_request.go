package models

import (
	"encoding/json"
)

// UpdateRecipientRequest represents a UpdateRecipientRequest struct.
// Request for updating a Recipient
type UpdateRecipientRequest struct {
	// Name
	Name string `json:"name"`
	// Email
	Email string `json:"email"`
	// Description
	Description string `json:"description"`
	// Type
	Type string `json:"type"`
	// Status
	Status string `json:"status"`
	// Metadata
	Metadata map[string]string `json:"metadata"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateRecipientRequest.
// It customizes the JSON marshaling process for UpdateRecipientRequest objects.
func (u *UpdateRecipientRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateRecipientRequest object to a map representation for JSON marshaling.
func (u *UpdateRecipientRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = u.Name
	structMap["email"] = u.Email
	structMap["description"] = u.Description
	structMap["type"] = u.Type
	structMap["status"] = u.Status
	structMap["metadata"] = u.Metadata
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateRecipientRequest.
// It customizes the JSON unmarshaling process for UpdateRecipientRequest objects.
func (u *UpdateRecipientRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name        string            `json:"name"`
		Email       string            `json:"email"`
		Description string            `json:"description"`
		Type        string            `json:"type"`
		Status      string            `json:"status"`
		Metadata    map[string]string `json:"metadata"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Name = temp.Name
	u.Email = temp.Email
	u.Description = temp.Description
	u.Type = temp.Type
	u.Status = temp.Status
	u.Metadata = temp.Metadata
	return nil
}
