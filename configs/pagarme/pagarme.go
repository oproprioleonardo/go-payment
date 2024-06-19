package pagarme

import (
	"net/http"
	"os"
	"pagarmeapisdk"
)

var Client pagarmeapisdk.ClientInterface

func init() {
	config := pagarmeapisdk.CreateConfiguration(
		pagarmeapisdk.WithServiceRefererName("lab-ia-order-consumer"),
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

		pagarmeapisdk.WithBasicAuthUserName(os.Getenv("PAGARME_API_KEY")),
	)
	Client = pagarmeapisdk.NewClient(config)
}
