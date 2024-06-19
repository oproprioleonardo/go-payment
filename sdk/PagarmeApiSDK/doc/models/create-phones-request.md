
# Create Phones Request

## Structure

`CreatePhonesRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `HomePhone` | [`*models.CreatePhoneRequest`](../../doc/models/create-phone-request.md) | Optional | - |
| `MobilePhone` | [`*models.CreatePhoneRequest`](../../doc/models/create-phone-request.md) | Optional | - |

## Example (as JSON)

```json
{
  "home_phone": {
    "country_code": "country_code0",
    "number": "number2",
    "area_code": "area_code0",
    "Type": "Type0"
  },
  "mobile_phone": {
    "country_code": "country_code0",
    "number": "number8",
    "area_code": "area_code0",
    "Type": "Type0"
  }
}
```

