package views

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Success Views
func OkView(ctx *fiber.Ctx, data interface{}) error {
	return ctx.
		Status(200).
		JSON(fiber.Map{
			"msg":  "ok",
			"body": data,
		})

}

func CreatedView(ctx *fiber.Ctx, data interface{}) error {
	return ctx.
		Status(201).
		JSON(fiber.Map{
			"msg":  "ok",
			"body": data,
		})

}

// Error Views
func InvalidParams(ctx *fiber.Ctx) error {
	return ctx.
		Status(400).
		JSON(fiber.Map{
			"msg": "invalid params",
		})
}

func InternalServerError(ctx *fiber.Ctx, err error) error {
	log.Println(err)
	return ctx.
		Status(500).
		JSON(fiber.Map{
			"msg": "something went wrong",
		})
}

func RecordNotFound(ctx *fiber.Ctx) error {
	return ctx.
		Status(404).
		JSON(fiber.Map{
			"msg": "not found",
		})
}

func UnAuthorisedView(ctx *fiber.Ctx) error {
	return ctx.
		Status(401).
		JSON(fiber.Map{
			"msg": "unauthorised",
		})
}

func ForbiddenView(ctx *fiber.Ctx) error {
	return ctx.
		Status(403).
		JSON(fiber.Map{
			"msg": "forbidden",
		})
}
