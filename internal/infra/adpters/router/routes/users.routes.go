package routes

import (
	"github.com/KaiBennington2/RabbitGo/internal/infra/adpters/modules/users"
	hdl "github.com/KaiBennington2/RabbitGo/internal/infra/handlers/rest/users"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r fiber.Router, m *users.UserModule) {

	protectedUserRoutes := r.Group("/users")
	{
		protectedUserRoutes.Post("/", hdl.Save(m.UseCases.Commands))
		protectedUserRoutes.Delete("/:user_id-:x_act", hdl.Delete(m.UseCases.Commands))
		protectedUserRoutes.Delete("/:user_id", hdl.LogicalDelete(m.UseCases.Commands))
	}

}
