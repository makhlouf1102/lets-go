package auth_middleware

import (
	"log"
	"net/http"
)

type TokenRefresher struct {
	handler http.Handler
}

func (t *TokenRefresher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// do something before
	t.handler.ServeHTTP(w, r)
	log.Printf("token verified")
}

func NewTokenRefresher(handlerToWrap http.Handler) *TokenRefresher {
	return &TokenRefresher{handlerToWrap}
}
