package models

import (
	"encoding/json"
)

// GetSubscriptionBoletoResponse represents a GetSubscriptionBoletoResponse struct.
// Response object for getting a boleto
type GetSubscriptionBoletoResponse struct {
	// Interest
	Interest Optional[GetInterestResponse] `json:"interest"`
	// Fine
	Fine                Optional[GetFineResponse] `json:"fine"`
	MaxDaysToPayPastDue Optional[int]             `json:"max_days_to_pay_past_due"`
}

// MarshalJSON implements the json.Marshaler interface for GetSubscriptionBoletoResponse.
// It customizes the JSON marshaling process for GetSubscriptionBoletoResponse objects.
func (g *GetSubscriptionBoletoResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetSubscriptionBoletoResponse object to a map representation for JSON marshaling.
func (g *GetSubscriptionBoletoResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Interest.IsValueSet() {
		structMap["interest"] = g.Interest.Value()
	}
	if g.Fine.IsValueSet() {
		structMap["fine"] = g.Fine.Value()
	}
	if g.MaxDaysToPayPastDue.IsValueSet() {
		structMap["max_days_to_pay_past_due"] = g.MaxDaysToPayPastDue.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetSubscriptionBoletoResponse.
// It customizes the JSON unmarshaling process for GetSubscriptionBoletoResponse objects.
func (g *GetSubscriptionBoletoResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Interest            Optional[GetInterestResponse] `json:"interest"`
		Fine                Optional[GetFineResponse]     `json:"fine"`
		MaxDaysToPayPastDue Optional[int]                 `json:"max_days_to_pay_past_due"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Interest = temp.Interest
	g.Fine = temp.Fine
	g.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
	return nil
}
