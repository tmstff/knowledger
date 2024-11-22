package api

import "github.com/gofiber/fiber/v2"

type feed struct {
	app *fiber.App
}

func NewFeed(app *fiber.App) *feed {
	return &feed{app}
}

func (f feed) SetupRoutes() fiber.Router {
	return f.app.Get("/feeds", func(c *fiber.Ctx) error {
		return c.JSON([]string{})
	})
}
