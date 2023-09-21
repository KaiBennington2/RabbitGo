package users

import (
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"sync"
)

const (
	RabbitMBConn string = "mb_rabbitmq_conn"
)

type UserModule struct {
	Repos    UserRepos
	Subs     UserEventSubscribers
	Services UserServices
	UseCases UserUseCases
}

var onceMod sync.Once
var modInstance *UserModule
var InitMODULE = func(
	conn map[string]ports.IConnectionDB,
	brokers map[string]ports.IBrokerConn,
) *UserModule {
	onceMod.Do(func() {

		pubs := LoadEvtPUBS(brokers[RabbitMBConn])
		repos := LoadREPOSITORIES(conn)
		services := LoadSERVICES(*repos, *pubs)
		useCases := LoadUSECASES(*services)
		subs := LoadEvtSUBS(brokers[RabbitMBConn], useCases.Events)

		inst := &UserModule{
			Repos:    *repos,
			Subs:     *subs,
			Services: *services,
			UseCases: *useCases,
		}
		modInstance = inst
	})
	return modInstance
}
