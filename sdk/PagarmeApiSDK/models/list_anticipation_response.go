package models

import (
	"encoding/json"
)

// ListAnticipationResponse represents a ListAnticipationResponse struct.
// Anticipations
type ListAnticipationResponse struct {
	// Anticipations
	Data Optional[[]GetAnticipationResponse] `json:"data"`
	// Paging
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListAnticipationResponse.
// It customizes the JSON marshaling process for ListAnticipationResponse objects.
func (l *ListAnticipationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListAnticipationResponse object to a map representation for JSON marshaling.
func (l *ListAnticipationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListAnticipationResponse.
// It customizes the JSON unmarshaling process for ListAnticipationResponse objects.
func (l *ListAnticipationResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetAnticipationResponse] `json:"data"`
		Paging Optional[PagingResponse]            `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
