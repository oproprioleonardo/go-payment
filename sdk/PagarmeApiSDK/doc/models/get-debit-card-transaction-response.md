
# Get Debit Card Transaction Response

Response object for getting a debit card transaction

## Structure

`GetDebitCardTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StatementDescriptor` | `Optional[string]` | Optional | Text that will appear on the debit card's statement |
| `AcquirerName` | `Optional[string]` | Optional | Acquirer name |
| `AcquirerAffiliationCode` | `Optional[string]` | Optional | Aquirer affiliation code |
| `AcquirerTid` | `Optional[string]` | Optional | Acquirer TID |
| `AcquirerNsu` | `Optional[string]` | Optional | Acquirer NSU |
| `AcquirerAuthCode` | `Optional[string]` | Optional | Acquirer authorization code |
| `OperationType` | `Optional[string]` | Optional | Operation type |
| `Card` | [`Optional[models.GetCardResponse]`](../../doc/models/get-card-response.md) | Optional | Card data |
| `AcquirerMessage` | `Optional[string]` | Optional | Acquirer message |
| `AcquirerReturnCode` | `Optional[string]` | Optional | Acquirer Return Code |
| `Mpi` | `Optional[string]` | Optional | Merchant Plugin |
| `Eci` | `Optional[string]` | Optional | Electronic Commerce Indicator (ECI) |
| `AuthenticationType` | `Optional[string]` | Optional | Authentication type |
| `ThreedAuthenticationUrl` | `Optional[string]` | Optional | 3D-S Authentication Url |
| `FundingSource` | `Optional[string]` | Optional | Identify when a card is prepaid, credit or debit. |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id8",
  "amount": 40,
  "status": "status6",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "statement_descriptor": "statement_descriptor0",
  "acquirer_name": "acquirer_name4",
  "acquirer_affiliation_code": "acquirer_affiliation_code8",
  "acquirer_tid": "acquirer_tid0",
  "acquirer_nsu": "acquirer_nsu0"
}
```

