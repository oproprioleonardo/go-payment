package models

import (
	"encoding/json"
)

// CreateAddressRequest represents a CreateAddressRequest struct.
// Request for creating a new Address
type CreateAddressRequest struct {
	// Street
	Street string `json:"street"`
	// Number
	Number string `json:"number"`
	// The zip code containing only numbers. No special characters or spaces.
	ZipCode string `json:"zip_code"`
	// Neighborhood
	Neighborhood string `json:"neighborhood"`
	// City
	City string `json:"city"`
	// State
	State string `json:"state"`
	// Country. Must be entered using ISO 3166-1 alpha-2 format. See https://pt.wikipedia.org/wiki/ISO_3166-1_alfa-2
	Country string `json:"country"`
	// Complement
	Complement string `json:"complement"`
	// Metadata
	Metadata Optional[map[string]string] `json:"metadata"`
	// Line 1 for address
	Line1 string `json:"line_1"`
	// Line 2 for address
	Line2 string `json:"line_2"`
}

// MarshalJSON implements the json.Marshaler interface for CreateAddressRequest.
// It customizes the JSON marshaling process for CreateAddressRequest objects.
func (c *CreateAddressRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateAddressRequest object to a map representation for JSON marshaling.
func (c *CreateAddressRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["street"] = c.Street
	structMap["number"] = c.Number
	structMap["zip_code"] = c.ZipCode
	structMap["neighborhood"] = c.Neighborhood
	structMap["city"] = c.City
	structMap["state"] = c.State
	structMap["country"] = c.Country
	structMap["complement"] = c.Complement
	if c.Metadata.IsValueSet() {
		structMap["metadata"] = c.Metadata.Value()
	}
	structMap["line_1"] = c.Line1
	structMap["line_2"] = c.Line2
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateAddressRequest.
// It customizes the JSON unmarshaling process for CreateAddressRequest objects.
func (c *CreateAddressRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Street       string                      `json:"street"`
		Number       string                      `json:"number"`
		ZipCode      string                      `json:"zip_code"`
		Neighborhood string                      `json:"neighborhood"`
		City         string                      `json:"city"`
		State        string                      `json:"state"`
		Country      string                      `json:"country"`
		Complement   string                      `json:"complement"`
		Metadata     Optional[map[string]string] `json:"metadata"`
		Line1        string                      `json:"line_1"`
		Line2        string                      `json:"line_2"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Street = temp.Street
	c.Number = temp.Number
	c.ZipCode = temp.ZipCode
	c.Neighborhood = temp.Neighborhood
	c.City = temp.City
	c.State = temp.State
	c.Country = temp.Country
	c.Complement = temp.Complement
	c.Metadata = temp.Metadata
	c.Line1 = temp.Line1
	c.Line2 = temp.Line2
	return nil
}
