package web

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/luigolima/go-temperature-by-cep-challenge/internal/usecase"
)

type WeatherUseCaseInterface interface {
	Execute(zipcode, apiKey string) (*usecase.WeatherOutputDTO, error)
}

type WeatherHandler struct {
	UseCase WeatherUseCaseInterface
}

func NewWeatherHandler(useCase WeatherUseCaseInterface) *WeatherHandler {
	return &WeatherHandler{
		UseCase: useCase,
	}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	zipcode := chi.URLParam(r, "zipcode")
	apiKey := os.Getenv("WEATHER_API_KEY")

	output, err := h.UseCase.Execute(zipcode, apiKey)
	if err != nil {
		if err.Error() == "invalid zipcode" {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}
		if err.Error() == "can not find zipcode" {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

