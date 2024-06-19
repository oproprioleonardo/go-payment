
# Get Withdraw Response

## Structure

`GetWithdrawResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `GatewayId` | `Optional[string]` | Optional | - |
| `Amount` | `Optional[int]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |
| `Metadata` | `Optional[[]string]` | Optional | - |
| `Fee` | `Optional[int]` | Optional | - |
| `FundingDate` | `Optional[time.Time]` | Optional | - |
| `FundingEstimatedDate` | `Optional[time.Time]` | Optional | - |
| `Type` | `Optional[string]` | Optional | - |
| `Source` | [`Optional[models.GetWithdrawSourceResponse]`](../../doc/models/get-withdraw-source-response.md) | Optional | - |
| `Target` | [`Optional[models.GetWithdrawTargetResponse]`](../../doc/models/get-withdraw-target-response.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "id6",
  "gateway_id": "gateway_id4",
  "amount": 78,
  "status": "status8",
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

