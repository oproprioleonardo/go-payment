package models

import (
	"encoding/json"
)

// PagingResponse represents a PagingResponse struct.
// Object used for returning lists of objects with pagination
type PagingResponse struct {
	// Total number of pages
	Total Optional[int] `json:"total"`
	// Previous page
	Previous Optional[string] `json:"previous"`
	// Next page
	Next Optional[string] `json:"next"`
}

// MarshalJSON implements the json.Marshaler interface for PagingResponse.
// It customizes the JSON marshaling process for PagingResponse objects.
func (p *PagingResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PagingResponse object to a map representation for JSON marshaling.
func (p *PagingResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.Total.IsValueSet() {
		structMap["total"] = p.Total.Value()
	}
	if p.Previous.IsValueSet() {
		structMap["previous"] = p.Previous.Value()
	}
	if p.Next.IsValueSet() {
		structMap["next"] = p.Next.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PagingResponse.
// It customizes the JSON unmarshaling process for PagingResponse objects.
func (p *PagingResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Total    Optional[int]    `json:"total"`
		Previous Optional[string] `json:"previous"`
		Next     Optional[string] `json:"next"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Total = temp.Total
	p.Previous = temp.Previous
	p.Next = temp.Next
	return nil
}
