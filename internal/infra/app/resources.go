package app

import (
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
)

type Resources struct {
	DatabasesMap map[string]ports.IConnectionDB
	BrokersMap   map[string]ports.IBrokerConn
}

func InitAppResources(config *types.Setting) (*Resources, error) {

	// ------------------------------------------------------------------------------------------------------------ //
	// ------------------- The initial persistence connections required for the APP are loaded -------------------- //
	// ------------------------------------------------------------------------------------------------------------ //
	dbConnectionsMap, err := ports.GetConnections(getDatabaseConnection, config.DbConfig...)
	if err != nil {
		return nil, err
	}

	// ------------------------------------------------------------------------------------------------------------ //
	// ----------------- The initial message brokers credentials required for the APP are loaded ------------------ //
	// ------------------------------------------------------------------------------------------------------------ //
	mbConnectionsMap, err := ports.GetMBrokerConnections(getMBrokerConnection, config.BrokerConfig...)
	if err != nil {
		return nil, err
	}

	return &Resources{
		DatabasesMap: dbConnectionsMap,
		BrokersMap:   mbConnectionsMap,
	}, nil
}
