# Invoices

```go
invoicesController := client.InvoicesController()
```

## Class Name

`InvoicesController`

## Methods

* [Update Invoice Metadata](../../doc/controllers/invoices.md#update-invoice-metadata)
* [Get Partial Invoice](../../doc/controllers/invoices.md#get-partial-invoice)
* [Cancel Invoice](../../doc/controllers/invoices.md#cancel-invoice)
* [Create Invoice](../../doc/controllers/invoices.md#create-invoice)
* [Get Invoices](../../doc/controllers/invoices.md#get-invoices)
* [Get Invoice](../../doc/controllers/invoices.md#get-invoice)
* [Update Invoice Status](../../doc/controllers/invoices.md#update-invoice-status)


# Update Invoice Metadata

Updates the metadata from an invoice

```go
UpdateInvoiceMetadata(
    ctx context.Context,
    invoiceId string,
    request models.UpdateMetadataRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `invoiceId` | `string` | Template, Required | The invoice id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the invoice metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
ctx := context.Background()
invoiceId := "invoice_id0"

request := models.UpdateMetadataRequest{
    Metadata: map[string]string{
"key0" : "metadata3",
},
}

apiResponse, err := invoicesController.UpdateInvoiceMetadata(ctx, invoiceId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Partial Invoice

```go
GetPartialInvoice(
    ctx context.Context,
    subscriptionId string) (
    models.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

apiResponse, err := invoicesController.GetPartialInvoice(ctx, subscriptionId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Cancel Invoice

Cancels an invoice

```go
CancelInvoice(
    ctx context.Context,
    invoiceId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `invoiceId` | `string` | Template, Required | Invoice id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
ctx := context.Background()
invoiceId := "invoice_id0"

apiResponse, err := invoicesController.CancelInvoice(ctx, invoiceId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Invoice

Create an Invoice

```go
CreateInvoice(
    ctx context.Context,
    subscriptionId string,
    cycleId string,
    request *models.CreateInvoiceRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `cycleId` | `string` | Template, Required | Cycle Id |
| `request` | [`*models.CreateInvoiceRequest`](../../doc/models/create-invoice-request.md) | Body, Optional | - |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
cycleId := "cycle_id6"

apiResponse, err := invoicesController.CreateInvoice(ctx, subscriptionId, cycleId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Invoices

Gets all invoices

```go
GetInvoices(
    ctx context.Context,
    page *int,
    size *int,
    code *string,
    customerId *string,
    subscriptionId *string,
    createdSince *time.Time,
    createdUntil *time.Time,
    status *string,
    dueSince *time.Time,
    dueUntil *time.Time,
    customerDocument *string) (
    models.ApiResponse[models.ListInvoicesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `code` | `*string` | Query, Optional | Filter for Invoice's code |
| `customerId` | `*string` | Query, Optional | Filter for Invoice's customer id |
| `subscriptionId` | `*string` | Query, Optional | Filter for Invoice's subscription id |
| `createdSince` | `*time.Time` | Query, Optional | Filter for Invoice's creation date start range |
| `createdUntil` | `*time.Time` | Query, Optional | Filter for Invoices creation date end range |
| `status` | `*string` | Query, Optional | Filter for Invoice's status |
| `dueSince` | `*time.Time` | Query, Optional | Filter for Invoice's due date start range |
| `dueUntil` | `*time.Time` | Query, Optional | Filter for Invoice's due date end range |
| `customerDocument` | `*string` | Query, Optional | - |

## Response Type

[`models.ListInvoicesResponse`](../../doc/models/list-invoices-response.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := invoicesController.GetInvoices(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Invoice

Gets an invoice

```go
GetInvoice(
    ctx context.Context,
    invoiceId string) (
    models.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `invoiceId` | `string` | Template, Required | Invoice Id |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
ctx := context.Background()
invoiceId := "invoice_id0"

apiResponse, err := invoicesController.GetInvoice(ctx, invoiceId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Invoice Status

Updates the status from an invoice

```go
UpdateInvoiceStatus(
    ctx context.Context,
    invoiceId string,
    request models.UpdateInvoiceStatusRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetInvoiceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `invoiceId` | `string` | Template, Required | Invoice Id |
| `request` | [`models.UpdateInvoiceStatusRequest`](../../doc/models/update-invoice-status-request.md) | Body, Required | Request for updating an invoice's status |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetInvoiceResponse`](../../doc/models/get-invoice-response.md)

## Example Usage

```go
ctx := context.Background()
invoiceId := "invoice_id0"

request := models.UpdateInvoiceStatusRequest{
    Status: "status8",
}

apiResponse, err := invoicesController.UpdateInvoiceStatus(ctx, invoiceId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

