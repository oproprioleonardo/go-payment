
# Get Payable Response

Response object for getting an payable

## Structure

`GetPayableResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[int64]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `Amount` | `Optional[int]` | Optional | - |
| `Fee` | `Optional[int]` | Optional | - |
| `AnticipationFee` | `Optional[int]` | Optional | - |
| `FraudCoverageFee` | `Optional[int]` | Optional | - |
| `Installment` | `Optional[int]` | Optional | - |
| `GatewayId` | `Optional[int]` | Optional | - |
| `ChargeId` | `Optional[string]` | Optional | - |
| `SplitId` | `Optional[string]` | Optional | - |
| `BulkAnticipationId` | `Optional[string]` | Optional | - |
| `AnticipationId` | `Optional[string]` | Optional | - |
| `RecipientId` | `Optional[string]` | Optional | - |
| `OriginatorModel` | `Optional[string]` | Optional | - |
| `OriginatorModelId` | `Optional[string]` | Optional | - |
| `PaymentDate` | `Optional[time.Time]` | Optional | - |
| `OriginalPaymentDate` | `Optional[time.Time]` | Optional | - |
| `Type` | `Optional[string]` | Optional | - |
| `PaymentMethod` | `Optional[string]` | Optional | - |
| `AccrualAt` | `Optional[time.Time]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `LiquidationArrangementId` | `Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "id": 134,
  "status": "status8",
  "amount": 24,
  "fee": 190,
  "anticipation_fee": 118
}
```

