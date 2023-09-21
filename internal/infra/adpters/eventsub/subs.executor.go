package evtsub

import (
	"context"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
)

var _ ports.ISubscribersExecutor = (*SubExecutor)(nil)

type SubExecutor struct {
	conn    ports.IBrokerConn
	service types.SubHandleFunc
	opts    types.ConsumerOpts
}

func NewSubscriberExecutor(broker ports.IBrokerConn, service types.SubHandleFunc, opts types.ConsumerOpts) *SubExecutor {
	return &SubExecutor{
		conn:    broker,
		service: service,
		opts:    opts,
	}
}

func (exc *SubExecutor) Execute(traceHandler types.SubTraceHandler) {
	ch, err := exc.conn.GetConnection()
	if err != nil {
		traceHandler.Errs <- err
	}

	if err = ch.DeclareQueue(context.Background(), exc.opts.QueueOpts); err != nil {
		traceHandler.Errs <- err
	}

	if err = ch.BindQueue(context.Background(),
		exc.opts.ExchangeName,
		exc.opts.RoutingKey,
		exc.opts.QueueOpts,
	); err != nil {
		traceHandler.Errs <- err
	}

	ch.Subscribe(context.Background(),
		exc.opts.QueueOpts.Name,
		exc.service,
		traceHandler,
		exc.opts.SubsOpts,
	)
}
