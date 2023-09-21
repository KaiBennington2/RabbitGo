package ports

import (
	"context"
	"github.com/KaiBennington2/RabbitGo/internal/domain/entity"
)

type ICreateUser interface {
	Execute(ctx context.Context, e *entity.User) (*string, error)
}

type IDeleteUser interface {
	Execute(ctx context.Context, id *string) (bool, error)
}

type IDeletePermanentUser interface {
	IDeleteUser
}

type IDeleteBusinessUser interface {
	IDeleteUser
}
