package admin

import (
	"net/http"
	"template1/pkg/services/analytics"
	"template1/pkg/services/stats"
	"template1/pkg/services/uploads"
	"template1/pkg/view"
	"template1/pkg/view/components/dashboard"
	"template1/pkg/view/components/forms"
	"template1/pkg/view/pages"

	"github.com/labstack/echo/v4"
)

func getDashboardHandler(c echo.Context) error {
	as := analytics.Service{}

	// TODO: give this better name, after I get rid of these redundant props
	awp := dashboard.AnalyticsWidgetProps{
		VisitsToday: as.TotalVisits(),
		VisitsTotal: as.TotalVisits(),
	}

	// this is here, cause of error handlign
	stats, err := stats.Get()
	if err != nil {
		return err
	}

	// same here
	// ups means uploads
	ups, err := uploads.Get()
	if err != nil {
		return err
	}

	props := pages.DashboardProps{
		AnalyticsWidgetProps: awp,
		Stats:                stats,
		Uploads:              ups,
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
	stats, err := stats.Get()
	if err != nil {
		return err
	}

	w := dashboard.SystemStatsWidget(stats)
	return view.Render(c, http.StatusOK, w)

}

func getFormsUpload(c echo.Context) error {
	form := forms.UploadAsset()
	return view.Render(c, http.StatusOK, form)
}
