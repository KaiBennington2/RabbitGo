package command

import (
	"context"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/dto"
	"github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/ports"
	"github.com/KaiBennington2/RabbitGo/internal/domain/entity"
	inPorts "github.com/KaiBennington2/RabbitGo/internal/domain/ports"
)

var _ ports.ICommands = (*CmdUseCases)(nil)

type CmdUseCases struct {
	registerUser       inPorts.ICreateUser
	deleteUser         inPorts.IDeletePermanentUser
	deleteBusinessUser inPorts.IDeleteBusinessUser
}

func NewUserCmdUseCases(
	registerUser inPorts.ICreateUser,
	deleteUser inPorts.IDeletePermanentUser,
	deleteBusinessUser inPorts.IDeleteBusinessUser,
) *CmdUseCases {

	return &CmdUseCases{
		registerUser:       registerUser,
		deleteUser:         deleteUser,
		deleteBusinessUser: deleteBusinessUser,
	}
}

// Create Use case to create a new user-type entity in the system
func (uc *CmdUseCases) Create(ctx context.Context, bodyPrev dto.CreateUserDTO) (*string, error) {
	var rsp *string

	rsp, err := uc.registerUser.Execute(ctx, &entity.User{
		ID:   bodyPrev.ID,
		Code: bodyPrev.Code,
		Name: bodyPrev.Name,
	})
	return rsp, err
}

// DeletedPermanent use case to perform a permanent deletion of a user-type entity in the system.
func (uc *CmdUseCases) DeletedPermanent(ctx context.Context, codePrev string) (*string, error) {
	var rsp *string
	var wasDeleted bool

	wasDeleted, err := uc.deleteUser.Execute(ctx, &codePrev)
	if err != nil {
		return rsp, err
	}

	if !wasDeleted {
		return rsp, fmt.Errorf("could not be deleted")
	}

	m := "The user was successfully deleted."
	return &m, nil
}

// DeletedBusiness use case to perform a logical deletion to a user-type entity in different processes of the system.
func (uc *CmdUseCases) DeletedBusiness(ctx context.Context, codePrev string) (*string, error) {
	var rsp *string
	var wasDeleted bool

	wasDeleted, err := uc.deleteBusinessUser.Execute(ctx, &codePrev)
	if err != nil {
		return rsp, err
	}

	if !wasDeleted {
		return rsp, fmt.Errorf("could not be deleted")
	}

	m := "The user was successfully deleted."
	return &m, nil
}
