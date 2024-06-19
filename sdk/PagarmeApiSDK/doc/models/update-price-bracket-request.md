
# Update Price Bracket Request

Request for updating a price bracket

## Structure

`UpdatePriceBracketRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StartQuantity` | `int` | Required | Start quantity of the bracket |
| `Price` | `int` | Required | Price |
| `EndQuantity` | `*int` | Optional | End quantity of the bracket |
| `OveragePrice` | `*int` | Optional | Overage price |

## Example (as JSON)

```json
{
  "start_quantity": 154,
  "price": 164,
  "end_quantity": 162,
  "overage_price": 176
}
```

