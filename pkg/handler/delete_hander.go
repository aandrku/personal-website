package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func deleteHandler(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
