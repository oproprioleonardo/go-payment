package models

import (
	"encoding/json"
)

// UpdateChargePaymentMethodRequest represents a UpdateChargePaymentMethodRequest struct.
// Request for updating the payment method of a charge
type UpdateChargePaymentMethodRequest struct {
	// Indicates if the payment method from the subscription must also be updated
	UpdateSubscription bool `json:"update_subscription"`
	// The new payment method
	PaymentMethod string `json:"payment_method"`
	// Credit card data
	CreditCard CreateCreditCardPaymentRequest `json:"credit_card"`
	// Debit card data
	DebitCard CreateDebitCardPaymentRequest `json:"debit_card"`
	// Boleto data
	Boleto CreateBoletoPaymentRequest `json:"boleto"`
	// Voucher data
	Voucher CreateVoucherPaymentRequest `json:"voucher"`
	// Cash data
	Cash CreateCashPaymentRequest `json:"cash"`
	// Bank Transfer data
	BankTransfer CreateBankTransferPaymentRequest `json:"bank_transfer"`
	PrivateLabel CreatePrivateLabelPaymentRequest `json:"private_label"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateChargePaymentMethodRequest.
// It customizes the JSON marshaling process for UpdateChargePaymentMethodRequest objects.
func (u *UpdateChargePaymentMethodRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateChargePaymentMethodRequest object to a map representation for JSON marshaling.
func (u *UpdateChargePaymentMethodRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["update_subscription"] = u.UpdateSubscription
	structMap["payment_method"] = u.PaymentMethod
	structMap["credit_card"] = u.CreditCard
	structMap["debit_card"] = u.DebitCard
	structMap["boleto"] = u.Boleto
	structMap["voucher"] = u.Voucher
	structMap["cash"] = u.Cash
	structMap["bank_transfer"] = u.BankTransfer
	structMap["private_label"] = u.PrivateLabel
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateChargePaymentMethodRequest.
// It customizes the JSON unmarshaling process for UpdateChargePaymentMethodRequest objects.
func (u *UpdateChargePaymentMethodRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		UpdateSubscription bool                             `json:"update_subscription"`
		PaymentMethod      string                           `json:"payment_method"`
		CreditCard         CreateCreditCardPaymentRequest   `json:"credit_card"`
		DebitCard          CreateDebitCardPaymentRequest    `json:"debit_card"`
		Boleto             CreateBoletoPaymentRequest       `json:"boleto"`
		Voucher            CreateVoucherPaymentRequest      `json:"voucher"`
		Cash               CreateCashPaymentRequest         `json:"cash"`
		BankTransfer       CreateBankTransferPaymentRequest `json:"bank_transfer"`
		PrivateLabel       CreatePrivateLabelPaymentRequest `json:"private_label"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.UpdateSubscription = temp.UpdateSubscription
	u.PaymentMethod = temp.PaymentMethod
	u.CreditCard = temp.CreditCard
	u.DebitCard = temp.DebitCard
	u.Boleto = temp.Boleto
	u.Voucher = temp.Voucher
	u.Cash = temp.Cash
	u.BankTransfer = temp.BankTransfer
	u.PrivateLabel = temp.PrivateLabel
	return nil
}
