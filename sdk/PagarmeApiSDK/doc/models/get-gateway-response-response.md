
# Get Gateway Response Response

The Transaction Gateway Response

## Structure

`GetGatewayResponseResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Code` | `Optional[string]` | Optional | The error code |
| `Errors` | [`Optional[[]models.GetGatewayErrorResponse]`](../../doc/models/get-gateway-error-response.md) | Optional | The gateway response errors list |

## Example (as JSON)

```json
{
  "code": "code6",
  "errors": [
    {
      "message": "message0"
    },
    {
      "message": "message0"
    }
  ]
}
```

