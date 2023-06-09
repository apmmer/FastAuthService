package routers

import (
	"net/http"

	"auth_service_api/configs"
	"auth_service_api/internal/handlers"
	"auth_service_api/internal/middlewares"

	_ "auth_service_api/docs"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Returns a mux.NewRouter() with added endpoints and declared middlewares.
func GetRouter() http.Handler {
	router := mux.NewRouter()
	router.Use(middlewares.ApiKeyMiddleware)

	router.HandleFunc("/api/healthcheck", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/api/users", handlers.RegisterUser).Methods("POST")
	router.HandleFunc("/api/users", handlers.GetUsersList).Methods("GET")
	router.HandleFunc("/api/users/{id}", handlers.GetUserById).Methods("GET")
	router.HandleFunc("/api/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/refresh", handlers.RefreshTokens).Methods("POST")
	router.HandleFunc("/api/validate", handlers.ValidateAccess).Methods("POST")
	router.HandleFunc("/api/logout", handlers.Logout).Methods("POST")

	// Swagger endpoint
	router.PathPrefix(configs.MainSettings.SwaggerUrl).Handler(httpSwagger.WrapHandler)

	// wrap our router in middleware for CORS
	corsHandler := cors.Default().Handler(router)

	return corsHandler
}
