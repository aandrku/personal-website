package handler

import (
	"net/http"
	"template1/pkg/view"
	"template1/pkg/view/components"

	"github.com/labstack/echo/v4"
)

func getHomeHandler(c echo.Context) error {
	component := components.HomeWindow()
	return view.Render(c, http.StatusOK, component)
}
