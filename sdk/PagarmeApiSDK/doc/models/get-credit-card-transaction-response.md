
# Get Credit Card Transaction Response

Response object for getting a credit card transaction

## Structure

`GetCreditCardTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StatementDescriptor` | `Optional[string]` | Optional | Text that will appear on the credit card's statement |
| `AcquirerName` | `*string` | Optional | Acquirer name |
| `AcquirerAffiliationCode` | `Optional[string]` | Optional | Aquirer affiliation code |
| `AcquirerTid` | `*string` | Optional | Acquirer TID |
| `AcquirerNsu` | `*string` | Optional | Acquirer NSU |
| `AcquirerAuthCode` | `Optional[string]` | Optional | Acquirer authorization code |
| `OperationType` | `Optional[string]` | Optional | Operation type |
| `Card` | [`Optional[models.GetCardResponse]`](../../doc/models/get-card-response.md) | Optional | Card data |
| `AcquirerMessage` | `Optional[string]` | Optional | Acquirer message |
| `AcquirerReturnCode` | `Optional[string]` | Optional | Acquirer Return Code |
| `Installments` | `Optional[int]` | Optional | Number of installments |
| `ThreedAuthenticationUrl` | `Optional[string]` | Optional | 3D-S authentication Url |
| `FundingSource` | `Optional[string]` | Optional | Identify when a card is prepaid, credit or debit. |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id8",
  "amount": 40,
  "status": "status6",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "statement_descriptor": "statement_descriptor2",
  "acquirer_name": "acquirer_name6",
  "acquirer_affiliation_code": "acquirer_affiliation_code6",
  "acquirer_tid": "acquirer_tid8",
  "acquirer_nsu": "acquirer_nsu8"
}
```

