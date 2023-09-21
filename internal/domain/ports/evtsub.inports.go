package ports

import (
	"context"
	"github.com/KaiBennington2/RabbitGo/internal/domain/entity"
)

type ISavedUserEvent interface {
	Execute(ctx context.Context, e *entity.User) error
}
