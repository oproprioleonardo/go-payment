package models

import (
	"encoding/json"
)

// CreateCardOptionsRequest represents a CreateCardOptionsRequest struct.
// Options for creating the card
type CreateCardOptionsRequest struct {
	// Indicates if the card should be verified before creation. If true, executes an authorization before saving the card.
	VerifyCard bool `json:"verify_card"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCardOptionsRequest.
// It customizes the JSON marshaling process for CreateCardOptionsRequest objects.
func (c *CreateCardOptionsRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCardOptionsRequest object to a map representation for JSON marshaling.
func (c *CreateCardOptionsRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["verify_card"] = c.VerifyCard
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCardOptionsRequest.
// It customizes the JSON unmarshaling process for CreateCardOptionsRequest objects.
func (c *CreateCardOptionsRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		VerifyCard bool `json:"verify_card"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.VerifyCard = temp.VerifyCard
	return nil
}
