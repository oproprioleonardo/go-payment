# Subscriptions

```go
subscriptionsController := client.SubscriptionsController()
```

## Class Name

`SubscriptionsController`

## Methods

* [Renew Subscription](../../doc/controllers/subscriptions.md#renew-subscription)
* [Update Subscription Card](../../doc/controllers/subscriptions.md#update-subscription-card)
* [Delete Usage](../../doc/controllers/subscriptions.md#delete-usage)
* [Create Discount](../../doc/controllers/subscriptions.md#create-discount)
* [Create an Usage](../../doc/controllers/subscriptions.md#create-an-usage)
* [Update Current Cycle Status](../../doc/controllers/subscriptions.md#update-current-cycle-status)
* [Delete Discount](../../doc/controllers/subscriptions.md#delete-discount)
* [Get Subscription Items](../../doc/controllers/subscriptions.md#get-subscription-items)
* [Update Subscription Payment Method](../../doc/controllers/subscriptions.md#update-subscription-payment-method)
* [Get Subscription Item](../../doc/controllers/subscriptions.md#get-subscription-item)
* [Get Subscriptions](../../doc/controllers/subscriptions.md#get-subscriptions)
* [Cancel Subscription](../../doc/controllers/subscriptions.md#cancel-subscription)
* [Create Increment](../../doc/controllers/subscriptions.md#create-increment)
* [Create Usage](../../doc/controllers/subscriptions.md#create-usage)
* [Get Discount by Id](../../doc/controllers/subscriptions.md#get-discount-by-id)
* [Create Subscription](../../doc/controllers/subscriptions.md#create-subscription)
* [Get Increment by Id](../../doc/controllers/subscriptions.md#get-increment-by-id)
* [Update Subscription Affiliation Id](../../doc/controllers/subscriptions.md#update-subscription-affiliation-id)
* [Update Subscription Metadata](../../doc/controllers/subscriptions.md#update-subscription-metadata)
* [Delete Increment](../../doc/controllers/subscriptions.md#delete-increment)
* [Get Subscription Cycles](../../doc/controllers/subscriptions.md#get-subscription-cycles)
* [Get Discounts](../../doc/controllers/subscriptions.md#get-discounts)
* [Update Subscription Billing Date](../../doc/controllers/subscriptions.md#update-subscription-billing-date)
* [Delete Subscription Item](../../doc/controllers/subscriptions.md#delete-subscription-item)
* [Get Increments](../../doc/controllers/subscriptions.md#get-increments)
* [Update Subscription Due Days](../../doc/controllers/subscriptions.md#update-subscription-due-days)
* [Update Subscription Start At](../../doc/controllers/subscriptions.md#update-subscription-start-at)
* [Update Subscription Item](../../doc/controllers/subscriptions.md#update-subscription-item)
* [Create Subscription Item](../../doc/controllers/subscriptions.md#create-subscription-item)
* [Get Subscription](../../doc/controllers/subscriptions.md#get-subscription)
* [Get Usages](../../doc/controllers/subscriptions.md#get-usages)
* [Update Latest Period End At](../../doc/controllers/subscriptions.md#update-latest-period-end-at)
* [Update Subscription Minium Price](../../doc/controllers/subscriptions.md#update-subscription-minium-price)
* [Get Subscription Cycle by Id](../../doc/controllers/subscriptions.md#get-subscription-cycle-by-id)
* [Get Usage Report](../../doc/controllers/subscriptions.md#get-usage-report)
* [Update Split Subscription](../../doc/controllers/subscriptions.md#update-split-subscription)


# Renew Subscription

```go
RenewSubscription(
    ctx context.Context,
    subscriptionId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetPeriodResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | - |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPeriodResponse`](../../doc/models/get-period-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

apiResponse, err := subscriptionsController.RenewSubscription(ctx, subscriptionId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Card

Updates the credit card from a subscription

```go
UpdateSubscriptionCard(
    ctx context.Context,
    subscriptionId string,
    request models.UpdateSubscriptionCardRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`models.UpdateSubscriptionCardRequest`](../../doc/models/update-subscription-card-request.md) | Body, Required | Request for updating a card |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

requestCard := models.CreateCardRequest{
    Type:             models.ToPointer("credit"),
}

request := models.UpdateSubscriptionCardRequest{
    CardId: "card_id2",
    Card:   requestCard,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionCard(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Usage

Deletes a usage

```go
DeleteUsage(
    ctx context.Context,
    subscriptionId string,
    itemId string,
    usageId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetUsageResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `itemId` | `string` | Template, Required | The subscription item id |
| `usageId` | `string` | Template, Required | The usage id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetUsageResponse`](../../doc/models/get-usage-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
itemId := "item_id0"
usageId := "usage_id0"

apiResponse, err := subscriptionsController.DeleteUsage(ctx, subscriptionId, itemId, usageId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Discount

Creates a discount

```go
CreateDiscount(
    ctx context.Context,
    subscriptionId string,
    request models.CreateDiscountRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetDiscountResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`models.CreateDiscountRequest`](../../doc/models/create-discount-request.md) | Body, Required | Request for creating a discount |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetDiscountResponse`](../../doc/models/get-discount-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

request := models.CreateDiscountRequest{
    Value:        float64(185.28),
    DiscountType: "discount_type4",
    ItemId:       "item_id6",
}

apiResponse, err := subscriptionsController.CreateDiscount(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create an Usage

Create Usage

```go
CreateAnUsage(
    ctx context.Context,
    subscriptionId string,
    itemId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetUsageResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `itemId` | `string` | Template, Required | Item id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetUsageResponse`](../../doc/models/get-usage-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
itemId := "item_id0"

apiResponse, err := subscriptionsController.CreateAnUsage(ctx, subscriptionId, itemId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Current Cycle Status

```go
UpdateCurrentCycleStatus(
    ctx context.Context,
    subscriptionId string,
    request models.UpdateCurrentCycleStatusRequest,
    idempotencyKey *string) (
    http.Response,
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `request` | [`models.UpdateCurrentCycleStatusRequest`](../../doc/models/update-current-cycle-status-request.md) | Body, Required | Request for updating the end date of the subscription current status |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

``

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

request := models.UpdateCurrentCycleStatusRequest{
    Status: "status8",
}

resp, err := subscriptionsController.UpdateCurrentCycleStatus(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    fmt.Println(resp.StatusCode)
}
```


# Delete Discount

Deletes a discount

```go
DeleteDiscount(
    ctx context.Context,
    subscriptionId string,
    discountId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetDiscountResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `discountId` | `string` | Template, Required | Discount Id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetDiscountResponse`](../../doc/models/get-discount-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
discountId := "discount_id8"

apiResponse, err := subscriptionsController.DeleteDiscount(ctx, subscriptionId, discountId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscription Items

Get Subscription Items

```go
GetSubscriptionItems(
    ctx context.Context,
    subscriptionId string,
    page *int,
    size *int,
    name *string,
    code *string,
    status *string,
    description *string,
    createdSince *string,
    createdUntil *string) (
    models.ApiResponse[models.ListSubscriptionItemsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `name` | `*string` | Query, Optional | The item name |
| `code` | `*string` | Query, Optional | Identification code in the client system |
| `status` | `*string` | Query, Optional | The item statis |
| `description` | `*string` | Query, Optional | The item description |
| `createdSince` | `*string` | Query, Optional | Filter for item's creation date start range |
| `createdUntil` | `*string` | Query, Optional | Filter for item's creation date end range |

## Response Type

[`models.ListSubscriptionItemsResponse`](../../doc/models/list-subscription-items-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

apiResponse, err := subscriptionsController.GetSubscriptionItems(ctx, subscriptionId, nil, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Payment Method

Updates the payment method from a subscription

```go
UpdateSubscriptionPaymentMethod(
    ctx context.Context,
    subscriptionId string,
    request models.UpdateSubscriptionPaymentMethodRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`models.UpdateSubscriptionPaymentMethodRequest`](../../doc/models/update-subscription-payment-method-request.md) | Body, Required | Request for updating the paymentmethod from a subscription |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

requestCard := models.CreateCardRequest{
    Type:             models.ToPointer("credit"),
}

request := models.UpdateSubscriptionPaymentMethodRequest{
    PaymentMethod: "payment_method4",
    CardId:        "card_id2",
    Card:          requestCard,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionPaymentMethod(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscription Item

Get Subscription Item

```go
GetSubscriptionItem(
    ctx context.Context,
    subscriptionId string,
    itemId string) (
    models.ApiResponse[models.GetSubscriptionItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `itemId` | `string` | Template, Required | Item id |

## Response Type

[`models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
itemId := "item_id0"

apiResponse, err := subscriptionsController.GetSubscriptionItem(ctx, subscriptionId, itemId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscriptions

Gets all subscriptions

```go
GetSubscriptions(
    ctx context.Context,
    page *int,
    size *int,
    code *string,
    billingType *string,
    customerId *string,
    planId *string,
    cardId *string,
    status *string,
    nextBillingSince *time.Time,
    nextBillingUntil *time.Time,
    createdSince *time.Time,
    createdUntil *time.Time) (
    models.ApiResponse[models.ListSubscriptionsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `code` | `*string` | Query, Optional | Filter for subscription's code |
| `billingType` | `*string` | Query, Optional | Filter for subscription's billing type |
| `customerId` | `*string` | Query, Optional | Filter for subscription's customer id |
| `planId` | `*string` | Query, Optional | Filter for subscription's plan id |
| `cardId` | `*string` | Query, Optional | Filter for subscription's card id |
| `status` | `*string` | Query, Optional | Filter for subscription's status |
| `nextBillingSince` | `*time.Time` | Query, Optional | Filter for subscription's next billing date start range |
| `nextBillingUntil` | `*time.Time` | Query, Optional | Filter for subscription's next billing date end range |
| `createdSince` | `*time.Time` | Query, Optional | Filter for subscription's creation date start range |
| `createdUntil` | `*time.Time` | Query, Optional | Filter for subscriptions creation date end range |

## Response Type

[`models.ListSubscriptionsResponse`](../../doc/models/list-subscriptions-response.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := subscriptionsController.GetSubscriptions(ctx, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Cancel Subscription

Cancels a subscription

```go
CancelSubscription(
    ctx context.Context,
    subscriptionId string,
    request *models.CreateCancelSubscriptionRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`*models.CreateCancelSubscriptionRequest`](../../doc/models/create-cancel-subscription-request.md) | Body, Optional | Request for cancelling a subscription |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

request := models.CreateCancelSubscriptionRequest{
    CancelPendingInvoices: true,
}

apiResponse, err := subscriptionsController.CancelSubscription(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Increment

Creates a increment

```go
CreateIncrement(
    ctx context.Context,
    subscriptionId string,
    request models.CreateIncrementRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetIncrementResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`models.CreateIncrementRequest`](../../doc/models/create-increment-request.md) | Body, Required | Request for creating a increment |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetIncrementResponse`](../../doc/models/get-increment-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

request := models.CreateIncrementRequest{
    Value:         float64(185.28),
    IncrementType: "increment_type8",
    ItemId:        "item_id6",
}

apiResponse, err := subscriptionsController.CreateIncrement(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Usage

Creates a usage

```go
CreateUsage(
    ctx context.Context,
    subscriptionId string,
    itemId string,
    body models.CreateUsageRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetUsageResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `itemId` | `string` | Template, Required | Item id |
| `body` | [`models.CreateUsageRequest`](../../doc/models/create-usage-request.md) | Body, Required | Request for creating a usage |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetUsageResponse`](../../doc/models/get-usage-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
itemId := "item_id0"

bodyUsedAt, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
if err != nil {
    log.Fatalln(err)
}
body := models.CreateUsageRequest{
    Quantity:    156,
    Description: "description4",
    UsedAt:      bodyUsedAt,
}

apiResponse, err := subscriptionsController.CreateUsage(ctx, subscriptionId, itemId, &body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Discount by Id

```go
GetDiscountById(
    ctx context.Context,
    subscriptionId string,
    discountId string) (
    models.ApiResponse[models.GetDiscountResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `discountId` | `string` | Template, Required | - |

## Response Type

[`models.GetDiscountResponse`](../../doc/models/get-discount-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
discountId := "discountId0"

apiResponse, err := subscriptionsController.GetDiscountById(ctx, subscriptionId, discountId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Subscription

Creates a new subscription

```go
CreateSubscription(
    ctx context.Context,
    body models.CreateSubscriptionRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`models.CreateSubscriptionRequest`](../../doc/models/create-subscription-request.md) | Body, Required | Request for creating a subscription |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()

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

bodyCard := models.CreateCardRequest{
    Type:             models.ToPointer("credit"),
}

bodyPricingScheme := models.CreatePricingSchemeRequest{
    SchemeType:    "scheme_type8",
}

bodyItems0PricingScheme := models.CreatePricingSchemeRequest{
    SchemeType:    "scheme_type8",
}


bodyItems0Discounts0 := models.CreateDiscountRequest{
    Value:        float64(90.66),
    DiscountType: "discount_type2",
    ItemId:       "item_id4",
}
bodyItems0Discounts := []models.CreateDiscountRequest{bodyItems0Discounts0}
bodyItems0 := models.CreateSubscriptionItemRequest{
    Description:   "description2",
    Id:            "id8",
    PlanItemId:    "plan_item_id8",
    Name:          "name8",
    PricingScheme: bodyItems0PricingScheme,
    Discounts:     bodyItems0Discounts,
}

bodyItems := []models.CreateSubscriptionItemRequest{bodyItems0}
bodyShippingAddress := models.CreateAddressRequest{
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

bodyShipping := models.CreateShippingRequest{
    Amount:                52,
    Description:           "description6",
    RecipientName:         "recipient_name2",
    RecipientPhone:        "recipient_phone6",
    AddressId:             "address_id6",
    Type:                  "type6",
    Address:               bodyShippingAddress,
}


bodyDiscounts0 := models.CreateDiscountRequest{
    Value:        float64(90.66),
    DiscountType: "discount_type2",
    ItemId:       "item_id4",
}
bodyDiscounts := []models.CreateDiscountRequest{bodyDiscounts0}

bodyIncrements0 := models.CreateIncrementRequest{
    Value:         float64(252.86),
    IncrementType: "increment_type6",
    ItemId:        "item_id6",
}
bodyIncrements := []models.CreateIncrementRequest{bodyIncrements0}
body := models.CreateSubscriptionRequest{
    Code:                 "code4",
    PaymentMethod:        "payment_method4",
    BillingType:          "billing_type0",
    StatementDescriptor:  "statement_descriptor6",
    Description:          "description4",
    Currency:             "currency6",
    Interval:             "interval6",
    IntervalCount:        170,
    Metadata:             map[string]string{
"key0" : "metadata7",
"key1" : "metadata8",
},
    Customer:             bodyCustomer,
    Card:                 bodyCard,
    PricingScheme:        bodyPricingScheme,
    Items:                bodyItems,
    Shipping:             bodyShipping,
    Discounts:            bodyDiscounts,
    Increments:           bodyIncrements,
}

apiResponse, err := subscriptionsController.CreateSubscription(ctx, &body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Increment by Id

```go
GetIncrementById(
    ctx context.Context,
    subscriptionId string,
    incrementId string) (
    models.ApiResponse[models.GetIncrementResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription Id |
| `incrementId` | `string` | Template, Required | The increment Id |

## Response Type

[`models.GetIncrementResponse`](../../doc/models/get-increment-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
incrementId := "increment_id8"

apiResponse, err := subscriptionsController.GetIncrementById(ctx, subscriptionId, incrementId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Affiliation Id

```go
UpdateSubscriptionAffiliationId(
    ctx context.Context,
    subscriptionId string,
    request models.UpdateSubscriptionAffiliationIdRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | - |
| `request` | [`models.UpdateSubscriptionAffiliationIdRequest`](../../doc/models/update-subscription-affiliation-id-request.md) | Body, Required | Request for updating a subscription affiliation id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

request := models.UpdateSubscriptionAffiliationIdRequest{
    GatewayAffiliationId: "gateway_affiliation_id2",
}

apiResponse, err := subscriptionsController.UpdateSubscriptionAffiliationId(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Metadata

Updates the metadata from a subscription

```go
UpdateSubscriptionMetadata(
    ctx context.Context,
    subscriptionId string,
    request models.UpdateMetadataRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the subscrption metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

request := models.UpdateMetadataRequest{
    Metadata: map[string]string{
"key0" : "metadata3",
},
}

apiResponse, err := subscriptionsController.UpdateSubscriptionMetadata(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Increment

Deletes a increment

```go
DeleteIncrement(
    ctx context.Context,
    subscriptionId string,
    incrementId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetIncrementResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `incrementId` | `string` | Template, Required | Increment id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetIncrementResponse`](../../doc/models/get-increment-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
incrementId := "increment_id8"

apiResponse, err := subscriptionsController.DeleteIncrement(ctx, subscriptionId, incrementId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscription Cycles

```go
GetSubscriptionCycles(
    ctx context.Context,
    subscriptionId string,
    page string,
    size string) (
    models.ApiResponse[models.ListCyclesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `page` | `string` | Query, Required | Page number |
| `size` | `string` | Query, Required | Page size |

## Response Type

[`models.ListCyclesResponse`](../../doc/models/list-cycles-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
page := "page8"
size := "size0"

apiResponse, err := subscriptionsController.GetSubscriptionCycles(ctx, subscriptionId, page, size)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Discounts

```go
GetDiscounts(
    ctx context.Context,
    subscriptionId string,
    page int,
    size int) (
    models.ApiResponse[models.ListDiscountsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `page` | `int` | Query, Required | Page number |
| `size` | `int` | Query, Required | Page size |

## Response Type

[`models.ListDiscountsResponse`](../../doc/models/list-discounts-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
page := 30
size := 18

apiResponse, err := subscriptionsController.GetDiscounts(ctx, subscriptionId, page, size)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Billing Date

Updates the billing date from a subscription

```go
UpdateSubscriptionBillingDate(
    ctx context.Context,
    subscriptionId string,
    request models.UpdateSubscriptionBillingDateRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `request` | [`models.UpdateSubscriptionBillingDateRequest`](../../doc/models/update-subscription-billing-date-request.md) | Body, Required | Request for updating the subscription billing date |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

requestNextBillingAt, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
if err != nil {
    log.Fatalln(err)
}
request := models.UpdateSubscriptionBillingDateRequest{
    NextBillingAt: requestNextBillingAt,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionBillingDate(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Subscription Item

Deletes a subscription item

```go
DeleteSubscriptionItem(
    ctx context.Context,
    subscriptionId string,
    subscriptionItemId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `subscriptionItemId` | `string` | Template, Required | Subscription item id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
subscriptionItemId := "subscription_item_id4"

apiResponse, err := subscriptionsController.DeleteSubscriptionItem(ctx, subscriptionId, subscriptionItemId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Increments

```go
GetIncrements(
    ctx context.Context,
    subscriptionId string,
    page *int,
    size *int) (
    models.ApiResponse[models.ListIncrementsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |

## Response Type

[`models.ListIncrementsResponse`](../../doc/models/list-increments-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

apiResponse, err := subscriptionsController.GetIncrements(ctx, subscriptionId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Due Days

Updates the boleto due days from a subscription

```go
UpdateSubscriptionDueDays(
    ctx context.Context,
    subscriptionId string,
    request models.UpdateSubscriptionDueDaysRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `request` | [`models.UpdateSubscriptionDueDaysRequest`](../../doc/models/update-subscription-due-days-request.md) | Body, Required | - |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

request := models.UpdateSubscriptionDueDaysRequest{
    BoletoDueDays: 226,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionDueDays(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Start At

Updates the start at date from a subscription

```go
UpdateSubscriptionStartAt(
    ctx context.Context,
    subscriptionId string,
    request models.UpdateSubscriptionStartAtRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `request` | [`models.UpdateSubscriptionStartAtRequest`](../../doc/models/update-subscription-start-at-request.md) | Body, Required | Request for updating the subscription start date |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

requestStartAt, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
if err != nil {
    log.Fatalln(err)
}
request := models.UpdateSubscriptionStartAtRequest{
    StartAt: requestStartAt,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionStartAt(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Item

Updates a subscription item

```go
UpdateSubscriptionItem(
    ctx context.Context,
    subscriptionId string,
    itemId string,
    body models.UpdateSubscriptionItemRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `itemId` | `string` | Template, Required | Item id |
| `body` | [`models.UpdateSubscriptionItemRequest`](../../doc/models/update-subscription-item-request.md) | Body, Required | Request for updating a subscription item |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
itemId := "item_id0"

bodyPricingSchemePriceBrackets0 := models.UpdatePriceBracketRequest{
    StartQuantity: 144,
    Price:         174,
}

bodyPricingSchemePriceBrackets := []models.UpdatePriceBracketRequest{bodyPricingSchemePriceBrackets0}
bodyPricingScheme := models.UpdatePricingSchemeRequest{
    SchemeType:    "scheme_type8",
    PriceBrackets: bodyPricingSchemePriceBrackets,
}

body := models.UpdateSubscriptionItemRequest{
    Description:   "description4",
    Status:        "status2",
    Name:          "name6",
    PricingScheme: bodyPricingScheme,
}

apiResponse, err := subscriptionsController.UpdateSubscriptionItem(ctx, subscriptionId, itemId, &body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Subscription Item

Creates a new Subscription item

```go
CreateSubscriptionItem(
    ctx context.Context,
    subscriptionId string,
    request models.CreateSubscriptionItemRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |
| `request` | [`models.CreateSubscriptionItemRequest`](../../doc/models/create-subscription-item-request.md) | Body, Required | Request for creating a subscription item |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionItemResponse`](../../doc/models/get-subscription-item-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
requestPricingScheme := models.CreatePricingSchemeRequest{
    SchemeType:    "scheme_type8",
}


requestDiscounts0 := models.CreateDiscountRequest{
    Value:        float64(90.66),
    DiscountType: "discount_type2",
    ItemId:       "item_id4",
}
requestDiscounts := []models.CreateDiscountRequest{requestDiscounts0}
request := models.CreateSubscriptionItemRequest{
    Description:   "description6",
    Id:            "id6",
    PlanItemId:    "plan_item_id6",
    Name:          "name6",
    PricingScheme: requestPricingScheme,
    Discounts:     requestDiscounts,
}


apiResponse, err := subscriptionsController.CreateSubscriptionItem(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscription

Gets a subscription

```go
GetSubscription(
    ctx context.Context,
    subscriptionId string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription id |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

apiResponse, err := subscriptionsController.GetSubscription(ctx, subscriptionId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Usages

Lists all usages from a subscription item

```go
GetUsages(
    ctx context.Context,
    subscriptionId string,
    itemId string,
    page *int,
    size *int,
    code *string,
    group *string,
    usedSince *time.Time,
    usedUntil *time.Time) (
    models.ApiResponse[models.ListUsagesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `itemId` | `string` | Template, Required | The subscription item id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `code` | `*string` | Query, Optional | Identification code in the client system |
| `group` | `*string` | Query, Optional | Identification group in the client system |
| `usedSince` | `*time.Time` | Query, Optional | - |
| `usedUntil` | `*time.Time` | Query, Optional | - |

## Response Type

[`models.ListUsagesResponse`](../../doc/models/list-usages-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
itemId := "item_id0"

apiResponse, err := subscriptionsController.GetUsages(ctx, subscriptionId, itemId, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Latest Period End At

```go
UpdateLatestPeriodEndAt(
    ctx context.Context,
    subscriptionId string,
    request models.UpdateCurrentCycleEndDateRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | - |
| `request` | [`models.UpdateCurrentCycleEndDateRequest`](../../doc/models/update-current-cycle-end-date-request.md) | Body, Required | Request for updating the end date of the current signature cycle |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

request := models.UpdateCurrentCycleEndDateRequest{
}

apiResponse, err := subscriptionsController.UpdateLatestPeriodEndAt(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Subscription Minium Price

Atualização do valor mínimo da assinatura

```go
UpdateSubscriptionMiniumPrice(
    ctx context.Context,
    subscriptionId string,
    request models.UpdateSubscriptionMinimumPriceRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | Subscription Id |
| `request` | [`models.UpdateSubscriptionMinimumPriceRequest`](../../doc/models/update-subscription-minimum-price-request.md) | Body, Required | Request da requisição com o valor mínimo que será configurado |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"

request := models.UpdateSubscriptionMinimumPriceRequest{
}

apiResponse, err := subscriptionsController.UpdateSubscriptionMiniumPrice(ctx, subscriptionId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Subscription Cycle by Id

```go
GetSubscriptionCycleById(
    ctx context.Context,
    subscriptionId string,
    cycleId string) (
    models.ApiResponse[models.GetPeriodResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription id |
| `cycleId` | `string` | Template, Required | - |

## Response Type

[`models.GetPeriodResponse`](../../doc/models/get-period-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
cycleId := "cycleId0"

apiResponse, err := subscriptionsController.GetSubscriptionCycleById(ctx, subscriptionId, cycleId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Usage Report

```go
GetUsageReport(
    ctx context.Context,
    subscriptionId string,
    periodId string) (
    models.ApiResponse[models.GetUsageReportResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `subscriptionId` | `string` | Template, Required | The subscription Id |
| `periodId` | `string` | Template, Required | The period Id |

## Response Type

[`models.GetUsageReportResponse`](../../doc/models/get-usage-report-response.md)

## Example Usage

```go
ctx := context.Background()
subscriptionId := "subscription_id0"
periodId := "period_id0"

apiResponse, err := subscriptionsController.GetUsageReport(ctx, subscriptionId, periodId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Split Subscription

```go
UpdateSplitSubscription(
    ctx context.Context,
    id string,
    request models.UpdateSubscriptionSplitRequest) (
    models.ApiResponse[models.GetSubscriptionResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `id` | `string` | Template, Required | Subscription's id |
| `request` | [`models.UpdateSubscriptionSplitRequest`](../../doc/models/update-subscription-split-request.md) | Body, Required | - |

## Response Type

[`models.GetSubscriptionResponse`](../../doc/models/get-subscription-response.md)

## Example Usage

```go
ctx := context.Background()
id := "id0"

requestRules0 := models.CreateSplitRequest{
    Type:        "type2",
    Amount:      118,
    RecipientId: "recipient_id2",
}

requestRules := []models.CreateSplitRequest{requestRules0}
request := models.UpdateSubscriptionSplitRequest{
    Enabled: false,
    Rules:   requestRules,
}

apiResponse, err := subscriptionsController.UpdateSplitSubscription(ctx, id, &request)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

