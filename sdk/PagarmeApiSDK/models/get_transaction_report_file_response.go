package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetTransactionReportFileResponse represents a GetTransactionReportFileResponse struct.
type GetTransactionReportFileResponse struct {
	Name Optional[string]    `json:"name"`
	Date Optional[time.Time] `json:"date"`
}

// MarshalJSON implements the json.Marshaler interface for GetTransactionReportFileResponse.
// It customizes the JSON marshaling process for GetTransactionReportFileResponse objects.
func (g *GetTransactionReportFileResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetTransactionReportFileResponse object to a map representation for JSON marshaling.
func (g *GetTransactionReportFileResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.Date.IsValueSet() {
		var DateVal *string = nil
		if g.Date.Value() != nil {
			val := g.Date.Value().Format(time.RFC3339)
			DateVal = &val
		}
		structMap["date"] = DateVal
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetTransactionReportFileResponse.
// It customizes the JSON unmarshaling process for GetTransactionReportFileResponse objects.
func (g *GetTransactionReportFileResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name Optional[string] `json:"name"`
		Date Optional[string] `json:"date"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Name = temp.Name
	g.Date.ShouldSetValue(temp.Date.IsValueSet())
	if temp.Date.Value() != nil {
		DateVal, err := time.Parse(time.RFC3339, (*temp.Date.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse date as % s format.", time.RFC3339)
		}
		g.Date.SetValue(&DateVal)
	}
	return nil
}
