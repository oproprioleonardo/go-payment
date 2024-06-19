
# Get Customer Response

Response object for getting a customer

## Structure

`GetCustomerResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `Optional[string]` | Optional | - |
| `Name` | `Optional[string]` | Optional | - |
| `Email` | `Optional[string]` | Optional | - |
| `Delinquent` | `Optional[bool]` | Optional | - |
| `CreatedAt` | `Optional[time.Time]` | Optional | - |
| `UpdatedAt` | `Optional[time.Time]` | Optional | - |
| `Document` | `Optional[string]` | Optional | - |
| `Type` | `Optional[string]` | Optional | - |
| `FbAccessToken` | `Optional[string]` | Optional | - |
| `Address` | [`Optional[models.GetAddressResponse]`](../../doc/models/get-address-response.md) | Optional | - |
| `Metadata` | `Optional[map[string]string]` | Optional | - |
| `Phones` | [`Optional[models.GetPhonesResponse]`](../../doc/models/get-phones-response.md) | Optional | - |
| `FbId` | `Optional[int64]` | Optional | - |
| `Code` | `Optional[string]` | Optional | Código de referência do cliente no sistema da loja. Max: 52 caracteres |
| `DocumentType` | `Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "id": "id4",
  "name": "name4",
  "email": "email2",
  "delinquent": false,
  "created_at": "2016-03-13T12:52:32.123Z"
}
```

