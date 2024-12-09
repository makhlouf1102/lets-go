package logger_middleware

import (
	"lets-go/libs/token"
	"log"
	"net/http"
	"time"
)

type AdminGuar struct {
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (a *AdminGuar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// var userRole []string
	// if protectedData, ok := r.Context().Value("protected_data").(token.TokenClaim); !ok {
	// 	http.Error(w, "server error", http.StatusInternalServerError)
	// 	return
	// } 
	// userRole = protectedData["userRoles"].([]string)
	start := time.Now()
	a.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *AdminGuar {
	return &AdminGuar{handlerToWrap}
}
