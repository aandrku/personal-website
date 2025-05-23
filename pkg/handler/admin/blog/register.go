package blog

import "github.com/labstack/echo/v4"

func Register(g *echo.Group) {
	c := NewController()

	// widget
	g.GET("", c.getBlogManagementWidget)

	// forms
	g.GET("/forms/update-title/:postid", c.getEditPostTitleForm)
	g.GET("/forms/update-short-desc/:postid", c.getEditPostShortDescriptionForm)
	g.GET("/forms/update-content/:postid", c.getEditContentForm)
	g.GET("/forms/delete/:postid", c.getDeleteForm)
	g.GET("/post/create-form", c.getCreateForm)

	g.POST("/post", c.createPost)
	g.POST("/post/update-title/:postid", c.updatePostTitle)
	g.POST("/post/update-short-desc/:postid", c.updatePostShortDescription)
	g.POST("/post/update-content/:postid", c.updatePostContent)
	g.DELETE("/post/:postid", c.deletePost)
}
