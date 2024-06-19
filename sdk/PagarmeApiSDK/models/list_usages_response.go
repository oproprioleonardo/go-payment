package models

import (
	"encoding/json"
)

// ListUsagesResponse represents a ListUsagesResponse struct.
// Response model for listing the usages from a subscription item
type ListUsagesResponse struct {
	// The usage objects
	Data Optional[[]GetUsageResponse] `json:"data"`
	// Paging object
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListUsagesResponse.
// It customizes the JSON marshaling process for ListUsagesResponse objects.
func (l *ListUsagesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListUsagesResponse object to a map representation for JSON marshaling.
func (l *ListUsagesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListUsagesResponse.
// It customizes the JSON unmarshaling process for ListUsagesResponse objects.
func (l *ListUsagesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetUsageResponse] `json:"data"`
		Paging Optional[PagingResponse]     `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
