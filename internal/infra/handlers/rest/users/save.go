package users

import (
	"github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/dto"
	"github.com/KaiBennington2/RabbitGo/internal/app/usecases/users/ports"
	"github.com/gofiber/fiber/v2"
)

func Save(cmdUCases ports.ICommands) fiber.Handler {
	return func(c *fiber.Ctx) error {
		inputDto := new(dto.CreateUserDTO)
		if err := c.BodyParser(inputDto); err != nil {
			return err
		}

		rsp, err := cmdUCases.Create(c.Context(), *inputDto)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(rsp)
	}
}
