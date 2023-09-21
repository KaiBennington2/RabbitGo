package ports

import (
	"context"
	"github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/dto"
)

type ICommands interface {
	Create(ctx context.Context, bodyPrev dto.CreateUserDTO) (*string, error)
	DeletedPermanent(ctx context.Context, codePrev string) (*string, error)
	DeletedBusiness(ctx context.Context, codePrev string) (*string, error)
}
