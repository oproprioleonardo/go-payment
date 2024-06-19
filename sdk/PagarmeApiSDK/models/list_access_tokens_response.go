package models

import (
	"encoding/json"
)

// ListAccessTokensResponse represents a ListAccessTokensResponse struct.
// Response object for listing access tokens
type ListAccessTokensResponse struct {
	// The access token objects
	Data Optional[[]GetAccessTokenResponse] `json:"data"`
	// Paging object
	Paging Optional[PagingResponse] `json:"paging"`
}

// MarshalJSON implements the json.Marshaler interface for ListAccessTokensResponse.
// It customizes the JSON marshaling process for ListAccessTokensResponse objects.
func (l *ListAccessTokensResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListAccessTokensResponse object to a map representation for JSON marshaling.
func (l *ListAccessTokensResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Data.IsValueSet() {
		structMap["data"] = l.Data.Value()
	}
	if l.Paging.IsValueSet() {
		structMap["paging"] = l.Paging.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListAccessTokensResponse.
// It customizes the JSON unmarshaling process for ListAccessTokensResponse objects.
func (l *ListAccessTokensResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Data   Optional[[]GetAccessTokenResponse] `json:"data"`
		Paging Optional[PagingResponse]           `json:"paging"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.Data = temp.Data
	l.Paging = temp.Paging
	return nil
}
