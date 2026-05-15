package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/luigolima/go-temperature-by-cep-challenge/internal/infra/web"
	"github.com/luigolima/go-temperature-by-cep-challenge/internal/usecase"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(httprate.LimitByIP(10, 1*time.Minute))

	weatherUseCase := usecase.NewWeatherUseCase()
	weatherHandler := web.NewWeatherHandler(weatherUseCase)

	r.Get("/{zipcode}", weatherHandler.GetWeather)

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
