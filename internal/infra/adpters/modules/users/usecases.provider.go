package users

import (
	"github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/command"
	events "github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/event"
	"github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/ports"
	"sync"
)

type UserEvents struct {
	Crud ports.IEvents
}

type UserUseCases struct {
	Commands ports.ICommands
	Events   UserEvents
}

var onceUCASES sync.Once
var ucsInstance *UserUseCases
var LoadUSECASES = func(
	services UserServices,
) *UserUseCases {
	onceUCASES.Do(func() {
		inst := &UserUseCases{
			//Inject each instance its implementation here.
			Commands: command.NewUserCmdUseCases(
				services.registerUser,
				services.deleteUser,
				services.deleteBusinessUser,
			),
			Events: UserEvents{
				Crud: events.NewUserEvtUseCases(services.registeredUser),
			},
		}
		ucsInstance = inst
	})
	return ucsInstance
}

// ----------------------------------------------------------------------------------------------------------------- //
