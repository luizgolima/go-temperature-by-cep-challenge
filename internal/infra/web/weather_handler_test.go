package web

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/luigolima/go-temperature-by-cep-challenge/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockWeatherUseCase struct {
	mock.Mock
}

func (m *MockWeatherUseCase) Execute(zipcode, apiKey string) (*usecase.WeatherOutputDTO, error) {
	args := m.Called(zipcode, apiKey)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*usecase.WeatherOutputDTO), args.Error(1)
}

func TestGetWeatherHandler(t *testing.T) {
	t.Run("should return 200 and weather data on success", func(t *testing.T) {
		mockUseCase := new(MockWeatherUseCase)
		handler := NewWeatherHandler(mockUseCase)

		expectedOutput := &usecase.WeatherOutputDTO{
			TempC: 28.5,
			TempF: 83.3,
			TempK: 301.65,
		}

		mockUseCase.On("Execute", "01153000", mock.Anything).Return(expectedOutput, nil)

		r := chi.NewRouter()
		r.Get("/{zipcode}", handler.GetWeather)

		req, _ := http.NewRequest("GET", "/01153000", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		
		var actualOutput usecase.WeatherOutputDTO
		json.Unmarshal(rr.Body.Bytes(), &actualOutput)
		assert.Equal(t, expectedOutput.TempC, actualOutput.TempC)
		assert.Equal(t, expectedOutput.TempF, actualOutput.TempF)
		assert.Equal(t, expectedOutput.TempK, actualOutput.TempK)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return 422 when zipcode is invalid", func(t *testing.T) {
		mockUseCase := new(MockWeatherUseCase)
		handler := NewWeatherHandler(mockUseCase)

		mockUseCase.On("Execute", "123", mock.Anything).Return(nil, errors.New("invalid zipcode"))

		r := chi.NewRouter()
		r.Get("/{zipcode}", handler.GetWeather)

		req, _ := http.NewRequest("GET", "/123", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Contains(t, rr.Body.String(), "invalid zipcode")
		mockUseCase.AssertExpectations(t)
	})

	t.Run("should return 404 when zipcode is not found", func(t *testing.T) {
		mockUseCase := new(MockWeatherUseCase)
		handler := NewWeatherHandler(mockUseCase)

		mockUseCase.On("Execute", "99999999", mock.Anything).Return(nil, errors.New("can not find zipcode"))

		r := chi.NewRouter()
		r.Get("/{zipcode}", handler.GetWeather)

		req, _ := http.NewRequest("GET", "/99999999", nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
		assert.Contains(t, rr.Body.String(), "can not find zipcode")
		mockUseCase.AssertExpectations(t)
	})
}
