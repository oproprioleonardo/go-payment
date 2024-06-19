package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetAnticipationResponse represents a GetAnticipationResponse struct.
// Anticipation
type GetAnticipationResponse struct {
	// Id
	Id Optional[string] `json:"id"`
	// Requested amount
	RequestedAmount Optional[int] `json:"requested_amount"`
	// Approved amount
	ApprovedAmount Optional[int] `json:"approved_amount"`
	// Recipient
	Recipient Optional[GetRecipientResponse] `json:"recipient"`
	// Anticipation id on the gateway
	Pgid Optional[string] `json:"pgid"`
	// Creation date
	CreatedAt Optional[time.Time] `json:"created_at"`
	// Last update date
	UpdatedAt Optional[time.Time] `json:"updated_at"`
	// Payment date
	PaymentDate Optional[time.Time] `json:"payment_date"`
	// Status
	Status Optional[string] `json:"status"`
	// Timeframe
	Timeframe Optional[string] `json:"timeframe"`
}

// MarshalJSON implements the json.Marshaler interface for GetAnticipationResponse.
// It customizes the JSON marshaling process for GetAnticipationResponse objects.
func (g *GetAnticipationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetAnticipationResponse object to a map representation for JSON marshaling.
func (g *GetAnticipationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.RequestedAmount.IsValueSet() {
		structMap["requested_amount"] = g.RequestedAmount.Value()
	}
	if g.ApprovedAmount.IsValueSet() {
		structMap["approved_amount"] = g.ApprovedAmount.Value()
	}
	if g.Recipient.IsValueSet() {
		structMap["recipient"] = g.Recipient.Value()
	}
	if g.Pgid.IsValueSet() {
		structMap["pgid"] = g.Pgid.Value()
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
	if g.PaymentDate.IsValueSet() {
		var PaymentDateVal *string = nil
		if g.PaymentDate.Value() != nil {
			val := g.PaymentDate.Value().Format(time.RFC3339)
			PaymentDateVal = &val
		}
		structMap["payment_date"] = PaymentDateVal
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.Timeframe.IsValueSet() {
		structMap["timeframe"] = g.Timeframe.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetAnticipationResponse.
// It customizes the JSON unmarshaling process for GetAnticipationResponse objects.
func (g *GetAnticipationResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id              Optional[string]               `json:"id"`
		RequestedAmount Optional[int]                  `json:"requested_amount"`
		ApprovedAmount  Optional[int]                  `json:"approved_amount"`
		Recipient       Optional[GetRecipientResponse] `json:"recipient"`
		Pgid            Optional[string]               `json:"pgid"`
		CreatedAt       Optional[string]               `json:"created_at"`
		UpdatedAt       Optional[string]               `json:"updated_at"`
		PaymentDate     Optional[string]               `json:"payment_date"`
		Status          Optional[string]               `json:"status"`
		Timeframe       Optional[string]               `json:"timeframe"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.RequestedAmount = temp.RequestedAmount
	g.ApprovedAmount = temp.ApprovedAmount
	g.Recipient = temp.Recipient
	g.Pgid = temp.Pgid
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
	g.PaymentDate.ShouldSetValue(temp.PaymentDate.IsValueSet())
	if temp.PaymentDate.Value() != nil {
		PaymentDateVal, err := time.Parse(time.RFC3339, (*temp.PaymentDate.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse payment_date as % s format.", time.RFC3339)
		}
		g.PaymentDate.SetValue(&PaymentDateVal)
	}
	g.Status = temp.Status
	g.Timeframe = temp.Timeframe
	return nil
}
