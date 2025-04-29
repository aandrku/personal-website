package handler

import (
	"net/http"
	"template1/pkg/view"
	"template1/pkg/view/pages/root"

	"github.com/labstack/echo/v4"
)

func rootHandler(c echo.Context) error {
	page := root.Page()

	return view.Render(c, http.StatusOK, page)
}
