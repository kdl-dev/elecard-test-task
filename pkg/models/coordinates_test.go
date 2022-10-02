package models

import "testing"

func TestNewCoordinates(t *testing.T) {
	x1 := "123.245"
	y1 := "346.999"

	x2 := "123,245"
	y2 := "346,999"

	x3 := "Hi"
	y3 := "Elecard"

	coord, err := NewCoordinates(x1, y1)
	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	if coord.X.String() != x1 || coord.Y.String() != y1 {
		t.Errorf("wrong coordinates: %v, expected (%s;%s)\n", coord, x1, y1)
	}

	coord, err = NewCoordinates(x2, y2)
	if err == nil {
		t.Errorf("expected error, but none (%v)\n", coord)
	}

	coord, err = NewCoordinates(x3, y3)
	if err == nil {
		t.Errorf("expected error, but none (%v)\n", coord)
	}
}
