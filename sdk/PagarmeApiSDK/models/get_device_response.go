package models

import (
	"encoding/json"
)

// GetDeviceResponse represents a GetDeviceResponse struct.
// Response object for geetting an order device
type GetDeviceResponse struct {
	// Device's platform name
	Platform Optional[string] `json:"platform"`
}

// MarshalJSON implements the json.Marshaler interface for GetDeviceResponse.
// It customizes the JSON marshaling process for GetDeviceResponse objects.
func (g *GetDeviceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetDeviceResponse object to a map representation for JSON marshaling.
func (g *GetDeviceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Platform.IsValueSet() {
		structMap["platform"] = g.Platform.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetDeviceResponse.
// It customizes the JSON unmarshaling process for GetDeviceResponse objects.
func (g *GetDeviceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Platform Optional[string] `json:"platform"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Platform = temp.Platform
	return nil
}
