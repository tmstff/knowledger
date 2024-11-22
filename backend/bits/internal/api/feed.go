package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"tmstff/knowledger/bits/internal/entities"
)

type feed struct {
	app     *fiber.App
	results []entities.Feed // TODO factor out!
}

func NewFeed(app *fiber.App) *feed {
	return &feed{app: app, results: []entities.Feed{}}
}

func (f feed) SetupRoutes() fiber.Router {
	f.app.Get("/feeds", func(c *fiber.Ctx) error {
		return c.JSON(f.results)
	})
	f.app.Post("/feeds", func(c *fiber.Ctx) error {
		var feed entities.Feed
		err := json.Unmarshal(c.Body(), &feed)
		f.results = append(f.results, feed)

		if err != nil {
			return err
		}

		return c.Status(201).JSON(feed)
	})
	return f.app
}
