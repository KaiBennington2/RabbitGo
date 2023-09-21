package types

import (
	"time"
)

type ExchangeT string

const (
	Direct  ExchangeT = `direct`
	Topic   ExchangeT = `topic`
	Fanout  ExchangeT = `fanout`
	Headers ExchangeT = `headers`
)

// --------------------------------------------------------------------------------------------------- //

type DeliveryMsg struct {
	Headers         DynamicMap
	ContentType     string
	ContentEncoding string
	DeliveryMode    uint8
	Priority        uint8
	CorrelationId   string
	ReplyTo         string
	Expiration      string
	MessageId       string
	Timestamp       time.Time
	Type            string
	UserId          string
	AppId           string
	Body            []byte
	ExchangeName    string
	RoutingKey      string
}

// ---------------------------------------------------------------------------------------------------- //

type BindOpts struct {
	NoWait bool
	Args   DynamicMap
}

func NewBindOpts(opts ...BindOpts) BindOpts {
	if len(opts) > 0 {
		return opts[0]
	}
	return BindOpts{Args: DynamicMap{}}
}

// --------------------------------------------------------------------------------------------------- //

type ExchangeOpts struct {
	Name       string
	Kind       ExchangeT
	Durable    bool
	AutoDelete bool
	Internal   bool
	BindOpts
}

func NewExchangeOpts(opts ...ExchangeOpts) ExchangeOpts {
	if len(opts) > 0 {
		return opts[0]
	}
	return ExchangeOpts{Kind: Topic, Durable: true, BindOpts: BindOpts{Args: DynamicMap{}}}
}

// --------------------------------------------------------------------------------------------------- //

type PublishOpts struct {
	Mandatory bool
	Immediate bool
}

func NewPublishOpts(opts ...PublishOpts) PublishOpts {
	if len(opts) > 0 {
		return opts[0]
	}
	return PublishOpts{}
}

// --------------------------------------------------------------------------------------------------- //

type QueueOpts struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	BindOpts
}

func NewQueueOpts(opts ...QueueOpts) QueueOpts {
	if len(opts) > 0 {
		return opts[0]
	}
	return QueueOpts{
		Durable: true,
		BindOpts: BindOpts{
			Args: DynamicMap{
				"x-dead-letter-exchange":    "deadLettering",
				"x-dead-letter-routing-key": "deadLettering.logs.fails",
				"x-max-retry":               10,
			},
		},
	}
}

// --------------------------------------------------------------------------------------------------- //

type SubHandleFunc struct {
	Handler       func([]byte) (any, error)
	HandlerReturn chan<- any
}

type SubTraceHandler struct {
	DoneMsg    chan<- []byte
	DiscardMsg chan<- []byte
	Errs       chan<- error
}

type SubscribeOpts struct {
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	WithDeadL bool
	Qos       int
	BindOpts
}

func NewSubscribeOpts(opts ...SubscribeOpts) SubscribeOpts {
	if len(opts) > 0 {
		return opts[0]
	}
	return SubscribeOpts{
		WithDeadL: true,
		BindOpts: BindOpts{
			Args: DynamicMap{},
		},
	}
}

// --------------------------------------------------------------------------------------------------- //

type ConsumerOpts struct {
	ExchangeName string
	RoutingKey   string
	QueueOpts    QueueOpts
	SubsOpts     SubscribeOpts
}
