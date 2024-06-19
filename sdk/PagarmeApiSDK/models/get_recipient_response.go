package models

import (
	"encoding/json"
	"log"
	"time"
)

// GetRecipientResponse represents a GetRecipientResponse struct.
// Recipient response
type GetRecipientResponse struct {
	// Id
	Id Optional[string] `json:"id"`
	// Name
	Name Optional[string] `json:"name"`
	// Email
	Email Optional[string] `json:"email"`
	// Document
	Document Optional[string] `json:"document"`
	// Description
	Description Optional[string] `json:"description"`
	// Type
	Type Optional[string] `json:"type"`
	// Status
	Status Optional[string] `json:"status"`
	// Creation date
	CreatedAt Optional[time.Time] `json:"created_at"`
	// Last update date
	UpdatedAt Optional[time.Time] `json:"updated_at"`
	// Deletion date
	DeletedAt Optional[time.Time] `json:"deleted_at"`
	// Default bank account
	DefaultBankAccount Optional[GetBankAccountResponse] `json:"default_bank_account"`
	// Info about the recipient on the gateway
	GatewayRecipients Optional[[]GetGatewayRecipientResponse] `json:"gateway_recipients"`
	// Metadata
	Metadata                      Optional[map[string]string]                `json:"metadata"`
	AutomaticAnticipationSettings Optional[GetAutomaticAnticipationResponse] `json:"automatic_anticipation_settings"`
	TransferSettings              Optional[GetTransferSettingsResponse]      `json:"transfer_settings"`
	// Recipient code
	Code Optional[string] `json:"code"`
	// Payment mode
	PaymentMode Optional[string] `json:"payment_mode"`
}

// MarshalJSON implements the json.Marshaler interface for GetRecipientResponse.
// It customizes the JSON marshaling process for GetRecipientResponse objects.
func (g *GetRecipientResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GetRecipientResponse object to a map representation for JSON marshaling.
func (g *GetRecipientResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if g.Id.IsValueSet() {
		structMap["id"] = g.Id.Value()
	}
	if g.Name.IsValueSet() {
		structMap["name"] = g.Name.Value()
	}
	if g.Email.IsValueSet() {
		structMap["email"] = g.Email.Value()
	}
	if g.Document.IsValueSet() {
		structMap["document"] = g.Document.Value()
	}
	if g.Description.IsValueSet() {
		structMap["description"] = g.Description.Value()
	}
	if g.Type.IsValueSet() {
		structMap["type"] = g.Type.Value()
	}
	if g.Status.IsValueSet() {
		structMap["status"] = g.Status.Value()
	}
	if g.CreatedAt.IsValueSet() {
		var CreatedAtVal *string = nil
		if g.CreatedAt.Value() != nil {
			val := g.CreatedAt.Value().Format(time.RFC3339)
			CreatedAtVal = &val
		}
		structMap["created_at"] = CreatedAtVal
	}
	if g.UpdatedAt.IsValueSet() {
		var UpdatedAtVal *string = nil
		if g.UpdatedAt.Value() != nil {
			val := g.UpdatedAt.Value().Format(time.RFC3339)
			UpdatedAtVal = &val
		}
		structMap["updated_at"] = UpdatedAtVal
	}
	if g.DeletedAt.IsValueSet() {
		var DeletedAtVal *string = nil
		if g.DeletedAt.Value() != nil {
			val := g.DeletedAt.Value().Format(time.RFC3339)
			DeletedAtVal = &val
		}
		structMap["deleted_at"] = DeletedAtVal
	}
	if g.DefaultBankAccount.IsValueSet() {
		structMap["default_bank_account"] = g.DefaultBankAccount.Value()
	}
	if g.GatewayRecipients.IsValueSet() {
		structMap["gateway_recipients"] = g.GatewayRecipients.Value()
	}
	if g.Metadata.IsValueSet() {
		structMap["metadata"] = g.Metadata.Value()
	}
	if g.AutomaticAnticipationSettings.IsValueSet() {
		structMap["automatic_anticipation_settings"] = g.AutomaticAnticipationSettings.Value()
	}
	if g.TransferSettings.IsValueSet() {
		structMap["transfer_settings"] = g.TransferSettings.Value()
	}
	if g.Code.IsValueSet() {
		structMap["code"] = g.Code.Value()
	}
	if g.PaymentMode.IsValueSet() {
		structMap["payment_mode"] = g.PaymentMode.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetRecipientResponse.
// It customizes the JSON unmarshaling process for GetRecipientResponse objects.
func (g *GetRecipientResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                            Optional[string]                           `json:"id"`
		Name                          Optional[string]                           `json:"name"`
		Email                         Optional[string]                           `json:"email"`
		Document                      Optional[string]                           `json:"document"`
		Description                   Optional[string]                           `json:"description"`
		Type                          Optional[string]                           `json:"type"`
		Status                        Optional[string]                           `json:"status"`
		CreatedAt                     Optional[string]                           `json:"created_at"`
		UpdatedAt                     Optional[string]                           `json:"updated_at"`
		DeletedAt                     Optional[string]                           `json:"deleted_at"`
		DefaultBankAccount            Optional[GetBankAccountResponse]           `json:"default_bank_account"`
		GatewayRecipients             Optional[[]GetGatewayRecipientResponse]    `json:"gateway_recipients"`
		Metadata                      Optional[map[string]string]                `json:"metadata"`
		AutomaticAnticipationSettings Optional[GetAutomaticAnticipationResponse] `json:"automatic_anticipation_settings"`
		TransferSettings              Optional[GetTransferSettingsResponse]      `json:"transfer_settings"`
		Code                          Optional[string]                           `json:"code"`
		PaymentMode                   Optional[string]                           `json:"payment_mode"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Id = temp.Id
	g.Name = temp.Name
	g.Email = temp.Email
	g.Document = temp.Document
	g.Description = temp.Description
	g.Type = temp.Type
	g.Status = temp.Status
	g.CreatedAt.ShouldSetValue(temp.CreatedAt.IsValueSet())
	if temp.CreatedAt.Value() != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, (*temp.CreatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		g.CreatedAt.SetValue(&CreatedAtVal)
	}
	g.UpdatedAt.ShouldSetValue(temp.UpdatedAt.IsValueSet())
	if temp.UpdatedAt.Value() != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, (*temp.UpdatedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		g.UpdatedAt.SetValue(&UpdatedAtVal)
	}
	g.DeletedAt.ShouldSetValue(temp.DeletedAt.IsValueSet())
	if temp.DeletedAt.Value() != nil {
		DeletedAtVal, err := time.Parse(time.RFC3339, (*temp.DeletedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse deleted_at as % s format.", time.RFC3339)
		}
		g.DeletedAt.SetValue(&DeletedAtVal)
	}
	g.DefaultBankAccount = temp.DefaultBankAccount
	g.GatewayRecipients = temp.GatewayRecipients
	g.Metadata = temp.Metadata
	g.AutomaticAnticipationSettings = temp.AutomaticAnticipationSettings
	g.TransferSettings = temp.TransferSettings
	g.Code = temp.Code
	g.PaymentMode = temp.PaymentMode
	return nil
}
