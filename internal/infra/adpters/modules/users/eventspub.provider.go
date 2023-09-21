package users

import (
	"github.com/KaiBennington2/RabbitGo/internal/domain/ports"
	evts "github.com/KaiBennington2/RabbitGo/internal/infra/adpters/eventpub/users"
	sdkPorts "github.com/KaiBennington2/RabbitGo/sdk/ports"
	"sync"
)

type UserEventsPub struct {
	Publisher ports.IUserEventPub
}

var oncePub sync.Once
var pubInstance *UserEventsPub
var LoadEvtPUBS = func(
	mbConn sdkPorts.IBrokerConn,
) *UserEventsPub {
	oncePub.Do(func() {
		inst := &UserEventsPub{
			//Inject each instance its implementation here.
			Publisher: evts.NewUserEventsPub(mbConn),
		}
		pubInstance = inst
	})
	return pubInstance

}
