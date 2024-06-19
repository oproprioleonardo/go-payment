
# Get Period Response

Response object for getting a period

## Structure

`GetPeriodResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartAt` | `Optional[time.Time]` | Optional | - |
| `EndAt` | `Optional[time.Time]` | Optional | - |
| `Id` | `Optional[string]` | Optional | - |
| `BillingAt` | `Optional[time.Time]` | Optional | - |
| `Subscription` | [`Optional[models.GetSubscriptionResponse]`](../../doc/models/get-subscription-response.md) | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `Duration` | `Optional[int]` | Optional | - |
| `CreatedAt` | `Optional[string]` | Optional | - |
| `UpdatedAt` | `Optional[string]` | Optional | - |
| `Cycle` | `Optional[int]` | Optional | - |

## Example (as JSON)

```json
{
  "start_at": "2016-03-13T12:52:32.123Z",
  "end_at": "2016-03-13T12:52:32.123Z",
  "id": "id0",
  "billing_at": "2016-03-13T12:52:32.123Z",
  "subscription": {
    "id": "id4",
    "code": "code2",
    "start_at": "2016-03-13T12:52:32.123Z",
    "interval": "interval2",
    "interval_count": 234
  }
}
```

