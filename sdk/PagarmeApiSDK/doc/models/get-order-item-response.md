
# Get Order Item Response

Response object for getting an order item

## Structure

`GetOrderItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | Id |
| `Type` | `Optional[string]` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `Amount` | `Optional[int]` | Optional | - |
| `Quantity` | `Optional[int]` | Optional | - |
| `Category` | `Optional[string]` | Optional | Category |
| `Code` | `Optional[string]` | Optional | Code |
| `Status` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id8",
  "type": "type8",
  "description": "description8",
  "amount": 224,
  "quantity": 82
}
```

