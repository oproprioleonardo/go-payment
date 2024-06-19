package models

import (
	"encoding/json"
)

// CreateGooglePayIntermediateSigningKeyRequest represents a CreateGooglePayIntermediateSigningKeyRequest struct.
// The GooglePay Intermediate Signing Key Request
type CreateGooglePayIntermediateSigningKeyRequest struct {
	// Uma mensagem codificada em Base64 com a descrição de pagamento da chave.
	SignedKey Optional[string] `json:"signed_key"`
	// Verifica se a origem da chave de assinatura intermediária é o Google. É codificada em Base64 e criada usando o ECDSA.
	Signatures Optional[[]string] `json:"signatures"`
}

// MarshalJSON implements the json.Marshaler interface for CreateGooglePayIntermediateSigningKeyRequest.
// It customizes the JSON marshaling process for CreateGooglePayIntermediateSigningKeyRequest objects.
func (c *CreateGooglePayIntermediateSigningKeyRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateGooglePayIntermediateSigningKeyRequest object to a map representation for JSON marshaling.
func (c *CreateGooglePayIntermediateSigningKeyRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.SignedKey.IsValueSet() {
		structMap["signed_key"] = c.SignedKey.Value()
	}
	if c.Signatures.IsValueSet() {
		structMap["signatures"] = c.Signatures.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateGooglePayIntermediateSigningKeyRequest.
// It customizes the JSON unmarshaling process for CreateGooglePayIntermediateSigningKeyRequest objects.
func (c *CreateGooglePayIntermediateSigningKeyRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		SignedKey  Optional[string]   `json:"signed_key"`
		Signatures Optional[[]string] `json:"signatures"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.SignedKey = temp.SignedKey
	c.Signatures = temp.Signatures
	return nil
}
