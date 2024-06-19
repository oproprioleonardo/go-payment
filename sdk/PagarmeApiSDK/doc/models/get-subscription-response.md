
# Get Subscription Response

## Structure

`GetSubscriptionResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Code` | `Optional[string]` | Optional | - |
| `StartAt` | `Optional[time.Time]` | Optional | - |
| `Interval` | `Optional[string]` | Optional | - |
| `IntervalCount` | `Optional[int]` | Optional | - |
| `BillingType` | `Optional[string]` | Optional | - |
| `CurrentCycle` | [`Optional[models.GetPeriodResponse]`](../../doc/models/get-period-response.md) | Optional | - |
| `PaymentMethod` | `Optional[string]` | Optional | - |
| `Currency` | `Optional[string]` | Optional | - |
| `Installments` | `Optional[int]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |
| `Customer` | [`Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `Card` | [`Optional[models.GetCardResponse]`](../../doc/models/get-card-response.md) | Optional | - |
| `Items` | [`Optional[[]models.GetSubscriptionItemResponse]`](../../doc/models/get-subscription-item-response.md) | Optional | - |
| `StatementDescriptor` | `Optional[string]` | Optional | - |
| `Metadata` | `Optional[map[string]string]` | Optional | - |
| `Setup` | [`Optional[models.GetSetupResponse]`](../../doc/models/get-setup-response.md) | Optional | - |
| `GatewayAffiliationId` | `Optional[string]` | Optional | Affiliation Code |
| `NextBillingAt` | `Optional[time.Time]` | Optional | - |
| `BillingDay` | `Optional[int]` | Optional | - |
| `MinimumPrice` | `Optional[int]` | Optional | - |
| `CanceledAt` | `Optional[time.Time]` | Optional | - |
| `Discounts` | [`Optional[[]models.GetDiscountResponse]`](../../doc/models/get-discount-response.md) | Optional | Subscription discounts |
| `Increments` | [`Optional[[]models.GetIncrementResponse]`](../../doc/models/get-increment-response.md) | Optional | Subscription increments |
| `BoletoDueDays` | `Optional[int]` | Optional | Days until boleto expires |
| `Split` | [`Optional[models.GetSubscriptionSplitResponse]`](../../doc/models/get-subscription-split-response.md) | Optional | Subscription's split response |
| `Boleto` | [`Optional[models.GetSubscriptionBoletoResponse]`](../../doc/models/get-subscription-boleto-response.md) | Optional | - |
| `ManualBilling` | `Optional[bool]` | Optional | - |

## Example (as JSON)

```json
{
  "boleto": {
    "interest": {
      "days": 2,
      "type": "percentage",
      "amount": 20
    },
    "fine": {
      "days": 2,
      "type": "flat",
      "amount": 10
    },
    "max_days_to_pay_past_due": 2
  },
  "id": "id4",
  "code": "code2",
  "start_at": "2016-03-13T12:52:32.123Z",
  "interval": "interval2",
  "interval_count": 224
}
```

