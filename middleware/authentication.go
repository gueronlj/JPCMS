package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/JPCMS/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

var secretKey = []byte(os.Getenv("SECRET"))

func GenerateToken(c echo.Context) error {
	// Set custom claims
	claims := &models.JwtCustomClaims{
		"Jon Snow",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// VerifyToken verifies a token JWT validate
