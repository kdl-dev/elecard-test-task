package models

import (
	"testing"
)

func TestNewRadius(t *testing.T) {
	r1 := "5"
	r2 := "-5"
	r3 := "Hi Elecard"

	radius, err := NewRadius(r1)
	if err != nil {
		t.Errorf("unexpected error: %v\n", err)
	}

	if radius.R.String() != r1 {
		t.Errorf("wrong radius: %s, expected %s\n", radius.R, r1)
	}

	radius, err = NewRadius(r2)
	if err == nil {
		t.Errorf("expected error, but none (%v)\n", radius)
	}

	radius, err = NewRadius(r3)
	if err == nil {
		t.Errorf("expected error, but none (%v)\n", radius)
	}
}
