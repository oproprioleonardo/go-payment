
# Get Anticipation Response

Anticipation

## Structure

`GetAnticipationResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | Id |
| `RequestedAmount` | `Optional[int]` | Optional | Requested amount |
| `ApprovedAmount` | `Optional[int]` | Optional | Approved amount |
| `Recipient` | [`Optional[models.GetRecipientResponse]`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `Pgid` | `Optional[string]` | Optional | Anticipation id on the gateway |
| `CreatedAt` | `Optional[time.Time]` | Optional | Creation date |
| `UpdatedAt` | `Optional[time.Time]` | Optional | Last update date |
| `PaymentDate` | `Optional[time.Time]` | Optional | Payment date |
| `Status` | `Optional[string]` | Optional | Status |
| `Timeframe` | `Optional[string]` | Optional | Timeframe |

## Example (as JSON)

```json
{
  "id": "id8",
  "requested_amount": 130,
  "approved_amount": 184,
  "recipient": {
    "id": "id8",
    "name": "name8",
    "email": "email8",
    "document": "document8",
    "description": "description2"
  },
  "pgid": "pgid4"
}
```

