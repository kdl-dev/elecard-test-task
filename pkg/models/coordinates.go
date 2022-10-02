package models

import (
	"encoding/json"
	"strconv"
)

type Coordinates struct {
	X json.Number `json:"x"`
	Y json.Number `json:"y"`
}

func NewCoordinates(x, y string) (Coordinates, error) {
	if _, err := strconv.ParseFloat(x, 64); err != nil {
		return Coordinates{}, err
	}

	if _, err := strconv.ParseFloat(y, 64); err != nil {
		return Coordinates{}, err
	}

	return Coordinates{
		X: json.Number(x),
		Y: json.Number(y),
	}, nil
}
