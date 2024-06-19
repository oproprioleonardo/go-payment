package models

import (
	"encoding/json"
	"fmt"
)

// GetVoucherTransactionResponse represents a GetVoucherTransactionResponse struct.
// Response for voucher transactions
type GetVoucherTransactionResponse struct {
	GetTransactionResponse
	// Text that will appear on the voucher's statement
	StatementDescriptor Optional[string] `json:"statement_descriptor"`
	// Acquirer name
	AcquirerName Optional[string] `json:"acquirer_name"`
	// Acquirer affiliation code
	AcquirerAffiliationCode Optional[string] `json:"acquirer_affiliation_code"`
	// Acquirer TID
	AcquirerTid Optional[string] `json:"acquirer_tid"`
	// Acquirer NSU
	AcquirerNsu Optional[string] `json:"acquirer_nsu"`
	// Acquirer authorization code
	AcquirerAuthCode Optional[string] `json:"acquirer_auth_code"`
	// acquirer_message
	AcquirerMessage Optional[string] `json:"acquirer_message"`
	// Acquirer return code
	AcquirerReturnCode Optional[string] `json:"acquirer_return_code"`
	// Operation type
	OperationType Optional[string] `json:"operation_type"`
	// Card data
	Card Optional[GetCardResponse] `json:"card"`
}

