
# List Cycles Response

Response object for listing subscription cycles

## Structure

`ListCyclesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Data` | [`Optional[[]models.GetPeriodResponse]`](../../doc/models/get-period-response.md) | Optional | The subscription cycles objects |
| `Paging` | [`Optional[models.PagingResponse]`](../../doc/models/paging-response.md) | Optional | Paging object |

## Example (as JSON)

```json
{
  "data": [
    {
      "start_at": "2016-03-13T12:52:32.123Z",
      "end_at": "2016-03-13T12:52:32.123Z",
      "id": "id0",
      "billing_at": "2016-03-13T12:52:32.123Z",
      "subscription": {
        "id": "id4",
        "code": "code2",
        "start_at": "2016-03-13T12:52:32.123Z",
        "interval": "interval2",
        "interval_count": 234
      }
    },
    {
      "start_at": "2016-03-13T12:52:32.123Z",
      "end_at": "2016-03-13T12:52:32.123Z",
      "id": "id0",
      "billing_at": "2016-03-13T12:52:32.123Z",
      "subscription": {
        "id": "id4",
        "code": "code2",
        "start_at": "2016-03-13T12:52:32.123Z",
        "interval": "interval2",
        "interval_count": 234
      }
    }
  ],
  "paging": {
    "total": 6,
    "previous": "previous2",
    "next": "next8"
  }
}
```

