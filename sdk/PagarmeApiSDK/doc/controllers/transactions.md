# Transactions

```go
transactionsController := client.TransactionsController()
```

## Class Name

`TransactionsController`


# Get Transaction

```go
GetTransaction(
    ctx context.Context,
    transactionId string) (
    models.ApiResponse[models.GetTransactionResponseInterface],
    error)
```

## Parameters

| Parameter | Type | Tags | Description |
|  --- | --- | --- | --- |
| `transactionId` | `string` | Template, Required | - |

## Response Type

[`models.GetTransactionResponseInterface`](../../doc/models/get-transaction-response.md)

## Example Usage

```go
ctx := context.Background()
transactionId := "transaction_id8"

apiResponse, err := transactionsController.GetTransaction(ctx, transactionId)
if err != nil {
    log.Fatalln(err)
} else {
    // Printing the result and response
    fmt.Println(apiResponse.Data)
    fmt.Println(apiResponse.Response.StatusCode)
}
```

