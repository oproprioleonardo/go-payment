package models

import (
	"encoding/json"
)

// GetPhonesResponse represents a GetPhonesResponse struct.
type GetPhonesResponse struct {
	HomePhone   Optional[GetPhoneResponse] `json:"home_phone"`
	MobilePhone Optional[GetPhoneResponse] `json:"mobile_phone"`
}

// MarshalJSON implements the json.Marshaler interface for GetPhonesResponse.
// It customizes the JSON marshaling process for GetPhonesResponse objects.
func (g *GetPhonesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetPhonesResponse object to a map representation for JSON marshaling.
func (g *GetPhonesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.HomePhone.IsValueSet() {
		structMap["home_phone"] = g.HomePhone.Value()
	}
	if g.MobilePhone.IsValueSet() {
		structMap["mobile_phone"] = g.MobilePhone.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPhonesResponse.
// It customizes the JSON unmarshaling process for GetPhonesResponse objects.
func (g *GetPhonesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		HomePhone   Optional[GetPhoneResponse] `json:"home_phone"`
		MobilePhone Optional[GetPhoneResponse] `json:"mobile_phone"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.HomePhone = temp.HomePhone
	g.MobilePhone = temp.MobilePhone
	return nil
}
