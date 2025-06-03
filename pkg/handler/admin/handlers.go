package admin

import (
	"log"
	"net/http"

	"github.com/aandrku/portfolio-v2/pkg/services/blog"
	"github.com/aandrku/portfolio-v2/pkg/services/markdown"
	"github.com/aandrku/portfolio-v2/pkg/services/project"
	"github.com/aandrku/portfolio-v2/pkg/services/stats"
	"github.com/aandrku/portfolio-v2/pkg/services/uploads"
	"github.com/aandrku/portfolio-v2/pkg/store/fs"
	"github.com/aandrku/portfolio-v2/pkg/view"
	"github.com/aandrku/portfolio-v2/pkg/view/admin"
	"github.com/labstack/echo/v4"
)

// getDashboardPage serves dashboard page to the client.
func getDashboardPage(c echo.Context) error {
	bs := blog.NewService()
	ps := project.NewManager(fs.Store{})

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

	posts, err := bs.Posts()
	if err != nil {
		return err
	}

	projects, err := ps.Projects()
	if err != nil {
		log.Print(err)
		return err
	}

	props := admin.DashboardProps{
		Stats:    stats,
		Uploads:  ups,
		Posts:    posts,
		Projects: projects,
	}

	dashboard := admin.Dashboard(props)

	return view.Render(c, http.StatusOK, dashboard)
}

// getStatsWidget serves stats widget to the client.
func getStatsWidget(c echo.Context) error {
	stats, err := stats.Get()
	if err != nil {
		return err
	}

	w := admin.SystemStatsWidget(stats)
	return view.Render(c, http.StatusOK, w)

}

// getUploadWidget serves upload widget to the client.
func getUploadWidget(c echo.Context) error {
	ups, err := uploads.Get()
	if err != nil {
		return err
	}
	w := admin.UploadsManagementWidget(ups)
	return view.Render(c, http.StatusOK, w)
}

// getUploadForm serves upload form to the client.
func getUploadForm(c echo.Context) error {
	form := admin.UploadAssetForm()
	return view.Render(c, http.StatusOK, form)
}

// getUploadDeleteForm server upload-delete form to the client.
func getUploadDeleteForm(c echo.Context) error {
	filename := c.Param("filename")
	form := admin.DeleteUploadForm(filename)
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
