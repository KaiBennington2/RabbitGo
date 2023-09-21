package services

import (
	"context"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/internal/domain/ports"
)

var _ ports.IDeletePermanentUser = (*DeleteService)(nil)

type DeleteService struct {
	userRepository ports.IUserRepository
}

func NewUserDeleteService(userRepository ports.IUserRepository) *DeleteService {
	return &DeleteService{
		userRepository: userRepository,
	}
}

func (srv *DeleteService) Execute(ctx context.Context, code *string) (bool, error) {
	if ok, err := srv.existsById(ctx, code, true); err != nil || !ok {
		if err != nil {
			return false, err
		}
		return false, err
	}

	wasDeleted, err := srv.userRepository.Delete(ctx, code)
	if err != nil {
		return false, err
	}
	return wasDeleted, nil
}

func (srv *DeleteService) existsById(ctx context.Context, code *string, withoutLogicDeleted ...bool) (bool, error) {
	if code == nil || *code == "" {
		return false, fmt.Errorf("not exists")
	}

	exist, err := srv.userRepository.ExistsById(ctx, code, withoutLogicDeleted...)
	if err != nil {
		return false, err
	}
	return exist, nil
}
