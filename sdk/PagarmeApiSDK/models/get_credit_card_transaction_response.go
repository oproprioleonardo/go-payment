package models

import (
	"encoding/json"
	"fmt"
)

// GetCreditCardTransactionResponse represents a GetCreditCardTransactionResponse struct.
// Response object for getting a credit card transaction
type GetCreditCardTransactionResponse struct {
	GetTransactionResponse
	// Text that will appear on the credit card's statement
	StatementDescriptor Optional[string] `json:"statement_descriptor"`
	// Acquirer name
	AcquirerName *string `json:"acquirer_name,omitempty"`
	// Aquirer affiliation code
	AcquirerAffiliationCode Optional[string] `json:"acquirer_affiliation_code"`
	// Acquirer TID
	AcquirerTid *string `json:"acquirer_tid,omitempty"`
	// Acquirer NSU
	AcquirerNsu *string `json:"acquirer_nsu,omitempty"`
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
	// 3D-S authentication Url
	ThreedAuthenticationUrl Optional[string] `json:"threed_authentication_url"`
	// Identify when a card is prepaid, credit or debit.
	FundingSource Optional[string] `json:"funding_source"`
}

// MarshalJSON implements the json.Marshaler interface for GetCreditCardTransactionResponse.
// It customizes the JSON marshaling process for GetCreditCardTransactionResponse objects.
func (g *GetCreditCardTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetCreditCardTransactionResponse object to a map representation for JSON marshaling.
func (g *GetCreditCardTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.GetTransactionResponse.TransactionType != nil {
		structMap["transaction_type"] = *g.GetTransactionResponse.TransactionType
	} else {
		structMap["transaction_type"] = "credit_card"
	}
	if g.StatementDescriptor.IsValueSet() {
		structMap["statement_descriptor"] = g.StatementDescriptor.Value()
	}
	if g.AcquirerName != nil {
		structMap["acquirer_name"] = g.AcquirerName
	}
	if g.AcquirerAffiliationCode.IsValueSet() {
		structMap["acquirer_affiliation_code"] = g.AcquirerAffiliationCode.Value()
	}
	if g.AcquirerTid != nil {
		structMap["acquirer_tid"] = g.AcquirerTid
	}
	if g.AcquirerNsu != nil {
		structMap["acquirer_nsu"] = g.AcquirerNsu
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
	if g.ThreedAuthenticationUrl.IsValueSet() {
		structMap["threed_authentication_url"] = g.ThreedAuthenticationUrl.Value()
	}
	if g.FundingSource.IsValueSet() {
		structMap["funding_source"] = g.FundingSource.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCreditCardTransactionResponse.
// It customizes the JSON unmarshaling process for GetCreditCardTransactionResponse objects.
func (g *GetCreditCardTransactionResponse) UnmarshalJSON(input []byte) error {
	g.GetTransactionResponse.UnmarshalJSON(input)
	temp := &struct {
		StatementDescriptor     Optional[string]          `json:"statement_descriptor"`
		AcquirerName            *string                   `json:"acquirer_name,omitempty"`
		AcquirerAffiliationCode Optional[string]          `json:"acquirer_affiliation_code"`
		AcquirerTid             *string                   `json:"acquirer_tid,omitempty"`
		AcquirerNsu             *string                   `json:"acquirer_nsu,omitempty"`
		AcquirerAuthCode        Optional[string]          `json:"acquirer_auth_code"`
		OperationType           Optional[string]          `json:"operation_type"`
		Card                    Optional[GetCardResponse] `json:"card"`
		AcquirerMessage         Optional[string]          `json:"acquirer_message"`
		AcquirerReturnCode      Optional[string]          `json:"acquirer_return_code"`
		Installments            Optional[int]             `json:"installments"`
		ThreedAuthenticationUrl Optional[string]          `json:"threed_authentication_url"`
		FundingSource           Optional[string]          `json:"funding_source"`
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
	g.ThreedAuthenticationUrl = temp.ThreedAuthenticationUrl
	g.FundingSource = temp.FundingSource
	return nil
}

// statement_descriptor returns the statement_descriptor from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetStatementDescriptor() Optional[string] {
	return g.StatementDescriptor
}

// acquirer_name returns the acquirer_name from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerName() *string {
	return g.AcquirerName
}

// acquirer_affiliation_code returns the acquirer_affiliation_code from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerAffiliationCode() Optional[string] {
	return g.AcquirerAffiliationCode
}

// acquirer_tid returns the acquirer_tid from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerTid() *string {
	return g.AcquirerTid
}

// acquirer_nsu returns the acquirer_nsu from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerNsu() *string {
	return g.AcquirerNsu
}

// acquirer_auth_code returns the acquirer_auth_code from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerAuthCode() Optional[string] {
	return g.AcquirerAuthCode
}

// operation_type returns the operation_type from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetOperationType() Optional[string] {
	return g.OperationType
}

// card returns the card from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetCard() Optional[GetCardResponse] {
	return g.Card
}

// acquirer_message returns the acquirer_message from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerMessage() Optional[string] {
	return g.AcquirerMessage
}

// acquirer_return_code returns the acquirer_return_code from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetAcquirerReturnCode() Optional[string] {
	return g.AcquirerReturnCode
}

// installments returns the installments from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetInstallments() Optional[int] {
	return g.Installments
}

// threed_authentication_url returns the threed_authentication_url from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetThreedAuthenticationUrl() Optional[string] {
	return g.ThreedAuthenticationUrl
}

// funding_source returns the funding_source from the GetCreditCardTransactionResponse.
func (g *GetCreditCardTransactionResponse) GetFundingSource() Optional[string] {
	return g.FundingSource
}

// UnmarshalGetCreditCardTransactionResponse is a utility function used to unmarshal JSON into a GetCreditCardTransactionResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetCreditCardTransactionResponseInterface and an error, if any.
func UnmarshalGetCreditCardTransactionResponse(
	inputType string,
	input []byte) (
	GetCreditCardTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getCreditCardTransactionResponse, ok := getTransactionResponse.(GetCreditCardTransactionResponseInterface); ok {
		return getCreditCardTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getCreditCardTransactionResponse %v", getTransactionResponse)
}

// ToGetCreditCardTransactionResponseArray converts an array of GetCreditCardTransactionResponseField to an array of GetCreditCardTransactionResponseInterface.
func ToGetCreditCardTransactionResponseArray(getCreditCardTransactionResponse []GetCreditCardTransactionResponseField) []GetCreditCardTransactionResponseInterface {
	var items []GetCreditCardTransactionResponseInterface
	for _, temp := range getCreditCardTransactionResponse {
		items = append(items, temp.GetCreditCardTransactionResponseInterface)
	}
	return items
}

// GetCreditCardTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetCreditCardTransactionResponseField struct {
	GetCreditCardTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCreditCardTransactionResponseField.
// It customizes the JSON unmarshaling process for GetCreditCardTransactionResponseField objects.
func (g *GetCreditCardTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetCreditCardTransactionResponseInterface, err = UnmarshalGetCreditCardTransactionResponse("credit_card", input)
	return err
}

// GetCreditCardTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetCreditCardTransactionResponseSlice struct {
	Slice []GetCreditCardTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetCreditCardTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetCreditCardTransactionResponseSlice objects.
func (ga *GetCreditCardTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetCreditCardTransactionResponseInterface{}
		for _, getCreditCardTransactionResponse := range temp {
			if getCreditCardTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetCreditCardTransactionResponseInterface
			obj, err := UnmarshalGetCreditCardTransactionResponse("credit_card", getCreditCardTransactionResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
