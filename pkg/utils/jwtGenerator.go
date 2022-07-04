package utils

import (
	"time"

	"entgo.io/ent/examples/version/ent/user"
	"github.com/golang-jwt/jwt"
)

// GenerateNewAccessToken func for generate a new Access token.
func GenerateNewAccessToken(uid int, uuid string, name string, picture string, status user.Status, tokenExp time.Time, jwtSecret string) (string, error) {
	// Set secret key from .env.example.example file.
	secret := jwtSecret

	// Create a new claims.
	claims := jwt.MapClaims{}

	// Set public claims:
	claims["exp"] = tokenExp.Unix()
	claims["uid"] = uid
	claims["uuid"] = uuid
	claims["name"] = name
	claims["status"] = status
	claims["picture"] = picture

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
