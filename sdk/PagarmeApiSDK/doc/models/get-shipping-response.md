
# Get Shipping Response

Response object for getting the shipping data

## Structure

`GetShippingResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `Optional[int]` | Optional | - |
| `Description` | `Optional[string]` | Optional | - |
| `RecipientName` | `Optional[string]` | Optional | - |
| `RecipientPhone` | `Optional[string]` | Optional | - |
| `Address` | [`Optional[models.GetAddressResponse]`](../../doc/models/get-address-response.md) | Optional | - |
| `MaxDeliveryDate` | `Optional[time.Time]` | Optional | Data máxima de entrega |
| `EstimatedDeliveryDate` | `Optional[time.Time]` | Optional | Prazo estimado de entrega |
| `Type` | `Optional[string]` | Optional | Shipping Type |

## Example (as JSON)

```json
{
  "amount": 214,
  "description": "description8",
  "recipient_name": "recipient_name6",
  "recipient_phone": "recipient_phone0",
  "address": {
    "id": "id6",
    "street": "street6",
    "number": "number4",
    "complement": "complement2",
    "zip_code": "zip_code0"
  }
}
```

