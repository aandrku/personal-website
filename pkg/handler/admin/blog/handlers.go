package blog

import (
	"net/http"

	"github.com/aandrku/portfolio-v2/pkg/model"
	"github.com/aandrku/portfolio-v2/pkg/services/blog"
	"github.com/aandrku/portfolio-v2/pkg/services/markdown"
	"github.com/aandrku/portfolio-v2/pkg/view"
	"github.com/aandrku/portfolio-v2/pkg/view/admin"
	"github.com/labstack/echo/v4"
)

func NewController() *Controller {
	s := blog.NewService()

	return &Controller{
		service: s,
	}
}

type Controller struct {
	service *blog.Service
}

// getPosts serves a blog management widget to the client.
func (ct *Controller) getPosts(c echo.Context) error {
	posts, err := ct.service.Posts()
	if err != nil {
		return err
	}

	w := admin.PostList(posts)
	return view.Render(c, http.StatusOK, w)
}

// getEditPostTitleForm serves an edit post title form.
func (ct *Controller) getEditPostTitleForm(c echo.Context) error {

	postID := c.Param("postid")
	post, err := ct.service.FindPost(postID)
	if err != nil {
		return err
	}

	form := admin.EditPostTitleForm(post)

	return view.Render(c, http.StatusOK, form)
}

// getEditPostDateForm serves an edit post date form.
func (ct *Controller) getEditPostShortDescriptionForm(c echo.Context) error {
	postID := c.Param("postid")
	post, err := ct.service.FindPost(postID)
	if err != nil {
		return err
	}

	form := admin.EditPostShortDescriptionForm(post)

	return view.Render(c, http.StatusOK, form)
}

// getEditContentForm serves an update blog content form.
func (ct *Controller) getEditContentForm(c echo.Context) error {
	postID := c.Param("postid")
	post, err := ct.service.FindPost(postID)
	if err != nil {
		return err
	}
	html, err := markdown.ToHTML(post.Content)
	if err != nil {
		return nil
	}

	form := admin.EditPostContentForm(html, post)

	return view.Render(c, http.StatusOK, form)
}

// getDeleteForm serves a delete blog form.
func (ct *Controller) getDeleteForm(c echo.Context) error {
	postID := c.Param("postid")
	post, err := ct.service.FindPost(postID)
	if err != nil {
		return err
	}

	form := admin.DeletePostForm(post)

	return view.Render(c, http.StatusOK, form)
}

func (ct *Controller) getCreateForm(c echo.Context) error {
	form := admin.CreatePostForm()

	return view.Render(c, http.StatusOK, form)
}

func (ct *Controller) createPost(c echo.Context) error {
	t := c.FormValue("title")
	sh := c.FormValue("short")
	cnt := c.FormValue("markdown")

	post := model.NewPost(t, sh, cnt)

	if err := ct.service.CreatePost(post); err != nil {
		return nil
	}

	c.Response().Header().Add("HX-Trigger", "updateBlog")
	return c.NoContent(http.StatusOK)
}

func (ct *Controller) updatePostTitle(c echo.Context) error {
	postID := c.Param("postid")

	title := c.FormValue("title")
	post, err := ct.service.FindPost(postID)
	if err != nil {
		return err
	}

	post.Title = title

	if err := ct.service.UpdatePost(post); err != nil {
		return err
	}

	c.Response().Header().Add("HX-Trigger", "updateBlog")
	return c.NoContent(http.StatusOK)
}

func (ct *Controller) updatePostShortDescription(c echo.Context) error {
	postID := c.Param("postid")

	sd := c.FormValue("short-description")
	post, err := ct.service.FindPost(postID)
	if err != nil {
		return err
	}

	post.ShortDesc = sd

	if err := ct.service.UpdatePost(post); err != nil {
		return err
	}

	c.Response().Header().Add("HX-Trigger", "updateBlog")
	return c.NoContent(http.StatusOK)
}

func (ct *Controller) updatePostContent(c echo.Context) error {
	postID := c.Param("postid")

	content := c.FormValue("markdown")
	post, err := ct.service.FindPost(postID)
	if err != nil {
		return err
	}

	post.Content = content

	if err := ct.service.UpdatePost(post); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (ct *Controller) deletePost(c echo.Context) error {
	postID := c.Param("postid")

	if err := ct.service.DeletePost(postID); err != nil {
		return err
	}

	c.Response().Header().Add("HX-Trigger", "updateBlog")
	return c.NoContent(http.StatusOK)
}
