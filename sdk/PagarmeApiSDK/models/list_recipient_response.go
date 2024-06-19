package models

import (
	"encoding/json"
)

// ListRecipientResponse represents a ListRecipientResponse struct.
// Response for the listing recipient method
type ListRecipientResponse struct {
	// Recipients
	Data Optional[[]GetRecipientResponse] `json:"data"`
	// Paging
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListRecipientResponse.
// It customizes the JSON marshaling process for ListRecipientResponse objects.
func (l *ListRecipientResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListRecipientResponse object to a map representation for JSON marshaling.
func (l *ListRecipientResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListRecipientResponse.
// It customizes the JSON unmarshaling process for ListRecipientResponse objects.
func (l *ListRecipientResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetRecipientResponse] `json:"data"`
		Paging Optional[PagingResponse]         `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
