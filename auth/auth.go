package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

var secretKey = []byte(os.Getenv("SECRET"))

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

// func createToken(username string) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
// 		jwt.MapClaims{
// 			"username": username,
// 			"exp":      time.Now().Add(time.Hour * 24).Unix(),
// 		})

// 	tokenString, err := token.SignedString(secretKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

func CheckAuth(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return ("Missing authorization header")
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		return ("Invalid token")
	}
	return ("success")
}

func Login(c echo.Context) error {
	// username := c.FormValue("username")
	// password := c.FormValue("password")
	// // Throws unauthorized error
	// if username != "jon" || password != "shhh!" {
	// 	return echo.ErrUnauthorized
	// }

	// Set custom claims
	claims := JwtCustomClaims{
		"jon",
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"token": tokenString,
	})
}
