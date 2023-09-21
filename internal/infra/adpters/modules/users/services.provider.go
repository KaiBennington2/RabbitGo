package users

import (
	"github.com/KaiBennington2/RabbitGo/internal/domain/events"
	"github.com/KaiBennington2/RabbitGo/internal/domain/ports"
	"github.com/KaiBennington2/RabbitGo/internal/domain/services"
	"sync"
)

type UserServices struct {
	// -- Each user service --
	registerUser       ports.ICreateUser
	deleteUser         ports.IDeletePermanentUser
	deleteBusinessUser ports.IDeleteBusinessUser
	// -- Each user event service
	registeredUser ports.ISavedUserEvent
}

var onceSRVCS sync.Once
var srvInstance *UserServices
var LoadSERVICES = func(
	repos UserRepos,
	evtPubs UserEventsPub,
) *UserServices {
	onceSRVCS.Do(func() {
		inst := &UserServices{
			//Inject each instance its implementation here.
			registerUser:       services.NewUserRegisterService(repos.mysql, evtPubs.Publisher),
			deleteUser:         services.NewUserDeleteService(repos.mysql),
			deleteBusinessUser: services.NewUserBusinessDeleteService(repos.mysql),
			//
			registeredUser: events.NewUserRegisteredEvtService(repos.mongo),
		}
		srvInstance = inst
	})
	return srvInstance
}

// ----------------------------------------------------------------------------------------------------------------- //
