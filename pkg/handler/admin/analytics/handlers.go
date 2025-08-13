package analytics

import (
	"net/http"

	"github.com/aandrku/personal-website/pkg/services/analytics"
	"github.com/aandrku/personal-website/pkg/view"
	"github.com/aandrku/personal-website/pkg/view/admin"
	"github.com/labstack/echo/v4"
)

func newController() Controller {
	return Controller{
		service: analytics.Service{},
	}
}

type Controller struct {
	service analytics.Service
}

// / getAnalyticsWidget serves analytics to the client.
func (ct Controller) getAnalytics(c echo.Context) error {
	s := analytics.Service{}

	p := admin.AnalyticsProps{
		VisitsToday: s.TotalVisits(),
		VisitsTotal: s.TotalVisits(),
	}

	w := admin.Analytics(p)

	return view.Render(c, http.StatusOK, w)
}
