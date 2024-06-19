package models

import (
	"encoding/json"
)

// CreateGooglePayRequest represents a CreateGooglePayRequest struct.
// The GooglePay Token Payment Request
type CreateGooglePayRequest struct {
	// Informação sobre a versão do token. Único valor aceito é EC_v2
	Version Optional[string] `json:"version"`
	// Dados de pagamento criptografados. Corresponde ao encryptedMessage do token Google.
	Data Optional[string] `json:"data"`
	// The GooglePay intermediate signing key request
	IntermediateSigningKey Optional[CreateGooglePayIntermediateSigningKeyRequest] `json:"intermediate_signing_key"`
	// Assinatura dos dados de pagamento. Verifica se a origem da mensagem é o Google. Corresponde ao signature do token Google.
	Signature          Optional[string] `json:"signature"`
	SignedMessage      Optional[string] `json:"signed_message"`
	MerchantIdentifier Optional[string] `json:"merchant_identifier"`
}

// MarshalJSON implements the json.Marshaler interface for CreateGooglePayRequest.
// It customizes the JSON marshaling process for CreateGooglePayRequest objects.
func (c *CreateGooglePayRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateGooglePayRequest object to a map representation for JSON marshaling.
func (c *CreateGooglePayRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Version.IsValueSet() {
		structMap["version"] = c.Version.Value()
	}
	if c.Data.IsValueSet() {
		structMap["data"] = c.Data.Value()
	}
	if c.IntermediateSigningKey.IsValueSet() {
		structMap["intermediate_signing_key"] = c.IntermediateSigningKey.Value()
	}
	if c.Signature.IsValueSet() {
		structMap["signature"] = c.Signature.Value()
	}
	if c.SignedMessage.IsValueSet() {
		structMap["signed_message"] = c.SignedMessage.Value()
	}
	if c.MerchantIdentifier.IsValueSet() {
		structMap["merchant_identifier"] = c.MerchantIdentifier.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateGooglePayRequest.
// It customizes the JSON unmarshaling process for CreateGooglePayRequest objects.
func (c *CreateGooglePayRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Version                Optional[string]                                       `json:"version"`
		Data                   Optional[string]                                       `json:"data"`
		IntermediateSigningKey Optional[CreateGooglePayIntermediateSigningKeyRequest] `json:"intermediate_signing_key"`
		Signature              Optional[string]                                       `json:"signature"`
		SignedMessage          Optional[string]                                       `json:"signed_message"`
		MerchantIdentifier     Optional[string]                                       `json:"merchant_identifier"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Version = temp.Version
	c.Data = temp.Data
	c.IntermediateSigningKey = temp.IntermediateSigningKey
	c.Signature = temp.Signature
	c.SignedMessage = temp.SignedMessage
	c.MerchantIdentifier = temp.MerchantIdentifier
	return nil
}
