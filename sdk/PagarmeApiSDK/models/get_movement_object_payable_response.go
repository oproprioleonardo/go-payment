package models

import (
	"encoding/json"
	"fmt"
)

// GetMovementObjectPayableResponse represents a GetMovementObjectPayableResponse struct.
type GetMovementObjectPayableResponse struct {
	GetMovementObjectBaseResponse
	Fee                      Optional[string] `json:"fee"`
	AnticipationFee          string           `json:"anticipation_fee"`
	FraudCoverageFee         string           `json:"fraud_coverage_fee"`
	Installment              string           `json:"installment"`
	SplitId                  string           `json:"split_id"`
	BulkAnticipationId       string           `json:"bulk_anticipation_id"`
	AnticipationId           string           `json:"anticipation_id"`
	RecipientId              string           `json:"recipient_id"`
	OriginatorModel          string           `json:"originator_model"`
	OriginatorModelId        string           `json:"originator_model_id"`
	PaymentDate              string           `json:"payment_date"`
	OriginalPaymentDate      string           `json:"original_payment_date"`
	PaymentMethod            string           `json:"payment_method"`
	AccrualAt                string           `json:"accrual_at"`
	LiquidationArrangementId string           `json:"liquidation_arrangement_id"`
}

// MarshalJSON implements the json.Marshaler interface for GetMovementObjectPayableResponse.
// It customizes the JSON marshaling process for GetMovementObjectPayableResponse objects.
func (g *GetMovementObjectPayableResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetMovementObjectPayableResponse object to a map representation for JSON marshaling.
func (g *GetMovementObjectPayableResponse) toMap() map[string]any {
	structMap := g.GetMovementObjectBaseResponse.toMap()
	if g.GetMovementObjectBaseResponse.Object != nil {
		structMap["object"] = *g.GetMovementObjectBaseResponse.Object
	} else {
		structMap["object"] = "payable"
	}
	if g.Fee.IsValueSet() {
		structMap["fee"] = g.Fee.Value()
	}
	structMap["anticipation_fee"] = g.AnticipationFee
	structMap["fraud_coverage_fee"] = g.FraudCoverageFee
	structMap["installment"] = g.Installment
	structMap["split_id"] = g.SplitId
	structMap["bulk_anticipation_id"] = g.BulkAnticipationId
	structMap["anticipation_id"] = g.AnticipationId
	structMap["recipient_id"] = g.RecipientId
	structMap["originator_model"] = g.OriginatorModel
	structMap["originator_model_id"] = g.OriginatorModelId
	structMap["payment_date"] = g.PaymentDate
	structMap["original_payment_date"] = g.OriginalPaymentDate
	structMap["payment_method"] = g.PaymentMethod
	structMap["accrual_at"] = g.AccrualAt
	structMap["liquidation_arrangement_id"] = g.LiquidationArrangementId
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectPayableResponse.
// It customizes the JSON unmarshaling process for GetMovementObjectPayableResponse objects.
func (g *GetMovementObjectPayableResponse) UnmarshalJSON(input []byte) error {
	g.GetMovementObjectBaseResponse.UnmarshalJSON(input)
	temp := &struct {
		Fee                      Optional[string] `json:"fee"`
		AnticipationFee          string           `json:"anticipation_fee"`
		FraudCoverageFee         string           `json:"fraud_coverage_fee"`
		Installment              string           `json:"installment"`
		SplitId                  string           `json:"split_id"`
		BulkAnticipationId       string           `json:"bulk_anticipation_id"`
		AnticipationId           string           `json:"anticipation_id"`
		RecipientId              string           `json:"recipient_id"`
		OriginatorModel          string           `json:"originator_model"`
		OriginatorModelId        string           `json:"originator_model_id"`
		PaymentDate              string           `json:"payment_date"`
		OriginalPaymentDate      string           `json:"original_payment_date"`
		PaymentMethod            string           `json:"payment_method"`
		AccrualAt                string           `json:"accrual_at"`
		LiquidationArrangementId string           `json:"liquidation_arrangement_id"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Fee = temp.Fee
	g.AnticipationFee = temp.AnticipationFee
	g.FraudCoverageFee = temp.FraudCoverageFee
	g.Installment = temp.Installment
	g.SplitId = temp.SplitId
	g.BulkAnticipationId = temp.BulkAnticipationId
	g.AnticipationId = temp.AnticipationId
	g.RecipientId = temp.RecipientId
	g.OriginatorModel = temp.OriginatorModel
	g.OriginatorModelId = temp.OriginatorModelId
	g.PaymentDate = temp.PaymentDate
	g.OriginalPaymentDate = temp.OriginalPaymentDate
	g.PaymentMethod = temp.PaymentMethod
	g.AccrualAt = temp.AccrualAt
	g.LiquidationArrangementId = temp.LiquidationArrangementId
	return nil
}

// fee returns the fee from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetFee() Optional[string] {
	return g.Fee
}

// anticipation_fee returns the anticipation_fee from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetAnticipationFee() string {
	return g.AnticipationFee
}

// fraud_coverage_fee returns the fraud_coverage_fee from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetFraudCoverageFee() string {
	return g.FraudCoverageFee
}

// installment returns the installment from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetInstallment() string {
	return g.Installment
}

// split_id returns the split_id from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetSplitId() string {
	return g.SplitId
}

// bulk_anticipation_id returns the bulk_anticipation_id from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetBulkAnticipationId() string {
	return g.BulkAnticipationId
}

// anticipation_id returns the anticipation_id from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetAnticipationId() string {
	return g.AnticipationId
}

// recipient_id returns the recipient_id from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetRecipientId() string {
	return g.RecipientId
}

// originator_model returns the originator_model from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetOriginatorModel() string {
	return g.OriginatorModel
}

// originator_model_id returns the originator_model_id from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetOriginatorModelId() string {
	return g.OriginatorModelId
}

// payment_date returns the payment_date from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetPaymentDate() string {
	return g.PaymentDate
}

// original_payment_date returns the original_payment_date from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetOriginalPaymentDate() string {
	return g.OriginalPaymentDate
}

// payment_method returns the payment_method from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetPaymentMethod() string {
	return g.PaymentMethod
}

// accrual_at returns the accrual_at from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetAccrualAt() string {
	return g.AccrualAt
}

// liquidation_arrangement_id returns the liquidation_arrangement_id from the GetMovementObjectPayableResponse.
func (g *GetMovementObjectPayableResponse) GetLiquidationArrangementId() string {
	return g.LiquidationArrangementId
}

// UnmarshalGetMovementObjectPayableResponse is a utility function used to unmarshal JSON into a GetMovementObjectPayableResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetMovementObjectPayableResponseInterface and an error, if any.
func UnmarshalGetMovementObjectPayableResponse(
	inputType string,
	input []byte) (
	GetMovementObjectPayableResponseInterface,
	error) {
	getMovementObjectBaseResponse, err := UnmarshalGetMovementObjectBaseResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getMovementObjectPayableResponse, ok := getMovementObjectBaseResponse.(GetMovementObjectPayableResponseInterface); ok {
		return getMovementObjectPayableResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getMovementObjectPayableResponse %v", getMovementObjectBaseResponse)
}

// ToGetMovementObjectPayableResponseArray converts an array of GetMovementObjectPayableResponseField to an array of GetMovementObjectPayableResponseInterface.
func ToGetMovementObjectPayableResponseArray(getMovementObjectPayableResponse []GetMovementObjectPayableResponseField) []GetMovementObjectPayableResponseInterface {
	var items []GetMovementObjectPayableResponseInterface
	for _, temp := range getMovementObjectPayableResponse {
		items = append(items, temp.GetMovementObjectPayableResponseInterface)
	}
	return items
}

// GetMovementObjectPayableResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetMovementObjectPayableResponseField struct {
	GetMovementObjectPayableResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectPayableResponseField.
// It customizes the JSON unmarshaling process for GetMovementObjectPayableResponseField objects.
func (g *GetMovementObjectPayableResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetMovementObjectPayableResponseInterface, err = UnmarshalGetMovementObjectPayableResponse("payable", input)
	return err
}

// GetMovementObjectPayableResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetMovementObjectPayableResponseSlice struct {
	Slice []GetMovementObjectPayableResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetMovementObjectPayableResponseSlice.
// It customizes the JSON unmarshaling process for GetMovementObjectPayableResponseSlice objects.
func (ga *GetMovementObjectPayableResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetMovementObjectPayableResponseInterface{}
		for _, getMovementObjectPayableResponse := range temp {
			if getMovementObjectPayableResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetMovementObjectPayableResponseInterface
			obj, err := UnmarshalGetMovementObjectPayableResponse("payable", getMovementObjectPayableResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
