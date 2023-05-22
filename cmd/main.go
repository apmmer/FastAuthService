package main

import (
	"AuthService/configs"
	"AuthService/database"
	"AuthService/internal/routers"
	"log"
	"net/http"
)

// @title Auth Service API
// @version v.1.0.0
// @description API server for auth stuff

// @host localhost:8080
// @BasePath /
// @schemes https

// @securityDefinitions.apikey JWTAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Api-Key

// @security JWTAuth
// @security ApiKeyAuth
func main() {
	database.InitDB(configs.MainSettings.UsersDBURL)
	router := routers.GetRouter()

	log.Println("AuthService is running on :8080 with HTTPS")
	log.Fatal(
		http.ListenAndServeTLS(
			":8080",
			configs.MainSettings.СertFileLocation,
			configs.MainSettings.CertKeyLocation,
			router,
		),
	)
}
