
# Get Subscription Item Response

## Structure

`GetSubscriptionItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |
| `PricingScheme` | [`Optional[models.GetPricingSchemeResponse]`](../../doc/models/get-pricing-scheme-response.md) | Optional | - |
| `Discounts` | [`Optional[[]models.GetDiscountResponse]`](../../doc/models/get-discount-response.md) | Optional | - |
| `Increments` | [`Optional[[]models.GetIncrementResponse]`](../../doc/models/get-increment-response.md) | Optional | - |
| `Subscription` | [`Optional[models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | - |
| `Name` | `Optional[string]` | Optional | Item name |
| `Quantity` | `Optional[int]` | Optional | - |
| `Cycles` | `Optional[int]` | Optional | - |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id2",
  "description": "description8",
  "status": "status6",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z"
}
```

