package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetUsageResponse represents a GetUsageResponse struct.
// Response object for getting a usage
type GetUsageResponse struct {
	// Id
	Id Optional[string] `json:"id"`
	// Quantity
	Quantity Optional[int] `json:"quantity"`
	// Description
	Description Optional[string] `json:"description"`
	// Used at
	UsedAt Optional[time.Time] `json:"used_at"`
	// Creation date
	CreatedAt Optional[time.Time] `json:"created_at"`
	// Status
	Status    Optional[string]    `json:"status"`
	DeletedAt Optional[time.Time] `json:"deleted_at"`
	// Subscription item
	SubscriptionItem Optional[GetSubscriptionItemResponse] `json:"subscription_item"`
	// Identification code in the client system
	Code Optional[string] `json:"code"`
	// Identification group in the client system
	Group Optional[string] `json:"group"`
	// Field used in item scheme type 'Percent'
	Amount Optional[int] `json:"amount"`
}

// MarshalJSON implements the json.Marshaler interface for GetUsageResponse.
// It customizes the JSON marshaling process for GetUsageResponse objects.
func (g *GetUsageResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetUsageResponse object to a map representation for JSON marshaling.
func (g *GetUsageResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Quantity.IsValueSet() {
		structMap["quantity"] = g.Quantity.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.UsedAt.IsValueSet() {
		var UsedAtVal *string = nil
		if g.UsedAt.Value() != nil {
			val := g.UsedAt.Value().Format(time.RFC3339)
			UsedAtVal = &val
		}
		structMap["used_at"] = UsedAtVal
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.DeletedAt.IsValueSet() {
		var DeletedAtVal *string = nil
		if g.DeletedAt.Value() != nil {
			val := g.DeletedAt.Value().Format(time.RFC3339)
			DeletedAtVal = &val
		}
		structMap["deleted_at"] = DeletedAtVal
	}
	if g.SubscriptionItem.IsValueSet() {
		structMap["subscription_item"] = g.SubscriptionItem.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.Group.IsValueSet() {
		structMap["group"] = g.Group.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetUsageResponse.
// It customizes the JSON unmarshaling process for GetUsageResponse objects.
func (g *GetUsageResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id               Optional[string]                      `json:"id"`
		Quantity         Optional[int]                         `json:"quantity"`
		Description      Optional[string]                      `json:"description"`
		UsedAt           Optional[string]                      `json:"used_at"`
		CreatedAt        Optional[string]                      `json:"created_at"`
		Status           Optional[string]                      `json:"status"`
		DeletedAt        Optional[string]                      `json:"deleted_at"`
		SubscriptionItem Optional[GetSubscriptionItemResponse] `json:"subscription_item"`
		Code             Optional[string]                      `json:"code"`
		Group            Optional[string]                      `json:"group"`
		Amount           Optional[int]                         `json:"amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Quantity = temp.Quantity
	g.Description = temp.Description
	g.UsedAt.ShouldSetValue(temp.UsedAt.IsValueSet())
	if temp.UsedAt.Value() != nil {
		UsedAtVal, err := time.Parse(time.RFC3339, (*temp.UsedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse used_at as % s format.", time.RFC3339)
		}
		g.UsedAt.SetValue(&UsedAtVal)
	}
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.Status = temp.Status
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	g.SubscriptionItem = temp.SubscriptionItem
	g.Code = temp.Code
	g.Group = temp.Group
	g.Amount = temp.Amount
	return nil
}
