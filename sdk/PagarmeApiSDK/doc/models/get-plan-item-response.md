
# Get Plan Item Response

Response object for getting a plan item

## Structure

`GetPlanItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Name` | `Optional[string]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |
| `PricingScheme` | [`Optional[models.GetPricingSchemeResponse]`](../../doc/models/get-pricing-scheme-response.md) | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `Plan` | [`Optional[models.GetPlanResponse]`](../../doc/models/get-plan-response.md) | Optional | - |
| `Quantity` | `Optional[int]` | Optional | - |
| `Cycles` | `Optional[int]` | Optional | - |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id8",
  "name": "name8",
  "status": "status0",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z"
}
```

