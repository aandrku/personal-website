package site

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aandrku/personal-website/pkg/services/about"
	"github.com/aandrku/personal-website/pkg/services/analytics"
	"github.com/aandrku/personal-website/pkg/services/auth"
	"github.com/aandrku/personal-website/pkg/services/email"
	"github.com/aandrku/personal-website/pkg/services/project"
	"github.com/aandrku/personal-website/pkg/store/fs"
	"github.com/aandrku/personal-website/pkg/view"
	"github.com/aandrku/personal-website/pkg/view/home"
	"github.com/aandrku/personal-website/pkg/view/shared"
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

// getIndex serves index page to the client.
func getIndex(c echo.Context) error {
	as := analytics.Service{}
	as.IncrementVisits()

	page := home.Index()

	return view.Render(c, http.StatusOK, page)
}

func getLogin(c echo.Context) error {
	page := home.LoginPage()

	_, err := auth.GetCurrentOTP()
	if err != nil {
		return err
	}

	return view.Render(c, http.StatusOK, page)
}

func postLogin(c echo.Context) error {
	OTP := c.FormValue("otp")

	cOTP, err := auth.GetCurrentOTP()
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	if OTP != cOTP {
		fmt.Printf("cotp %s, otp %s", cOTP, OTP)
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	secretKey := os.Getenv("JWT_KEY")

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin": true,
	})

	// Sign and get the complete encoded token as a string using the secret
	authToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Printf("failed to sing token %v", err)
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	cookie := &http.Cookie{
		Name:     "auth",
		Value:    authToken,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		MaxAge:   86400,
		Secure:   false,
		HttpOnly: true,
	}

	c.SetCookie(cookie)
	return c.Redirect(http.StatusSeeOther, "/admin/dashboard")
}

// getAboutWindow serves about window to the client.
func getAboutWindow(c echo.Context) error {
	info, err := about.GetInfo()
	if err != nil {
		log.Fatalf("failed to get about info: %v\n", err)
	}
	component := home.AboutWindow(info)
	return view.Render(c, http.StatusOK, component)
}

func getProject(c echo.Context) error {
	id := c.Param("id")
	s := project.NewManager(fs.Store{})

	p, err := s.FindProject(id)
	if err != nil {
		return err
	}

	page := home.ProjectPage(p)

	return view.Render(c, http.StatusOK, page)
}

// getHomeWindow serves home window to the client.
func getHomeWindow(c echo.Context) error {
	component := home.HomeWindow()
	return view.Render(c, http.StatusOK, component)
}

func getProjectsWindow(c echo.Context) error {
	s := project.NewManager(fs.Store{})
	p, err := s.Projects()
	if err != nil {
		return err
	}

	component := home.ProjectsWindow(p)
	return view.Render(c, http.StatusOK, component)
}

// getContactWindow serves contact window to the client.
func getContactWindow(c echo.Context) error {
	component := home.ContactWindow()
	return view.Render(c, http.StatusOK, component)
}

func postContact(c echo.Context) error {
	n := c.FormValue("name")
	e := c.FormValue("email")
	msg := c.FormValue("message")

	if err := email.SendContact(n, e, msg); err != nil {
		ntf := shared.Notification("Due to my skill issues sending your email failed:(")
		return view.Render(c, http.StatusOK, ntf)
	}
	ntf := shared.Notification("Success! Your email was send and I'll get back to you once I see it.")

	return view.Render(c, http.StatusOK, ntf)
}

// getDelete serves empty http response to the client.
//
// This handler is used for removal of html elements, while using HTMX.
func getDelete(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
