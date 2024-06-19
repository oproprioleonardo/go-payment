
# Get Split Response

Split response

## Structure

`GetSplitResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | `Optional[string]` | Optional | Type |
| `Amount` | `Optional[int]` | Optional | Amount |
| `Recipient` | [`Optional[models.GetRecipientResponse]`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `GatewayId` | `Optional[string]` | Optional | The split rule gateway id |
| `Options` | [`Optional[models.GetSplitOptionsResponse]`](../../doc/models/get-split-options-response.md) | Optional | - |
| `Id` | `Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "type": "type0",
  "amount": 252,
  "recipient": {
    "id": "id8",
    "name": "name8",
    "email": "email8",
    "document": "document8",
    "description": "description2"
  },
  "gateway_id": "gateway_id0",
  "options": {
    "liable": false,
    "charge_processing_fee": false,
    "charge_remainder_fee": "charge_remainder_fee0"
  }
}
```

