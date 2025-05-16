package admin

import (
	"net/http"
	"template1/pkg/services/analytics"
	"template1/pkg/services/stats"
	"template1/pkg/view"
	"template1/pkg/view/components/dashboard"
	"template1/pkg/view/pages"

	"github.com/labstack/echo/v4"
)

func getDashboardHandler(c echo.Context) error {
	as := analytics.Service{}
	ss := stats.Service{}

	awp := dashboard.AnalyticsWidgetProps{
		VisitsToday: as.TotalVisits(),
		VisitsTotal: as.TotalVisits(),
	}

	stats, err := ss.Stats()
	if err != nil {
		return err
	}

	props := pages.DashboardProps{
		AnalyticsWidgetProps: awp,
		Stats:                stats,
	}
	dashboard := pages.Dashboard(props)

	return view.Render(c, http.StatusOK, dashboard)
}

func getAnalyticsHandler(c echo.Context) error {
	s := analytics.Service{}

	p := dashboard.AnalyticsWidgetProps{
		VisitsToday: s.TotalVisits(),
		VisitsTotal: s.TotalVisits(),
	}

	w := dashboard.AnalyticsWidget(p)

	return view.Render(c, http.StatusOK, w)
}

func getStatsHandler(c echo.Context) error {
	s := stats.Service{}

	stats, err := s.Stats()
	if err != nil {
		return err
	}

	w := dashboard.SystemStatsWidget(stats)
	return view.Render(c, http.StatusOK, w)

}
