package models

import (
	"encoding/json"
)

// CreateEmvDataDecryptRequest represents a CreateEmvDataDecryptRequest struct.
type CreateEmvDataDecryptRequest struct {
	// Emv Decrypt cipher type
	Cipher string `json:"cipher"`
	// Dukpt data request
	Dukpt *CreateEmvDataDukptDecryptRequest `json:"dukpt,omitempty"`
	// Encrypted tags list
	Tags []CreateEmvDataTlvDecryptRequest `json:"tags"`
}

// MarshalJSON implements the json.Marshaler interface for CreateEmvDataDecryptRequest.
// It customizes the JSON marshaling process for CreateEmvDataDecryptRequest objects.
func (c *CreateEmvDataDecryptRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateEmvDataDecryptRequest object to a map representation for JSON marshaling.
func (c *CreateEmvDataDecryptRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["cipher"] = c.Cipher
	if c.Dukpt != nil {
		structMap["dukpt"] = c.Dukpt
	}
	structMap["tags"] = c.Tags
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateEmvDataDecryptRequest.
// It customizes the JSON unmarshaling process for CreateEmvDataDecryptRequest objects.
func (c *CreateEmvDataDecryptRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Cipher string                            `json:"cipher"`
		Dukpt  *CreateEmvDataDukptDecryptRequest `json:"dukpt,omitempty"`
		Tags   []CreateEmvDataTlvDecryptRequest  `json:"tags"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Cipher = temp.Cipher
	c.Dukpt = temp.Dukpt
	c.Tags = temp.Tags
	return nil
}
