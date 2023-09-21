package services

import (
	"context"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/internal/domain/ports"
)

var _ ports.IDeleteBusinessUser = (*DeleteService)(nil)

type BusinessDeleteService struct {
	userRepository ports.IUserRepository
}

func NewUserBusinessDeleteService(userRepository ports.IUserRepository) *BusinessDeleteService {
	return &BusinessDeleteService{
		userRepository: userRepository,
	}
}

func (srv *BusinessDeleteService) Execute(ctx context.Context, code *string) (bool, error) {
	if ok, err := srv.existsById(ctx, code); err != nil || !ok {
		if err != nil {
			return false, err
		}
		return false, err
	}

	wasDeleted, err := srv.userRepository.LogicalDelete(ctx, code)
	if err != nil {
		return false, err
	}
	return wasDeleted, nil
}

func (srv *BusinessDeleteService) existsById(ctx context.Context, code *string, withoutLogicDeleted ...bool) (bool, error) {
	if code == nil || *code == "" {
		return false, fmt.Errorf("not exists")
	}

	exist, err := srv.userRepository.ExistsById(ctx, code, withoutLogicDeleted...)
	if err != nil {
		return false, err
	}
	return exist, nil
}
