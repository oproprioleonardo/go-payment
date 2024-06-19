package models

import (
	"encoding/json"
)

// CreateDeviceRequest represents a CreateDeviceRequest struct.
// Request for creating a device
type CreateDeviceRequest struct {
	// Device's platform
	Platform *string `json:"platform,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateDeviceRequest.
// It customizes the JSON marshaling process for CreateDeviceRequest objects.
func (c *CreateDeviceRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateDeviceRequest object to a map representation for JSON marshaling.
func (c *CreateDeviceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Platform != nil {
		structMap["platform"] = c.Platform
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateDeviceRequest.
// It customizes the JSON unmarshaling process for CreateDeviceRequest objects.
func (c *CreateDeviceRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Platform *string `json:"platform,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Platform = temp.Platform
	return nil
}
