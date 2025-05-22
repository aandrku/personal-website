package admin

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	// dashboard
	g.GET("/dashboard", getDashboardPage)

	// analytics
	g.GET("/analytics", getAnalyticsWidget)

	// statistics
	g.GET("/stats", getStatsWidget)

	// uploads
	g.GET("/upload/form", getUploadForm)
	g.GET("/upload/delete-form/:filename", getUploadDeleteForm)
	g.POST("/upload", postUpload)
	g.DELETE("/upload/:filename", deleteUpload)
}
