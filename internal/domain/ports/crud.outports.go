package ports

import (
	"context"
	"github.com/KaiBennington2/RabbitGo/internal/domain/entity"
)

// IUserRepository is an interface for interacting with the user-type entity repository.
type IUserRepository interface {
	Save(ctx context.Context, e *entity.User) (*string, error)
	Delete(ctx context.Context, code *string) (bool, error)
	LogicalDelete(ctx context.Context, code *string) (bool, error)
	ExistsById(ctx context.Context, code *string, withoutLogicDeleted ...bool) (bool, error)
}
