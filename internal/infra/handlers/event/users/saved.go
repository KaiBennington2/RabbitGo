package users

import (
	"context"
	"github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	"log"
	"time"
)

func Saved(evtUCases ports.IEvents) types.SubHandleFunc {
	handlerResponse := make(chan any)
	handlerError := make(chan error)
	go func() {
		for {
			select {
			case rsp := <-handlerResponse:
				if rsp != nil {
					log.Printf("Handler response:: %v\n", rsp)
				}
			case err := <-handlerError:
				log.Printf("Error Handler:: %v\n", err)
			}
		}
	}()

	return types.SubHandleFunc{
		Handler: func(msg []byte) (any, error) {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
			defer cancel()
			err := evtUCases.Saved(ctx, msg)
			return nil, err
		},
		HandlerReturn: handlerResponse,
	}
}
