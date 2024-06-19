
# Get Pricing Scheme Response

Response object for getting a pricing scheme

## Structure

`GetPricingSchemeResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Price` | `Optional[int]` | Optional | - |
| `SchemeType` | `Optional[string]` | Optional | - |
| `PriceBrackets` | [`Optional[[]models.GetPriceBracketResponse]`](../../doc/models/get-price-bracket-response.md) | Optional | - |
| `MinimumPrice` | `Optional[int]` | Optional | - |
| `Percentage` | `Optional[float64]` | Optional | percentual value used in pricing_scheme Percent |

## Example (as JSON)

```json
{
  "price": 182,
  "scheme_type": "scheme_type8",
  "price_brackets": [
    {
      "start_quantity": 144,
      "price": 174,
      "end_quantity": 152,
      "overage_price": 166
    }
  ],
  "minimum_price": 170,
  "percentage": 166.36
}
```

