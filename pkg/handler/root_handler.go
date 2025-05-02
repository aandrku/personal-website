package handler

import (
	"net/http"
	"template1/pkg/view"
	"template1/pkg/view/pages"

	"github.com/labstack/echo/v4"
)

func indexHandler(c echo.Context) error {
	page := pages.Index()

	return view.Render(c, http.StatusOK, page)
}
