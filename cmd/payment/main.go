package main

import (
	"context"
	"encoding/json"
	"github.com/imersao17/payment/internal/entity"
	"github.com/imersao17/payment/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

// {"order_id": "123", "card_hash": "hash", "total": 100}
func main() {
	ctx := context.Background()
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			panic(err)
		}
	}(ch)

	msgs := make(chan amqp.Delivery)
	go func() {
		err := rabbitmq.Consume(ch, msgs, "orders")
		if err != nil {
			panic(err)
		}
	}()

	for msg := range msgs {
		var orderRequest entity.OrderRequest

		if err := json.Unmarshal(msg.Body, &orderRequest); err != nil {
			slog.Error(err.Error())
			break
		}

		response, err := orderRequest.Process()
		if err != nil {
			slog.Error(err.Error())
			break
		}

		responseJSON, err := json.Marshal(response)
		if err != nil {
			slog.Error(err.Error())
			break
		}

		if err := rabbitmq.Publish(ctx, ch, string(responseJSON), "amq.direct"); err != nil {
			slog.Error(err.Error())
			break
		}

		if err := msg.Ack(false); err != nil {
			return
		}

		slog.Info("Order processed")
	}

}
