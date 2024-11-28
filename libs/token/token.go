package token

import (
	"errors"
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

func Parse(tokenString string, secretkey string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC and not tampered
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(secretkey), nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractClaims(token *jwt.Token, secretKey []byte) (jwt.MapClaims, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, errors.New("unable to extract claims")
}

func CreateAccessToken(userID string) (string, error) {
	return CreateToken(userID, []byte(env.Get("TOKEN_ACCESS_SECRET")), accessTokenExpiresIN)
}

func CreateRefreshToken(userID string) (string, error) {
	return CreateToken(userID, []byte(env.Get("TOKEN_REFRESH_SECRET")), refreshTokenExpiresIN)
}