// MarshalJSON implements the json.Marshaler interface for GetVoucherTransactionResponse.
// It customizes the JSON marshaling process for GetVoucherTransactionResponse objects.
func (g *GetVoucherTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetVoucherTransactionResponse object to a map representation for JSON marshaling.
func (g *GetVoucherTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.GetTransactionResponse.TransactionType != nil {
		structMap["transaction_type"] = *g.GetTransactionResponse.TransactionType
	} else {
		structMap["transaction_type"] = "voucher"
	}
	if g.StatementDescriptor.IsValueSet() {
		structMap["statement_descriptor"] = g.StatementDescriptor.Value()
	}
	if g.AcquirerName.IsValueSet() {
		structMap["acquirer_name"] = g.AcquirerName.Value()
	}
	if g.AcquirerAffiliationCode.IsValueSet() {
		structMap["acquirer_affiliation_code"] = g.AcquirerAffiliationCode.Value()
	}
	if g.AcquirerTid.IsValueSet() {
		structMap["acquirer_tid"] = g.AcquirerTid.Value()
	}
	if g.AcquirerNsu.IsValueSet() {
		structMap["acquirer_nsu"] = g.AcquirerNsu.Value()
	}
	if g.AcquirerAuthCode.IsValueSet() {
		structMap["acquirer_auth_code"] = g.AcquirerAuthCode.Value()
	}
	if g.AcquirerMessage.IsValueSet() {
		structMap["acquirer_message"] = g.AcquirerMessage.Value()
	}
	if g.AcquirerReturnCode.IsValueSet() {
		structMap["acquirer_return_code"] = g.AcquirerReturnCode.Value()
	}
	if g.OperationType.IsValueSet() {
		structMap["operation_type"] = g.OperationType.Value()
	}
	if g.Card.IsValueSet() {
		structMap["card"] = g.Card.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetVoucherTransactionResponse.
// It customizes the JSON unmarshaling process for GetVoucherTransactionResponse objects.
func (g *GetVoucherTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		StatementDescriptor     Optional[string]          `json:"statement_descriptor"`
		AcquirerName            Optional[string]          `json:"acquirer_name"`
		AcquirerAffiliationCode Optional[string]          `json:"acquirer_affiliation_code"`
		AcquirerTid             Optional[string]          `json:"acquirer_tid"`
		AcquirerNsu             Optional[string]          `json:"acquirer_nsu"`
		AcquirerAuthCode        Optional[string]          `json:"acquirer_auth_code"`
		AcquirerMessage         Optional[string]          `json:"acquirer_message"`
		AcquirerReturnCode      Optional[string]          `json:"acquirer_return_code"`
		OperationType           Optional[string]          `json:"operation_type"`
		Card                    Optional[GetCardResponse] `json:"card"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.StatementDescriptor = temp.StatementDescriptor
	g.AcquirerName = temp.AcquirerName
	g.AcquirerAffiliationCode = temp.AcquirerAffiliationCode
	g.AcquirerTid = temp.AcquirerTid
	g.AcquirerNsu = temp.AcquirerNsu
	g.AcquirerAuthCode = temp.AcquirerAuthCode
	g.AcquirerMessage = temp.AcquirerMessage
	g.AcquirerReturnCode = temp.AcquirerReturnCode
	g.OperationType = temp.OperationType
	g.Card = temp.Card
	return nil
}

// statement_descriptor returns the statement_descriptor from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetStatementDescriptor() Optional[string] {
	return g.StatementDescriptor
}

// acquirer_name returns the acquirer_name from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerName() Optional[string] {
	return g.AcquirerName
}

// acquirer_affiliation_code returns the acquirer_affiliation_code from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerAffiliationCode() Optional[string] {
	return g.AcquirerAffiliationCode
}

// acquirer_tid returns the acquirer_tid from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerTid() Optional[string] {
	return g.AcquirerTid
}

// acquirer_nsu returns the acquirer_nsu from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerNsu() Optional[string] {
	return g.AcquirerNsu
}

// acquirer_auth_code returns the acquirer_auth_code from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerAuthCode() Optional[string] {
	return g.AcquirerAuthCode
}

// acquirer_message returns the acquirer_message from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerMessage() Optional[string] {
	return g.AcquirerMessage
}

// acquirer_return_code returns the acquirer_return_code from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetAcquirerReturnCode() Optional[string] {
	return g.AcquirerReturnCode
}

// operation_type returns the operation_type from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetOperationType() Optional[string] {
	return g.OperationType
}

// card returns the card from the GetVoucherTransactionResponse.
func (g *GetVoucherTransactionResponse) GetCard() Optional[GetCardResponse] {
	return g.Card
}

// UnmarshalGetVoucherTransactionResponse is a utility function used to unmarshal JSON into a GetVoucherTransactionResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetVoucherTransactionResponseInterface and an error, if any.
func UnmarshalGetVoucherTransactionResponse(
	inputType string,
	input []byte) (
	GetVoucherTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getVoucherTransactionResponse, ok := getTransactionResponse.(GetVoucherTransactionResponseInterface); ok {
		return getVoucherTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getVoucherTransactionResponse %v", getTransactionResponse)
}

// ToGetVoucherTransactionResponseArray converts an array of GetVoucherTransactionResponseField to an array of GetVoucherTransactionResponseInterface.
func ToGetVoucherTransactionResponseArray(getVoucherTransactionResponse []GetVoucherTransactionResponseField) []GetVoucherTransactionResponseInterface {
	var items []GetVoucherTransactionResponseInterface
	for _, temp := range getVoucherTransactionResponse {
		items = append(items, temp.GetVoucherTransactionResponseInterface)
	}
	return items
}

// GetVoucherTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetVoucherTransactionResponseField struct {
	GetVoucherTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetVoucherTransactionResponseField.
// It customizes the JSON unmarshaling process for GetVoucherTransactionResponseField objects.
func (g *GetVoucherTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetVoucherTransactionResponseInterface, err = UnmarshalGetVoucherTransactionResponse("voucher", input)
	return err
}

// GetVoucherTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetVoucherTransactionResponseSlice struct {
	Slice []GetVoucherTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetVoucherTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetVoucherTransactionResponseSlice objects.
func (ga *GetVoucherTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetVoucherTransactionResponseInterface{}
		for _, getVoucherTransactionResponse := range temp {
			if getVoucherTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetVoucherTransactionResponseInterface
			obj, err := UnmarshalGetVoucherTransactionResponse("voucher", getVoucherTransactionResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
