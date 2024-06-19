
# Get Order Response

Response object for getting an Order

## Structure

`GetOrderResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Code` | `Optional[string]` | Optional | - |
| `Amount` | `Optional[int]` | Optional | - |
| `Currency` | `Optional[string]` | Optional | - |
| `Closed` | `Optional[bool]` | Optional | Indicates whether the order is closed |
| `Items` | [`Optional[[]models.GetOrderItemResponse]`](../../doc/models/get-order-item-response.md) | Optional | - |
| `Customer` | [`Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |
| `ClosedAt` | `Optional[time.Time]` | Optional | - |
| `Charges` | [`Optional[[]models.GetChargeResponse]`](../../doc/models/get-charge-response.md) | Optional | - |
| `InvoiceUrl` | `Optional[string]` | Optional | - |
| `Shipping` | [`Optional[models.GetShippingResponse]`](../../doc/models/get-shipping-response.md) | Optional | - |
| `Metadata` | `Optional[map[string]string]` | Optional | - |
| `Checkouts` | [`Optional[[]models.GetCheckoutPaymentResponse]`](../../doc/models/get-checkout-payment-response.md) | Optional | Checkout Payment Settings Response |
| `Ip` | `Optional[string]` | Optional | Ip address |
| `SessionId` | `Optional[string]` | Optional | Session id |
| `Location` | [`Optional[models.GetLocationResponse]`](../../doc/models/get-location-response.md) | Optional | Location |
| `Device` | [`Optional[models.GetDeviceResponse]`](../../doc/models/get-device-response.md) | Optional | Device's informations |

## Example (as JSON)

```json
{
  "id": "id6",
  "code": "code4",
  "amount": 64,
  "currency": "currency6",
  "closed": false
}
```

