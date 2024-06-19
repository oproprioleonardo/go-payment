
# Get Pix Transaction Response

Response object when getting a pix transaction

## Structure

`GetPixTransactionResponse`

## Inherits From

[`GetTransactionResponse`](../../doc/models/get-transaction-response.md)

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `QrCode` | `Optional[string]` | Optional | - |
| `QrCodeUrl` | `Optional[string]` | Optional | - |
| `ExpiresAt` | `Optional[time.Time]` | Optional | - |
| `AdditionalInformation` | [`Optional[[]models.PixAdditionalInformation]`](../../doc/models/pix-additional-information.md) | Optional | - |
| `EndToEndId` | `Optional[string]` | Optional | - |
| `Payer` | [`Optional[models.GetPixPayerResponse]`](../../doc/models/get-pix-payer-response.md) | Optional | - |
| `PixProviderTid` | `Optional[string]` | Optional | Pix provider TID |

## Example (as JSON)

```json
{
  "gateway_id": "gateway_id8",
  "amount": 40,
  "status": "status6",
  "success": false,
  "created_at": "2016-03-13T12:52:32.123Z",
  "qr_code": "qr_code0",
  "qr_code_url": "qr_code_url6",
  "expires_at": "2016-03-13T12:52:32.123Z",
  "additional_information": [
    {
      "Name": "Name0",
      "Value": "Value2"
    },
    {
      "Name": "Name0",
      "Value": "Value2"
    }
  ],
  "end_to_end_id": "end_to_end_id6"
}
```

