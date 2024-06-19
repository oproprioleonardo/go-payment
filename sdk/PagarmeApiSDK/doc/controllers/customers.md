# Customers

```go
customersController := client.CustomersController()
```

## Class Name

`CustomersController`

## Methods

* [Update Card](../../doc/controllers/customers.md#update-card)
* [Update Address](../../doc/controllers/customers.md#update-address)
* [Delete Access Token](../../doc/controllers/customers.md#delete-access-token)
* [Create Customer](../../doc/controllers/customers.md#create-customer)
* [Create Address](../../doc/controllers/customers.md#create-address)
* [Delete Access Tokens](../../doc/controllers/customers.md#delete-access-tokens)
* [Get Address](../../doc/controllers/customers.md#get-address)
* [Delete Address](../../doc/controllers/customers.md#delete-address)
* [Create Card](../../doc/controllers/customers.md#create-card)
* [Get Customers](../../doc/controllers/customers.md#get-customers)
* [Update Customer](../../doc/controllers/customers.md#update-customer)
* [Create Access Token](../../doc/controllers/customers.md#create-access-token)
* [Get Access Tokens](../../doc/controllers/customers.md#get-access-tokens)
* [Get Cards](../../doc/controllers/customers.md#get-cards)
* [Renew Card](../../doc/controllers/customers.md#renew-card)
* [Get Access Token](../../doc/controllers/customers.md#get-access-token)
* [Update Customer Metadata](../../doc/controllers/customers.md#update-customer-metadata)
* [Delete Card](../../doc/controllers/customers.md#delete-card)
* [Get Addresses](../../doc/controllers/customers.md#get-addresses)
* [Get Customer](../../doc/controllers/customers.md#get-customer)
* [Get Card](../../doc/controllers/customers.md#get-card)


# Update Card

Updates a card

```go
UpdateCard(
    ctx context.Context,
    customerId string,
    cardId string,
    request models.UpdateCardRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetCardResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |
| `cardId` | `string` | Template, Required | Card id |
| `request` | [`models.UpdateCardRequest`](../../doc/models/update-card-request.md) | Body, Required | Request for updating a card |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetCardResponse`](../../doc/models/get-card-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
cardId := "card_id4"

requestBillingAddress := models.CreateAddressRequest{
    Street:       "street8",
    Number:       "number4",
    ZipCode:      "zip_code2",
    Neighborhood: "neighborhood4",
    City:         "city2",
    State:        "state6",
    Country:      "country2",
    Complement:   "complement6",
    Line1:        "line_18",
    Line2:        "line_26",
}

request := models.UpdateCardRequest{
    HolderName:       "holder_name2",
    ExpMonth:         10,
    ExpYear:          30,
    Metadata:         map[string]string{
"key0" : "metadata3",
},
    Label:            "label6",
    BillingAddress:   requestBillingAddress,
}

apiResponse, err := customersController.UpdateCard(ctx, customerId, cardId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Address

Updates an address

```go
UpdateAddress(
    ctx context.Context,
    customerId string,
    addressId string,
    request models.UpdateAddressRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetAddressResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |
| `addressId` | `string` | Template, Required | Address Id |
| `request` | [`models.UpdateAddressRequest`](../../doc/models/update-address-request.md) | Body, Required | Request for updating an address |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetAddressResponse`](../../doc/models/get-address-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
addressId := "address_id0"

request := models.UpdateAddressRequest{
    Number:     "number4",
    Complement: "complement2",
    Metadata:   map[string]string{
"key0" : "metadata3",
},
    Line2:      "line_24",
}

apiResponse, err := customersController.UpdateAddress(ctx, customerId, addressId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Access Token

Delete a customer's access token

```go
DeleteAccessToken(
    ctx context.Context,
    customerId string,
    tokenId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetAccessTokenResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |
| `tokenId` | `string` | Template, Required | Token Id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetAccessTokenResponse`](../../doc/models/get-access-token-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
tokenId := "token_id6"

apiResponse, err := customersController.DeleteAccessToken(ctx, customerId, tokenId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Customer

Creates a new customer

```go
CreateCustomer(
    ctx context.Context,
    request models.CreateCustomerRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetCustomerResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `request` | [`models.CreateCustomerRequest`](../../doc/models/create-customer-request.md) | Body, Required | Request for creating a customer |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetCustomerResponse`](../../doc/models/get-customer-response.md)

## Example Usage

```go
ctx := context.Background()
requestAddress := models.CreateAddressRequest{
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

requestPhones := models.CreatePhonesRequest{
}

request := models.CreateCustomerRequest{
    Name:         "{\n    \"name\": \"Tony Stark\"\n}",
    Email:        "email0",
    Document:     "document0",
    Type:         "type4",
    Metadata:     map[string]string{
"key0" : "metadata3",
},
    Code:         "code4",
    Address:      requestAddress,
    Phones:       requestPhones,
}


apiResponse, err := customersController.CreateCustomer(ctx, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Address

Creates a new address for a customer

```go
CreateAddress(
    ctx context.Context,
    customerId string,
    request models.CreateAddressRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetAddressResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |
| `request` | [`models.CreateAddressRequest`](../../doc/models/create-address-request.md) | Body, Required | Request for creating an address |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetAddressResponse`](../../doc/models/get-address-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
request := models.CreateAddressRequest{
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


apiResponse, err := customersController.CreateAddress(ctx, customerId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Access Tokens

Delete a Customer's access tokens

```go
DeleteAccessTokens(
    ctx context.Context,
    customerId string) (
    models.ApiResponse[models.ListAccessTokensResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |

## Response Type

[`models.ListAccessTokensResponse`](../../doc/models/list-access-tokens-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"

apiResponse, err := customersController.DeleteAccessTokens(ctx, customerId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Address

Get a customer's address

```go
GetAddress(
    ctx context.Context,
    customerId string,
    addressId string) (
    models.ApiResponse[models.GetAddressResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer id |
| `addressId` | `string` | Template, Required | Address Id |

## Response Type

[`models.GetAddressResponse`](../../doc/models/get-address-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
addressId := "address_id0"

apiResponse, err := customersController.GetAddress(ctx, customerId, addressId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Address

Delete a Customer's address

```go
DeleteAddress(
    ctx context.Context,
    customerId string,
    addressId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetAddressResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |
| `addressId` | `string` | Template, Required | Address Id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetAddressResponse`](../../doc/models/get-address-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
addressId := "address_id0"

apiResponse, err := customersController.DeleteAddress(ctx, customerId, addressId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Card

Creates a new card for a customer

```go
CreateCard(
    ctx context.Context,
    customerId string,
    request models.CreateCardRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetCardResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer id |
| `request` | [`models.CreateCardRequest`](../../doc/models/create-card-request.md) | Body, Required | Request for creating a card |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetCardResponse`](../../doc/models/get-card-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
request := models.CreateCardRequest{
    Type:             models.ToPointer("credit"),
}


apiResponse, err := customersController.CreateCard(ctx, customerId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Customers

Get all Customers

```go
GetCustomers(
    ctx context.Context,
    name *string,
    document *string,
    page *int,
    size *int,
    email *string,
    code *string) (
    models.ApiResponse[models.ListCustomersResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `name` | `*string` | Query, Optional | Name of the Customer |
| `document` | `*string` | Query, Optional | Document of the Customer |
| `page` | `*int` | Query, Optional | Current page the the search<br>**Default**: `1` |
| `size` | `*int` | Query, Optional | Quantity pages of the search<br>**Default**: `10` |
| `email` | `*string` | Query, Optional | Customer's email |
| `code` | `*string` | Query, Optional | Customer's code |

## Response Type

[`models.ListCustomersResponse`](../../doc/models/list-customers-response.md)

## Example Usage

```go
ctx := context.Background()
page := 1
size := 10

apiResponse, err := customersController.GetCustomers(ctx, nil, nil, &page, &size, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Customer

Updates a customer

```go
UpdateCustomer(
    ctx context.Context,
    customerId string,
    request models.UpdateCustomerRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetCustomerResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer id |
| `request` | [`models.UpdateCustomerRequest`](../../doc/models/update-customer-request.md) | Body, Required | Request for updating a customer |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetCustomerResponse`](../../doc/models/get-customer-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"

request := models.UpdateCustomerRequest{
}

apiResponse, err := customersController.UpdateCustomer(ctx, customerId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Access Token

Creates a access token for a customer

```go
CreateAccessToken(
    ctx context.Context,
    customerId string,
    request models.CreateAccessTokenRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetAccessTokenResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |
| `request` | [`models.CreateAccessTokenRequest`](../../doc/models/create-access-token-request.md) | Body, Required | Request for creating a access token |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetAccessTokenResponse`](../../doc/models/get-access-token-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"

request := models.CreateAccessTokenRequest{
}

apiResponse, err := customersController.CreateAccessToken(ctx, customerId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Access Tokens

Get all access tokens from a customer

```go
GetAccessTokens(
    ctx context.Context,
    customerId string,
    page *int,
    size *int) (
    models.ApiResponse[models.ListAccessTokensResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |

## Response Type

[`models.ListAccessTokensResponse`](../../doc/models/list-access-tokens-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"

apiResponse, err := customersController.GetAccessTokens(ctx, customerId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Cards

Get all cards from a customer

```go
GetCards(
    ctx context.Context,
    customerId string,
    page *int,
    size *int) (
    models.ApiResponse[models.ListCardsResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |

## Response Type

[`models.ListCardsResponse`](../../doc/models/list-cards-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"

apiResponse, err := customersController.GetCards(ctx, customerId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Renew Card

Renew a card

```go
RenewCard(
    ctx context.Context,
    customerId string,
    cardId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetCardResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer id |
| `cardId` | `string` | Template, Required | Card Id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetCardResponse`](../../doc/models/get-card-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
cardId := "card_id4"

apiResponse, err := customersController.RenewCard(ctx, customerId, cardId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Access Token

Get a Customer's access token

```go
GetAccessToken(
    ctx context.Context,
    customerId string,
    tokenId string) (
    models.ApiResponse[models.GetAccessTokenResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |
| `tokenId` | `string` | Template, Required | Token Id |

## Response Type

[`models.GetAccessTokenResponse`](../../doc/models/get-access-token-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
tokenId := "token_id6"

apiResponse, err := customersController.GetAccessToken(ctx, customerId, tokenId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Customer Metadata

Updates the metadata a customer

```go
UpdateCustomerMetadata(
    ctx context.Context,
    customerId string,
    request models.UpdateMetadataRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetCustomerResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | The customer id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Request for updating the customer metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetCustomerResponse`](../../doc/models/get-customer-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"

request := models.UpdateMetadataRequest{
    Metadata: map[string]string{
"key0" : "metadata3",
},
}

apiResponse, err := customersController.UpdateCustomerMetadata(ctx, customerId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Delete Card

Delete a customer's card

```go
DeleteCard(
    ctx context.Context,
    customerId string,
    cardId string,
    idempotencyKey *string) (
    models.ApiResponse[models.GetCardResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |
| `cardId` | `string` | Template, Required | Card Id |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetCardResponse`](../../doc/models/get-card-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
cardId := "card_id4"

apiResponse, err := customersController.DeleteCard(ctx, customerId, cardId, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Addresses

Gets all adressess from a customer

```go
GetAddresses(
    ctx context.Context,
    customerId string,
    page *int,
    size *int) (
    models.ApiResponse[models.ListAddressesResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |

## Response Type

[`models.ListAddressesResponse`](../../doc/models/list-addresses-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"

apiResponse, err := customersController.GetAddresses(ctx, customerId, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Customer

Get a customer

```go
GetCustomer(
    ctx context.Context,
    customerId string) (
    models.ApiResponse[models.GetCustomerResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer Id |

## Response Type

[`models.GetCustomerResponse`](../../doc/models/get-customer-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"

apiResponse, err := customersController.GetCustomer(ctx, customerId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Card

Get a customer's card

```go
GetCard(
    ctx context.Context,
    customerId string,
    cardId string) (
    models.ApiResponse[models.GetCardResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `customerId` | `string` | Template, Required | Customer id |
| `cardId` | `string` | Template, Required | Card id |

## Response Type

[`models.GetCardResponse`](../../doc/models/get-card-response.md)

## Example Usage

```go
ctx := context.Background()
customerId := "customer_id8"
cardId := "card_id4"

apiResponse, err := customersController.GetCard(ctx, customerId, cardId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

