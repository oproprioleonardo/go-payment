
# Get Invoice Response

Response object for getting an invoice

## Structure

`GetInvoiceResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Code` | `Optional[string]` | Optional | - |
| `Url` | `Optional[string]` | Optional | - |
| `Amount` | `Optional[int]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `PaymentMethod` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `Items` | [`Optional[[]models.GetInvoiceItemResponse]`](../../doc/models/get-invoice-item-response.md) | Optional | - |
| `Customer` | [`Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `Charge` | [`Optional[models.GetChargeResponse]`](../../doc/models/get-charge-response.md) | Optional | - |
| `Installments` | `Optional[int]` | Optional | - |
| `BillingAddress` | [`Optional[models.GetBillingAddressResponse]`](../../doc/models/get-billing-address-response.md) | Optional | - |
| `Subscription` | [`Optional[models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | - |
| `Cycle` | [`Optional[models.GetPeriodResponse]`](../../doc/models/get-period-response.md) | Optional | - |
| `Shipping` | [`Optional[models.GetShippingResponse]`](../../doc/models/get-shipping-response.md) | Optional | - |
| `Metadata` | `Optional[map[string]string]` | Optional | - |
| `DueAt` | `Optional[time.Time]` | Optional | - |
| `CanceledAt` | `Optional[time.Time]` | Optional | - |
| `BillingAt` | `Optional[time.Time]` | Optional | - |
| `SeenAt` | `Optional[time.Time]` | Optional | - |
| `TotalDiscount` | `Optional[int]` | Optional | Total discounted value |
| `TotalIncrement` | `Optional[int]` | Optional | Total discounted value |
| `SubscriptionId` | `Optional[string]` | Optional | Subscription Id |

## Example (as JSON)

```json
{
  "id": "id0",
  "code": "code8",
  "url": "url4",
  "amount": 168,
  "status": "status8"
}
```

