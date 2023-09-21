package events

import (
	"context"
	"errors"
	"github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/ports"
	inPorts "github.com/KaiBennington2/RabbitGo/internal/domain/ports"
)

var _ ports.IEvents = (*EvtUseCases)(nil)

type EvtUseCases struct {
	registeredEvt inPorts.ISavedUserEvent
}

func NewUserEvtUseCases(registeredEvt inPorts.ISavedUserEvent) *EvtUseCases {
	return &EvtUseCases{
		registeredEvt: registeredEvt,
	}
}

func (uc *EvtUseCases) Saved(_ context.Context, message []byte) error {
	if string(message) == "error" {
		return errors.New("Error de caso de uso. ")
	}
	return nil
}
