# Payables

```go
payablesController := client.PayablesController()
```

## Class Name

`PayablesController`

## Methods

* [Get Payables](../../doc/controllers/payables.md#get-payables)
* [Get Payable by Id](../../doc/controllers/payables.md#get-payable-by-id)


# Get Payables

```go
GetPayables(
    ctx context.Context,
    mType *string,
    splitId *string,
    bulkAnticipationId *string,
    installment *int,
    status *string,
    recipientId *string,
    amount *int,
    chargeId *string,
    paymentDateUntil *string,
    paymentDateSince *time.Time,
    updatedUntil *time.Time,
    updatedSince *time.Time,
    createdUntil *time.Time,
    createdSince *time.Time,
    liquidationArrangementId *string,
    page *int,
    size *int,
    gatewayId *int64) (
    models.ApiResponse[models.ListPayablesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `mType` | `*string` | Query, Optional | - |
| `splitId` | `*string` | Query, Optional | - |
| `bulkAnticipationId` | `*string` | Query, Optional | - |
| `installment` | `*int` | Query, Optional | - |
| `status` | `*string` | Query, Optional | - |
| `recipientId` | `*string` | Query, Optional | - |
| `amount` | `*int` | Query, Optional | - |
| `chargeId` | `*string` | Query, Optional | - |
| `paymentDateUntil` | `*string` | Query, Optional | - |
| `paymentDateSince` | `*time.Time` | Query, Optional | - |
| `updatedUntil` | `*time.Time` | Query, Optional | - |
| `updatedSince` | `*time.Time` | Query, Optional | - |
| `createdUntil` | `*time.Time` | Query, Optional | - |
| `createdSince` | `*time.Time` | Query, Optional | - |
| `liquidationArrangementId` | `*string` | Query, Optional | - |
| `page` | `*int` | Query, Optional | - |
| `size` | `*int` | Query, Optional | - |
| `gatewayId` | `*int64` | Query, Optional | - |

## Response Type

[`models.ListPayablesResponse`](../../doc/models/list-payables-response.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := payablesController.GetPayables(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Payable by Id

```go
GetPayableById(
    ctx context.Context,
    id int64) (
    models.ApiResponse[models.GetPayableResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `int64` | Template, Required | - |

## Response Type

[`models.GetPayableResponse`](../../doc/models/get-payable-response.md)

## Example Usage

```go
ctx := context.Background()
id := int64(112)

apiResponse, err := payablesController.GetPayableById(ctx, id)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

