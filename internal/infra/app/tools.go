package app

import (
	"github.com/KaiBennington2/RabbitGo/internal/infra/adpters/conn"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	"github.com/KaiBennington2/RabbitGo/sdk/web"
	"strings"
)

func getDatabaseConnection(motorName string, dbConfig types.DbConfig) ports.IConnectionDB {
	switch strings.ToLower(motorName) {
	case "mysql":
		return conn.NewMySqlConnection(&dbConfig)
	case "mongodb":
		return conn.NewMongoConnection(&dbConfig)
	default:
		return conn.NewMySqlConnection(&dbConfig)
	}
}

func getMBrokerConnection(brokerName string, mbConfig types.BrokerConfig) ports.IBrokerConn {
	switch strings.ToLower(brokerName) {
	case "rabbitmq":
		return conn.NewRabbitMQ(&mbConfig)
	default:
		return web.NewRabbitMQConn(&mbConfig)
	}
}
