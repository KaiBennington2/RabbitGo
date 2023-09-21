package users

import (
	evts "github.com/KaiBennington2/RabbitGo/internal/infra/adpters/eventsub/users"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"sync"
)

type UserEventSubscribers struct {
	SavedEvent ports.ISubscribersExecutor
}

var onceSub sync.Once
var subInstance *UserEventSubscribers
var LoadEvtSUBS = func(
	mbConn ports.IBrokerConn,
	events UserEvents,
) *UserEventSubscribers {
	onceSub.Do(func() {
		Subscribers := evts.NewUserEventsSub(mbConn)
		inst := &UserEventSubscribers{
			SavedEvent: Subscribers.SavedEvent(events.Crud),
		}
		subInstance = inst
	})
	return subInstance

}
