
# Get Invoice Item Response

Response object for getting an invoice item

## Structure

`GetInvoiceItemResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `Optional[int]` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `PricingScheme` | [`Optional[models.GetPricingSchemeResponse]`](../../doc/models/get-pricing-scheme-response.md) | Optional | - |
| `PriceBracket` | [`Optional[models.GetPriceBracketResponse]`](../../doc/models/get-price-bracket-response.md) | Optional | - |
| `Quantity` | `Optional[int]` | Optional | - |
| `Name` | `Optional[string]` | Optional | - |
| `SubscriptionItemId` | `Optional[string]` | Optional | Subscription Item Id |

## Example (as JSON)

```json
{
  "amount": 212,
  "description": "description6",
  "pricing_scheme": {
    "price": 166,
    "scheme_type": "scheme_type8",
    "price_brackets": [
      {
        "start_quantity": 144,
        "price": 174,
        "end_quantity": 152,
        "overage_price": 166
      },
      {
        "start_quantity": 144,
        "price": 174,
        "end_quantity": 152,
        "overage_price": 166
      },
      {
        "start_quantity": 144,
        "price": 174,
        "end_quantity": 152,
        "overage_price": 166
      }
    ],
    "minimum_price": 6,
    "percentage": 251.76
  },
  "price_bracket": {
    "start_quantity": 164,
    "price": 154,
    "end_quantity": 172,
    "overage_price": 186
  },
  "quantity": 70
}
```

