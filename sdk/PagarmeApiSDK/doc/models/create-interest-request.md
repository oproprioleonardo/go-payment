
# Create Interest Request

Interest Request

## Structure

`CreateInterestRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Days` | `int` | Required | Days |
| `Type` | `string` | Required | Type |
| `Amount` | `int` | Required | Amount |

## Example (as JSON)

```json
{
  "days": 4,
  "type": "\"percentage\" or \"flat\"",
  "amount": 78
}
```

