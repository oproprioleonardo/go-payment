package models

import (
	"encoding/json"
)

// GetBillingAddressResponse represents a GetBillingAddressResponse struct.
// Response object for getting a billing address
type GetBillingAddressResponse struct {
	Street       Optional[string] `json:"street"`
	Number       Optional[string] `json:"number"`
	ZipCode      Optional[string] `json:"zip_code"`
	Neighborhood Optional[string] `json:"neighborhood"`
	City         Optional[string] `json:"city"`
	State        Optional[string] `json:"state"`
	Country      Optional[string] `json:"country"`
	Complement   Optional[string] `json:"complement"`
	// Line 1 for address
	Line1 Optional[string] `json:"line_1"`
	// Line 2 for address
	Line2 Optional[string] `json:"line_2"`
}

// MarshalJSON implements the json.Marshaler interface for GetBillingAddressResponse.
// It customizes the JSON marshaling process for GetBillingAddressResponse objects.
func (g *GetBillingAddressResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetBillingAddressResponse object to a map representation for JSON marshaling.
func (g *GetBillingAddressResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Street.IsValueSet() {
		structMap["street"] = g.Street.Value()
	}
	if g.Number.IsValueSet() {
		structMap["number"] = g.Number.Value()
	}
	if g.ZipCode.IsValueSet() {
		structMap["zip_code"] = g.ZipCode.Value()
	}
	if g.Neighborhood.IsValueSet() {
		structMap["neighborhood"] = g.Neighborhood.Value()
	}
	if g.City.IsValueSet() {
		structMap["city"] = g.City.Value()
	}
	if g.State.IsValueSet() {
		structMap["state"] = g.State.Value()
	}
	if g.Country.IsValueSet() {
		structMap["country"] = g.Country.Value()
	}
	if g.Complement.IsValueSet() {
		structMap["complement"] = g.Complement.Value()
	}
	if g.Line1.IsValueSet() {
		structMap["line_1"] = g.Line1.Value()
	}
	if g.Line2.IsValueSet() {
		structMap["line_2"] = g.Line2.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBillingAddressResponse.
// It customizes the JSON unmarshaling process for GetBillingAddressResponse objects.
func (g *GetBillingAddressResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Street       Optional[string] `json:"street"`
		Number       Optional[string] `json:"number"`
		ZipCode      Optional[string] `json:"zip_code"`
		Neighborhood Optional[string] `json:"neighborhood"`
		City         Optional[string] `json:"city"`
		State        Optional[string] `json:"state"`
		Country      Optional[string] `json:"country"`
		Complement   Optional[string] `json:"complement"`
		Line1        Optional[string] `json:"line_1"`
		Line2        Optional[string] `json:"line_2"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Street = temp.Street
	g.Number = temp.Number
	g.ZipCode = temp.ZipCode
	g.Neighborhood = temp.Neighborhood
	g.City = temp.City
	g.State = temp.State
	g.Country = temp.Country
	g.Complement = temp.Complement
	g.Line1 = temp.Line1
	g.Line2 = temp.Line2
	return nil
}
