package models

import (
	"encoding/json"
	"fmt"
)

// GetMovementObjectRefundResponse represents a GetMovementObjectRefundResponse struct.
// Generic response object for getting a MovementObjectRefund.
type GetMovementObjectRefundResponse struct {
	GetMovementObjectBaseResponse
	FraudCoverageFee     Optional[string] `json:"fraud_coverage_fee"`
	ChargeFeeRecipientId Optional[string] `json:"charge_fee_recipient_id"`
	BankAccountId        Optional[string] `json:"bank_account_id"`
	LocalTransactionId   Optional[string] `json:"local_transaction_id"`
	UpdatedAt            Optional[string] `json:"updated_at"`
}

// MarshalJSON implements the json.Marshaler interface for GetMovementObjectRefundResponse.
// It customizes the JSON marshaling process for GetMovementObjectRefundResponse objects.
func (g *GetMovementObjectRefundResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetMovementObjectRefundResponse object to a map representation for JSON marshaling.
func (g *GetMovementObjectRefundResponse) toMap() map[string]any {
	structMap := g.GetMovementObjectBaseResponse.toMap()
	if g.GetMovementObjectBaseResponse.Object != nil {
		structMap["object"] = *g.GetMovementObjectBaseResponse.Object
	} else {
		structMap["object"] = "refund"
	}
	if g.FraudCoverageFee.IsValueSet() {
		structMap["fraud_coverage_fee"] = g.FraudCoverageFee.Value()
	}
	if g.ChargeFeeRecipientId.IsValueSet() {
		structMap["charge_fee_recipient_id"] = g.ChargeFeeRecipientId.Value()
	}
	if g.BankAccountId.IsValueSet() {
		structMap["bank_account_id"] = g.BankAccountId.Value()
	}
	if g.LocalTransactionId.IsValueSet() {
		structMap["local_transaction_id"] = g.LocalTransactionId.Value()
	}
	if g.UpdatedAt.IsValueSet() {
		structMap["updated_at"] = g.UpdatedAt.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectRefundResponse.
// It customizes the JSON unmarshaling process for GetMovementObjectRefundResponse objects.
func (g *GetMovementObjectRefundResponse) UnmarshalJSON(input []byte) error {
	g.GetMovementObjectBaseResponse.UnmarshalJSON(input)
	temp := &struct {
		FraudCoverageFee     Optional[string] `json:"fraud_coverage_fee"`
		ChargeFeeRecipientId Optional[string] `json:"charge_fee_recipient_id"`
		BankAccountId        Optional[string] `json:"bank_account_id"`
		LocalTransactionId   Optional[string] `json:"local_transaction_id"`
		UpdatedAt            Optional[string] `json:"updated_at"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.FraudCoverageFee = temp.FraudCoverageFee
	g.ChargeFeeRecipientId = temp.ChargeFeeRecipientId
	g.BankAccountId = temp.BankAccountId
	g.LocalTransactionId = temp.LocalTransactionId
	g.UpdatedAt = temp.UpdatedAt
	return nil
}

// fraud_coverage_fee returns the fraud_coverage_fee from the GetMovementObjectRefundResponse.
func (g *GetMovementObjectRefundResponse) GetFraudCoverageFee() Optional[string] {
	return g.FraudCoverageFee
}

// charge_fee_recipient_id returns the charge_fee_recipient_id from the GetMovementObjectRefundResponse.
func (g *GetMovementObjectRefundResponse) GetChargeFeeRecipientId() Optional[string] {
	return g.ChargeFeeRecipientId
}

// bank_account_id returns the bank_account_id from the GetMovementObjectRefundResponse.
func (g *GetMovementObjectRefundResponse) GetBankAccountId() Optional[string] {
	return g.BankAccountId
}

// local_transaction_id returns the local_transaction_id from the GetMovementObjectRefundResponse.
func (g *GetMovementObjectRefundResponse) GetLocalTransactionId() Optional[string] {
	return g.LocalTransactionId
}

// updated_at returns the updated_at from the GetMovementObjectRefundResponse.
func (g *GetMovementObjectRefundResponse) GetUpdatedAt() Optional[string] {
	return g.UpdatedAt
}

// UnmarshalGetMovementObjectRefundResponse is a utility function used to unmarshal JSON into a GetMovementObjectRefundResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetMovementObjectRefundResponseInterface and an error, if any.
func UnmarshalGetMovementObjectRefundResponse(
	inputType string,
	input []byte) (
	GetMovementObjectRefundResponseInterface,
	error) {
	getMovementObjectBaseResponse, err := UnmarshalGetMovementObjectBaseResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getMovementObjectRefundResponse, ok := getMovementObjectBaseResponse.(GetMovementObjectRefundResponseInterface); ok {
		return getMovementObjectRefundResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getMovementObjectRefundResponse %v", getMovementObjectBaseResponse)
}

// ToGetMovementObjectRefundResponseArray converts an array of GetMovementObjectRefundResponseField to an array of GetMovementObjectRefundResponseInterface.
func ToGetMovementObjectRefundResponseArray(getMovementObjectRefundResponse []GetMovementObjectRefundResponseField) []GetMovementObjectRefundResponseInterface {
	var items []GetMovementObjectRefundResponseInterface
	for _, temp := range getMovementObjectRefundResponse {
		items = append(items, temp.GetMovementObjectRefundResponseInterface)
	}
	return items
}

// GetMovementObjectRefundResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetMovementObjectRefundResponseField struct {
	GetMovementObjectRefundResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectRefundResponseField.
// It customizes the JSON unmarshaling process for GetMovementObjectRefundResponseField objects.
func (g *GetMovementObjectRefundResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetMovementObjectRefundResponseInterface, err = UnmarshalGetMovementObjectRefundResponse("refund", input)
	return err
}

// GetMovementObjectRefundResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetMovementObjectRefundResponseSlice struct {
	Slice []GetMovementObjectRefundResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectRefundResponseSlice.
// It customizes the JSON unmarshaling process for GetMovementObjectRefundResponseSlice objects.
func (ga *GetMovementObjectRefundResponseSlice) UnmarshalJSON(input []byte) error {
	tempStruct := struct {
		Slice []json.RawMessage
	}{}
	temp := []json.RawMessage{}
	err := json.Unmarshal(input, &tempStruct)
	if err != nil {
		err := json.Unmarshal(input, &temp)
		if err != nil {
			return err
		}
	} else {
		temp = tempStruct.Slice
	}

	ga.Slice = nil
	if temp != nil {
		ga.Slice = []GetMovementObjectRefundResponseInterface{}
		for _, getMovementObjectRefundResponse := range temp {
			if getMovementObjectRefundResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetMovementObjectRefundResponseInterface
			obj, err := UnmarshalGetMovementObjectRefundResponse("refund", getMovementObjectRefundResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
