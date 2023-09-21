package web

import (
	"fmt"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"strconv"
)

type rabbitMQ struct {
	config *types.BrokerConfig
	conn   *amqp.Connection
}

func NewRabbitMQConn(config *types.BrokerConfig) *rabbitMQ {
	return &rabbitMQ{
		config: config,
	}
}

var _ ports.IBrokerConn = (*rabbitMQ)(nil)

func (rb *rabbitMQ) GetConnection() (ports.IChannel, error) {
	if rb.config == nil {
		return nil, fmt.Errorf("config is empty")
	}

	config := rb.config
	dialString := fmt.Sprintf(
		"%s://%s:%s@%s:%s/",
		os.ExpandEnv(config.Driver),
		os.ExpandEnv(config.Username),
		os.ExpandEnv(config.Password),
		os.ExpandEnv(config.Host),
		os.ExpandEnv(strconv.Itoa(int(config.Port))),
	)

	conn, err := amqp.Dial(dialString)
	if err != nil {
		return nil, err
	}
	rb.conn = conn
	log.Println("RabbitMQ connection established")
	return NewChannel(rb.conn), nil
}

func (rb *rabbitMQ) Close() error {
	return rb.conn.Close()
}

func (rb *rabbitMQ) IsClosed() bool {
	return rb.conn.IsClosed()
}
