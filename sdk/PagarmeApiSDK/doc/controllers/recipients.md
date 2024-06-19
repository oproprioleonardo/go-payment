# Recipients

```go
recipientsController := client.RecipientsController()
```

## Class Name

`RecipientsController`

## Methods

* [Update Recipient](../../doc/controllers/recipients.md#update-recipient)
* [Create Anticipation](../../doc/controllers/recipients.md#create-anticipation)
* [Get Anticipation Limits](../../doc/controllers/recipients.md#get-anticipation-limits)
* [Get Recipients](../../doc/controllers/recipients.md#get-recipients)
* [Get Withdraw by Id](../../doc/controllers/recipients.md#get-withdraw-by-id)
* [Update Recipient Default Bank Account](../../doc/controllers/recipients.md#update-recipient-default-bank-account)
* [Update Recipient Metadata](../../doc/controllers/recipients.md#update-recipient-metadata)
* [Get Transfers](../../doc/controllers/recipients.md#get-transfers)
* [Get Transfer](../../doc/controllers/recipients.md#get-transfer)
* [Create Withdraw](../../doc/controllers/recipients.md#create-withdraw)
* [Update Automatic Anticipation Settings](../../doc/controllers/recipients.md#update-automatic-anticipation-settings)
* [Get Anticipation](../../doc/controllers/recipients.md#get-anticipation)
* [Update Recipient Transfer Settings](../../doc/controllers/recipients.md#update-recipient-transfer-settings)
* [Get Anticipations](../../doc/controllers/recipients.md#get-anticipations)
* [Get Recipient](../../doc/controllers/recipients.md#get-recipient)
* [Get Balance](../../doc/controllers/recipients.md#get-balance)
* [Get Withdrawals](../../doc/controllers/recipients.md#get-withdrawals)
* [Create Transfer](../../doc/controllers/recipients.md#create-transfer)
* [Create Recipient](../../doc/controllers/recipients.md#create-recipient)
* [Get Recipient by Code](../../doc/controllers/recipients.md#get-recipient-by-code)
* [Get Default Recipient](../../doc/controllers/recipients.md#get-default-recipient)


# Update Recipient

Updates a recipient

```go
UpdateRecipient(
    ctx context.Context,
    recipientId string,
    request models.UpdateRecipientRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `request` | [`models.UpdateRecipientRequest`](../../doc/models/update-recipient-request.md) | Body, Required | Recipient data |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

request := models.UpdateRecipientRequest{
    Name:        "name6",
    Email:       "email0",
    Description: "description6",
    Type:        "type4",
    Status:      "status8",
    Metadata:    map[string]string{
"key0" : "metadata3",
},
}

apiResponse, err := recipientsController.UpdateRecipient(ctx, recipientId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Anticipation

Creates an anticipation

```go
CreateAnticipation(
    ctx context.Context,
    recipientId string,
    request models.CreateAnticipationRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetAnticipationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `request` | [`models.CreateAnticipationRequest`](../../doc/models/create-anticipation-request.md) | Body, Required | Anticipation data |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetAnticipationResponse`](../../doc/models/get-anticipation-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

requestPaymentDate, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
if err != nil {
    log.Fatalln(err)
}
request := models.CreateAnticipationRequest{
    Amount:      242,
    Timeframe:   "timeframe8",
    PaymentDate: requestPaymentDate,
}

apiResponse, err := recipientsController.CreateAnticipation(ctx, recipientId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Anticipation Limits

Gets the anticipation limits for a recipient

```go
GetAnticipationLimits(
    ctx context.Context,
    recipientId string,
    timeframe string,
    paymentDate time.Time) (
    models.ApiResponse[models.GetAnticipationLimitResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `timeframe` | `string` | Query, Required | Timeframe |
| `paymentDate` | `time.Time` | Query, Required | Anticipation payment date |

## Response Type

[`models.GetAnticipationLimitResponse`](../../doc/models/get-anticipation-limit-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"
timeframe := "timeframe2"
paymentDate, err := time.Parse(time.RFC3339, "2016-03-13T12:52:32.123Z")
if err != nil {
    log.Fatalln(err)
}

apiResponse, err := recipientsController.GetAnticipationLimits(ctx, recipientId, timeframe, paymentDate)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Recipients

Retrieves paginated recipients information

```go
GetRecipients(
    ctx context.Context,
    page *int,
    size *int) (
    models.ApiResponse[models.ListRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |

## Response Type

[`models.ListRecipientResponse`](../../doc/models/list-recipient-response.md)

## Example Usage

```go
ctx := context.Background()

apiResponse, err := recipientsController.GetRecipients(ctx, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Withdraw by Id

```go
GetWithdrawById(
    ctx context.Context,
    recipientId string,
    withdrawalId string) (
    models.ApiResponse[models.GetWithdrawResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | - |
| `withdrawalId` | `string` | Template, Required | - |

## Response Type

[`models.GetWithdrawResponse`](../../doc/models/get-withdraw-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"
withdrawalId := "withdrawal_id2"

apiResponse, err := recipientsController.GetWithdrawById(ctx, recipientId, withdrawalId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Recipient Default Bank Account

Updates the default bank account from a recipient

```go
UpdateRecipientDefaultBankAccount(
    ctx context.Context,
    recipientId string,
    request models.UpdateRecipientBankAccountRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `request` | [`models.UpdateRecipientBankAccountRequest`](../../doc/models/update-recipient-bank-account-request.md) | Body, Required | Bank account data |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

requestBankAccount := models.CreateBankAccountRequest{
    HolderName:        "holder_name0",
    HolderType:        "holder_type6",
    HolderDocument:    "holder_document8",
    Bank:              "bank2",
    BranchNumber:      "branch_number0",
    AccountNumber:     "account_number4",
    AccountCheckDigit: "account_check_digit0",
    Type:              "type6",
    Metadata:          map[string]string{
"key0" : "metadata1",
"key1" : "metadata0",
},
}

request := models.UpdateRecipientBankAccountRequest{
    PaymentMode: "bank_transfer",
    BankAccount: requestBankAccount,
}

apiResponse, err := recipientsController.UpdateRecipientDefaultBankAccount(ctx, recipientId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Recipient Metadata

Updates recipient metadata

```go
UpdateRecipientMetadata(
    ctx context.Context,
    recipientId string,
    request models.UpdateMetadataRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `request` | [`models.UpdateMetadataRequest`](../../doc/models/update-metadata-request.md) | Body, Required | Metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

request := models.UpdateMetadataRequest{
    Metadata: map[string]string{
"key0" : "metadata3",
},
}

apiResponse, err := recipientsController.UpdateRecipientMetadata(ctx, recipientId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Transfers

Gets a paginated list of transfers for the recipient

```go
GetTransfers(
    ctx context.Context,
    recipientId string,
    page *int,
    size *int,
    status *string,
    createdSince *time.Time,
    createdUntil *time.Time) (
    models.ApiResponse[models.ListTransferResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `status` | `*string` | Query, Optional | Filter for transfer status |
| `createdSince` | `*time.Time` | Query, Optional | Filter for start range of transfer creation date |
| `createdUntil` | `*time.Time` | Query, Optional | Filter for end range of transfer creation date |

## Response Type

[`models.ListTransferResponse`](../../doc/models/list-transfer-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

apiResponse, err := recipientsController.GetTransfers(ctx, recipientId, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Transfer

Gets a transfer

```go
GetTransfer(
    ctx context.Context,
    recipientId string,
    transferId string) (
    models.ApiResponse[models.GetTransferResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `transferId` | `string` | Template, Required | Transfer id |

## Response Type

[`models.GetTransferResponse`](../../doc/models/get-transfer-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"
transferId := "transfer_id6"

apiResponse, err := recipientsController.GetTransfer(ctx, recipientId, transferId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Withdraw

```go
CreateWithdraw(
    ctx context.Context,
    recipientId string,
    request models.CreateWithdrawRequest) (
    models.ApiResponse[models.GetWithdrawResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | - |
| `request` | [`models.CreateWithdrawRequest`](../../doc/models/create-withdraw-request.md) | Body, Required | - |

## Response Type

[`models.GetWithdrawResponse`](../../doc/models/get-withdraw-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

request := models.CreateWithdrawRequest{
    Amount:   242,
}

apiResponse, err := recipientsController.CreateWithdraw(ctx, recipientId, &request)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Automatic Anticipation Settings

Updates recipient metadata

```go
UpdateAutomaticAnticipationSettings(
    ctx context.Context,
    recipientId string,
    request models.UpdateAutomaticAnticipationSettingsRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `request` | [`models.UpdateAutomaticAnticipationSettingsRequest`](../../doc/models/update-automatic-anticipation-settings-request.md) | Body, Required | Metadata |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

request := models.UpdateAutomaticAnticipationSettingsRequest{
}

apiResponse, err := recipientsController.UpdateAutomaticAnticipationSettings(ctx, recipientId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Anticipation

Gets an anticipation

```go
GetAnticipation(
    ctx context.Context,
    recipientId string,
    anticipationId string) (
    models.ApiResponse[models.GetAnticipationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `anticipationId` | `string` | Template, Required | Anticipation id |

## Response Type

[`models.GetAnticipationResponse`](../../doc/models/get-anticipation-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"
anticipationId := "anticipation_id0"

apiResponse, err := recipientsController.GetAnticipation(ctx, recipientId, anticipationId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Update Recipient Transfer Settings

```go
UpdateRecipientTransferSettings(
    ctx context.Context,
    recipientId string,
    request models.UpdateTransferSettingsRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient Identificator |
| `request` | [`models.UpdateTransferSettingsRequest`](../../doc/models/update-transfer-settings-request.md) | Body, Required | - |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

request := models.UpdateTransferSettingsRequest{
    TransferEnabled:  "transfer_enabled2",
    TransferInterval: "transfer_interval6",
    TransferDay:      "transfer_day6",
}

apiResponse, err := recipientsController.UpdateRecipientTransferSettings(ctx, recipientId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Anticipations

Retrieves a paginated list of anticipations from a recipient

```go
GetAnticipations(
    ctx context.Context,
    recipientId string,
    page *int,
    size *int,
    status *string,
    timeframe *string,
    paymentDateSince *time.Time,
    paymentDateUntil *time.Time,
    createdSince *time.Time,
    createdUntil *time.Time) (
    models.ApiResponse[models.ListAnticipationResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |
| `page` | `*int` | Query, Optional | Page number |
| `size` | `*int` | Query, Optional | Page size |
| `status` | `*string` | Query, Optional | Filter for anticipation status |
| `timeframe` | `*string` | Query, Optional | Filter for anticipation timeframe |
| `paymentDateSince` | `*time.Time` | Query, Optional | Filter for start range for anticipation payment date |
| `paymentDateUntil` | `*time.Time` | Query, Optional | Filter for end range for anticipation payment date |
| `createdSince` | `*time.Time` | Query, Optional | Filter for start range for anticipation creation date |
| `createdUntil` | `*time.Time` | Query, Optional | Filter for end range for anticipation creation date |

## Response Type

[`models.ListAnticipationResponse`](../../doc/models/list-anticipation-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

apiResponse, err := recipientsController.GetAnticipations(ctx, recipientId, nil, nil, nil, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Recipient

Retrieves recipient information

```go
GetRecipient(
    ctx context.Context,
    recipientId string) (
    models.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipiend id |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

apiResponse, err := recipientsController.GetRecipient(ctx, recipientId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Balance

Get balance information for a recipient

```go
GetBalance(
    ctx context.Context,
    recipientId string) (
    models.ApiResponse[models.GetBalanceResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient id |

## Response Type

[`models.GetBalanceResponse`](../../doc/models/get-balance-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

apiResponse, err := recipientsController.GetBalance(ctx, recipientId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Withdrawals

Gets a paginated list of transfers for the recipient

```go
GetWithdrawals(
    ctx context.Context,
    recipientId string,
    page *int,
    size *int,
    status *string,
    createdSince *time.Time,
    createdUntil *time.Time) (
    models.ApiResponse[models.ListWithdrawals],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | - |
| `page` | `*int` | Query, Optional | - |
| `size` | `*int` | Query, Optional | - |
| `status` | `*string` | Query, Optional | - |
| `createdSince` | `*time.Time` | Query, Optional | - |
| `createdUntil` | `*time.Time` | Query, Optional | - |

## Response Type

[`models.ListWithdrawals`](../../doc/models/list-withdrawals.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

apiResponse, err := recipientsController.GetWithdrawals(ctx, recipientId, nil, nil, nil, nil, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Transfer

Creates a transfer for a recipient

```go
CreateTransfer(
    ctx context.Context,
    recipientId string,
    request models.CreateTransferRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetTransferResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `recipientId` | `string` | Template, Required | Recipient Id |
| `request` | [`models.CreateTransferRequest`](../../doc/models/create-transfer-request.md) | Body, Required | Transfer data |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetTransferResponse`](../../doc/models/get-transfer-response.md)

## Example Usage

```go
ctx := context.Background()
recipientId := "recipient_id0"

request := models.CreateTransferRequest{
    Amount:   242,
    Metadata: map[string]string{
"key0" : "metadata3",
},
}

apiResponse, err := recipientsController.CreateTransfer(ctx, recipientId, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Create Recipient

Creates a new recipient

```go
CreateRecipient(
    ctx context.Context,
    request models.CreateRecipientRequest,
    idempotencyKey *string) (
    models.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `request` | [`models.CreateRecipientRequest`](../../doc/models/create-recipient-request.md) | Body, Required | Recipient data |
| `idempotencyKey` | `*string` | Header, Optional | - |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
ctx := context.Background()

requestDefaultBankAccount := models.CreateBankAccountRequest{
    HolderName:        "holder_name4",
    HolderType:        "holder_type0",
    HolderDocument:    "holder_document2",
    Bank:              "bank6",
    BranchNumber:      "branch_number4",
    AccountNumber:     "account_number8",
    AccountCheckDigit: "account_check_digit4",
    Type:              "type2",
    Metadata:          map[string]string{
"key0" : "metadata5",
"key1" : "metadata4",
"key2" : "metadata3",
},
}

request := models.CreateRecipientRequest{
    Name:               "name6",
    Email:              "email0",
    Description:        "description6",
    Document:           "document0",
    Type:               "type4",
    Metadata:           map[string]string{
"key0" : "metadata3",
},
    Code:               "code4",
    PaymentMode:        "bank_transfer",
    DefaultBankAccount: requestDefaultBankAccount,
}

apiResponse, err := recipientsController.CreateRecipient(ctx, &request, nil)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Recipient by Code

Retrieves recipient information

```go
GetRecipientByCode(
    ctx context.Context,
    code string) (
    models.ApiResponse[models.GetRecipientResponse],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `code` | `string` | Template, Required | Recipient code |

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
ctx := context.Background()
code := "code8"

apiResponse, err := recipientsController.GetRecipientByCode(ctx, code)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```


# Get Default Recipient

```go
GetDefaultRecipient(
    ctx context.Context) (
    models.ApiResponse[models.GetRecipientResponse],
    error)
```

## Response Type

[`models.GetRecipientResponse`](../../doc/models/get-recipient-response.md)

## Example Usage

```go
ctx := context.Background()
apiResponse, err := recipientsController.GetDefaultRecipient(ctx)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

