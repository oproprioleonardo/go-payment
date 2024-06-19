package models

import (
	"encoding/json"
)

// CreateCheckoutPaymentRequest represents a CreateCheckoutPaymentRequest struct.
// Checkout payment request
type CreateCheckoutPaymentRequest struct {
	// Accepted Payment Methods
	AcceptedPaymentMethods []string `json:"accepted_payment_methods"`
	// Accepted Multi Payment Methods
	AcceptedMultiPaymentMethods []interface{} `json:"accepted_multi_payment_methods"`
	// Success url
	SuccessUrl string `json:"success_url"`
	// Default payment method
	DefaultPaymentMethod *string `json:"default_payment_method,omitempty"`
	// Gateway Affiliation Id
	GatewayAffiliationId *string `json:"gateway_affiliation_id,omitempty"`
	// Credit Card payment request
	CreditCard *CreateCheckoutCreditCardPaymentRequest `json:"credit_card,omitempty"`
	// Debit Card payment request
	DebitCard *CreateCheckoutDebitCardPaymentRequest `json:"debit_card,omitempty"`
	// Boleto payment request
	Boleto *CreateCheckoutBoletoPaymentRequest `json:"boleto,omitempty"`
	// Customer is editable?
	CustomerEditable *bool `json:"customer_editable,omitempty"`
	// Time in minutes for expiration
	ExpiresIn *int `json:"expires_in,omitempty"`
	// Skip postpay success screen?
	SkipCheckoutSuccessPage bool `json:"skip_checkout_success_page"`
	// Billing Address is editable?
	BillingAddressEditable bool `json:"billing_address_editable"`
	// Billing Address
	BillingAddress CreateAddressRequest `json:"billing_address"`
	// Bank Transfer payment request
	BankTransfer *CreateCheckoutBankTransferRequest `json:"bank_transfer,omitempty"`
	// Accepted Brands
	AcceptedBrands []string `json:"accepted_brands"`
	// Pix payment request
	Pix *CreateCheckoutPixPaymentRequest `json:"pix,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCheckoutPaymentRequest.
// It customizes the JSON marshaling process for CreateCheckoutPaymentRequest objects.
func (c *CreateCheckoutPaymentRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCheckoutPaymentRequest object to a map representation for JSON marshaling.
func (c *CreateCheckoutPaymentRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["accepted_payment_methods"] = c.AcceptedPaymentMethods
	structMap["accepted_multi_payment_methods"] = c.AcceptedMultiPaymentMethods
	structMap["success_url"] = c.SuccessUrl
	if c.DefaultPaymentMethod != nil {
		structMap["default_payment_method"] = c.DefaultPaymentMethod
	}
	if c.GatewayAffiliationId != nil {
		structMap["gateway_affiliation_id"] = c.GatewayAffiliationId
	}
	if c.CreditCard != nil {
		structMap["credit_card"] = c.CreditCard
	}
	if c.DebitCard != nil {
		structMap["debit_card"] = c.DebitCard
	}
	if c.Boleto != nil {
		structMap["boleto"] = c.Boleto
	}
	if c.CustomerEditable != nil {
		structMap["customer_editable"] = c.CustomerEditable
	}
	if c.ExpiresIn != nil {
		structMap["expires_in"] = c.ExpiresIn
	}
	structMap["skip_checkout_success_page"] = c.SkipCheckoutSuccessPage
	structMap["billing_address_editable"] = c.BillingAddressEditable
	structMap["billing_address"] = c.BillingAddress
	if c.BankTransfer != nil {
		structMap["bank_transfer"] = c.BankTransfer
	}
	structMap["accepted_brands"] = c.AcceptedBrands
	if c.Pix != nil {
		structMap["pix"] = c.Pix
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCheckoutPaymentRequest.
// It customizes the JSON unmarshaling process for CreateCheckoutPaymentRequest objects.
func (c *CreateCheckoutPaymentRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		AcceptedPaymentMethods      []string                                `json:"accepted_payment_methods"`
		AcceptedMultiPaymentMethods []interface{}                           `json:"accepted_multi_payment_methods"`
		SuccessUrl                  string                                  `json:"success_url"`
		DefaultPaymentMethod        *string                                 `json:"default_payment_method,omitempty"`
		GatewayAffiliationId        *string                                 `json:"gateway_affiliation_id,omitempty"`
		CreditCard                  *CreateCheckoutCreditCardPaymentRequest `json:"credit_card,omitempty"`
		DebitCard                   *CreateCheckoutDebitCardPaymentRequest  `json:"debit_card,omitempty"`
		Boleto                      *CreateCheckoutBoletoPaymentRequest     `json:"boleto,omitempty"`
		CustomerEditable            *bool                                   `json:"customer_editable,omitempty"`
		ExpiresIn                   *int                                    `json:"expires_in,omitempty"`
		SkipCheckoutSuccessPage     bool                                    `json:"skip_checkout_success_page"`
		BillingAddressEditable      bool                                    `json:"billing_address_editable"`
		BillingAddress              CreateAddressRequest                    `json:"billing_address"`
		BankTransfer                *CreateCheckoutBankTransferRequest      `json:"bank_transfer,omitempty"`
		AcceptedBrands              []string                                `json:"accepted_brands"`
		Pix                         *CreateCheckoutPixPaymentRequest        `json:"pix,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.AcceptedPaymentMethods = temp.AcceptedPaymentMethods
	c.AcceptedMultiPaymentMethods = temp.AcceptedMultiPaymentMethods
	c.SuccessUrl = temp.SuccessUrl
	c.DefaultPaymentMethod = temp.DefaultPaymentMethod
	c.GatewayAffiliationId = temp.GatewayAffiliationId
	c.CreditCard = temp.CreditCard
	c.DebitCard = temp.DebitCard
	c.Boleto = temp.Boleto
	c.CustomerEditable = temp.CustomerEditable
	c.ExpiresIn = temp.ExpiresIn
	c.SkipCheckoutSuccessPage = temp.SkipCheckoutSuccessPage
	c.BillingAddressEditable = temp.BillingAddressEditable
	c.BillingAddress = temp.BillingAddress
	c.BankTransfer = temp.BankTransfer
	c.AcceptedBrands = temp.AcceptedBrands
	c.Pix = temp.Pix
	return nil
}
