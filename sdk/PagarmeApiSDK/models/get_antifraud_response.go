package models

import (
	"encoding/json"
)

// GetAntifraudResponse represents a GetAntifraudResponse struct.
type GetAntifraudResponse struct {
	Status        Optional[string] `json:"status"`
	ReturnCode    Optional[string] `json:"return_code"`
	ReturnMessage Optional[string] `json:"return_message"`
	ProviderName  Optional[string] `json:"provider_name"`
	Score         Optional[string] `json:"score"`
}

// MarshalJSON implements the json.Marshaler interface for GetAntifraudResponse.
// It customizes the JSON marshaling process for GetAntifraudResponse objects.
func (g *GetAntifraudResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetAntifraudResponse object to a map representation for JSON marshaling.
func (g *GetAntifraudResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.ReturnCode.IsValueSet() {
		structMap["return_code"] = g.ReturnCode.Value()
	}
	if g.ReturnMessage.IsValueSet() {
		structMap["return_message"] = g.ReturnMessage.Value()
	}
	if g.ProviderName.IsValueSet() {
		structMap["provider_name"] = g.ProviderName.Value()
	}
	if g.Score.IsValueSet() {
		structMap["score"] = g.Score.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAntifraudResponse.
// It customizes the JSON unmarshaling process for GetAntifraudResponse objects.
func (g *GetAntifraudResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Status        Optional[string] `json:"status"`
		ReturnCode    Optional[string] `json:"return_code"`
		ReturnMessage Optional[string] `json:"return_message"`
		ProviderName  Optional[string] `json:"provider_name"`
		Score         Optional[string] `json:"score"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Status = temp.Status
	g.ReturnCode = temp.ReturnCode
	g.ReturnMessage = temp.ReturnMessage
	g.ProviderName = temp.ProviderName
	g.Score = temp.Score
	return nil
}
