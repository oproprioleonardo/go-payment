
# Get Voucher Transaction Response

Response for voucher transactions

## Structure

`GetVoucherTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `StatementDescriptor` | `Optional[string]` | Optional | Text that will appear on the voucher's statement |
| `AcquirerName` | `Optional[string]` | Optional | Acquirer name |
| `AcquirerAffiliationCode` | `Optional[string]` | Optional | Acquirer affiliation code |
| `AcquirerTid` | `Optional[string]` | Optional | Acquirer TID |
| `AcquirerNsu` | `Optional[string]` | Optional | Acquirer NSU |
| `AcquirerAuthCode` | `Optional[string]` | Optional | Acquirer authorization code |
| `AcquirerMessage` | `Optional[string]` | Optional | acquirer_message |
| `AcquirerReturnCode` | `Optional[string]` | Optional | Acquirer return code |
| `OperationType` | `Optional[string]` | Optional | Operation type |
| `Card` | [`Optional[models.GetCardResponse]`](../../doc/models/get-card-response.md) | Optional | Card data |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id8",
  "amount": 40,
  "status": "status6",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "statement_descriptor": "statement_descriptor6",
  "acquirer_name": "acquirer_name0",
  "acquirer_affiliation_code": "acquirer_affiliation_code2",
  "acquirer_tid": "acquirer_tid4",
  "acquirer_nsu": "acquirer_nsu4"
}
```

