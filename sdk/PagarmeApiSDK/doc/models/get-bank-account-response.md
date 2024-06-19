
# Get Bank Account Response

## Structure

`GetBankAccountResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | Id |
| `HolderName` | `Optional[string]` | Optional | Holder name |
| `HolderType` | `Optional[string]` | Optional | Holder type |
| `Bank` | `Optional[string]` | Optional | Bank |
| `BranchNumber` | `Optional[string]` | Optional | Branch number |
| `BranchCheckDigit` | `Optional[string]` | Optional | Branch check digit |
| `AccountNumber` | `Optional[string]` | Optional | Account number |
| `AccountCheckDigit` | `Optional[string]` | Optional | Account check digit |
| `Type` | `Optional[string]` | Optional | Bank account type |
| `Status` | `Optional[string]` | Optional | Bank account status |
| `CreatedAt` | `Optional[time.Time]` | Optional | Creation date |
| `UpdatedAt` | `Optional[time.Time]` | Optional | Last update date |
| `DeletedAt` | `Optional[time.Time]` | Optional | Deletion date |
| `Recipient` | [`Optional[models.GetRecipientResponse]`](../../doc/models/get-recipient-response.md) | Optional | Recipient |
| `Metadata` | `Optional[map[string]string]` | Optional | Metadata |
| `PixKey` | `Optional[string]` | Optional | Pix Key |

## Example (as JSON)

```json
{
  "id": "id6",
  "holder_name": "holder_name2",
  "holder_type": "holder_type8",
  "bank": "bank4",
  "branch_number": "branch_number2"
}
```

