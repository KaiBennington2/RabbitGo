package events

import (
	"context"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/internal/domain/entity"
	"github.com/KaiBennington2/RabbitGo/internal/domain/ports"
)

type RegisteredEvtService struct {
	userRepository ports.IUserRepository
}

func NewUserRegisteredEvtService(userRepository ports.IUserRepository) *RegisteredEvtService {
	return &RegisteredEvtService{
		userRepository: userRepository,
	}
}

func (srv *RegisteredEvtService) Execute(ctx context.Context, e *entity.User) error {
	if e == nil {
		return fmt.Errorf("No user-type entity was found to be processed. ")
	}
	if _, err := srv.userRepository.Save(ctx, e); err != nil {
		return err
	}
	return nil
}
