package token

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
	"lets-go/libs/env"
)

func CreateToken(userID string) (string, error) {
    // Set token claims
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(), // Set expiration to 24 hours
    }

    // Create the token using the claims and signing method
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign and return the token as a string
	return token.SignedString([]byte(env.Get("SECRET_KEY")))
}
