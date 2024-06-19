
# Get Usage Response

Response object for getting a usage

## Structure

`GetUsageResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | Id |
| `Quantity` | `Optional[int]` | Optional | Quantity |
| `Description` | `Optional[string]` | Optional | Description |
| `UsedAt` | `Optional[time.Time]` | Optional | Used at |
| `CreatedAt` | `Optional[time.Time]` | Optional | Creation date |
| `Status` | `Optional[string]` | Optional | Status |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |
| `SubscriptionItem` | [`Optional[models.GetSubscriptionItemResponse]`](../../doc/models/get-subscription-item-response.md) | Optional | Subscription item |
| `Code` | `Optional[string]` | Optional | Identification code in the client system |
| `Group` | `Optional[string]` | Optional | Identification group in the client system |
| `Amount` | `Optional[int]` | Optional | Field used in item scheme type 'Percent' |

## Example (as JSON)

```json
{
  "id": "id2",
  "quantity": 34,
  "description": "description2",
  "used_at": "2016-03-13T12:52:32.123Z",
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

