module github.com/ia/lab/payment

go 1.20

replace pagarmeapisdk => ./sdk/PagarmeApiSDK

require (
	github.com/apimatic/go-core-runtime v0.0.13
	github.com/rabbitmq/amqp091-go v1.10.0
	golang.org/x/exp v0.0.0-20240119083558-1b970713d09a
	pagarmeapisdk v0.0.0
)
