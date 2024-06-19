package models

import (
	"encoding/json"
)

// GetBalanceResponse represents a GetBalanceResponse struct.
// Balance
type GetBalanceResponse struct {
	// Currency
	Currency Optional[string] `json:"currency"`
	// Amount available for transferring
	AvailableAmount Optional[int64] `json:"available_amount"`
	// Recipient
	Recipient          Optional[GetRecipientResponse] `json:"recipient"`
	TransferredAmount  Optional[int64]                `json:"transferred_amount"`
	WaitingFundsAmount Optional[int64]                `json:"waiting_funds_amount"`
}

// MarshalJSON implements the json.Marshaler interface for GetBalanceResponse.
// It customizes the JSON marshaling process for GetBalanceResponse objects.
func (g *GetBalanceResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetBalanceResponse object to a map representation for JSON marshaling.
func (g *GetBalanceResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Currency.IsValueSet() {
		structMap["currency"] = g.Currency.Value()
	}
	if g.AvailableAmount.IsValueSet() {
		structMap["available_amount"] = g.AvailableAmount.Value()
	}
	if g.Recipient.IsValueSet() {
		structMap["recipient"] = g.Recipient.Value()
	}
	if g.TransferredAmount.IsValueSet() {
		structMap["transferred_amount"] = g.TransferredAmount.Value()
	}
	if g.WaitingFundsAmount.IsValueSet() {
		structMap["waiting_funds_amount"] = g.WaitingFundsAmount.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBalanceResponse.
// It customizes the JSON unmarshaling process for GetBalanceResponse objects.
func (g *GetBalanceResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Currency           Optional[string]               `json:"currency"`
		AvailableAmount    Optional[int64]                `json:"available_amount"`
		Recipient          Optional[GetRecipientResponse] `json:"recipient"`
		TransferredAmount  Optional[int64]                `json:"transferred_amount"`
		WaitingFundsAmount Optional[int64]                `json:"waiting_funds_amount"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Currency = temp.Currency
	g.AvailableAmount = temp.AvailableAmount
	g.Recipient = temp.Recipient
	g.TransferredAmount = temp.TransferredAmount
	g.WaitingFundsAmount = temp.WaitingFundsAmount
	return nil
}
