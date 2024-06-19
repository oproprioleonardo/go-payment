
# List Plans Response

Response object for listing plans

## Structure

`ListPlansResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`Optional[[]models.GetPlanResponse]`](../../doc/models/get-plan-response.md) | Optional | The plan objects |
| `Paging` | [`Optional[models.PagingResponse]`](../../doc/models/paging-response.md) | Optional | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "id": "id0",
      "name": "name0",
      "description": "description0",
      "url": "url4",
      "statement_descriptor": "statement_descriptor0"
    },
    {
      "id": "id0",
      "name": "name0",
      "description": "description0",
      "url": "url4",
      "statement_descriptor": "statement_descriptor0"
    },
    {
      "id": "id0",
      "name": "name0",
      "description": "description0",
      "url": "url4",
      "statement_descriptor": "statement_descriptor0"
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

