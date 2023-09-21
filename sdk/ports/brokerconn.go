package ports

import (
	"errors"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
)

type IBrokerConn interface {
	GetConnection() (IChannel, error)
	Close() error
	IsClosed() bool
}

type ISubscribersExecutor interface {
	Execute(traceHandler types.SubTraceHandler)
}

func GetMBrokerConnections(
	fn func(brokerName string, mbConfig types.BrokerConfig) IBrokerConn,
	mbSettings ...types.BrokerConfig,
) (map[string]IBrokerConn, error) {
	if mbSettings == nil {
		return nil, errors.New(`No messaging broker configurations were found to initiate connections. `)
	}

	resp := make(map[string]IBrokerConn)
	for _, config := range mbSettings {
		tempConn := fn(config.BrokerName, config)
		_, err := tempConn.GetConnection()
		if err != nil {
			return nil, err
		}
		err = tempConn.Close()
		if err != nil {
			return nil, err
		}
		resp[fmt.Sprintf("%s_conn", config.ConnectionName)] = tempConn
	}
	return resp, nil
}
