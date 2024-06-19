
# Get Charge Response

Response object for getting a charge

## Structure

`GetChargeResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Code` | `Optional[string]` | Optional | - |
| `GatewayId` | `Optional[string]` | Optional | - |
| `Amount` | `Optional[int]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `Currency` | `Optional[string]` | Optional | - |
| `PaymentMethod` | `Optional[string]` | Optional | - |
| `DueAt` | `Optional[time.Time]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |
| `LastTransaction` | [`Optional[models.GetTransactionResponseInterface]`](../../doc/models/get-transaction-response.md) | Optional | - |
| `Invoice` | [`Optional[models.GetInvoiceResponse]`](../../doc/models/get-invoice-response.md) | Optional | - |
| `Order` | [`Optional[models.GetOrderResponse]`](../../doc/models/get-order-response.md) | Optional | - |
| `Customer` | [`Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `Metadata` | `Optional[map[string]string]` | Optional | - |
| `PaidAt` | `Optional[time.Time]` | Optional | - |
| `CanceledAt` | `Optional[time.Time]` | Optional | - |
| `CanceledAmount` | `Optional[int]` | Optional | Canceled Amount |
| `PaidAmount` | `Optional[int]` | Optional | Paid amount |
| `InterestAndFinePaid` | `Optional[int]` | Optional | interest and fine paid |
| `RecurrencyCycle` | `Optional[string]` | Optional | Defines whether the card has been used one or more times. |

## Example (as JSON)

```json
{
  "recurrency_cycle": "\"first\" or \"subsequent\"",
  "id": "id0",
  "code": "code8",
  "gateway_id": "gateway_id0",
  "amount": 164,
  "status": "status2"
}
```

