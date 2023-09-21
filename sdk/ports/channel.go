package ports

import (
	"context"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
)

type IChannel interface {
	BindExchange(
		ctx context.Context, routingKey string,
		sourceOpts *types.ExchangeOpts, targetOpts *types.ExchangeOpts,
		bindOpts ...types.BindOpts,
	) error

	DeclareExchange(
		ctx context.Context,
		opts types.ExchangeOpts,
	) error

	Publish(
		ctx context.Context,
		message types.DeliveryMsg,
		opts ...types.PublishOpts,
	) error

	DeclareQueue(
		ctx context.Context,
		opts types.QueueOpts,
	) error

	BindQueue(
		ctx context.Context,
		exchangeName string,
		routingKey string,
		opts ...types.QueueOpts,
	) error

	Subscribe(
		ctx context.Context,
		queueName string,
		msgHandler types.SubHandleFunc,
		traceHandler types.SubTraceHandler,
		opts ...types.SubscribeOpts,
	)
}
