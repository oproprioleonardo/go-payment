
# Get Card Response

Response object for getting a credit card

## Structure

`GetCardResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `LastFourDigits` | `Optional[string]` | Optional | - |
| `Brand` | `Optional[string]` | Optional | - |
| `HolderName` | `Optional[string]` | Optional | - |
| `ExpMonth` | `Optional[int]` | Optional | - |
| `ExpYear` | `Optional[int]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |
| `BillingAddress` | [`Optional[models.GetBillingAddressResponse]`](../../doc/models/get-billing-address-response.md) | Optional | - |
| `Customer` | [`Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `Metadata` | `Optional[map[string]string]` | Optional | - |
| `Type` | `Optional[string]` | Optional | Card type |
| `HolderDocument` | `Optional[string]` | Optional | Document number for the card's holder |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |
| `FirstSixDigits` | `Optional[string]` | Optional | First six digits |
| `Label` | `Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id4",
  "last_four_digits": "last_four_digits0",
  "brand": "brand8",
  "holder_name": "holder_name0",
  "exp_month": 52
}
```

