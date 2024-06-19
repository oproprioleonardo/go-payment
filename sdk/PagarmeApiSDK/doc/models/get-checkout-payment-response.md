
# Get Checkout Payment Response

Resposta das configurações de pagamento do checkout

## Structure

`GetCheckoutPaymentResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Amount` | `Optional[int]` | Optional | Valor em centavos |
| `DefaultPaymentMethod` | `Optional[string]` | Optional | Meio de pagamento padrão no checkout |
| `SuccessUrl` | `Optional[string]` | Optional | Url de redirecionamento de sucesso após o checkou |
| `PaymentUrl` | `Optional[string]` | Optional | Url para pagamento usando o checkout |
| `GatewayAffiliationId` | `Optional[string]` | Optional | Código da afiliação onde o pagamento será processado no gateway |
| `AcceptedPaymentMethods` | `Optional[[]string]` | Optional | Meios de pagamento aceitos no checkout |
| `Status` | `Optional[string]` | Optional | Status do checkout |
| `SkipCheckoutSuccessPage` | `Optional[bool]` | Optional | Pular tela de sucesso pós-pagamento? |
| `CreatedAt` | `Optional[time.Time]` | Optional | Data de criação |
| `UpdatedAt` | `Optional[time.Time]` | Optional | Data de atualização |
| `CanceledAt` | `Optional[time.Time]` | Optional | Data de cancelamento |
| `CustomerEditable` | `Optional[bool]` | Optional | Torna o objeto customer editável |
| `Customer` | [`Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | Dados do comprador |
| `Billingaddress` | [`Optional[models.GetAddressResponse]`](../../doc/models/get-address-response.md) | Optional | Dados do endereço de cobrança |
| `CreditCard` | [`Optional[models.GetCheckoutCreditCardPaymentResponse]`](../../doc/models/get-checkout-credit-card-payment-response.md) | Optional | Configurações de cartão de crédito |
| `Boleto` | [`Optional[models.GetCheckoutBoletoPaymentResponse]`](../../doc/models/get-checkout-boleto-payment-response.md) | Optional | Configurações de boleto |
| `BillingAddressEditable` | `Optional[bool]` | Optional | Indica se o billing address poderá ser editado |
| `Shipping` | [`Optional[models.GetShippingResponse]`](../../doc/models/get-shipping-response.md) | Optional | Configurações  de entrega |
| `Shippable` | `Optional[bool]` | Optional | Indica se possui entrega |
| `ClosedAt` | `Optional[time.Time]` | Optional | Data de fechamento |
| `ExpiresAt` | `Optional[time.Time]` | Optional | Data de expiração |
| `Currency` | `Optional[string]` | Optional | Moeda |
| `DebitCard` | [`Optional[models.GetCheckoutDebitCardPaymentResponse]`](../../doc/models/get-checkout-debit-card-payment-response.md) | Optional | Configurações de cartão de débito |
| `BankTransfer` | [`Optional[models.GetCheckoutBankTransferPaymentResponse]`](../../doc/models/get-checkout-bank-transfer-payment-response.md) | Optional | Bank transfer payment response |
| `AcceptedBrands` | `Optional[[]string]` | Optional | Accepted Brands |
| `Pix` | [`Optional[models.GetCheckoutPixPaymentResponse]`](../../doc/models/get-checkout-pix-payment-response.md) | Optional | Pix payment response |

## Example (as JSON)

```json
{
  "id": "id6",
  "amount": 148,
  "default_payment_method": "default_payment_method6",
  "success_url": "success_url8",
  "payment_url": "payment_url0"
}
```

