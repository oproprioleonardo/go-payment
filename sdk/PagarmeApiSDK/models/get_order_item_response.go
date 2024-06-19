package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetOrderItemResponse represents a GetOrderItemResponse struct.
// Response object for getting an order item
type GetOrderItemResponse struct {
	// Id
	Id          Optional[string] `json:"id"`
	Type        Optional[string] `json:"type"`
	Description Optional[string] `json:"description"`
	Amount      Optional[int]    `json:"amount"`
	Quantity    Optional[int]    `json:"quantity"`
	// Category
	Category Optional[string] `json:"category"`
	// Code
	Code      Optional[string]    `json:"code"`
	Status    Optional[string]    `json:"status"`
	CreatedAt Optional[time.Time] `json:"created_at"`
	UpdatedAt Optional[time.Time] `json:"updated_at"`
}

// MarshalJSON implements the json.Marshaler interface for GetOrderItemResponse.
// It customizes the JSON marshaling process for GetOrderItemResponse objects.
func (g *GetOrderItemResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetOrderItemResponse object to a map representation for JSON marshaling.
func (g *GetOrderItemResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Quantity.IsValueSet() {
		structMap["quantity"] = g.Quantity.Value()
	}
	if g.Category.IsValueSet() {
		structMap["category"] = g.Category.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.UpdatedAt.IsValueSet() {
		var UpdatedAtVal *string = nil
		if g.UpdatedAt.Value() != nil {
			val := g.UpdatedAt.Value().Format(time.RFC3339)
			UpdatedAtVal = &val
		}
		structMap["updated_at"] = UpdatedAtVal
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOrderItemResponse.
// It customizes the JSON unmarshaling process for GetOrderItemResponse objects.
func (g *GetOrderItemResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id          Optional[string] `json:"id"`
		Type        Optional[string] `json:"type"`
		Description Optional[string] `json:"description"`
		Amount      Optional[int]    `json:"amount"`
		Quantity    Optional[int]    `json:"quantity"`
		Category    Optional[string] `json:"category"`
		Code        Optional[string] `json:"code"`
		Status      Optional[string] `json:"status"`
		CreatedAt   Optional[string] `json:"created_at"`
		UpdatedAt   Optional[string] `json:"updated_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Type = temp.Type
	g.Description = temp.Description
	g.Amount = temp.Amount
	g.Quantity = temp.Quantity
	g.Category = temp.Category
	g.Code = temp.Code
	g.Status = temp.Status
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.UpdatedAt.ShouldSetValue(temp.UpdatedAt.IsValueSet())
	if temp.UpdatedAt.Value() != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, (*temp.UpdatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		g.UpdatedAt.SetValue(&UpdatedAtVal)
	}
	return nil
}
