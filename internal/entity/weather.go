package entity

import (
	"errors"
	"regexp"
)

type Weather struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

func NewWeather(tempC float64) *Weather {
	return &Weather{
		TempC: tempC,
		TempF: tempC*1.8 + 32,
		TempK: tempC + 273.15,
	}
}

func ValidateZipCode(zipcode string) error {
	match, _ := regexp.MatchString(`^[0-9]{8}$`, zipcode)
	if !match {
		return errors.New("invalid zipcode")
	}
	return nil
}
