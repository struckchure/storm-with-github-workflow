package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestReturnCurrentDate(t *testing.T) {
	currentDate := ReturnCurrentDate()

	assert.NotEmpty(t, currentDate)
}

func TestRootHandler(t *testing.T) {
	app := fiber.New()
	app.Get("/", RootHandler)

	resp, err := app.Test(httptest.NewRequest("GET", "/", nil), -1)
	assert.NoError(t, err)

	respBody := map[string]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	assert.NoError(t, err)

	assert.NoError(t, nil, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	assert.Equal(t, "gofiber", respBody["app"])
	assert.NotEmpty(t, respBody["datetime"])
}
