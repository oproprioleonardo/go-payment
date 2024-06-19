
# Get Discount Response

Response object for getting a discount

## Structure

`GetDiscountResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Value` | `Optional[float64]` | Optional | - |
| `DiscountType` | `Optional[string]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `Cycles` | `Optional[int]` | Optional | - |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `Subscription` | [`Optional[models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | - |
| `SubscriptionItem` | [`Optional[models.GetSubscriptionItemResponse]`](../../doc/models/get-subscription-item-response.md) | Optional | The subscription item |

## Example (as JSON)

```json
{
  "id": "id4",
  "value": 139.66,
  "discount_type": "discount_type2",
  "status": "status6",
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

