package models

import (
	"encoding/json"
)

// CreateEmvDataTlvDecryptRequest represents a CreateEmvDataTlvDecryptRequest struct.
type CreateEmvDataTlvDecryptRequest struct {
	// Emv tag
	Tag string `json:"tag"`
	// Emv lenght
	Lenght string `json:"lenght"`
	// Emv value
	Value string `json:"value"`
}

// MarshalJSON implements the json.Marshaler interface for CreateEmvDataTlvDecryptRequest.
// It customizes the JSON marshaling process for CreateEmvDataTlvDecryptRequest objects.
func (c *CreateEmvDataTlvDecryptRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateEmvDataTlvDecryptRequest object to a map representation for JSON marshaling.
func (c *CreateEmvDataTlvDecryptRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["tag"] = c.Tag
	structMap["lenght"] = c.Lenght
	structMap["value"] = c.Value
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateEmvDataTlvDecryptRequest.
// It customizes the JSON unmarshaling process for CreateEmvDataTlvDecryptRequest objects.
func (c *CreateEmvDataTlvDecryptRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Tag    string `json:"tag"`
		Lenght string `json:"lenght"`
		Value  string `json:"value"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Tag = temp.Tag
	c.Lenght = temp.Lenght
	c.Value = temp.Value
	return nil
}
