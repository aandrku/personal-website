package blog

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	c := NewController()

	// widget
	g.GET("/widget", c.getBlogManagementWidget)

	// forms
	g.GET("/update-title-form/:postid", c.getEditPostTitleForm)
	g.GET("/update-short-desc-form/:postid", c.getEditPostShortDescriptionForm)
	g.GET("/update-content-form/:postid", c.getEditContentForm)
	g.GET("/delete-form/:postid", c.getDeleteForm)
	g.GET("/create-form", c.getCreateForm)

	g.POST("/posts", c.createPost)
	g.POST("/posts/:postid/title", c.updatePostTitle)
	g.POST("/posts/:postid/short-desc", c.updatePostShortDescription)
	g.POST("/posts/:postid/content", c.updatePostContent)
	g.DELETE("/posts/:postid", c.deletePost)
}
