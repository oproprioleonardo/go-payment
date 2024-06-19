package models

import (
	"encoding/json"
)

// GetGatewayResponseResponse represents a GetGatewayResponseResponse struct.
// The Transaction Gateway Response
type GetGatewayResponseResponse struct {
	// The error code
	Code Optional[string] `json:"code"`
	// The gateway response errors list
	Errors Optional[[]GetGatewayErrorResponse] `json:"errors"`
}

// MarshalJSON implements the json.Marshaler interface for GetGatewayResponseResponse.
// It customizes the JSON marshaling process for GetGatewayResponseResponse objects.
func (g *GetGatewayResponseResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetGatewayResponseResponse object to a map representation for JSON marshaling.
func (g *GetGatewayResponseResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.Errors.IsValueSet() {
		structMap["errors"] = g.Errors.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetGatewayResponseResponse.
// It customizes the JSON unmarshaling process for GetGatewayResponseResponse objects.
func (g *GetGatewayResponseResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Code   Optional[string]                    `json:"code"`
		Errors Optional[[]GetGatewayErrorResponse] `json:"errors"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Code = temp.Code
	g.Errors = temp.Errors
	return nil
}
