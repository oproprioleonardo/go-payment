package models

import (
	"encoding/json"
)

// GetSetupResponse represents a GetSetupResponse struct.
// Response object for getting the setup from a subscription
type GetSetupResponse struct {
	Id          Optional[string] `json:"id"`
	Description Optional[string] `json:"description"`
	Amount      Optional[int]    `json:"amount"`
	Status      Optional[string] `json:"status"`
}

// MarshalJSON implements the json.Marshaler interface for GetSetupResponse.
// It customizes the JSON marshaling process for GetSetupResponse objects.
func (g *GetSetupResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetSetupResponse object to a map representation for JSON marshaling.
func (g *GetSetupResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSetupResponse.
// It customizes the JSON unmarshaling process for GetSetupResponse objects.
func (g *GetSetupResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id          Optional[string] `json:"id"`
		Description Optional[string] `json:"description"`
		Amount      Optional[int]    `json:"amount"`
		Status      Optional[string] `json:"status"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Description = temp.Description
	g.Amount = temp.Amount
	g.Status = temp.Status
	return nil
}
