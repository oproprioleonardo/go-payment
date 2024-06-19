
# Get Address Response

Response object for getting an Address

## Structure

`GetAddressResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Street` | `Optional[string]` | Optional | - |
| `Number` | `Optional[string]` | Optional | - |
| `Complement` | `Optional[string]` | Optional | - |
| `ZipCode` | `Optional[string]` | Optional | - |
| `Neighborhood` | `Optional[string]` | Optional | - |
| `City` | `Optional[string]` | Optional | - |
| `State` | `Optional[string]` | Optional | - |
| `Country` | `Optional[string]` | Optional | - |
| `Status` | `Optional[string]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |
| `Customer` | [`Optional[models.GetCustomerResponse]`](../../doc/models/get-customer-response.md) | Optional | - |
| `Metadata` | `Optional[map[string]string]` | Optional | - |
| `Line1` | `Optional[string]` | Optional | Line 1 for address |
| `Line2` | `Optional[string]` | Optional | Line 2 for address |
| `DeletedAt` | `Optional[time.Time]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id2",
  "street": "street2",
  "number": "number0",
  "complement": "complement8",
  "zip_code": "zip_code6"
}
```

