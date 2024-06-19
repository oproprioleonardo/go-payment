
# Update Pricing Scheme Request

Request for updating a pricing scheme

## Structure

`UpdatePricingSchemeRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SchemeType` | `string` | Required | Scheme type |
| `PriceBrackets` | [`[]models.UpdatePriceBracketRequest`](../../doc/models/update-price-bracket-request.md) | Required | Price brackets |
| `Price` | `*int` | Optional | Price |
| `MinimumPrice` | `*int` | Optional | Minimum price |
| `Percentage` | `*float64` | Optional | percentual value used in pricing_scheme Percent |

## Example (as JSON)

```json
{
  "scheme_type": "scheme_type0",
  "price_brackets": [
    {
      "start_quantity": 144,
      "price": 174,
      "end_quantity": 152,
      "overage_price": 166
    }
  ],
  "price": 162,
  "minimum_price": 2,
  "percentage": 62.28
}
```

