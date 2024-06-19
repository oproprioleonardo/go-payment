
# Getting Started with PagarmeApiSDK

## Introduction

Pagarme API

### Requirements

The SDK requires **Go version 1.18 or above**.

## Building

### Install Dependencies

Resolve all the SDK dependencies, using the `go get` command.

## Installation

The following section explains how to use the pagarmeapisdk library in a new project.

### 1. Add SDK as a Dependency to the Application

- Add the following lines to your application's `go.mod` file:

```go
replace pagarmeapisdk => ".\\PagarmeApiSDK" // local path to the SDK

require pagarmeapisdk v0.0.0
```

- Resolve the dependencies in the updated `go.mod` file, using the `go get` command.

## Initialize the API Client

**_Note:_** Documentation for the client can be found [here.](doc/client.md)

The following parameters are configurable for the API Client:

| Parameter | Type | Description |
|  --- | --- | --- |
| `serviceRefererName` | `string` |  |
| `httpConfiguration` | [`HttpConfiguration`](doc/http-configuration.md) | Configurable http client options like timeout and retries. |
| `basicAuthUserName` | `string` | The username to use with basic authentication |
| `basicAuthPassword` | `string` | The password to use with basic authentication |

The API client can be initialized as follows:

```go
config := pagarmeapisdk.CreateConfiguration(
    pagarmeapisdk.WithServiceRefererName("ServiceRefererName"),
    pagarmeapisdk.WithHttpConfiguration(
        pagarmeapisdk.CreateHttpConfiguration(
            pagarmeapisdk.WithTimeout(0),
            pagarmeapisdk.WithTransport(http.DefaultTransport),
            pagarmeapisdk.WithRetryConfiguration(
                pagarmeapisdk.CreateRetryConfiguration(
                    pagarmeapisdk.WithMaxRetryAttempts(0),
                    pagarmeapisdk.WithRetryOnTimeout(true),
                    pagarmeapisdk.WithRetryInterval(1),
                    pagarmeapisdk.WithMaximumRetryWaitTime(0),
                    pagarmeapisdk.WithBackoffFactor(2),
                    pagarmeapisdk.WithHttpStatusCodesToRetry([]int64{408, 413, 429, 500, 502, 503, 504, 521, 522, 524}),
                    pagarmeapisdk.WithHttpMethodsToRetry([]string{"GET", "PUT"}),
                ),
            ),
        ),
    ),
    pagarmeapisdk.WithBasicAuthUserName("BasicAuthUserName"),
    pagarmeapisdk.WithBasicAuthPassword("BasicAuthPassword"),
)
client := pagarmeapisdk.NewClient(config)
```

## Authorization

This API uses `Basic Authentication`.

## API Errors

Here is the list of errors that the API might throw.

| HTTP Status Code | Error Description | Exception Class |
|  --- | --- | --- |
| 400 | Invalid request | [`ErrorException`](doc/models/error-exception.md) |
| 401 | Invalid API key | [`ErrorException`](doc/models/error-exception.md) |
| 404 | An informed resource was not found | [`ErrorException`](doc/models/error-exception.md) |
| 412 | Business validation error | [`ErrorException`](doc/models/error-exception.md) |
| 422 | Contract validation error | [`ErrorException`](doc/models/error-exception.md) |
| 500 | Internal server error | [`ErrorException`](doc/models/error-exception.md) |

## List of APIs

* [Subscriptions](doc/controllers/subscriptions.md)
* [Orders](doc/controllers/orders.md)
* [Plans](doc/controllers/plans.md)
* [Invoices](doc/controllers/invoices.md)
* [Customers](doc/controllers/customers.md)
* [Charges](doc/controllers/charges.md)
* [Recipients](doc/controllers/recipients.md)
* [Tokens](doc/controllers/tokens.md)
* [Transactions](doc/controllers/transactions.md)
* [Transfers](doc/controllers/transfers.md)
* [Payables](doc/controllers/payables.md)
* [Balance Operations](doc/controllers/balance-operations.md)

## Classes Documentation

* [HttpConfiguration](doc/http-configuration.md)
* [RetryConfiguration](doc/retry-configuration.md)

