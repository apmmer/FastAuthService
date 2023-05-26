package middlewares

import (
	"auth_service_api/configs"
	"auth_service_api/internal/handlers/handlers_utils"
	"log"
	"net/http"
	"strings"
)

// ApiKeyMiddleware checks provided api-key for the endpoint
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
			handlers_utils.ErrorResponse(w, "Missing API Key", http.StatusForbidden)
			return
		}
		// validate the api-key
		if headerApiKey != configs.MainSettings.ApiKey {
			handlers_utils.ErrorResponse(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
