package models

import (
	"encoding/json"
)

// CreateEmvDataDukptDecryptRequest represents a CreateEmvDataDukptDecryptRequest struct.
type CreateEmvDataDukptDecryptRequest struct {
	// Key serial number
	Ksn string `json:"ksn"`
}

// MarshalJSON implements the json.Marshaler interface for CreateEmvDataDukptDecryptRequest.
// It customizes the JSON marshaling process for CreateEmvDataDukptDecryptRequest objects.
func (c *CreateEmvDataDukptDecryptRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateEmvDataDukptDecryptRequest object to a map representation for JSON marshaling.
func (c *CreateEmvDataDukptDecryptRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["ksn"] = c.Ksn
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateEmvDataDukptDecryptRequest.
// It customizes the JSON unmarshaling process for CreateEmvDataDukptDecryptRequest objects.
func (c *CreateEmvDataDukptDecryptRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Ksn string `json:"ksn"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Ksn = temp.Ksn
	return nil
}
