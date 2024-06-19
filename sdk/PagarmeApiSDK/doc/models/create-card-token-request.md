
# Create Card Token Request

Card token data

## Structure

`CreateCardTokenRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Number` | `string` | Required | Credit card number |
| `HolderName` | `string` | Required | Holder name, as written on the card |
| `ExpMonth` | `int` | Required | The expiration month |
| `ExpYear` | `int` | Required | The expiration year, that can be informed with 2 or 4 digits |
| `Cvv` | `string` | Required | The card's security code |
| `Brand` | `string` | Required | Card brand |
| `Label` | `string` | Required | - |

## Example (as JSON)

```json
{
  "number": "number8",
  "holder_name": "holder_name6",
  "exp_month": 168,
  "exp_year": 208,
  "cvv": "cvv8",
  "brand": "brand4",
  "label": "label0"
}
```

