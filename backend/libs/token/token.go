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

type TokenClaim struct {
	UserID    string        `json:"user_id"`
	UserRoles []string      `json:"user_roles"`
	ExpiresIn time.Duration `json:"exp"`
}

func (tc *TokenClaim) ConvertToJWTClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"user_id":    tc.UserID,
		"user_roles": tc.UserRoles,
		"exp":        time.Now().Add(tc.ExpiresIn).Unix(),
	}
}

func CreateToken(claim TokenClaim, secretKey []byte) (string, error) {
	// Set token claims
	claims := claim.ConvertToJWTClaims()

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

func CreateAccessToken(userID string, listRoles []string) (string, error) {
	return CreateToken(TokenClaim{
		UserID:    userID,
		UserRoles: listRoles,
		ExpiresIn: accessTokenExpiresIN,
	}, []byte(env.Get("TOKEN_ACCESS_SECRET")))
}

func CreateRefreshToken(userID string, listRoles []string) (string, error) {
	return CreateToken(TokenClaim{
		UserID:    userID,
		UserRoles: listRoles,
		ExpiresIn: refreshTokenExpiresIN,
	}, []byte(env.Get("TOKEN_REFRESH_SECRET")))
}
