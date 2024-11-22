package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
	"tmstff/knowledger/bits/internal/util"
)

func TestGetUsersEndpoint(t *testing.T) {
	app := fiber.New()

	NewFeed(app).SetupRoutes()

	// Create a request
	req := httptest.NewRequest("GET", "/feeds", nil)

	// Test the handler
	res, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, res.StatusCode)
	
	s, err := util.StringOf(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, s, "[]")
}
