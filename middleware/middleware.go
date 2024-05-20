package middleware

import (
	"fmt"
	"net/http"

	"github.com/gueronlj/JPCMS/auth"
	"github.com/labstack/echo/v4"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if auth.CheckJWT(c.Request()) == "success" {
			fmt.Println("we mid!")
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, auth.CheckJWT(c.Request()))
	}
}
