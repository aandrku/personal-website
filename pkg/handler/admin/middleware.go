package admin

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("auth")
		if err != nil {
			log.Printf("cookie is not retrieved %v", err)
			c.Response().Header().Set("HX-Redirect", "/login")
			return c.NoContent(http.StatusOK)
		}

		authToken := cookie.Value

		token, err := jwt.Parse(authToken, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_KEY")), nil
		})
		if err != nil {
			log.Printf("failed to parse the token %v", err)
			c.Response().Header().Set("HX-Redirect", "/login")
			return c.NoContent(http.StatusOK)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			fmt.Print("claims not ok")
			c.Response().Header().Set("HX-Redirect", "/login")
			return c.NoContent(http.StatusOK)
		}

		admin, ok := claims["admin"].(bool)
		if !ok || !admin {
			fmt.Print("admin not okay")
			c.Response().Header().Set("HX-Redirect", "/login")
			return c.NoContent(http.StatusOK)
		}

		return next(c)
	}
}
