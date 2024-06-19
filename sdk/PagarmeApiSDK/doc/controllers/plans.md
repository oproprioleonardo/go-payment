# Plans

```go
plansController := client.PlansController()
```

## Class Name

`PlansController`

## Methods

* [Get Plan](../../doc/controllers/plans.md#get-plan)
* [Delete Plan](../../doc/controllers/plans.md#delete-plan)
* [Update Plan Metadata](../../doc/controllers/plans.md#update-plan-metadata)
* [Update Plan Item](../../doc/controllers/plans.md#update-plan-item)
* [Create Plan Item](../../doc/controllers/plans.md#create-plan-item)
* [Get Plan Item](../../doc/controllers/plans.md#get-plan-item)
* [Create Plan](../../doc/controllers/plans.md#create-plan)
* [Delete Plan Item](../../doc/controllers/plans.md#delete-plan-item)
* [Get Plans](../../doc/controllers/plans.md#get-plans)
* [Update Plan](../../doc/controllers/plans.md#update-plan)


# Get Plan

Gets a plan

```go
GetPlan(
    ctx context.Context,
    planId string) (
    models.ApiResponse[models.GetPlanResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |

## Response Type

[`models.GetPlanResponse`](../../doc/models/get-plan-response.md)

## Example Usage

```go
ctx := context.Background()
planId := "plan_id8"

apiResponse, err := plansController.GetPlan(ctx, planId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Plan

Deletes a plan

```go
DeletePlan(
    ctx context.Context,
    planId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetPlanResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanResponse`](../../doc/models/get-plan-response.md)

## Example Usage

```go
ctx := context.Background()
planId := "plan_id8"

apiResponse, err := plansController.DeletePlan(ctx, planId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Plan Metadata

Updates the metadata from a plan

```go
UpdatePlanMetadata(
    ctx context.Context,
    planId string,
    request models.UpdateMetadataRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetPlanResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | The plan id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the plan metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanResponse`](../../doc/models/get-plan-response.md)

## Example Usage

```go
ctx := context.Background()
planId := "plan_id8"

request := models.UpdateMetadataRequest{
    Metadata: map[string]string{
"key0" : "metadata3",
},
}

apiResponse, err := plansController.UpdatePlanMetadata(ctx, planId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Plan Item

Updates a plan item

```go
UpdatePlanItem(
    ctx context.Context,
    planId string,
    planItemId string,
    body models.UpdatePlanItemRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetPlanItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `planItemId` | `string` | Template, Required | Plan item id |
| `body` | [`models.UpdatePlanItemRequest`](../../doc/models/update-plan-item-request.md) | Body, Required | Request for updating the plan item |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanItemResponse`](../../doc/models/get-plan-item-response.md)

## Example Usage

```go
ctx := context.Background()
planId := "plan_id8"
planItemId := "plan_item_id0"

bodyPricingSchemePriceBrackets0 := models.UpdatePriceBracketRequest{
    StartQuantity: 144,
    Price:         174,
}

bodyPricingSchemePriceBrackets := []models.UpdatePriceBracketRequest{bodyPricingSchemePriceBrackets0}
bodyPricingScheme := models.UpdatePricingSchemeRequest{
    SchemeType:    "scheme_type8",
    PriceBrackets: bodyPricingSchemePriceBrackets,
}

body := models.UpdatePlanItemRequest{
    Name:          "name6",
    Description:   "description4",
    Status:        "status2",
    PricingScheme: bodyPricingScheme,
}

apiResponse, err := plansController.UpdatePlanItem(ctx, planId, planItemId, &body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Plan Item

Adds a new item to a plan

```go
CreatePlanItem(
    ctx context.Context,
    planId string,
    request models.CreatePlanItemRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetPlanItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `request` | [`models.CreatePlanItemRequest`](../../doc/models/create-plan-item-request.md) | Body, Required | Request for creating a plan item |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanItemResponse`](../../doc/models/get-plan-item-response.md)

## Example Usage

```go
ctx := context.Background()
planId := "plan_id8"

requestPricingScheme := models.CreatePricingSchemeRequest{
    SchemeType:    "scheme_type8",
}

request := models.CreatePlanItemRequest{
    Name:          "name6",
    Id:            "id6",
    Description:   "description6",
    PricingScheme: requestPricingScheme,
}

apiResponse, err := plansController.CreatePlanItem(ctx, planId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Plan Item

Gets a plan item

```go
GetPlanItem(
    ctx context.Context,
    planId string,
    planItemId string) (
    models.ApiResponse[models.GetPlanItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `planItemId` | `string` | Template, Required | Plan item id |

## Response Type

[`models.GetPlanItemResponse`](../../doc/models/get-plan-item-response.md)

## Example Usage

```go
ctx := context.Background()
planId := "plan_id8"
planItemId := "plan_item_id0"

apiResponse, err := plansController.GetPlanItem(ctx, planId, planItemId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Plan

Creates a new plan

```go
CreatePlan(
    ctx context.Context,
    body models.CreatePlanRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetPlanResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `body` | [`models.CreatePlanRequest`](../../doc/models/create-plan-request.md) | Body, Required | Request for creating a plan |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanResponse`](../../doc/models/get-plan-response.md)

## Example Usage

```go
ctx := context.Background()


bodyItems0PricingScheme := models.CreatePricingSchemeRequest{
    SchemeType:    "scheme_type8",
}

bodyItems0 := models.CreatePlanItemRequest{
    Name:          "name8",
    Id:            "id8",
    Description:   "description2",
    PricingScheme: bodyItems0PricingScheme,
}
bodyItems := []models.CreatePlanItemRequest{bodyItems0}
bodyPricingScheme := models.CreatePricingSchemeRequest{
    SchemeType:    "scheme_type8",
}

body := models.CreatePlanRequest{
    Name:                "name6",
    Description:         "description4",
    StatementDescriptor: "statement_descriptor6",
    Shippable:           false,
    PaymentMethods:      []string{"payment_methods9"},
    Installments:        []int{207},
    Currency:            "currency6",
    Interval:            "interval6",
    IntervalCount:       170,
    BillingDays:         []int{201, 200},
    BillingType:         "billing_type0",
    Metadata:            map[string]string{
"key0" : "metadata7",
"key1" : "metadata8",
},
    Items:               bodyItems,
    PricingScheme:       bodyPricingScheme,
}

apiResponse, err := plansController.CreatePlan(ctx, &body, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Plan Item

Removes an item from a plan

```go
DeletePlanItem(
    ctx context.Context,
    planId string,
    planItemId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetPlanItemResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `planItemId` | `string` | Template, Required | Plan item id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanItemResponse`](../../doc/models/get-plan-item-response.md)

## Example Usage

```go
ctx := context.Background()
planId := "plan_id8"
planItemId := "plan_item_id0"

apiResponse, err := plansController.DeletePlanItem(ctx, planId, planItemId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Plans

Gets all plans

```go
GetPlans(
    ctx context.Context,
    page *int,
    size *int,
    name *string,
    status *string,
    billingType *string,
    createdSince *time.Time,
    createdUntil *time.Time) (
    models.ApiResponse[models.ListPlansResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `name` | `*string` | Query, Optional | Filter for Plan's name |
| `status` | `*string` | Query, Optional | Filter for Plan's status |
| `billingType` | `*string` | Query, Optional | Filter for plan's billing type |
| `createdSince` | `*time.Time` | Query, Optional | Filter for plan's creation date start range |
| `createdUntil` | `*time.Time` | Query, Optional | Filter for plan's creation date end range |

## Response Type

[`models.ListPlansResponse`](../../doc/models/list-plans-response.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := plansController.GetPlans(ctx, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Plan

Updates a plan

```go
UpdatePlan(
    ctx context.Context,
    planId string,
    request models.UpdatePlanRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetPlanResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `planId` | `string` | Template, Required | Plan id |
| `request` | [`models.UpdatePlanRequest`](../../doc/models/update-plan-request.md) | Body, Required | Request for updating a plan |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetPlanResponse`](../../doc/models/get-plan-response.md)

## Example Usage

```go
ctx := context.Background()
planId := "plan_id8"

request := models.UpdatePlanRequest{
    Name:                "name6",
    Description:         "description6",
    Installments:        []int{151, 152},
    StatementDescriptor: "statement_descriptor6",
    Currency:            "currency6",
    Interval:            "interval4",
    IntervalCount:       114,
    PaymentMethods:      []string{"payment_methods1", "payment_methods0", "payment_methods9"},
    BillingType:         "billing_type0",
    Status:              "status8",
    Shippable:           false,
    BillingDays:         []int{115},
    Metadata:            map[string]string{
"key0" : "metadata3",
},
}

apiResponse, err := plansController.UpdatePlan(ctx, planId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

