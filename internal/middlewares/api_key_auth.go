package middlewares

import (
	"AuthService/configs"
	"AuthService/internal/general_utils"
	"log"
	"net/http"
	"strings"
)

// Checks provided api-key for each endpoint, excluding '/swagger'
func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request URL path starts with /swagger
		// If so, skip this middleware
		if strings.HasPrefix(r.URL.Path, "/swagger") {
			next.ServeHTTP(w, r)
			return
		}
		headerApiKey := r.Header.Get("X-Api-Key")
		log.Printf("ApiKeyMiddleware: headerApiKey=%s", headerApiKey)
		// Check for the header
		if headerApiKey == "" {
			general_utils.ErrorResponse(w, "Missing API Key", http.StatusForbidden)
			return
		}
		// validate the api-key
		if headerApiKey != configs.MainSettings.ApiKey {
			general_utils.ErrorResponse(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
