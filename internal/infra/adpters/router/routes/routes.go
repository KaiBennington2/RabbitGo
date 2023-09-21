package routes

import (
	"github.com/KaiBennington2/RabbitGo/internal/infra/adpters/modules/users"
	"github.com/KaiBennington2/RabbitGo/sdk/ports"
)

func InitRoutes(app ports.WebServer, umod *users.UserModule) {
	protected := app.Group("/protected")
	UserRoutes(protected, umod)
}
