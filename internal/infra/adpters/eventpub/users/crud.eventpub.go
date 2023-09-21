package users

import (
	"context"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/internal/domain/entity"
	outPorts "github.com/KaiBennington2/RabbitGo/internal/domain/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	"github.com/KaiBennington2/RabbitGo/sdk/utils"
	"github.com/google/uuid"
	"sync"
	"time"
)

const (
	defaultExchange     = `user`
	appId               = `ms-go-users`
	contentType         = `application/json`
	contentEncoding     = `UTF-8`
	persistenceMode     = 2
	errProcessingEntity = `An error occurred processing the user-type entity in a byte array: %v.\n `
)

var _ outPorts.IUserEventPub = (*EventsPublisher)(nil)
var dmInstance types.DeliveryMsg
var onceDMsg sync.Once

type EventsPublisher struct {
	conn ports.IBrokerConn
}

var defaultDeliveryMsg = func() types.DeliveryMsg {
	onceDMsg.Do(func() {
		inst := types.DeliveryMsg{
			ContentType:     contentType,
			ContentEncoding: contentEncoding,
			DeliveryMode:    persistenceMode,
			MessageId:       uuid.New().String(),
			Timestamp:       time.Now(),
			AppId:           appId,
		}
		dmInstance = inst
	})
	return dmInstance
}

func NewUserEventsPub(conn ports.IBrokerConn) *EventsPublisher {
	return &EventsPublisher{
		conn: conn,
	}
}

func (evt *EventsPublisher) SavedEvent(ctx context.Context, e *entity.User) error {
	data, err := utils.ToJson(e)
	if err != nil {
		return fmt.Errorf(errProcessingEntity, err)
	}

	ch, err := evt.conn.GetConnection()
	if err != nil {
		return err
	}

	dMsg := defaultDeliveryMsg()
	dMsg.CorrelationId = fmt.Sprintf("req-%s", e.Code)
	dMsg.Type = "saved-user"
	dMsg.Body = data
	dMsg.ExchangeName = defaultExchange
	dMsg.RoutingKey = "user.saved"
	return ch.Publish(ctx, dMsg)

}

func (evt *EventsPublisher) DeletedEvent(ctx context.Context, code string, isPermanent bool) error {
	type payload struct {
		code        string
		isPermanent bool
	}

	data, err := utils.ToJson(&payload{code, isPermanent})
	if err != nil {
		return fmt.Errorf(errProcessingEntity, err)
	}

	ch, err := evt.conn.GetConnection()
	if err != nil {
		return err
	}

	dMsg := defaultDeliveryMsg()
	dMsg.CorrelationId = fmt.Sprintf("req-%s", code)
	dMsg.Type = "deleted-user"
	dMsg.Body = data
	dMsg.ExchangeName = defaultExchange
	dMsg.RoutingKey = "user.deleted"
	return ch.Publish(ctx, dMsg)
}
