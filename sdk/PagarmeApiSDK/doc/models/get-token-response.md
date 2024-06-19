
# Get Token Response

Token data

## Structure

`GetTokenResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Type` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `ExpiresAt` | `Optional[string]` | Optional | - |
| `Card` | [`Optional[models.GetCardTokenResponse]`](../../doc/models/get-card-token-response.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "id8",
  "type": "type2",
  "created_at": "2016-03-13T12:52:32.123Z",
  "expires_at": "expires_at2",
  "card": {
    "last_four_digits": "last_four_digits2",
    "holder_name": "holder_name2",
    "holder_document": "holder_document0",
    "exp_month": 228,
    "exp_year": 68
  }
}
```

