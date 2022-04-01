package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Ninja struct {
	Name   string
	Weapon string
}

func getNinja(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

var ninja Ninja

func createNinja(ctx *fiber.Ctx) error {
	body := new(Ninja)
	err := ctx.BodyParser(body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		return err
	}

	ninja = Ninja{
		Name:   body.Name,
		Weapon: body.Weapon,
	}
	return ctx.Status(fiber.StatusOK).JSON(ninja)
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	//app.Use(limiter.New())

	app.Use(limiter.New(limiter.Config{
		Max:               20,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	app.Get("/dashboard", monitor.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	app.Get("/ninja", getNinja)
	app.Post("/ninja", createNinja)

	app.Listen(":8080")
}
