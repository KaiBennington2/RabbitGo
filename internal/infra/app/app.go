package app

import (
	"errors"
	"fmt"
	"github.com/KaiBennington2/RabbitGo/internal/infra/adpters/modules/users"
	"github.com/KaiBennington2/RabbitGo/internal/infra/adpters/router/routes"
	"github.com/KaiBennington2/RabbitGo/internal/infra/adpters/router/server"
	"github.com/KaiBennington2/RabbitGo/internal/infra/adpters/settings"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
	"github.com/KaiBennington2/RabbitGo/sdk/types"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartApplication() error {

	// ------------------------------------------------------------------------------------------------------------ //
	// ---------------------------- The APP initial configuration files are loaded -------------------------------- //
	// ------------------------------------------------------------------------------------------------------------ //
	appConfig := ports.ISetting(settings.NewSetting(types.Yaml))
	config, err := appConfig.Load()
	if err != nil || config == nil {
		if err != nil {
			return err
		}
		return errors.New("the configuration file is empty. ")
	}
	// ------------------------------------------------------------------------------------------------------------ //

	// ------------------------------------------------------------------------------------------------------------ //
	// ------------------------- All rsc related to the APP are initialized --------------------------- //
	// ------------------------------------------------------------------------------------------------------------ //
	rsc, err := InitAppResources(config)
	if err != nil {
		return err
	}
	// ------------------------------------------------------------------------------------------------------------ //

	// ------------------------------------------------------------------------------------------------------------ //
	// ------------------------ The Web Server where the APP will listen is initialized. -------------------------- //
	// ------------------------------------------------------------------------------------------------------------ //
	app := ports.WebServer(server.NewFiberServer(config.ServerConfig))
	// Init Logger Middleware
	app.Use(logger.New())

	// ------------------------------------------------------------------------------------------------------------ //
	// -------------------------- Section where the system modules are initialized. ------------------------------- //
	// ------------------------------------------------------------------------------------------------------------ //
	uMod := users.InitMODULE(rsc.DatabasesMap, rsc.BrokersMap)
	routes.InitRoutes(app, uMod)

	okChan := make(chan []byte)
	dsChan := make(chan []byte)
	errChan := make(chan error)

	tck := types.SubTraceHandler{
		DoneMsg:    okChan,
		DiscardMsg: dsChan,
		Errs:       errChan,
	}

	go func() {
		for {
			select {
			case msg := <-dsChan:
				fmt.Printf("mensaje descartado:: %v\n\n", string(msg))
			case msg := <-okChan:
				fmt.Printf("mensaje exitoso:: %v\n", string(msg))
			case err := <-errChan:
				fmt.Printf("mensaje erroneo:: %v\n", err)
			}
		}
	}()

	go uMod.Subs.SavedEvent.Execute(tck)
	// ------------------------------------------------------------------------------------------------------------ //

	// ------------------------------------------------------------------------------------------------------------ //
	// ---------------------- The server is up on the port indicated in the configuration. ------------------------ //
	// ------------------------------------------------------------------------------------------------------------ //
	return app.Start()
	// ------------------------------------------------------------------------------------------------------------ //
}
