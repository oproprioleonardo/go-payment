
# Get Transaction Response

Generic response object for getting a transaction.

## Structure

`GetTransactionResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `GatewayId` | `Optional[string]` | Optional | Gateway transaction id |
| `Amount` | `Optional[int]` | Optional | Amount in cents |
| `Status` | `Optional[string]` | Optional | Transaction status |
| `Success` | `Optional[bool]` | Optional | Indicates if the transaction ocurred successfuly |
| `CreatedAt` | `Optional[time.Time]` | Optional | Creation date |
| `UpdatedAt` | `Optional[time.Time]` | Optional | Last update date |
| `AttemptCount` | `Optional[int]` | Optional | Number of attempts tried |
| `MaxAttempts` | `Optional[int]` | Optional | Max attempts |
| `Splits` | [`Optional[[]models.GetSplitResponse]`](../../doc/models/get-split-response.md) | Optional | Splits |
| `NextAttempt` | `Optional[time.Time]` | Optional | Date and time of the next attempt |
| `TransactionType` | `*string` | Optional | - |
| `Id` | `Optional[string]` | Optional | Código da transação |
| `GatewayResponse` | [`Optional[models.GetGatewayResponseResponse]`](../../doc/models/get-gateway-response-response.md) | Optional | The Gateway Response |
| `AntifraudResponse` | [`Optional[models.GetAntifraudResponse]`](../../doc/models/get-antifraud-response.md) | Optional | - |
| `Metadata` | `Optional[map[string]string]` | Optional | - |
| `Split` | [`Optional[[]models.GetSplitResponse]`](../../doc/models/get-split-response.md) | Optional | - |
| `Interest` | [`Optional[models.GetInterestResponse]`](../../doc/models/get-interest-response.md) | Optional | - |
| `Fine` | [`Optional[models.GetFineResponse]`](../../doc/models/get-fine-response.md) | Optional | - |
| `MaxDaysToPayPastDue` | `Optional[int]` | Optional | - |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id8",
  "amount": 40,
  "status": "status6",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "qr_code": "qr_code0",
  "qr_code_url": "qr_code_url6",
  "expires_at": "2016-03-13T12:52:32.123Z",
  "additional_information": [
    {
      "Name": "Name0",
      "Value": "Value2"
    },
    {
      "Name": "Name0",
      "Value": "Value2"
    }
  ],
  "end_to_end_id": "end_to_end_id6"
}
```

