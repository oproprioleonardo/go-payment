
# Get Plan Response

Response object for getting a plan

## Structure

`GetPlanResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Name` | `Optional[string]` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `Url` | `Optional[string]` | Optional | - |
| `StatementDescriptor` | `Optional[string]` | Optional | - |
| `Interval` | `Optional[string]` | Optional | - |
| `IntervalCount` | `Optional[int]` | Optional | - |
| `BillingType` | `Optional[string]` | Optional | - |
| `PaymentMethods` | `Optional[[]string]` | Optional | - |
| `Installments` | `Optional[[]int]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `Currency` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |
| `Items` | [`Optional[[]models.GetPlanItemResponse]`](../../doc/models/get-plan-item-response.md) | Optional | - |
| `BillingDays` | `Optional[[]int]` | Optional | - |
| `Shippable` | `Optional[bool]` | Optional | - |
| `Metadata` | `Optional[map[string]string]` | Optional | - |
| `TrialPeriodDays` | `Optional[int]` | Optional | - |
| `MinimumPrice` | `Optional[int]` | Optional | - |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "name": "name0",
  "description": "description0",
  "url": "url4",
  "statement_descriptor": "statement_descriptor0"
}
```

