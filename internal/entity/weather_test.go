package entity

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewWeather(t *testing.T) {
	tempC := 28.5
	weather := NewWeather(tempC)

	assert.InDelta(t, 28.5, weather.TempC, 0.01)
	assert.InDelta(t, 83.3, weather.TempF, 0.01)
	assert.InDelta(t, 301.65, weather.TempK, 0.01)
}

func TestValidateZipCode(t *testing.T) {
	tests := []struct {
		zipcode string
		isValid bool
	}{
		{"12345678", true},
		{"1234567", false},
		{"123456789", false},
		{"12345a78", false},
		{"", false},
	}

	for _, tt := range tests {
		err := ValidateZipCode(tt.zipcode)
		if tt.isValid {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
			assert.Equal(t, "invalid zipcode", err.Error())
		}
	}
}
