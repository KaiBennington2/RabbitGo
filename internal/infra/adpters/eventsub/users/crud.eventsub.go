package users

import (
	aPorts "github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/ports"
	evtsub "github.com/KaiBennington2/RabbitGo/internal/infra/adpters/eventsub"
	evtHandler "github.com/KaiBennington2/RabbitGo/internal/infra/handlers/event/users"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
)

type EventsSubscriber struct {
	conn ports.IBrokerConn
}

func NewUserEventsSub(conn ports.IBrokerConn) *EventsSubscriber {
	return &EventsSubscriber{
		conn: conn,
	}
}

func (evt *EventsSubscriber) SavedEvent(events aPorts.IEvents) ports.ISubscribersExecutor {
	return evtsub.NewSubscriberExecutor(
		evt.conn,
		evtHandler.Saved(events),
		types.ConsumerOpts{
			ExchangeName: "user",
			RoutingKey:   "user.saved",
			SubsOpts: types.NewSubscribeOpts(
				types.SubscribeOpts{
					Consumer:  "ms-go-user_saved",
					WithDeadL: true,
					BindOpts:  types.BindOpts{Args: types.DynamicMap{}},
				},
			),
			QueueOpts: types.NewQueueOpts(
				types.QueueOpts{
					Name:    "user-saved",
					Durable: true,
					BindOpts: types.BindOpts{
						Args: types.DynamicMap{
							"x-dead-letter-exchange":    "deadLettering",
							"x-dead-letter-routing-key": "deadLettering.logs.fails",
							"x-max-retry":               10,
						},
					},
				},
			),
		})
}
