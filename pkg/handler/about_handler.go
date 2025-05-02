package handler

import (
	"net/http"
	"template1/pkg/view"
	"template1/pkg/view/components"

	"github.com/labstack/echo/v4"
)

func getAboutHandler(c echo.Context) error {
	content := view.Unsafe("<div>Hello, world</div>")

	component := components.AboutWindow(content)

	return view.Render(c, http.StatusOK, component)
}
