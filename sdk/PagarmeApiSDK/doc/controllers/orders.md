# Orders

```go
ordersController := client.OrdersController()
```

## Class Name

`OrdersController`

## Methods

* [Get Orders](../../doc/controllers/orders.md#get-orders)
* [Update Order Item](../../doc/controllers/orders.md#update-order-item)
* [Delete All Order Items](../../doc/controllers/orders.md#delete-all-order-items)
* [Delete Order Item](../../doc/controllers/orders.md#delete-order-item)
* [Close Order](../../doc/controllers/orders.md#close-order)
* [Create Order](../../doc/controllers/orders.md#create-order)
* [Create Order Item](../../doc/controllers/orders.md#create-order-item)
* [Get Order Item](../../doc/controllers/orders.md#get-order-item)
* [Update Order Metadata](../../doc/controllers/orders.md#update-order-metadata)
* [Get Order](../../doc/controllers/orders.md#get-order)


# Get Orders

Gets all orders

```go
GetOrders(
    ctx context.Context,
    page *int,
    size *int,
    code *string,
    status *string,
    createdSince *time.Time,
    createdUntil *time.Time,
    customerId *string) (
    models.ApiResponse[models.ListOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `code` | `*string` | Query, Optional | Filter for order's code |
| `status` | `*string` | Query, Optional | Filter for order's status |
| `createdSince` | `*time.Time` | Query, Optional | Filter for order's creation date start range |
| `createdUntil` | `*time.Time` | Query, Optional | Filter for order's creation date end range |
| `customerId` | `*string` | Query, Optional | Filter for order's customer id |

## Response Type

[`models.ListOrderResponse`](../../doc/models/list-order-response.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := ordersController.GetOrders(ctx, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Order Item

```go
UpdateOrderItem(
    ctx context.Context,
    orderId string,
    itemId string,
    request models.UpdateOrderItemRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `itemId` | `string` | Template, Required | Item Id |
| `request` | [`models.UpdateOrderItemRequest`](../../doc/models/update-order-item-request.md) | Body, Required | Item Model |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
ctx := context.Background()
orderId := "orderId2"
itemId := "itemId8"

request := models.UpdateOrderItemRequest{
    Amount:      242,
    Description: "description6",
    Quantity:    100,
    Category:    "category4",
}

apiResponse, err := ordersController.UpdateOrderItem(ctx, orderId, itemId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete All Order Items

```go
DeleteAllOrderItems(
    ctx context.Context,
    orderId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
ctx := context.Background()
orderId := "orderId2"

apiResponse, err := ordersController.DeleteAllOrderItems(ctx, orderId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Order Item

```go
DeleteOrderItem(
    ctx context.Context,
    orderId string,
    itemId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `itemId` | `string` | Template, Required | Item Id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
ctx := context.Background()
orderId := "orderId2"
itemId := "itemId8"

apiResponse, err := ordersController.DeleteOrderItem(ctx, orderId, itemId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Close Order

```go
CloseOrder(
    ctx context.Context,
    id string,
    request models.UpdateOrderStatusRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `string` | Template, Required | Order Id |
| `request` | [`models.UpdateOrderStatusRequest`](../../doc/models/update-order-status-request.md) | Body, Required | Update Order Model |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
ctx := context.Background()
id := "id0"

request := models.UpdateOrderStatusRequest{
    Status: "status8",
}

apiResponse, err := ordersController.CloseOrder(ctx, id, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Order

Creates a new Order

```go
CreateOrder(
    ctx context.Context,
    body models.CreateOrderRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`models.CreateOrderRequest`](../../doc/models/create-order-request.md) | Body, Required | Request for creating an order |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
ctx := context.Background()

bodyItems0 := models.CreateOrderItemRequest{
    Amount:      164,
    Description: "description2",
    Quantity:    22,
    Category:    "category6",
}

bodyItems := []models.CreateOrderItemRequest{bodyItems0}
bodyCustomerAddress := models.CreateAddressRequest{
    Street:       "street6",
    Number:       "number4",
    ZipCode:      "zip_code0",
    Neighborhood: "neighborhood2",
    City:         "city6",
    State:        "state2",
    Country:      "country0",
    Complement:   "complement2",
    Line1:        "line_10",
    Line2:        "line_24",
}

bodyCustomerPhones := models.CreatePhonesRequest{
}

bodyCustomer := models.CreateCustomerRequest{
    Name:         "{\n    \"name\": \"Tony Stark\"\n}",
    Email:        "email6",
    Document:     "document6",
    Type:         "type0",
    Metadata:     map[string]string{
"key0" : "metadata3",
},
    Code:         "code8",
    Address:      bodyCustomerAddress,
    Phones:       bodyCustomerPhones,
}

bodyPayments0 := models.CreatePaymentRequest{
    PaymentMethod:        "payment_method8",
}

bodyPayments := []models.CreatePaymentRequest{bodyPayments0}
body := models.CreateOrderRequest{
    Code:             "code4",
    Closed:           true,
    Items:            bodyItems,
    Customer:         bodyCustomer,
    Payments:         bodyPayments,
}

apiResponse, err := ordersController.CreateOrder(ctx, &body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Order Item

```go
CreateOrderItem(
    ctx context.Context,
    orderId string,
    request models.CreateOrderItemRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `request` | [`models.CreateOrderItemRequest`](../../doc/models/create-order-item-request.md) | Body, Required | Order Item Model |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
ctx := context.Background()
orderId := "orderId2"
request := models.CreateOrderItemRequest{
    Amount:      242,
    Description: "description6",
    Quantity:    100,
    Category:    "category4",
}


apiResponse, err := ordersController.CreateOrderItem(ctx, orderId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Order Item

```go
GetOrderItem(
    ctx context.Context,
    orderId string,
    itemId string) (
    models.ApiResponse[models.GetOrderItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order Id |
| `itemId` | `string` | Template, Required | Item Id |

## Response Type

[`models.GetOrderItemResponse`](../../doc/models/get-order-item-response.md)

## Example Usage

```go
ctx := context.Background()
orderId := "orderId2"
itemId := "itemId8"

apiResponse, err := ordersController.GetOrderItem(ctx, orderId, itemId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Order Metadata

Updates the metadata from an order

```go
UpdateOrderMetadata(
    ctx context.Context,
    orderId string,
    request models.UpdateMetadataRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | The order id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the order metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
ctx := context.Background()
orderId := "order_id6"

request := models.UpdateMetadataRequest{
    Metadata: map[string]string{
"key0" : "metadata3",
},
}

apiResponse, err := ordersController.UpdateOrderMetadata(ctx, orderId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Order

Gets an order

```go
GetOrder(
    ctx context.Context,
    orderId string) (
    models.ApiResponse[models.GetOrderResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `orderId` | `string` | Template, Required | Order id |

## Response Type

[`models.GetOrderResponse`](../../doc/models/get-order-response.md)

## Example Usage

```go
ctx := context.Background()
orderId := "order_id6"

apiResponse, err := ordersController.GetOrder(ctx, orderId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

