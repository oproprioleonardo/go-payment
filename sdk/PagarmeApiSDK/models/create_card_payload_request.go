package models

import (
	"encoding/json"
)

// CreateCardPayloadRequest represents a CreateCardPayloadRequest struct.
type CreateCardPayloadRequest struct {
	Type      Optional[string]                 `json:"type"`
	GooglePay Optional[CreateGooglePayRequest] `json:"google_pay"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCardPayloadRequest.
// It customizes the JSON marshaling process for CreateCardPayloadRequest objects.
func (c *CreateCardPayloadRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCardPayloadRequest object to a map representation for JSON marshaling.
func (c *CreateCardPayloadRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Type.IsValueSet() {
		structMap["type"] = c.Type.Value()
	}
	if c.GooglePay.IsValueSet() {
		structMap["google_pay"] = c.GooglePay.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCardPayloadRequest.
// It customizes the JSON unmarshaling process for CreateCardPayloadRequest objects.
func (c *CreateCardPayloadRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type      Optional[string]                 `json:"type"`
		GooglePay Optional[CreateGooglePayRequest] `json:"google_pay"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Type = temp.Type
	c.GooglePay = temp.GooglePay
	return nil
}
