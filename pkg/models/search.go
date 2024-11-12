// made by recanman
package models

import "github.com/google/uuid"

type Search struct {
	ID       string `json:"id"`
	Platform string `json:"platform" validate:"required"`
	Location string `json:"location" validate:"required"`

	// Either car_make or "Any" for all makes
	CarMake string `json:"car_make" validate:"required"`

	YearMin *int `json:"year_min"`
	YearMax *int `json:"year_max"`

	MileageMin *int `json:"mileage_min"`
	MileageMax *int `json:"mileage_max"`

	PriceMin *int `json:"price_min"`
	PriceMax *int `json:"price_max"`

	Radius *int `json:"radius"`
}

func NewID() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	return uuid.String(), nil
}
