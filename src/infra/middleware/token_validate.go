package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func ValidateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r.Header.Get("Authorization"))

		if isValidToken(token) {
			next(w, r)
			return
		}

		errorResponse := ErrorResponse{Message: "Unauthorized"}
		jsonResponse, err := json.Marshal(errorResponse)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(jsonResponse)
	}
}

func isValidToken(token string) bool {
	validToken := os.Getenv("AUTH_TOKEN")
	return token == validToken
}

func extractToken(authorizationHeader string) string {
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}
