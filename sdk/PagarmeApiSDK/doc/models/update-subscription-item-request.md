
# Update Subscription Item Request

Request for updating a subscription item

## Structure

`UpdateSubscriptionItemRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `string` | Required | Description |
| `Status` | `string` | Required | Status |
| `PricingScheme` | [`models.UpdatePricingSchemeRequest`](../../doc/models/update-pricing-scheme-request.md) | Required | Pricing scheme |
| `Name` | `string` | Required | Item name |
| `Cycles` | `*int` | Optional | Number of cycles that the item will be charged |
| `Quantity` | `*int` | Optional | Quantity |
| `MinimumPrice` | `*int` | Optional | Minimum price |

## Example (as JSON)

```json
{
  "description": "description2",
  "status": "status4",
  "pricing_scheme": {
    "scheme_type": "scheme_type8",
    "price_brackets": [
      {
        "start_quantity": 144,
        "price": 174,
        "end_quantity": 152,
        "overage_price": 166
      }
    ],
    "price": 166,
    "minimum_price": 6,
    "percentage": 251.76
  },
  "name": "name2",
  "cycles": 108,
  "quantity": 128,
  "minimum_price": 140
}
```

