
# Create Subscription Item Request

Request for creating a new subscription item

## Structure

`CreateSubscriptionItemRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `string` | Required | Item description |
| `PricingScheme` | [`models.CreatePricingSchemeRequest`](../../doc/models/create-pricing-scheme-request.md) | Required | Pricing scheme |
| `Id` | `string` | Required | Item id |
| `PlanItemId` | `string` | Required | Plan item id |
| `Discounts` | [`[]models.CreateDiscountRequest`](../../doc/models/create-discount-request.md) | Required | Discounts for the item |
| `Name` | `string` | Required | Item name |
| `Cycles` | `*int` | Optional | Number of cycles which the item will be charged |
| `Quantity` | `*int` | Optional | Quantity of items |
| `MinimumPrice` | `*int` | Optional | Minimum price |

## Example (as JSON)

```json
{
  "description": "description0",
  "pricing_scheme": {
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
    "price": 166,
    "minimum_price": 6,
    "percentage": 251.76
  },
  "id": "id0",
  "plan_item_id": "plan_item_id0",
  "discounts": [
    {
      "value": 90.66,
      "discount_type": "discount_type2",
      "item_id": "item_id4",
      "cycles": 126,
      "description": "description4"
    }
  ],
  "name": "name0",
  "cycles": 106,
  "quantity": 130,
  "minimum_price": 114
}
```

