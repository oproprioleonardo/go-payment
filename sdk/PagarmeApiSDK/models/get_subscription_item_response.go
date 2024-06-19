package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetSubscriptionItemResponse represents a GetSubscriptionItemResponse struct.
type GetSubscriptionItemResponse struct {
	Id            Optional[string]                   `json:"id"`
	Description   Optional[string]                   `json:"description"`
	Status        Optional[string]                   `json:"status"`
	CreatedAt     Optional[time.Time]                `json:"created_at"`
	UpdatedAt     Optional[time.Time]                `json:"updated_at"`
	PricingScheme Optional[GetPricingSchemeResponse] `json:"pricing_scheme"`
	Discounts     Optional[[]GetDiscountResponse]    `json:"discounts"`
	Increments    Optional[[]GetIncrementResponse]   `json:"increments"`
	Subscription  Optional[GetSubscriptionResponse]  `json:"subscription"`
	// Item name
	Name      Optional[string]    `json:"name"`
	Quantity  Optional[int]       `json:"quantity"`
	Cycles    Optional[int]       `json:"cycles"`
	DeletedAt Optional[time.Time] `json:"deleted_at"`
}

// MarshalJSON implements the json.Marshaler interface for GetSubscriptionItemResponse.
// It customizes the JSON marshaling process for GetSubscriptionItemResponse objects.
func (g *GetSubscriptionItemResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetSubscriptionItemResponse object to a map representation for JSON marshaling.
func (g *GetSubscriptionItemResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
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
	if g.PricingScheme.IsValueSet() {
		structMap["pricing_scheme"] = g.PricingScheme.Value()
	}
	if g.Discounts.IsValueSet() {
		structMap["discounts"] = g.Discounts.Value()
	}
	if g.Increments.IsValueSet() {
		structMap["increments"] = g.Increments.Value()
	}
	if g.Subscription.IsValueSet() {
		structMap["subscription"] = g.Subscription.Value()
	}
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.Quantity.IsValueSet() {
		structMap["quantity"] = g.Quantity.Value()
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
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSubscriptionItemResponse.
// It customizes the JSON unmarshaling process for GetSubscriptionItemResponse objects.
func (g *GetSubscriptionItemResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id            Optional[string]                   `json:"id"`
		Description   Optional[string]                   `json:"description"`
		Status        Optional[string]                   `json:"status"`
		CreatedAt     Optional[string]                   `json:"created_at"`
		UpdatedAt     Optional[string]                   `json:"updated_at"`
		PricingScheme Optional[GetPricingSchemeResponse] `json:"pricing_scheme"`
		Discounts     Optional[[]GetDiscountResponse]    `json:"discounts"`
		Increments    Optional[[]GetIncrementResponse]   `json:"increments"`
		Subscription  Optional[GetSubscriptionResponse]  `json:"subscription"`
		Name          Optional[string]                   `json:"name"`
		Quantity      Optional[int]                      `json:"quantity"`
		Cycles        Optional[int]                      `json:"cycles"`
		DeletedAt     Optional[string]                   `json:"deleted_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Description = temp.Description
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
	g.PricingScheme = temp.PricingScheme
	g.Discounts = temp.Discounts
	g.Increments = temp.Increments
	g.Subscription = temp.Subscription
	g.Name = temp.Name
	g.Quantity = temp.Quantity
	g.Cycles = temp.Cycles
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	return nil
}
