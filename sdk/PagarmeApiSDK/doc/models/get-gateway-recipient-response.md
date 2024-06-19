
# Get Gateway Recipient Response

Information about the recipient on the gateway

## Structure

`GetGatewayRecipientResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `Optional[string]` | Optional | Gateway name |
| `Status` | `Optional[string]` | Optional | Status of the recipient on the gateway |
| `Pgid` | `Optional[string]` | Optional | Recipient id on the gateway |
| `CreatedAt` | `Optional[string]` | Optional | Creation date |
| `UpdatedAt` | `Optional[string]` | Optional | Last update date |

## Example (as JSON)

```json
{
  "gateway": "gateway2",
  "status": "status4",
  "pgid": "pgid8",
  "created_at": "created_at0",
  "updated_at": "updated_at8"
}
```

