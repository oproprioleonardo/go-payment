
# Get Checkout Payment Settings Response

Checkout Payment Settings Response

## Structure

`GetCheckoutPaymentSettingsResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SuccessUrl` | `Optional[string]` | Optional | Success Url |
| `PaymentUrl` | `Optional[string]` | Optional | Payment Url |
| `AcceptedPaymentMethods` | `Optional[[]string]` | Optional | Accepted Payment Methods |
| `Status` | `Optional[string]` | Optional | Status |
| `Customer` | [`Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | Customer |
| `Amount` | `Optional[int]` | Optional | Payment amount |
| `DefaultPaymentMethod` | `Optional[string]` | Optional | Default Payment Method |
| `GatewayAffiliationId` | `Optional[string]` | Optional | Gateway Affiliation Id |

## Example (as JSON)

```json
{
  "success_url": "success_url0",
  "payment_url": "payment_url8",
  "accepted_payment_methods": [
    "accepted_payment_methods1",
    "accepted_payment_methods2"
  ],
  "status": "status0",
  "customer": {
    "id": "id0",
    "name": "name0",
    "email": "email6",
    "delinquent": false,
    "created_at": "2016-03-13T12:52:32.123Z"
  }
}
```

