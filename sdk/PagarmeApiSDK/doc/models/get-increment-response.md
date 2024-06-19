
# Get Increment Response

Response object for getting a increment

## Structure

`GetIncrementResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Value` | `Optional[float64]` | Optional | - |
| `IncrementType` | `Optional[string]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `Cycles` | `Optional[int]` | Optional | - |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `Subscription` | [`Optional[models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | - |
| `SubscriptionItem` | [`Optional[models.GetSubscriptionItemResponse]`](../../doc/models/get-subscription-item-response.md) | Optional | The Subscription Item |

## Example (as JSON)

```json
{
  "id": "id0",
  "value": 167.72,
  "increment_type": "increment_type2",
  "status": "status2",
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

