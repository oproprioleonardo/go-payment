package models

import (
	"encoding/json"
)

// GetPhoneResponse represents a GetPhoneResponse struct.
type GetPhoneResponse struct {
	CountryCode Optional[string] `json:"country_code"`
	Number      Optional[string] `json:"number"`
	AreaCode    Optional[string] `json:"area_code"`
}

// MarshalJSON implements the json.Marshaler interface for GetPhoneResponse.
// It customizes the JSON marshaling process for GetPhoneResponse objects.
func (g *GetPhoneResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetPhoneResponse object to a map representation for JSON marshaling.
func (g *GetPhoneResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.CountryCode.IsValueSet() {
		structMap["country_code"] = g.CountryCode.Value()
	}
	if g.Number.IsValueSet() {
		structMap["number"] = g.Number.Value()
	}
	if g.AreaCode.IsValueSet() {
		structMap["area_code"] = g.AreaCode.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPhoneResponse.
// It customizes the JSON unmarshaling process for GetPhoneResponse objects.
func (g *GetPhoneResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CountryCode Optional[string] `json:"country_code"`
		Number      Optional[string] `json:"number"`
		AreaCode    Optional[string] `json:"area_code"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.CountryCode = temp.CountryCode
	g.Number = temp.Number
	g.AreaCode = temp.AreaCode
	return nil
}
