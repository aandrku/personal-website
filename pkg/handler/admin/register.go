package admin

import (
	"github.com/aandrku/portfolio-v2/pkg/handler/admin/about"
	"github.com/aandrku/portfolio-v2/pkg/handler/admin/analytics"
	"github.com/aandrku/portfolio-v2/pkg/handler/admin/blog"
	"github.com/aandrku/portfolio-v2/pkg/handler/admin/projects"
	"github.com/labstack/echo/v4"
)

func Register(g *echo.Group) {
	g.Use(auth)
	// dashboard
	g.GET("/dashboard", getDashboardPage)

	// markdown preview
	g.POST("/markdown-preview", postMarkdownPreview)

	// analytics
	analytics.Register(g.Group("/analytics"))

	// about
	about.Register(g.Group("/about"))

	// statistics
	g.GET("/stats", getStats)

	// uploads
	g.GET("/uploads", getUploads)
	g.GET("/upload/form", getUploadForm)
	g.GET("/upload/delete-form/:filename", getUploadDeleteForm)
	g.POST("/upload", postUpload)
	g.DELETE("/upload/:filename", deleteUpload)

	// blog
	blog.Register(g.Group("/blog"))

	// projects
	projects.Register(g.Group("/projects"))
}
