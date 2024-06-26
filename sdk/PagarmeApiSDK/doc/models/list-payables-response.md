
# List Payables Response

Response object for listing payable objects

## Structure

`ListPayablesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`Optional[[]models.GetPayableResponse]`](../../doc/models/get-payable-response.md) | Optional | The payable object |
| `Paging` | [`Optional[models.PagingResponse]`](../../doc/models/paging-response.md) | Optional | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": 138,
      "status": "status2",
      "amount": 236,
      "fee": 194,
      "anticipation_fee": 142
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

