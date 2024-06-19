
# Get Boleto Transaction Response

Response object for getting a boleto transaction

## Structure

`GetBoletoTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Url` | `Optional[string]` | Optional | - |
| `Barcode` | `Optional[string]` | Optional | - |
| `NossoNumero` | `Optional[string]` | Optional | - |
| `Bank` | `Optional[string]` | Optional | - |
| `DocumentNumber` | `Optional[string]` | Optional | - |
| `Instructions` | `Optional[string]` | Optional | - |
| `BillingAddress` | [`Optional[models.GetBillingAddressResponse]`](../../doc/models/get-billing-address-response.md) | Optional | - |
| `DueAt` | `Optional[time.Time]` | Optional | - |
| `QrCode` | `Optional[string]` | Optional | - |
| `Line` | `Optional[string]` | Optional | - |
| `PdfPassword` | `Optional[string]` | Optional | - |
| `Pdf` | `Optional[string]` | Optional | - |
| `PaidAt` | `Optional[time.Time]` | Optional | - |
| `PaidAmount` | `Optional[string]` | Optional | - |
| `Type` | `Optional[string]` | Optional | - |
| `CreditAt` | `Optional[time.Time]` | Optional | - |
| `StatementDescriptor` | `Optional[string]` | Optional | Soft Descriptor |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id8",
  "amount": 40,
  "status": "status6",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "url": "url0",
  "barcode": "barcode4",
  "nosso_numero": "nosso_numero6",
  "bank": "bank4",
  "document_number": "document_number0"
}
```

