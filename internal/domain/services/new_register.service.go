package services

import (
	"context"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/internal/domain/entity"
	"github.com/KaiBennington2/RabbitGo/internal/domain/ports"
)

var _ ports.ICreateUser = (*RegisterService)(nil)

type RegisterService struct {
	userRepository ports.IUserRepository
	userEventPub   ports.IUserEventPub
}

func NewUserRegisterService(repository ports.IUserRepository, eventPub ports.IUserEventPub) *RegisterService {
	return &RegisterService{
		userRepository: repository,
		userEventPub:   eventPub,
	}
}

func (srv *RegisterService) Execute(ctx context.Context, e *entity.User) (*string, error) {
	var newUserID *string
	var err error

	if ok, err := srv.existsById(ctx, &e.Code); err != nil || ok {
		if ok {
			return newUserID, err
		}
		return newUserID, err
	}

	newUserID, err = srv.userRepository.Save(ctx, e)
	if err != nil {
		return newUserID, err
	}

	e.Code = *newUserID
	err = srv.userEventPub.SavedEvent(ctx, e)
	if err != nil {
		return newUserID, err
	}
	return newUserID, nil
}

func (srv *RegisterService) existsById(ctx context.Context, code *string, withoutLogicDeleted ...bool) (bool, error) {
	if code == nil || *code == "" {
		return false, fmt.Errorf("not exists")
	}

	exist, err := srv.userRepository.ExistsById(ctx, code, withoutLogicDeleted...)
	if err != nil {
		return false, err
	}
	return exist, nil
}
