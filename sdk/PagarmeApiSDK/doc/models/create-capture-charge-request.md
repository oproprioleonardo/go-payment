
# Create Capture Charge Request

Request for capturing a charge

## Structure

`CreateCaptureChargeRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `string` | Required | Code for the charge. Sending this field will update the code send on the charge and order creation. |
| `Amount` | `*int` | Optional | The amount that will be captured |
| `Split` | [`[]models.CreateSplitRequest`](../../doc/models/create-split-request.md) | Optional | Splits |
| `OperationReference` | `string` | Required | - |

## Example (as JSON)

```json
{
  "code": "code8",
  "amount": 96,
  "split": [
    {
      "type": "type2",
      "amount": 10,
      "recipient_id": "recipient_id2",
      "options": {
        "liable": false,
        "charge_processing_fee": false,
        "charge_remainder_fee": false
      },
      "split_rule_id": "split_rule_id0"
    },
    {
      "type": "type2",
      "amount": 10,
      "recipient_id": "recipient_id2",
      "options": {
        "liable": false,
        "charge_processing_fee": false,
        "charge_remainder_fee": false
      },
      "split_rule_id": "split_rule_id0"
    }
  ],
  "operation_reference": "operation_reference0"
}
```

