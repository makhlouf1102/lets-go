package token

import (
	"lets-go/libs/env"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	refreshTokenExpiresIN time.Duration = time.Hour * 24 * 7
	accessTokenExpiresIN  time.Duration = time.Minute * 15
)

func CreateToken(userID string, secretKey []byte, expireIN time.Duration) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expireIN).Unix(), // Set expiration to 24 hours
	}

	// Create the token using the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and return the token as a string
	return token.SignedString(secretKey)
}

func CreateAccessToken(userID string) (string, error) {
	return CreateToken(userID, []byte(env.Get("TOKEN_ACCESS_SECRET")), accessTokenExpiresIN)
}

func CreateRefreshToken(userID string) (string, error) {
	return CreateToken(userID, []byte(env.Get("TOKEN_REFRESH_SECRET")), refreshTokenExpiresIN)
}
