
# List Invoices Response

Response object for listing invoices

## Structure

`ListInvoicesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`Optional[[]models.GetInvoiceResponse]`](../../doc/models/get-invoice-response.md) | Optional | The Invoice objects |
| `Paging` | [`Optional[models.PagingResponse]`](../../doc/models/paging-response.md) | Optional | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id0",
      "code": "code8",
      "url": "url4",
      "amount": 236,
      "status": "status2"
    },
    {
      "id": "id0",
      "code": "code8",
      "url": "url4",
      "amount": 236,
      "status": "status2"
    },
    {
      "id": "id0",
      "code": "code8",
      "url": "url4",
      "amount": 236,
      "status": "status2"
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

