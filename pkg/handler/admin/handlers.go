package admin

import (
	"log"
	"net/http"

	"github.com/aandrku/portfolio-v2/pkg/services/analytics"
	"github.com/aandrku/portfolio-v2/pkg/services/markdown"
	"github.com/aandrku/portfolio-v2/pkg/services/stats"
	"github.com/aandrku/portfolio-v2/pkg/services/uploads"
	"github.com/aandrku/portfolio-v2/pkg/view"
	"github.com/aandrku/portfolio-v2/pkg/view/components/dashboard"
	"github.com/aandrku/portfolio-v2/pkg/view/components/forms"
	"github.com/aandrku/portfolio-v2/pkg/view/pages"

	"github.com/labstack/echo/v4"
)

// getDashboardPage serves dashboard page to the client.
func getDashboardPage(c echo.Context) error {
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

// / getAnalyticsWidget serves analytics widget to the client.
func getAnalyticsWidget(c echo.Context) error {
	s := analytics.Service{}

	p := dashboard.AnalyticsWidgetProps{
		VisitsToday: s.TotalVisits(),
		VisitsTotal: s.TotalVisits(),
	}

	w := dashboard.AnalyticsWidget(p)

	return view.Render(c, http.StatusOK, w)
}

// getStatsWidget serves stats widget to the client.
func getStatsWidget(c echo.Context) error {
	stats, err := stats.Get()
	if err != nil {
		return err
	}

	w := dashboard.SystemStatsWidget(stats)
	return view.Render(c, http.StatusOK, w)

}

// getUploadWidget serves upload widget to the client.
func getUploadWidget(c echo.Context) error {
	ups, err := uploads.Get()
	if err != nil {
		return err
	}
	w := dashboard.UploadsManagementWidget(ups)
	return view.Render(c, http.StatusOK, w)
}

// getUploadForm serves upload form to the client.
func getUploadForm(c echo.Context) error {
	form := forms.UploadAsset()
	return view.Render(c, http.StatusOK, form)
}

// getUploadDeleteForm server upload-delete form to the client.
func getUploadDeleteForm(c echo.Context) error {
	filename := c.Param("filename")
	form := forms.DeleteUploadConfirmationForm(filename)
	return view.Render(c, http.StatusOK, form)
}

// postUpload store a new upload on the server's filesystem.
func postUpload(c echo.Context) error {
	name := c.FormValue("filename")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	err = uploads.New(src, name)
	if err != nil {
		return err
	}

	c.Response().Header().Add("HX-Trigger", "updateUploads")
	return c.NoContent(http.StatusOK)
}

// deleteUpload deletes an upload from the server.
func deleteUpload(c echo.Context) error {
	filename := c.Param("filename")

	err := uploads.Remove(filename)
	if err != nil {
		log.Fatalf("%s file, %v", filename, err)
		return err
	}

	c.Response().Header().Add("HX-Trigger", "updateUploads")
	return c.NoContent(http.StatusOK)
}

func postMarkdownPreview(c echo.Context) error {
	md := c.FormValue("markdown")

	html, err := markdown.ToHTML(md)
	if err != nil {
		return err
	}

	t := view.Unsafe(html)

	return view.Render(c, http.StatusOK, t)
}
