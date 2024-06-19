package models

import (
	"encoding/json"
)

// GetUsageReportResponse represents a GetUsageReportResponse struct.
type GetUsageReportResponse struct {
	Url              Optional[string] `json:"url"`
	UsageReportUrl   Optional[string] `json:"usage_report_url"`
	GroupedReportUrl Optional[string] `json:"grouped_report_url"`
}

// MarshalJSON implements the json.Marshaler interface for GetUsageReportResponse.
// It customizes the JSON marshaling process for GetUsageReportResponse objects.
func (g *GetUsageReportResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetUsageReportResponse object to a map representation for JSON marshaling.
func (g *GetUsageReportResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Url.IsValueSet() {
		structMap["url"] = g.Url.Value()
	}
	if g.UsageReportUrl.IsValueSet() {
		structMap["usage_report_url"] = g.UsageReportUrl.Value()
	}
	if g.GroupedReportUrl.IsValueSet() {
		structMap["grouped_report_url"] = g.GroupedReportUrl.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetUsageReportResponse.
// It customizes the JSON unmarshaling process for GetUsageReportResponse objects.
func (g *GetUsageReportResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Url              Optional[string] `json:"url"`
		UsageReportUrl   Optional[string] `json:"usage_report_url"`
		GroupedReportUrl Optional[string] `json:"grouped_report_url"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Url = temp.Url
	g.UsageReportUrl = temp.UsageReportUrl
	g.GroupedReportUrl = temp.GroupedReportUrl
	return nil
}
