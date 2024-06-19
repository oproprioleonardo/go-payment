
# Get Recipient Response

Recipient response

## Structure

`GetRecipientResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | Id |
| `Name` | `Optional[string]` | Optional | Name |
| `Email` | `Optional[string]` | Optional | Email |
| `Document` | `Optional[string]` | Optional | Document |
| `Description` | `Optional[string]` | Optional | Description |
| `Type` | `Optional[string]` | Optional | Type |
| `Status` | `Optional[string]` | Optional | Status |
| `CreatedAt` | `Optional[time.Time]` | Optional | Creation date |
| `UpdatedAt` | `Optional[time.Time]` | Optional | Last update date |
| `DeletedAt` | `Optional[time.Time]` | Optional | Deletion date |
| `DefaultBankAccount` | [`Optional[models.GetBankAccountResponse]`](../../doc/models/get-bank-account-response.md) | Optional | Default bank account |
| `GatewayRecipients` | [`Optional[[]models.GetGatewayRecipientResponse]`](../../doc/models/get-gateway-recipient-response.md) | Optional | Info about the recipient on the gateway |
| `Metadata` | `Optional[map[string]string]` | Optional | Metadata |
| `AutomaticAnticipationSettings` | [`Optional[models.GetAutomaticAnticipationResponse]`](../../doc/models/get-automatic-anticipation-response.md) | Optional | - |
| `TransferSettings` | [`Optional[models.GetTransferSettingsResponse]`](../../doc/models/get-transfer-settings-response.md) | Optional | - |
| `Code` | `Optional[string]` | Optional | Recipient code |
| `PaymentMode` | `Optional[string]` | Optional | Payment mode<br>**Default**: `"bank_transfer"` |

## Example (as JSON)

```json
{
  "payment_mode": "bank_transfer",
  "id": "id4",
  "name": "name4",
  "email": "email2",
  "document": "document2",
  "description": "description6"
}
```

