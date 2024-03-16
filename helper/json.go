package helper

import (
	"github.com/gofiber/fiber/v2"
)

func ReadRequestBody(ctx *fiber.Ctx, result interface{}) {
	err := ctx.BodyParser(result)
	PanicIfError(err)
}

func WriteReponse(ctx *fiber.Ctx, response interface{}, status int) error {
	ctx.Set("Content-type", "application/json")
	err := ctx.Status(status).JSON(response)
	PanicIfError(err)
	return err
}
