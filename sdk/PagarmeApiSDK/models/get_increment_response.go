package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetIncrementResponse represents a GetIncrementResponse struct.
// Response object for getting a increment
type GetIncrementResponse struct {
	Id            Optional[string]                  `json:"id"`
	Value         Optional[float64]                 `json:"value"`
	IncrementType Optional[string]                  `json:"increment_type"`
	Status        Optional[string]                  `json:"status"`
	CreatedAt     Optional[time.Time]               `json:"created_at"`
	Cycles        Optional[int]                     `json:"cycles"`
	DeletedAt     Optional[time.Time]               `json:"deleted_at"`
	Description   Optional[string]                  `json:"description"`
	Subscription  Optional[GetSubscriptionResponse] `json:"subscription"`
	// The Subscription Item
	SubscriptionItem Optional[GetSubscriptionItemResponse] `json:"subscription_item"`
}

// MarshalJSON implements the json.Marshaler interface for GetIncrementResponse.
// It customizes the JSON marshaling process for GetIncrementResponse objects.
func (g *GetIncrementResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetIncrementResponse object to a map representation for JSON marshaling.
func (g *GetIncrementResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Value.IsValueSet() {
		structMap["value"] = g.Value.Value()
	}
	if g.IncrementType.IsValueSet() {
		structMap["increment_type"] = g.IncrementType.Value()
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
	if g.Cycles.IsValueSet() {
		structMap["cycles"] = g.Cycles.Value()
	}
	if g.DeletedAt.IsValueSet() {
		var DeletedAtVal *string = nil
		if g.DeletedAt.Value() != nil {
			val := g.DeletedAt.Value().Format(time.RFC3339)
			DeletedAtVal = &val
		}
		structMap["deleted_at"] = DeletedAtVal
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.Subscription.IsValueSet() {
		structMap["subscription"] = g.Subscription.Value()
	}
	if g.SubscriptionItem.IsValueSet() {
		structMap["subscription_item"] = g.SubscriptionItem.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetIncrementResponse.
// It customizes the JSON unmarshaling process for GetIncrementResponse objects.
func (g *GetIncrementResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id               Optional[string]                      `json:"id"`
		Value            Optional[float64]                     `json:"value"`
		IncrementType    Optional[string]                      `json:"increment_type"`
		Status           Optional[string]                      `json:"status"`
		CreatedAt        Optional[string]                      `json:"created_at"`
		Cycles           Optional[int]                         `json:"cycles"`
		DeletedAt        Optional[string]                      `json:"deleted_at"`
		Description      Optional[string]                      `json:"description"`
		Subscription     Optional[GetSubscriptionResponse]     `json:"subscription"`
		SubscriptionItem Optional[GetSubscriptionItemResponse] `json:"subscription_item"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Value = temp.Value
	g.IncrementType = temp.IncrementType
	g.Status = temp.Status
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.Cycles = temp.Cycles
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	g.Description = temp.Description
	g.Subscription = temp.Subscription
	g.SubscriptionItem = temp.SubscriptionItem
	return nil
}
