
# Get Three D Secure Response

3D-S payment authentication response

## Structure

`GetThreeDSecureResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mpi` | `Optional[string]` | Optional | MPI Vendor |
| `Eci` | `Optional[string]` | Optional | Electronic Commerce Indicator (ECI) (Opcional) |
| `Cavv` | `Optional[string]` | Optional | Online payment cryptogram, definido pelo 3-D Secure. |
| `TransactionId` | `Optional[string]` | Optional | Identificador da transação (XID) |
| `SuccessUrl` | `Optional[string]` | Optional | Url de redirecionamento de sucessso |

## Example (as JSON)

```json
{
  "mpi": "mpi2",
  "eci": "eci4",
  "cavv": "cavv0",
  "transaction_Id": "transaction_Id0",
  "success_url": "success_url6"
}
```

