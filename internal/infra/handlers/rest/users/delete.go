package users

import (
	"github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/ports"
	"github.com/gofiber/fiber/v2"
)

func LogicalDelete(cmdUCases ports.ICommands) fiber.Handler {
	return func(c *fiber.Ctx) error {

		rsp, err := cmdUCases.DeletedBusiness(c.Context(), c.Params("user_code"))
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(rsp)
	}
}

func Delete(cmdUCases ports.ICommands) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rsp, err := cmdUCases.DeletedPermanent(c.Context(), c.Params("user_code"))
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(rsp)
	}
}
