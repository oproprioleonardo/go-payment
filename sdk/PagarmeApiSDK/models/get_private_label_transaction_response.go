package models

import (
	"encoding/json"
	"fmt"
)

// GetPrivateLabelTransactionResponse represents a GetPrivateLabelTransactionResponse struct.
// Response object for getting a private label transaction
type GetPrivateLabelTransactionResponse struct {
	GetTransactionResponse
	// Text that will appear on the credit card's statement
	StatementDescriptor Optional[string] `json:"statement_descriptor"`
	// Acquirer name
	AcquirerName Optional[string] `json:"acquirer_name"`
	// Aquirer affiliation code
	AcquirerAffiliationCode Optional[string] `json:"acquirer_affiliation_code"`
	// Acquirer TID
	AcquirerTid Optional[string] `json:"acquirer_tid"`
	// Acquirer NSU
	AcquirerNsu Optional[string] `json:"acquirer_nsu"`
	// Acquirer authorization code
	AcquirerAuthCode Optional[string] `json:"acquirer_auth_code"`
	// Operation type
	OperationType Optional[string] `json:"operation_type"`
	// Card data
	Card Optional[GetCardResponse] `json:"card"`
	// Acquirer message
	AcquirerMessage Optional[string] `json:"acquirer_message"`
	// Acquirer Return Code
	AcquirerReturnCode Optional[string] `json:"acquirer_return_code"`
	// Number of installments
	Installments Optional[int] `json:"installments"`
}

// MarshalJSON implements the json.Marshaler interface for GetPrivateLabelTransactionResponse.
// It customizes the JSON marshaling process for GetPrivateLabelTransactionResponse objects.
func (g *GetPrivateLabelTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetPrivateLabelTransactionResponse object to a map representation for JSON marshaling.
func (g *GetPrivateLabelTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.GetTransactionResponse.TransactionType != nil {
		structMap["transaction_type"] = *g.GetTransactionResponse.TransactionType
	} else {
		structMap["transaction_type"] = "private_label"
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
	if g.OperationType.IsValueSet() {
		structMap["operation_type"] = g.OperationType.Value()
	}
	if g.Card.IsValueSet() {
		structMap["card"] = g.Card.Value()
	}
	if g.AcquirerMessage.IsValueSet() {
		structMap["acquirer_message"] = g.AcquirerMessage.Value()
	}
	if g.AcquirerReturnCode.IsValueSet() {
		structMap["acquirer_return_code"] = g.AcquirerReturnCode.Value()
	}
	if g.Installments.IsValueSet() {
		structMap["installments"] = g.Installments.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPrivateLabelTransactionResponse.
// It customizes the JSON unmarshaling process for GetPrivateLabelTransactionResponse objects.
func (g *GetPrivateLabelTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		StatementDescriptor     Optional[string]          `json:"statement_descriptor"`
		AcquirerName            Optional[string]          `json:"acquirer_name"`
		AcquirerAffiliationCode Optional[string]          `json:"acquirer_affiliation_code"`
		AcquirerTid             Optional[string]          `json:"acquirer_tid"`
		AcquirerNsu             Optional[string]          `json:"acquirer_nsu"`
		AcquirerAuthCode        Optional[string]          `json:"acquirer_auth_code"`
		OperationType           Optional[string]          `json:"operation_type"`
		Card                    Optional[GetCardResponse] `json:"card"`
		AcquirerMessage         Optional[string]          `json:"acquirer_message"`
		AcquirerReturnCode      Optional[string]          `json:"acquirer_return_code"`
		Installments            Optional[int]             `json:"installments"`
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
	g.OperationType = temp.OperationType
	g.Card = temp.Card
	g.AcquirerMessage = temp.AcquirerMessage
	g.AcquirerReturnCode = temp.AcquirerReturnCode
	g.Installments = temp.Installments
	return nil
}

// statement_descriptor returns the statement_descriptor from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetStatementDescriptor() Optional[string] {
	return g.StatementDescriptor
}

// acquirer_name returns the acquirer_name from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerName() Optional[string] {
	return g.AcquirerName
}

// acquirer_affiliation_code returns the acquirer_affiliation_code from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerAffiliationCode() Optional[string] {
	return g.AcquirerAffiliationCode
}

// acquirer_tid returns the acquirer_tid from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerTid() Optional[string] {
	return g.AcquirerTid
}

// acquirer_nsu returns the acquirer_nsu from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerNsu() Optional[string] {
	return g.AcquirerNsu
}

// acquirer_auth_code returns the acquirer_auth_code from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerAuthCode() Optional[string] {
	return g.AcquirerAuthCode
}

// operation_type returns the operation_type from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetOperationType() Optional[string] {
	return g.OperationType
}

// card returns the card from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetCard() Optional[GetCardResponse] {
	return g.Card
}

// acquirer_message returns the acquirer_message from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerMessage() Optional[string] {
	return g.AcquirerMessage
}

// acquirer_return_code returns the acquirer_return_code from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerReturnCode() Optional[string] {
	return g.AcquirerReturnCode
}

// installments returns the installments from the GetPrivateLabelTransactionResponse.
func (g *GetPrivateLabelTransactionResponse) GetInstallments() Optional[int] {
	return g.Installments
}

// UnmarshalGetPrivateLabelTransactionResponse is a utility function used to unmarshal JSON into a GetPrivateLabelTransactionResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetPrivateLabelTransactionResponseInterface and an error, if any.
func UnmarshalGetPrivateLabelTransactionResponse(
	inputType string,
	input []byte) (
	GetPrivateLabelTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getPrivateLabelTransactionResponse, ok := getTransactionResponse.(GetPrivateLabelTransactionResponseInterface); ok {
		return getPrivateLabelTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getPrivateLabelTransactionResponse %v", getTransactionResponse)
}

// ToGetPrivateLabelTransactionResponseArray converts an array of GetPrivateLabelTransactionResponseField to an array of GetPrivateLabelTransactionResponseInterface.
func ToGetPrivateLabelTransactionResponseArray(getPrivateLabelTransactionResponse []GetPrivateLabelTransactionResponseField) []GetPrivateLabelTransactionResponseInterface {
	var items []GetPrivateLabelTransactionResponseInterface
	for _, temp := range getPrivateLabelTransactionResponse {
		items = append(items, temp.GetPrivateLabelTransactionResponseInterface)
	}
	return items
}

// GetPrivateLabelTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetPrivateLabelTransactionResponseField struct {
	GetPrivateLabelTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPrivateLabelTransactionResponseField.
// It customizes the JSON unmarshaling process for GetPrivateLabelTransactionResponseField objects.
func (g *GetPrivateLabelTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetPrivateLabelTransactionResponseInterface, err = UnmarshalGetPrivateLabelTransactionResponse("private_label", input)
	return err
}

// GetPrivateLabelTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetPrivateLabelTransactionResponseSlice struct {
	Slice []GetPrivateLabelTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetPrivateLabelTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetPrivateLabelTransactionResponseSlice objects.
func (ga *GetPrivateLabelTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetPrivateLabelTransactionResponseInterface{}
		for _, getPrivateLabelTransactionResponse := range temp {
			if getPrivateLabelTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetPrivateLabelTransactionResponseInterface
			obj, err := UnmarshalGetPrivateLabelTransactionResponse("private_label", getPrivateLabelTransactionResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
