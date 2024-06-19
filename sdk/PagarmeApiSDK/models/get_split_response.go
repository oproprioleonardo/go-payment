package models

import (
	"encoding/json"
)

// GetSplitResponse represents a GetSplitResponse struct.
// Split response
type GetSplitResponse struct {
	// Type
	Type Optional[string] `json:"type"`
	// Amount
	Amount Optional[int] `json:"amount"`
	// Recipient
	Recipient Optional[GetRecipientResponse] `json:"recipient"`
	// The split rule gateway id
	GatewayId Optional[string]                  `json:"gateway_id"`
	Options   Optional[GetSplitOptionsResponse] `json:"options"`
	Id        Optional[string]                  `json:"id"`
}

// MarshalJSON implements the json.Marshaler interface for GetSplitResponse.
// It customizes the JSON marshaling process for GetSplitResponse objects.
func (g *GetSplitResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetSplitResponse object to a map representation for JSON marshaling.
func (g *GetSplitResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Recipient.IsValueSet() {
		structMap["recipient"] = g.Recipient.Value()
	}
	if g.GatewayId.IsValueSet() {
		structMap["gateway_id"] = g.GatewayId.Value()
	}
	if g.Options.IsValueSet() {
		structMap["options"] = g.Options.Value()
	}
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSplitResponse.
// It customizes the JSON unmarshaling process for GetSplitResponse objects.
func (g *GetSplitResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type      Optional[string]                  `json:"type"`
		Amount    Optional[int]                     `json:"amount"`
		Recipient Optional[GetRecipientResponse]    `json:"recipient"`
		GatewayId Optional[string]                  `json:"gateway_id"`
		Options   Optional[GetSplitOptionsResponse] `json:"options"`
		Id        Optional[string]                  `json:"id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Type = temp.Type
	g.Amount = temp.Amount
	g.Recipient = temp.Recipient
	g.GatewayId = temp.GatewayId
	g.Options = temp.Options
	g.Id = temp.Id
	return nil
}
