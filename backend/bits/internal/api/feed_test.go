package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
	"tmstff/knowledger/bits/internal/util"
)

func TestFeeds(t *testing.T) {
	// given
	app := setupApi()

	// when
	req := httptest.NewRequest("POST", "/feeds", strings.NewReader(`{"name": "my feed name"}`))
	res, err := app.Test(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, 201, res.StatusCode)

	assert.NoError(t, err)
	assert.JSONEq(t, `{"name": "my feed name"}`, util.ForceStringOf(t, res.Body))

	// when
	req = httptest.NewRequest("GET", "/feeds", nil)
	res, err = app.Test(req)

	// then
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)

	assert.NoError(t, err)
	assert.JSONEq(t, "[{\"name\": \"my feed name\"}]", util.ForceStringOf(t, res.Body))
}

func setupApi() *fiber.App {
	app := fiber.New()

	NewFeed(app).SetupRoutes()
	return app
}
