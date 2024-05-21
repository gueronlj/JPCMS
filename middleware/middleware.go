package middleware

import (
	"net/http"

	"github.com/gueronlj/JPCMS/auth"
	"github.com/labstack/echo/v4"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reason, response := auth.CheckJWT((c.Request()))
		if response {
			return next(c)
		}
		return c.JSON(http.StatusUnauthorized, reason)
	}
}
