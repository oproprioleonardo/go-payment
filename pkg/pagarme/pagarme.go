package pagarme

import (
	"context"
	"encoding/json"
	"github.com/ia/lab/payment/configs/pagarme"
	"github.com/ia/lab/payment/internal/entity"
	"golang.org/x/exp/slog"
	"pagarmeapisdk/models"
)

func ProcessOrder(ctx context.Context, o *entity.OrderRequest) (*entity.OrderResponse, error) {
	customerType := "individual"
	if o.Customer.Document.Type == "CNPJ" {
		customerType = "company"
	}
	phones := models.CreatePhonesRequest{}

	if o.Customer.Phone.Type == "landline" {
		phones.HomePhone = &models.CreatePhoneRequest{
			CountryCode: &o.Customer.Phone.CountryCode,
			AreaCode:    &o.Customer.Phone.AreaCode,
			Number:      &o.Customer.Phone.Number,
		}
	} else if o.Customer.Phone.Type == "mobile" {
		phones.MobilePhone = &models.CreatePhoneRequest{
			CountryCode: &o.Customer.Phone.CountryCode,
			AreaCode:    &o.Customer.Phone.AreaCode,
			Number:      &o.Customer.Phone.Number,
		}
	}

	address := models.CreateAddressRequest{
		City:    o.Customer.Address.City,
		Country: o.Customer.Address.Country,
		Line1:   o.Customer.Address.Line1,
		Line2:   o.Customer.Address.Line2,
		State:   o.Customer.Address.State,
		ZipCode: o.Customer.Address.ZipCode,
	}

	resp, err := pagarme.Client.OrdersController().CreateOrder(ctx, &models.CreateOrderRequest{
		Items: []models.CreateOrderItemRequest{
			{
				Amount:      o.Total,
				Description: "Saldo de Tokens",
				Quantity:    1,
			},
		},
		Customer: models.CreateCustomerRequest{
			Name:         o.Customer.Name,
			Email:        o.Customer.Email,
			Document:     o.Customer.Document.Value,
			DocumentType: &o.Customer.Document.Type,
			Type:         customerType,
			Address:      address,
			Phones:       phones,
		},
		Payments: []models.CreatePaymentRequest{
			createPaymentMethod(o, address),
		},
	}, nil)

	if err != nil {
		return nil, err
	}
	respJSON, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	slog.Info(string(respJSON))

	return o.Process()
}

func createPaymentMethod(o *entity.OrderRequest, address models.CreateAddressRequest) models.CreatePaymentRequest {
	amount := new(int)
	*amount = o.Total
	statementDescriptor := new(string)
	*statementDescriptor = "DAI - Tokens de IA"
	payment := models.CreatePaymentRequest{
		PaymentMethod: o.PaymentMethod,
		Amount:        amount,
	}
	if o.PaymentMethod == "pix" {
		expiresIn := new(int)
		*expiresIn = 259200
		payment.Pix = &models.CreatePixPaymentRequest{
			ExpiresIn: expiresIn,
		}
	} else if o.PaymentMethod == "credit_card" {
		cardHash := new(string)
		operationType := new(string)
		installments := new(int)
		*installments = 1
		*operationType = "auth_and_capture"
		*cardHash = o.CardHash
		payment.CreditCard = &models.CreateCreditCardPaymentRequest{
			CardToken:           cardHash,
			OperationType:       operationType,
			StatementDescriptor: statementDescriptor,
			Installments:        installments,
		}
	} else if o.PaymentMethod == "debit_card" {
		cardHash := new(string)
		*cardHash = o.CardHash
		payment.DebitCard = &models.CreateDebitCardPaymentRequest{
			CardToken:           cardHash,
			StatementDescriptor: statementDescriptor,
		}
	} else if o.PaymentMethod == "boleto" {
		bank := new(string)
		*bank = "001"
		payment.Boleto = &models.CreateBoletoPaymentRequest{
			Bank:                models.NewOptional[string](bank),
			StatementDescriptor: *statementDescriptor,
			DocumentNumber:      o.Customer.Document.Value,
			Instructions:        "Pagar até o vencimento",
			BillingAddress:      address,
		}
	}
	return payment
}
