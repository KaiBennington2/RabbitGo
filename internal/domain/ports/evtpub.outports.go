package ports

import (
	"context"
	"github.com/KaiBennington2/RabbitGo/internal/domain/entity"
)

type IUserEventPub interface {
	SavedEvent(ctx context.Context, e *entity.User) error
	DeletedEvent(ctx context.Context, code string, isPermanent bool) error
}
