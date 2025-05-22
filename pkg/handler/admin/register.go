package admin

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	g.GET("/dashboard", getDashboardHandler)
	g.GET("/analytics", getAnalyticsHandler)
	g.GET("/stats", getStatsHandler)
	g.GET("/forms/upload", getFormsUpload)
	g.GET("/forms/upload/delete/:filename", getFormsUploadDelete)
	g.DELETE("/upload/:filename", deleteUpload)
	g.POST("/upload", postUpload)
}
