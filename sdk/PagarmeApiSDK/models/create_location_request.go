package models

import (
	"encoding/json"
)

// CreateLocationRequest represents a CreateLocationRequest struct.
// Request for creating a location
type CreateLocationRequest struct {
	// Latitude
	Latitude string `json:"latitude"`
	// Longitude
	Longitude string `json:"longitude"`
}

// MarshalJSON implements the json.Marshaler interface for CreateLocationRequest.
// It customizes the JSON marshaling process for CreateLocationRequest objects.
func (c *CreateLocationRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateLocationRequest object to a map representation for JSON marshaling.
func (c *CreateLocationRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["latitude"] = c.Latitude
	structMap["longitude"] = c.Longitude
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateLocationRequest.
// It customizes the JSON unmarshaling process for CreateLocationRequest objects.
func (c *CreateLocationRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Latitude = temp.Latitude
	c.Longitude = temp.Longitude
	return nil
}
