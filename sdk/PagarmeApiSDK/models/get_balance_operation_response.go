package models

import (
	"encoding/json"
)

// GetBalanceOperationResponse represents a GetBalanceOperationResponse struct.
// Generic response object for getting a BalanceOperation.
type GetBalanceOperationResponse struct {
	Id               Optional[string]                       `json:"id"`
	Status           Optional[string]                       `json:"status"`
	BalanceAmount    Optional[string]                       `json:"balance_amount"`
	BalanceOldAmount Optional[string]                       `json:"balance_old_amount"`
	Type             Optional[string]                       `json:"type"`
	Amount           Optional[string]                       `json:"amount"`
	Fee              Optional[string]                       `json:"fee"`
	CreatedAt        Optional[string]                       `json:"created_at"`
	MovementObject   GetMovementObjectBaseResponseInterface `json:"movement_object,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for GetBalanceOperationResponse.
// It customizes the JSON marshaling process for GetBalanceOperationResponse objects.
func (g *GetBalanceOperationResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetBalanceOperationResponse object to a map representation for JSON marshaling.
func (g *GetBalanceOperationResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.BalanceAmount.IsValueSet() {
		structMap["balance_amount"] = g.BalanceAmount.Value()
	}
	if g.BalanceOldAmount.IsValueSet() {
		structMap["balance_old_amount"] = g.BalanceOldAmount.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.Amount.IsValueSet() {
		structMap["amount"] = g.Amount.Value()
	}
	if g.Fee.IsValueSet() {
		structMap["fee"] = g.Fee.Value()
	}
	if g.CreatedAt.IsValueSet() {
		structMap["created_at"] = g.CreatedAt.Value()
	}
	if g.MovementObject != nil {
		structMap["movement_object"] = g.MovementObject
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetBalanceOperationResponse.
// It customizes the JSON unmarshaling process for GetBalanceOperationResponse objects.
func (g *GetBalanceOperationResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id               Optional[string]                       `json:"id"`
		Status           Optional[string]                       `json:"status"`
		BalanceAmount    Optional[string]                       `json:"balance_amount"`
		BalanceOldAmount Optional[string]                       `json:"balance_old_amount"`
		Type             Optional[string]                       `json:"type"`
		Amount           Optional[string]                       `json:"amount"`
		Fee              Optional[string]                       `json:"fee"`
		CreatedAt        Optional[string]                       `json:"created_at"`
		MovementObject   GetMovementObjectBaseResponseInterface `json:"movement_object,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Status = temp.Status
	g.BalanceAmount = temp.BalanceAmount
	g.BalanceOldAmount = temp.BalanceOldAmount
	g.Type = temp.Type
	g.Amount = temp.Amount
	g.Fee = temp.Fee
	g.CreatedAt = temp.CreatedAt
	g.MovementObject = temp.MovementObject
	return nil
}
