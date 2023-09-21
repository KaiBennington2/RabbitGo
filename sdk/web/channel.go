package web

import (
	"context"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Channel struct {
	conn *amqp.Connection
}

var _ ports.IChannel = (*Channel)(nil)

func NewChannel(conn *amqp.Connection) *Channel {
	return &Channel{conn}
}

func (ch *Channel) BindExchange(_ context.Context,
	routingKey string,
	srcOpts *types.ExchangeOpts,
	trgOpts *types.ExchangeOpts,
	bindOpts ...types.BindOpts,
) error {
	c, err := ch.conn.Channel()
	defer c.Close()
	if err != nil {
		return err
	}

	err = excDeclare(c, *srcOpts)
	if err != nil {
		return err
	}

	err = excDeclare(c, *trgOpts)
	if err != nil {
		return err
	}

	// Make the link between exchanges in messaging broker
	bOpts := types.NewBindOpts(bindOpts...)
	err = c.ExchangeBind(
		trgOpts.Name,
		routingKey,
		srcOpts.Name,
		bOpts.NoWait,
		amqp.Table(bOpts.Args),
	)
	if err != nil {
		return err
	}
	return nil
}

func (ch *Channel) DeclareExchange(_ context.Context, opts types.ExchangeOpts) error {
	c, err := ch.conn.Channel()
	defer c.Close()
	if err != nil {
		return err
	}

	err = excDeclare(c, opts)
	if err != nil {
		return err
	}
	return nil
}

func (ch *Channel) Publish(ctx context.Context, message types.DeliveryMsg, opts ...types.PublishOpts) error {
	c, err := ch.conn.Channel()
	defer c.Close()
	if err != nil {
		return err
	}

	opt := types.NewPublishOpts(opts...)
	cx, cancel := context.WithCancel(ctx)
	defer cancel()

	return c.PublishWithContext(
		cx,
		message.ExchangeName,
		message.RoutingKey,
		opt.Mandatory,
		opt.Immediate,
		amqp.Publishing{
			Headers:         amqp.Table(message.Headers),
			ContentType:     message.ContentType,
			ContentEncoding: message.ContentEncoding,
			DeliveryMode:    message.DeliveryMode,
			Priority:        message.Priority,
			CorrelationId:   message.CorrelationId,
			ReplyTo:         message.ReplyTo,
			Expiration:      message.Expiration,
			MessageId:       message.MessageId,
			Timestamp:       message.Timestamp,
			Type:            message.Type,
			UserId:          message.UserId,
			AppId:           message.AppId,
			Body:            message.Body,
		},
	)
}

func (ch *Channel) DeclareQueue(_ context.Context, opts types.QueueOpts) error {
	c, err := ch.conn.Channel()
	defer c.Close()
	if err != nil {
		return err
	}

	queueOpts := types.NewQueueOpts(opts)
	if _, err = c.QueueDeclare(
		queueOpts.Name,
		queueOpts.Durable,
		queueOpts.AutoDelete,
		queueOpts.Exclusive,
		queueOpts.NoWait,
		amqp.Table(queueOpts.Args),
	); err != nil {
		return err
	}
	return nil
}

func (ch *Channel) BindQueue(ctx context.Context,
	exchangeName string,
	routingKey string,
	opts ...types.QueueOpts,
) error {
	c, err := ch.conn.Channel()
	defer c.Close()
	if err != nil {
		return err
	}

	queueOpts := types.NewQueueOpts(opts...)
	err = ch.DeclareQueue(ctx, queueOpts)
	if err != nil {
		return err
	}

	if err = c.QueueBind(
		queueOpts.Name,
		routingKey,
		exchangeName,
		queueOpts.NoWait,
		amqp.Table(queueOpts.Args),
	); err != nil {
		return err
	}
	return nil
}

func (ch *Channel) Subscribe(
	ctx context.Context,
	queueName string,
	msgHdl types.SubHandleFunc,
	traceHdl types.SubTraceHandler,
	opts ...types.SubscribeOpts,
) {
	c, err := ch.conn.Channel()
	defer c.Close()
	if err != nil {
		traceHdl.Errs <- err
	}

	opt := types.NewSubscribeOpts(opts...)

	if opt.Qos >= 0 {
		prefetchCount := opt.Qos * 4
		err = c.Qos(prefetchCount, 0, false)
		if err != nil {
			traceHdl.Errs <- err
		}
	}

	delivery, err := c.Consume(
		queueName,
		opt.Consumer,
		opt.AutoAck,
		opt.Exclusive,
		opt.NoLocal,
		opt.NoWait,
		amqp.Table(opt.Args),
	)
	if err != nil {
		traceHdl.Errs <- err
	}

	for msg := range delivery {

		rsp, err := msgHdl.Handler(msg.Body)
		if err == nil {
			msgHdl.HandlerReturn <- rsp
			traceHdl.DoneMsg <- msg.Body
			_ = msg.Ack(false)
			continue
		}

		traceHdl.Errs <- err
		if opt.WithDeadL {
			traceHdl.DiscardMsg <- msg.Body
			_ = msg.Reject(false)
			continue
		}

		traceHdl.DiscardMsg <- msg.Body
		_ = msg.Nack(false, true)

	} // endFor

}

func excDeclare(c *amqp.Channel, opts ...types.ExchangeOpts) error {
	o := types.NewExchangeOpts(opts...)
	if err := c.ExchangeDeclare(
		o.Name,
		string(o.Kind),
		o.Durable,
		o.AutoDelete,
		o.Internal,
		o.NoWait,
		amqp.Table(o.Args),
	); err != nil {
		return err
	}
	return nil
}
