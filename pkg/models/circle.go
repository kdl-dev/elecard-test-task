package models

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Radius struct {
	R json.Number `json:"radius"`
}

func NewRadius(radius string) (Radius, error) {
	value, err := strconv.ParseFloat(radius, 64)
	if err != nil {
		return Radius{}, err
	}

	if value < 0 {
		return Radius{}, errors.New("the radius value must be greater than or equal to zero")
	}

	return Radius{R: json.Number(radius)}, nil
}

type Circle struct {
	Coordinates
	Radius
}

func NewCircle(coordinates Coordinates, radius Radius) (Circle, error) {
	return Circle{
		Coordinates: coordinates,
		Radius:      radius,
	}, nil
}
