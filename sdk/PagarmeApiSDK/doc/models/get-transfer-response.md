
# Get Transfer Response

Transfer response

## Structure

`GetTransferResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | Id |
| `Amount` | `Optional[int]` | Optional | Transfer amount |
| `Status` | `Optional[string]` | Optional | Transfer status |
| `CreatedAt` | `Optional[time.Time]` | Optional | Transfer creation date |
| `UpdatedAt` | `Optional[time.Time]` | Optional | Transfer last update date |
| `BankAccount` | [`Optional[models.GetBankAccountResponse]`](../../doc/models/get-bank-account-response.md) | Optional | Bank account |
| `Metadata` | `Optional[map[string]string]` | Optional | Metadata |

## Example (as JSON)

```json
{
  "id": "id2",
  "amount": 146,
  "status": "status4",
  "created_at": "2016-03-13T12:52:32.123Z",
  "updated_at": "2016-03-13T12:52:32.123Z"
}
```

