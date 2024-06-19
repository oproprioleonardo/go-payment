package models

import (
	"encoding/json"
	"fmt"
)

// GetDebitCardTransactionResponse represents a GetDebitCardTransactionResponse struct.
// Response object for getting a debit card transaction
type GetDebitCardTransactionResponse struct {
	GetTransactionResponse
	// Text that will appear on the debit card's statement
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
	// Merchant Plugin
	Mpi Optional[string] `json:"mpi"`
	// Electronic Commerce Indicator (ECI)
	Eci Optional[string] `json:"eci"`
	// Authentication type
	AuthenticationType Optional[string] `json:"authentication_type"`
	// 3D-S Authentication Url
	ThreedAuthenticationUrl Optional[string] `json:"threed_authentication_url"`
	// Identify when a card is prepaid, credit or debit.
	FundingSource Optional[string] `json:"funding_source"`
}

// MarshalJSON implements the json.Marshaler interface for GetDebitCardTransactionResponse.
// It customizes the JSON marshaling process for GetDebitCardTransactionResponse objects.
func (g *GetDebitCardTransactionResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetDebitCardTransactionResponse object to a map representation for JSON marshaling.
func (g *GetDebitCardTransactionResponse) toMap() map[string]any {
	structMap := g.GetTransactionResponse.toMap()
	if g.GetTransactionResponse.TransactionType != nil {
		structMap["transaction_type"] = *g.GetTransactionResponse.TransactionType
	} else {
		structMap["transaction_type"] = "debit_card"
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
	if g.Mpi.IsValueSet() {
		structMap["mpi"] = g.Mpi.Value()
	}
	if g.Eci.IsValueSet() {
		structMap["eci"] = g.Eci.Value()
	}
	if g.AuthenticationType.IsValueSet() {
		structMap["authentication_type"] = g.AuthenticationType.Value()
	}
	if g.ThreedAuthenticationUrl.IsValueSet() {
		structMap["threed_authentication_url"] = g.ThreedAuthenticationUrl.Value()
	}
	if g.FundingSource.IsValueSet() {
		structMap["funding_source"] = g.FundingSource.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetDebitCardTransactionResponse.
// It customizes the JSON unmarshaling process for GetDebitCardTransactionResponse objects.
func (g *GetDebitCardTransactionResponse) UnmarshalJSON(input []byte) error {
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
		Mpi                     Optional[string]          `json:"mpi"`
		Eci                     Optional[string]          `json:"eci"`
		AuthenticationType      Optional[string]          `json:"authentication_type"`
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
	g.Mpi = temp.Mpi
	g.Eci = temp.Eci
	g.AuthenticationType = temp.AuthenticationType
	g.ThreedAuthenticationUrl = temp.ThreedAuthenticationUrl
	g.FundingSource = temp.FundingSource
	return nil
}

// statement_descriptor returns the statement_descriptor from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetStatementDescriptor() Optional[string] {
	return g.StatementDescriptor
}

// acquirer_name returns the acquirer_name from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerName() Optional[string] {
	return g.AcquirerName
}

// acquirer_affiliation_code returns the acquirer_affiliation_code from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerAffiliationCode() Optional[string] {
	return g.AcquirerAffiliationCode
}

// acquirer_tid returns the acquirer_tid from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerTid() Optional[string] {
	return g.AcquirerTid
}

// acquirer_nsu returns the acquirer_nsu from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerNsu() Optional[string] {
	return g.AcquirerNsu
}

// acquirer_auth_code returns the acquirer_auth_code from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerAuthCode() Optional[string] {
	return g.AcquirerAuthCode
}

// operation_type returns the operation_type from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetOperationType() Optional[string] {
	return g.OperationType
}

// card returns the card from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetCard() Optional[GetCardResponse] {
	return g.Card
}

// acquirer_message returns the acquirer_message from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerMessage() Optional[string] {
	return g.AcquirerMessage
}

// acquirer_return_code returns the acquirer_return_code from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAcquirerReturnCode() Optional[string] {
	return g.AcquirerReturnCode
}

// mpi returns the mpi from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetMpi() Optional[string] {
	return g.Mpi
}

// eci returns the eci from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetEci() Optional[string] {
	return g.Eci
}

// authentication_type returns the authentication_type from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetAuthenticationType() Optional[string] {
	return g.AuthenticationType
}

// threed_authentication_url returns the threed_authentication_url from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetThreedAuthenticationUrl() Optional[string] {
	return g.ThreedAuthenticationUrl
}

// funding_source returns the funding_source from the GetDebitCardTransactionResponse.
func (g *GetDebitCardTransactionResponse) GetFundingSource() Optional[string] {
	return g.FundingSource
}

// UnmarshalGetDebitCardTransactionResponse is a utility function used to unmarshal JSON into a GetDebitCardTransactionResponseInterface.
// It takes a inputType string, JSON input []byte and returns the corresponding GetDebitCardTransactionResponseInterface and an error, if any.
func UnmarshalGetDebitCardTransactionResponse(
	inputType string,
	input []byte) (
	GetDebitCardTransactionResponseInterface,
	error) {
	getTransactionResponse, err := UnmarshalGetTransactionResponse(inputType, input)
	if err != nil {
		return nil, err
	}

	if getDebitCardTransactionResponse, ok := getTransactionResponse.(GetDebitCardTransactionResponseInterface); ok {
		return getDebitCardTransactionResponse, nil
	}
	return nil, fmt.Errorf("Cannot unmarshal getDebitCardTransactionResponse %v", getTransactionResponse)
}

// ToGetDebitCardTransactionResponseArray converts an array of GetDebitCardTransactionResponseField to an array of GetDebitCardTransactionResponseInterface.
func ToGetDebitCardTransactionResponseArray(getDebitCardTransactionResponse []GetDebitCardTransactionResponseField) []GetDebitCardTransactionResponseInterface {
	var items []GetDebitCardTransactionResponseInterface
	for _, temp := range getDebitCardTransactionResponse {
		items = append(items, temp.GetDebitCardTransactionResponseInterface)
	}
	return items
}

// GetDebitCardTransactionResponseField is a utility struct that helps the go JSON deserializer invoke the proper deserialization logic.
type GetDebitCardTransactionResponseField struct {
	GetDebitCardTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetDebitCardTransactionResponseField.
// It customizes the JSON unmarshaling process for GetDebitCardTransactionResponseField objects.
func (g *GetDebitCardTransactionResponseField) UnmarshalJSON(input []byte) error {
	var err error
	g.GetDebitCardTransactionResponseInterface, err = UnmarshalGetDebitCardTransactionResponse("debit_card", input)
	return err
}

// GetDebitCardTransactionResponseSlice is a utility struct that helps in the unmarshalling of slices.
type GetDebitCardTransactionResponseSlice struct {
	Slice []GetDebitCardTransactionResponseInterface
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetDebitCardTransactionResponseSlice.
// It customizes the JSON unmarshaling process for GetDebitCardTransactionResponseSlice objects.
func (ga *GetDebitCardTransactionResponseSlice) UnmarshalJSON(input []byte) error {
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
		ga.Slice = []GetDebitCardTransactionResponseInterface{}
		for _, getDebitCardTransactionResponse := range temp {
			if getDebitCardTransactionResponse == nil {
				ga.Slice = append(ga.Slice, nil)
			}
			var obj GetDebitCardTransactionResponseInterface
			obj, err := UnmarshalGetDebitCardTransactionResponse("debit_card", getDebitCardTransactionResponse)
			if err != nil {
				return err
			}
			ga.Slice = append(ga.Slice, obj)
		}
	}
	return nil
}
