
# Create Card Request

Card data

## Structure

`CreateCardRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Number` | `*string` | Optional | Credit card number |
| `HolderName` | `*string` | Optional | Holder name, as written on the card |
| `ExpMonth` | `*int` | Optional | The expiration month |
| `ExpYear` | `*int` | Optional | The expiration year, that can be informed with 2 or 4 digits |
| `Cvv` | `*string` | Optional | The card's security code |
| `BillingAddress` | [`*models.CreateAddressRequest`](../../doc/models/create-address-request.md) | Optional | Card's billing address |
| `Brand` | `*string` | Optional | Card brand |
| `BillingAddressId` | `*string` | Optional | The address id for the billing address |
| `Metadata` | `map[string]string` | Optional | Metadata |
| `Type` | `*string` | Optional | Card type<br>**Default**: `"credit"` |
| `Options` | [`*models.CreateCardOptionsRequest`](../../doc/models/create-card-options-request.md) | Optional | Options for creating the card |
| `HolderDocument` | `*string` | Optional | Document number for the card's holder |
| `PrivateLabel` | `*bool` | Optional | Indicates whether it is a private label card |
| `Label` | `*string` | Optional | - |
| `Id` | `*string` | Optional | Identifier |
| `Token` | `*string` | Optional | token identifier |

## Example (as JSON)

```json
{
  "type": "credit",
  "number": "number0",
  "holder_name": "holder_name8",
  "exp_month": 92,
  "exp_year": 204,
  "cvv": "cvv0"
}
```

