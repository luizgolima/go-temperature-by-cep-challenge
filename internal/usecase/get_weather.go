package usecase

import (
	"github.com/luigolima/go-temperature-by-cep-challenge/internal/entity"
	"github.com/luigolima/go-temperature-by-cep-challenge/internal/infra/viacep"
	"github.com/luigolima/go-temperature-by-cep-challenge/internal/infra/weatherapi"
)

type WeatherOutputDTO struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type WeatherUseCase struct{}

func NewWeatherUseCase() *WeatherUseCase {
	return &WeatherUseCase{}
}

func (u *WeatherUseCase) Execute(zipcode, apiKey string) (*WeatherOutputDTO, error) {
	err := entity.ValidateZipCode(zipcode)
	if err != nil {
		return nil, err
	}

	city, err := viacep.GetCityByZipCode(zipcode)
	if err != nil {
		return nil, err
	}

	tempC, err := weatherapi.GetTemperatureByCity(city, apiKey)
	if err != nil {
		return nil, err
	}

	weather := entity.NewWeather(tempC)

	return &WeatherOutputDTO{
		TempC: weather.TempC,
		TempF: weather.TempF,
		TempK: weather.TempK,
	}, nil
}

