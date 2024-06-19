
# Get Balance Response

Balance

## Structure

`GetBalanceResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Currency` | `Optional[string]` | Optional | Currency |
| `AvailableAmount` | `Optional[int64]` | Optional | Amount available for transferring |
| `Recipient` | [`Optional[models.GetRecipientResponse]`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `TransferredAmount` | `Optional[int64]` | Optional | - |
| `WaitingFundsAmount` | `Optional[int64]` | Optional | - |

## Example (as JSON)

```json
{
  "currency": "currency2",
  "available_amount": 96,
  "recipient": {
    "id": "id8",
    "name": "name8",
    "email": "email8",
    "document": "document8",
    "description": "description2"
  },
  "transferred_amount": 142,
  "waiting_funds_amount": 174
}
```

