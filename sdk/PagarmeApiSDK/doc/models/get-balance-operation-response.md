
# Get Balance Operation Response

Generic response object for getting a BalanceOperation.

## Structure

`GetBalanceOperationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `BalanceAmount` | `Optional[string]` | Optional | - |
| `BalanceOldAmount` | `Optional[string]` | Optional | - |
| `Type` | `Optional[string]` | Optional | - |
| `Amount` | `Optional[string]` | Optional | - |
| `Fee` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[string]` | Optional | - |
| `MovementObject` | [`models.GetMovementObjectBaseResponseInterface`](../../doc/models/get-movement-object-base-response.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "id0",
  "status": "status2",
  "balance_amount": "balance_amount0",
  "balance_old_amount": "balance_old_amount8",
  "type": "type0"
}
```

