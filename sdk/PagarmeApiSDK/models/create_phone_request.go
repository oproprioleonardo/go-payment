package models

import (
	"encoding/json"
)

// CreatePhoneRequest represents a CreatePhoneRequest struct.
type CreatePhoneRequest struct {
	CountryCode *string          `json:"country_code,omitempty"`
	Number      *string          `json:"number,omitempty"`
	AreaCode    *string          `json:"area_code,omitempty"`
	Type        Optional[string] `json:"Type"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePhoneRequest.
// It customizes the JSON marshaling process for CreatePhoneRequest objects.
func (c *CreatePhoneRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreatePhoneRequest object to a map representation for JSON marshaling.
func (c *CreatePhoneRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.CountryCode != nil {
		structMap["country_code"] = c.CountryCode
	}
	if c.Number != nil {
		structMap["number"] = c.Number
	}
	if c.AreaCode != nil {
		structMap["area_code"] = c.AreaCode
	}
	if c.Type.IsValueSet() {
		structMap["Type"] = c.Type.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePhoneRequest.
// It customizes the JSON unmarshaling process for CreatePhoneRequest objects.
func (c *CreatePhoneRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CountryCode *string          `json:"country_code,omitempty"`
		Number      *string          `json:"number,omitempty"`
		AreaCode    *string          `json:"area_code,omitempty"`
		Type        Optional[string] `json:"Type"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CountryCode = temp.CountryCode
	c.Number = temp.Number
	c.AreaCode = temp.AreaCode
	c.Type = temp.Type
	return nil
}
