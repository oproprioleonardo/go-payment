package models

import (
	"encoding/json"
)

// GetGatewayRecipientResponse represents a GetGatewayRecipientResponse struct.
// Information about the recipient on the gateway
type GetGatewayRecipientResponse struct {
	// Gateway name
	Gateway Optional[string] `json:"gateway"`
	// Status of the recipient on the gateway
	Status Optional[string] `json:"status"`
	// Recipient id on the gateway
	Pgid Optional[string] `json:"pgid"`
	// Creation date
	CreatedAt Optional[string] `json:"created_at"`
	// Last update date
	UpdatedAt Optional[string] `json:"updated_at"`
}

// MarshalJSON implements the json.Marshaler interface for GetGatewayRecipientResponse.
// It customizes the JSON marshaling process for GetGatewayRecipientResponse objects.
func (g *GetGatewayRecipientResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetGatewayRecipientResponse object to a map representation for JSON marshaling.
func (g *GetGatewayRecipientResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Gateway.IsValueSet() {
		structMap["gateway"] = g.Gateway.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Pgid.IsValueSet() {
		structMap["pgid"] = g.Pgid.Value()
	}
	if g.CreatedAt.IsValueSet() {
		structMap["created_at"] = g.CreatedAt.Value()
	}
	if g.UpdatedAt.IsValueSet() {
		structMap["updated_at"] = g.UpdatedAt.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetGatewayRecipientResponse.
// It customizes the JSON unmarshaling process for GetGatewayRecipientResponse objects.
func (g *GetGatewayRecipientResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Gateway   Optional[string] `json:"gateway"`
		Status    Optional[string] `json:"status"`
		Pgid      Optional[string] `json:"pgid"`
		CreatedAt Optional[string] `json:"created_at"`
		UpdatedAt Optional[string] `json:"updated_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Gateway = temp.Gateway
	g.Status = temp.Status
	g.Pgid = temp.Pgid
	g.CreatedAt = temp.CreatedAt
	g.UpdatedAt = temp.UpdatedAt
	return nil
}
