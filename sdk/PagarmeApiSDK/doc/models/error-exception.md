
# Error Exception

Api Error Exception

## Structure

`ErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Message` | `*string` | Required | - |
| `Errors` | `interface{}` | Required | - |
| `Request` | `interface{}` | Required | - |

## Example (as JSON)

```json
{
  "message": "message4",
  "errors": {
    "key1": "val1",
    "key2": "val2"
  },
  "request": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

