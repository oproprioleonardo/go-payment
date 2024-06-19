
# Create Increment Request

Request for creating a new increment

## Structure

`CreateIncrementRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Value` | `float64` | Required | The increment value |
| `IncrementType` | `string` | Required | Increment type. Can be either flat or percentage. |
| `ItemId` | `string` | Required | The item where the increment will be applied |
| `Cycles` | `*int` | Optional | Number of cycles that the increment will be applied |
| `Description` | `*string` | Optional | Description |

## Example (as JSON)

```json
{
  "value": 72.04,
  "increment_type": "increment_type4",
  "item_id": "item_id8",
  "cycles": 196,
  "description": "description8"
}
```

