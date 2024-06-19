# Transfers

```go
transfersController := client.TransfersController()
```

## Class Name

`TransfersController`

## Methods

* [Get Transfer by Id](../../doc/controllers/transfers.md#get-transfer-by-id)
* [Create Transfer](../../doc/controllers/transfers.md#create-transfer)
* [Get Transfers](../../doc/controllers/transfers.md#get-transfers)


# Get Transfer by Id

```go
GetTransferById(
    ctx context.Context,
    transferId string) (
    models.ApiResponse[models.GetTransfer],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `transferId` | `string` | Template, Required | - |

## Response Type

[`models.GetTransfer`](../../doc/models/get-transfer.md)

## Example Usage

```go
ctx := context.Background()
transferId := "transfer_id6"

apiResponse, err := transfersController.GetTransferById(ctx, transferId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Transfer

```go
CreateTransfer(
    ctx context.Context,
    request models.CreateTransfer) (
    models.ApiResponse[models.GetTransfer],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `request` | [`models.CreateTransfer`](../../doc/models/create-transfer.md) | Body, Required | - |

## Response Type

[`models.GetTransfer`](../../doc/models/get-transfer.md)

## Example Usage

```go
ctx := context.Background()

request := models.CreateTransfer{
    Amount:   242,
    SourceId: "source_id0",
    TargetId: "target_id6",
}

apiResponse, err := transfersController.CreateTransfer(ctx, &request)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Transfers

Gets all transfers

```go
GetTransfers(
    ctx context.Context) (
    models.ApiResponse[models.ListTransfers],
    error)
```

## Response Type

[`models.ListTransfers`](../../doc/models/list-transfers.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := transfersController.GetTransfers(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

