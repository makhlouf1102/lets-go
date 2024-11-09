package auth

import (
	"net/http"
)

type RegisterRequest struct {
}

func Register(w http.ResponseWriter, r *http.Request) {}

func Login(w http.ResponseWriter, r *http.Request) {}
