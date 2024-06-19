package models

import (
	"encoding/json"
)

// ListSubscriptionItemsResponse represents a ListSubscriptionItemsResponse struct.
// Response model for listing subscription items
type ListSubscriptionItemsResponse struct {
	// The subscription items
	Data Optional[[]GetSubscriptionItemResponse] `json:"data"`
	// Paging object
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionItemsResponse.
// It customizes the JSON marshaling process for ListSubscriptionItemsResponse objects.
func (l *ListSubscriptionItemsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionItemsResponse object to a map representation for JSON marshaling.
func (l *ListSubscriptionItemsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionItemsResponse.
// It customizes the JSON unmarshaling process for ListSubscriptionItemsResponse objects.
func (l *ListSubscriptionItemsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetSubscriptionItemResponse] `json:"data"`
		Paging Optional[PagingResponse]                `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
