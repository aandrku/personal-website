package admin

import (
	"net/http"
	"template1/pkg/services/analytics"
	"template1/pkg/view"
	"template1/pkg/view/components/dashboard"
	"template1/pkg/view/pages"

	"github.com/labstack/echo/v4"
)

func getDashboardHandler(c echo.Context) error {
	as := analytics.Service{}

	awp := dashboard.AnalyticsWidgetProps{
		VisitsToday: as.TotalVisits(),
		VisitsTotal: as.TotalVisits(),
	}

	props := pages.DashboardProps{
		AnalyticsWidgetProps: awp,
	}
	dashboard := pages.Dashboard(props)

	return view.Render(c, http.StatusOK, dashboard)
}
