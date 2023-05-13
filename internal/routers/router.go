package routers

import (
	"net/http"

	"AuthService/internal/handlers"

	_ "AuthService/docs"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func GetRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/healthcheck", handlers.HealthCheck).Methods("GET")

	// Swagger endpoint
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	// wrap our router in middleware for CORS
	corsHandler := cors.Default().Handler(router)

	return corsHandler
}
